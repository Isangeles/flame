/*
 * character.go
 *
 * Copyright 2018-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/module/craft"
	"github.com/isangeles/flame/module/dialog"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/flag"
	"github.com/isangeles/flame/module/item"
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/quest"
	"github.com/isangeles/flame/module/serial"
	"github.com/isangeles/flame/module/skill"
	"github.com/isangeles/flame/module/training"
	"github.com/isangeles/flame/module/useaction"
)

// Character struct represents game character.
type Character struct {
	id, serial       string
	level            int
	hp               int
	mana             int
	exp              int
	live             bool
	agony            bool
	ai               bool
	sex              Gender
	race             *Race
	attitude         Attitude
	alignment        Alignment
	guild            Guild
	attributes       *Attributes
	resilience       objects.Resilience
	posX, posY       float64
	destX, destY     float64
	defX, defY       float64
	cooldown         int64 // millis
	areaID           string
	inventory        *item.Inventory
	equipment        *Equipment
	journal          *quest.Journal
	crafting         *craft.Crafting
	targets          []res.CharacterTargetData
	effects          map[string]*effect.Effect
	skills           map[string]*skill.Skill
	memory           map[string]*TargetMemory
	dialogs          map[string]*dialog.Dialog
	flags            map[string]flag.Flag
	trainings        []*training.TrainerTraining
	casted           useaction.Usable
	chatLog          *objects.Log
	privateLog       *objects.Log
	onSkillActivated func(s *skill.Skill)
	onChatSent       func(t string)
	onModifierTaken  func(m *effect.Modifier)
	onEffectTaken    func(e *effect.Effect)
}

const (
	baseExp  = 1000
	globalCD = 2000 // millis
)

// New creates new character from specified data.
func New(data res.CharacterData) *Character {
	c := Character{
		attributes: new(Attributes),
		inventory:  item.NewInventory(),
		effects:    make(map[string]*effect.Effect),
		skills:     make(map[string]*skill.Skill),
		memory:     make(map[string]*TargetMemory),
		dialogs:    make(map[string]*dialog.Dialog),
		flags:      make(map[string]flag.Flag),
		chatLog:    objects.NewLog(),
		privateLog: objects.NewLog(),
	}
	c.equipment = newEquipment(&c)
	c.journal = quest.NewJournal(&c)
	c.crafting = craft.NewCrafting(&c)
	c.Apply(data)
	// Register serial.
	serial.Register(&c)
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
		c.Interrupt() // interrupt current action
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
	// Journal && inventory.
	c.Journal().Update(delta)
	c.Inventory().Update(delta)
	c.Inventory().SetCapacity(c.Attributes().Lift())
	// Skills.
	for _, s := range c.skills {
		s.Update(delta)
	}
	// Effects.
	for serial, e := range c.effects {
		e.Update(delta)
		// Remove expired effects.
		if e.Time() <= 0 {
			delete(c.effects, serial)
		}
	}
	// Recipes.
	for _, r := range c.Crafting().Recipes() {
		r.Update(delta)
	}
	// Casting action.
	if c.Casted() != nil {
		time := c.Casted().UseAction().Cast() + delta
		c.Casted().UseAction().SetCast(time)
		if time >= c.Casted().UseAction().CastMax() {
			c.useCasted(c.Casted())
			c.casted = nil
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

// Level returns character level.
func (c *Character) Level() int {
	return c.level
}

// AI checks if character is controlled by AI.
func (c *Character) AI() bool {
	return c.ai
}

// SetAI marks or unmarks character as controled by AI.
func (c *Character) SetAI(ai bool) {
	c.ai = ai
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
	return baseExp * c.Level()
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

// SetGander sets character gender.
func (c *Character) SetGender(gender Gender) {
	c.sex = gender
}

// Race returns character race.
func (c *Character) Race() *Race {
	return c.race
}

// Attitude returns character attitude.
func (c *Character) Attitude() Attitude {
	return c.attitude
}

// SetAttitude sets character attitude.
func (c *Character) SetAttitude(att Attitude) {
	c.attitude = att
}

// AttitudeFor returns attitude for specified objects.
func (c *Character) AttitudeFor(o objects.Object) Attitude {
	mem := c.memory[o.ID()+o.Serial()]
	if mem != nil {
		return mem.Attitude
	}
	obChar, ok := o.(*Character)
	if !ok || !obChar.Live() {
		return Neutral
	}
	if obChar.Attitude() == Hostile {
		return Hostile
	}
	return c.Attitude()
}

// Guild returns character guild.
func (c *Character) Guild() Guild {
	return c.guild
}

// Attributes returns character attributes.
func (c *Character) Attributes() *Attributes {
	return c.attributes
}

// Alignment returns character alignment
func (c *Character) Alignment() Alignment {
	return c.alignment
}

// SetAlignment sets character alignment.
func (c *Character) SetAlignment(ali Alignment) {
	c.alignment = ali
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
// position of the character.
func (c *Character) SetPosition(x, y float64) {
	c.posX, c.posY = x, y
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
	c.skills[s.ID()] = s
}

// RemoveSkill removes specified skill.
func (c *Character) RemoveSkill(s *skill.Skill) {
	delete(c.skills, s.ID())
}

// Targets returns character targets.
func (c *Character) Targets() (targets []effect.Target) {
	for _, td := range c.targets {
		ob := serial.Object(td.ID, td.Serial)
		if ob == nil {
			continue
		}
		tar, ok := ob.(effect.Target)
		if ok {
			targets = append(targets, tar)
		}
	}
	return
}

// SetTarget sets specified 'targetable' as current
// target.
func (c *Character) SetTarget(t effect.Target) {
	if t == nil {
		c.targets = []res.CharacterTargetData{}
		return
	}
	tarData := res.CharacterTargetData{t.ID(), t.Serial()}
	c.targets = []res.CharacterTargetData{tarData}
}

// Moving checks whether character is moving.
func (c *Character) Moving() bool {
	if c.posX != c.destX || c.posY != c.destY {
		return true
	} else {
		return false
	}
}

// Fighting checks if character is in combat.
func (c *Character) Fighting() bool {
	if len(c.Targets()) > 0 {
		tar := c.Targets()[0]
		if c.AttitudeFor(tar) != Hostile {
			return false
		}
		tarPos, ok := tar.(objects.Positioner)
		if !ok {
			return false
		}
		return objects.Range(c, tarPos) <= c.SightRange()
	}
	return false
}

// ChatLog returns character speech log channel.
func (c *Character) ChatLog() *objects.Log {
	return c.chatLog
}

// PrivateLog returns character private log channel.
func (c *Character) PrivateLog() *objects.Log {
	return c.privateLog
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

// SetOnModifierTakenFunc sets function triggered after receiving new
// modifier.
func (c *Character) SetOnModifierTakenFunc(f func(m *effect.Modifier)) {
	c.onModifierTaken = f
}

// SetOnEffectTakenFunc sets function triggered after taking new
// effect.
func (c *Character) SetOnEffectTakenFunc(f func(e *effect.Effect)) {
	c.onEffectTaken = f
}

// Interrupt stops any acction(like skill
// casting) performed by character.
func (c *Character) Interrupt() {
	c.casted = nil
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
	log.Dbg.Printf("char: %s#%s: add flag: %s", c.ID(), c.Serial(), f)
}

// RemoveFlag removes specified flag.
func (c *Character) RemoveFlag(f flag.Flag) {
	delete(c.flags, f.ID())
	log.Dbg.Printf("char: %s#%s: remove flag: %s", c.ID(), c.Serial(), f)
}

// Journal returns quest journal.
func (c *Character) Journal() *quest.Journal {
	return c.journal
}

// Crafting returns character crafting object.
func (c *Character) Crafting() *craft.Crafting {
	return c.crafting
}

// Trainings returns all trainings.
func (c *Character) Trainings() []*training.TrainerTraining {
	return c.trainings
}

// AddTrainings adds specified training to
// character trainings list.
func (c *Character) AddTraining(t *training.TrainerTraining) {
	c.trainings = append(c.trainings, t)
}

// AreaID returns ID of character area.
func (c *Character) AreaID() string {
	return c.areaID
}

// SetAreaID sets specifid area ID as
// character area ID.
func (c *Character) SetAreaID(areaID string) {
	c.areaID = areaID
}

// Casted returns casted action.
func (c *Character) Casted() useaction.Usable {
	return c.casted
}

// levelup promotes character to next level.
func (c *Character) levelup() {
	c.level += 1
	c.SetHealth(c.MaxHealth())
	c.SetMana(c.MaxMana())
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
		e.SetSource(c.ID(), c.Serial())
		effects = append(effects, e)
	}
	return effects
}
