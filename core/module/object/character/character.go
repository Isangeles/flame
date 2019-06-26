/*
 * character.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 *
 *
 */

// character package provides game character struct and
// other types for game characters.
package character

import (
	"fmt"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/core/module/object/craft"
	"github.com/isangeles/flame/core/module/object/dialog"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/object/item"
	"github.com/isangeles/flame/core/module/object/skill"
	"github.com/isangeles/flame/core/module/object/quest"
	"github.com/isangeles/flame/core/module/flag"
	"github.com/isangeles/flame/core/module/serial"
	"github.com/isangeles/flame/log"
)

// Character struct represents game character.
type Character struct {
	id, serial, name string
	level            int
	hp, maxHP        int
	mana, maxMana    int
	exp, maxExp      int
	live             bool
	agony            bool
	sex              Gender
	race             Race
	attitude         Attitude
	alignment        Alignment
	guild            Guild
	attributes       Attributes
	resilience       object.Resilience
	posX, posY       float64
	destX, destY     float64
	defX, defY       float64
	cooldown         int64 // millis
	inventory        *item.Inventory
	equipment        *Equipment
	journal          *quest.Journal
	targets          []effect.Target
	effects          map[string]*effect.Effect
	skills           map[string]*skill.Skill
	memory           map[string]*TargetMemory
	dialogs          map[string]*dialog.Dialog
	recipes          map[string]*craft.Recipe
	flags            map[string]flag.Flag
	chatlog          chan string
	combatlog        chan string
	privlog          chan string
	onSkillActivated func(s *skill.Skill)
	onChatSent       func(t string)
}

const (
	base_exp  = 1000
	global_cd = 2000 // millis
)

// New creates new character from specified data.
func New(data res.CharacterBasicData) *Character {
	c := Character{
		id:        data.ID,
		serial:    data.Serial,
		name:      data.Name,
		sex:       Gender(data.Sex),
		race:      Race(data.Race),
		attitude:  Attitude(data.Attitude),
		alignment: Alignment(data.Alignment),
	}
	c.attributes = Attributes{data.Str, data.Con, data.Dex, data.Int, data.Wis}
	c.live = true
	c.inventory = item.NewInventory(c.Attributes().Lift())
	c.equipment = newEquipment(&c)
	c.journal = quest.NewJournal(&c)
	c.targets = make([]effect.Target, 1)
	c.effects = make(map[string]*effect.Effect)
	c.skills = make(map[string]*skill.Skill)
	c.memory = make(map[string]*TargetMemory)
	c.dialogs = make(map[string]*dialog.Dialog)
	c.flags = make(map[string]flag.Flag)
	c.recipes = make(map[string]*craft.Recipe)
	c.chatlog = make(chan string, 1)
	c.combatlog = make(chan string, 3)
	c.privlog = make(chan string, 3)
	// Set level.
	for i := 0; i < data.Level; i++ {
		oldMaxExp := c.MaxExperience()
		c.levelup()
		c.SetExperience(oldMaxExp)
	}
	// Add flags.
	for _, fd := range data.Flags {
		f := flag.Flag(fd.ID)
		c.flags[f.ID()] = f
	}
	return &c
}

// Update updates character.
func (c *Character) Update(delta int64) {
	// Global cooldown.
	if c.cooldown > 0 {
		c.cooldown -= delta
	}
	if c.cooldown < 0 {
		c.cooldown = 0
	}
	// Move to dest point.
	if c.Moving() {
		c.Interrupt() // interrupt current acction
		if c.posX < c.destX {
			c.posX += 1
		}
		if c.posX > c.destX {
			c.posX -= 1
		}
		if c.posY < c.destY {
			c.posY += 1
		}
		if c.posY > c.destY {
			c.posY -= 1
		}
	}
	// Check experience value.
	if c.Experience() >= c.MaxExperience() {
		c.levelup()
	}
	// Check health value.
	if c.Health() <= c.agonyHP() {
		c.agony = true
	} else if c.Agony() {
		c.agony = false
	}
	if c.Health() <= 0 {
		c.live = false
	} else if !c.Live() {
		c.live = true
	}
	// Journal.
	c.Journal().Update(delta)
	// Effects.
	for serial, e := range c.effects {
		e.Update(delta)
		// Remove expired effects.
		if e.Time() <= 0 {
			delete(c.effects, serial)
		}
	}
	// Skills.
	for _, s := range c.Skills() {
		s.Update(delta)
		if s.Casted() {
			s.Activate()
			if c.onSkillActivated != nil {
				c.onSkillActivated(s)
			}
			c.cooldown = global_cd
		}
	}
}

// ID returns character ID.
func (c *Character) ID() string {
	return c.id
}

// Serial returns serial value.
func (c *Character) Serial() string {
	return c.serial
}

// SerialID returns character ID and serial value
// in form: [ID]_[serial].
func (c *Character) SerialID() string {
	return fmt.Sprintf("%s_%s", c.ID(), c.serial)
}

// Name returns character name.
func (c *Character) Name() string {
	return c.name
}

// Level returns character level.
func (c *Character) Level() int {
	return c.level
}

// Health returns current value of health
// points.
func (c *Character) Health() int {
	return c.hp
}

// MaxHealth returns maximal value of
// health points.
func (c *Character) MaxHealth() int {
	return c.attributes.Health() + (Base_health * c.Level())
}

// Mana returns current value of mana
// points.
func (c *Character) Mana() int {
	return c.mana
}

// MaxMana returns maximal value of mana
// points.
func (c *Character) MaxMana() int {
	return c.attributes.Mana() + (Base_mana * c.Level() / 2)
}

// Experience returns current value of experience
// points.
func (c *Character) Experience() int {
	return c.exp
}

// MaxExperience returns maximal value of
// experience points.
func (c *Character) MaxExperience() int {
	return c.maxExp
}

// Live checks wheter character is live.
func (c *Character) Live() bool {
	return c.live
}

// Agony check wheter character is in
// agony state.
func (c *Character) Agony() bool {
	return c.agony
}

// Gender returns character gender.
func (c *Character) Gender() Gender {
	return c.sex
}

// Race returns character race.
func (c *Character) Race() Race {
	return c.race
}

// Attitude returns character attitude.
func (c *Character) Attitude() Attitude {
	return c.attitude
}

// AttitudeFor returns attitude for specified target.
func (c *Character) AttitudeFor(tar effect.Target) Attitude {
	mem := c.memory[tar.ID() + tar.Serial()]
	if mem != nil {
		return mem.Attitude
	}
	char, ok := tar.(*Character)
	if !ok {
		return Neutral
	}
	if char.Attitude() == Hostile {
		return Hostile
	}
	return c.Attitude()
}

// Guild returns character guild.
func (c *Character) Guild() Guild {
	return c.guild
}

// Attributes returns character attributes.
func (c *Character) Attributes() Attributes {
	return c.attributes
}

// Alignment returns character alignment
func (c *Character) Alignment() Alignment {
	return c.alignment
}

// Position returns current character position.
func (c *Character) Position() (float64, float64) {
	return c.posX, c.posY
}

// DestPoint returns current destination point position.
func (c *Character) DestPoint() (float64, float64) {
	return c.destX, c.destY
}

// DefaultPosition returns character default position.
func (c *Character) DefaultPosition() (float64, float64) {
	return c.defX, c.defY
}

// SightRange returns current sight range.
func (c *Character) SightRange() float64 {
	return c.attributes.Sight()
}

// Inventory returns character inventory.
func (c *Character) Inventory() *item.Inventory {
	return c.inventory
}

// Equipment returns character equipment.
func (c *Character) Equipment() *Equipment {
	return c.equipment
}

// SetName sets specified text as character
// display name.
func (c *Character) SetName(name string) {
	c.name = name
}

// SetHealth sets specified value as current
// amount of health points.
func (c *Character) SetHealth(hp int) {
	c.hp = hp
}

// SetMana sets specified value as current
// amount of mana points.
func (c *Character) SetMana(mana int) {
	c.mana = mana
}

// SetExperience sets specified value as current
// amount of experience points.
func (c *Character) SetExperience(exp int) {
	c.exp = exp
}

// SetPosition sets specified XY position as current
// position of character and current destination point.
func (c *Character) SetPosition(x, y float64) {
	c.posX, c.posY = x, y
	c.destX, c.destY = x, y
}

// SetDestPoint sets specified XY position as current
// destionation point of character.
func (c *Character) SetDestPoint(x, y float64) {
	c.destX, c.destY = x, y
}

// SetDefaultPosition sets specified XY position as
// default character position.
func (c *Character) SetDefaultPosition(x, y float64) {
	c.defX, c.defY = x, y
}

// SetSerial sets specified serial value for this
// character.
func (c *Character) SetSerial(serial string) {
	c.serial = serial
}

// Effects returns character all effects.
func (c *Character) Effects() []*effect.Effect {
	effs := make([]*effect.Effect, 0)
	for _, e := range c.effects {
		effs = append(effs, e)
	}
	return effs
}

// AddEffect add specified effect to character effects.
func (c *Character) AddEffect(e *effect.Effect) {
	e.SetTarget(c)
	c.effects[e.ID()+e.Serial()] = e
}

// RemoveEffect removes effect from character.
func (c *Character) RemoveEffect(e *effect.Effect) {
	delete(c.effects, e.ID()+e.Serial())
}

// Skills return all character skills.
func (c *Character) Skills() []*skill.Skill {
	skills := make([]*skill.Skill, 0)
	for _, s := range c.skills {
		skills = append(skills, s)
	}
	return skills
}

// AddSkill adds specified skill to characters
// skills.
func (c *Character) AddSkill(s *skill.Skill) {
	c.skills[s.ID()+s.Serial()] = s
}

// RemoveSkill removes specified skill.
func (c *Character) RemoveSkill(s *skill.Skill) {
	delete(c.skills, s.ID()+s.Serial())
}

// Targets returns character targets.
func (c *Character) Targets() []effect.Target {
	return c.targets
}

// SetTarget sets specified 'targetable' as current
// target.
func (c *Character) SetTarget(t effect.Target) {
	c.targets[0] = t
}

// Moving checks whether character is moving.
func (c *Character) Moving() bool {
	if c.posX != c.destX || c.posY != c.destY {
		return true
	} else {
		return false
	}
}

// Casting checks whether any of character
// skills is casted.
func (c *Character) Casting() bool {
	for _, s := range c.Skills() {
		if s.Casting() {
			return true
		}
	}
	return false
}

// Fighting checks if character is in combat.
func (c *Character) Fighting() bool {
	tar := c.Targets()[0]
	if tar != nil && c.AttitudeFor(tar) == Hostile {
		return object.Range(c, tar) <= c.SightRange()
	}
	return false
}

// CombatLog returns character combat log channel.
func (c *Character) CombatLog() chan string {
	return c.combatlog
}

// ChatLog returns character speech log channel.
func (c *Character) ChatLog() chan string {
	return c.chatlog
}

// PrivateLog returns character private log channel.
func (c *Character) PrivateLog() chan string {
	return c.privlog
}

// SetOnSkillActivatedFunc sets function triggered after
// activation one of character skills.
func (c *Character) SetOnSkillActivatedFunc(f func(s *skill.Skill)) {
	c.onSkillActivated = f
}

// SetOnChatSentFunc sets function triggered after sending text
// on character chat channel.
func (c *Character) SetOnChatSentFunc(f func(t string)) {
	c.onChatSent = f
}

// Interrupt stops any acction(like skill
// casting) performed by character.
func (c *Character) Interrupt() {
	for _, s := range c.skills {
		if s.Casting() {
			s.StopCast()
		}
	}
}

// SendChat sends specified text to character
// speech log channel.
func (c *Character) SendChat(t string) {
	select {
	case c.chatlog <- t:
		if c.onChatSent != nil {
			c.onChatSent(t)
		}
	default:
	}
}

// SendCombat sends specified text message to
// comabt log channel.
func (c *Character) SendCombat(t string) {
	select {
	case c.combatlog <- t:
	default:
	}
}

// SendPrivate sends specified text to character
// private log.
func (c *Character) SendPrivate(t string) {
	select {
	case c.privlog <- t:
	default:
	}
}

// Dialog returns dialog for specified character.
func (c *Character) Dialog(char Character) *dialog.Dialog {
	// TODO: find proper dialog for specified character.
	for _, d := range c.dialogs {
		return d
	}
	return nil
}

// Dialogs returns all character dialogs.
func (c *Character) Dialogs() (dls []*dialog.Dialog) {
	for _, d := range c.dialogs {
		dls = append(dls, d)
	}
	return
}

// AddDialog adds specified dialog to character and
// sets character as dialog owner.
func (c *Character) AddDialog(d *dialog.Dialog) {
	d.SetOwner(c)
	c.dialogs[d.ID()] = d
}

// Flags returns all active flags.
func (c *Character) Flags() (flags []flag.Flag) {
	for _, f := range c.flags {
		flags = append(flags, f)
	}
	return
}

// AddFlag adds specified flag.
func (c *Character) AddFlag(f flag.Flag) {
	c.flags[f.ID()] = f
	log.Dbg.Printf("char:%s_%s:add_flag:%s", c.ID(), c.Serial(), f)
}

// RemoveFlag removes specified flag.
func (c *Character) RemoveFlag(f flag.Flag) {
	delete(c.flags, f.ID())
	log.Dbg.Printf("char:%s_%s:remove_flag:%s", c.ID(), c.Serial(), f)
}

// Journal returns quest journal.
func (c *Character) Journal() *quest.Journal {
	return c.journal
}

// Recipes returns character recipes.
func (c *Character) Recipes() (recipes []*craft.Recipe) {
	for _, r := range c.recipes {
		recipes = append(recipes, r)
	}
	return
}

// AddRecipe adds specified recipe to character.
func (c *Character) AddRecipe(r *craft.Recipe) {
	c.recipes[r.ID()] = r
}

// levelup promotes character to next level.
func (c *Character) levelup() {
	c.level += 1
	c.SetHealth(c.MaxHealth())
	c.SetMana(c.MaxMana())
	c.maxExp = base_exp * c.Level()
}

// agonyHP returns value of health causing
// agony state.
func (c *Character) agonyHP() int {
	return 10 / 100 * c.MaxHealth()
}

// buildEffects creates new effects from specified
// data with character as a source.
func (c *Character) buildEffects(effectsData ...res.EffectData) []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	for _, ed := range effectsData {
		e := effect.New(ed)
		e.SetSource(c)
		serial.AssignSerial(e)
		effects = append(effects, e)
	}
	return effects
}
