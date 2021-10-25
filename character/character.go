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
	"sync"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/craft"
	"github.com/isangeles/flame/dialog"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/flag"
	"github.com/isangeles/flame/item"
	"github.com/isangeles/flame/objects"
	"github.com/isangeles/flame/quest"
	"github.com/isangeles/flame/serial"
	"github.com/isangeles/flame/skill"
	"github.com/isangeles/flame/training"
	"github.com/isangeles/flame/useaction"
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
	respawn          int64
	areaID           string
	inventory        *item.Inventory
	equipment        *Equipment
	journal          *quest.Journal
	crafting         *craft.Crafting
	targets          []res.SerialObjectData
	kills            []res.KillData
	effects          map[string]*effect.Effect
	skills           map[string]*skill.Skill
	memory           *sync.Map
	dialogs          *sync.Map
	flags            *sync.Map
	trainings        []*training.TrainerTraining
	casted           res.CastedObjectData
	chatLog          *objects.Log
	onSkillActivated func(s *skill.Skill)
	onChatSent       func(t string)
	onModifierTaken  func(m effect.Modifier)
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
		memory:     new(sync.Map),
		dialogs:    new(sync.Map),
		flags:      new(sync.Map),
		chatLog:    objects.NewLog(),
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
	if ob := c.Casted(); ob != nil {
		time := ob.UseAction().Cast() + delta
		ob.UseAction().SetCast(time)
		if time >= ob.UseAction().CastMax() {
			c.useCasted(ob)
			c.casted.ID = ""
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

// Health returns current value of health
// points.
func (c *Character) Health() int {
	return c.hp
}

// MaxHealth returns maximal value of
// health points.
func (c *Character) MaxHealth() int {
	return c.attributes.Health() + (BaseHealth * c.Level())
}

// Mana returns current value of mana
// points.
func (c *Character) Mana() int {
	return c.mana
}

// MaxMana returns maximal value of mana
// points.
func (c *Character) MaxMana() int {
	return c.attributes.Mana() + (BaseMana * c.Level() / 2)
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
	ob, _ := c.memory.Load(o.ID()+o.Serial())
	mem, ok := ob.(*TargetMemory)
	if ok {
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
	c.live = c.hp > 0
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
// position and destination point of the character.
func (c *Character) SetPosition(x, y float64) {
	c.posX, c.posY = x, y
	c.SetDestPoint(x, y)
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
	// Update ownerships.
	for _, s := range c.Skills() {
		s.UseAction().SetOwner(c)
	}
	for _, r := range c.Crafting().Recipes() {
		r.UseAction().SetOwner(c)
	}
	for _, t := range c.Trainings() {
		t.UseAction().SetOwner(c)
	}
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
	s.UseAction().SetOwner(c)
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
		c.targets = []res.SerialObjectData{}
		return
	}
	tarData := res.SerialObjectData{t.ID(), t.Serial()}
	c.targets = []res.SerialObjectData{tarData}
}

// Moving checks whether character is moving.
func (c *Character) Moving() bool {
	return c.Live() && (c.posX != c.destX || c.posY != c.destY)
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
func (c *Character) SetOnModifierTakenFunc(f func(m effect.Modifier)) {
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
	c.casted.ID = ""
}

// Dialog returns dialog for specified character.
func (c *Character) Dialog(char Character) (dial *dialog.Dialog) {
	// TODO: find proper dialog for specified character.
	findDialog := func(k, v interface{}) bool {
		d, ok := v.(*dialog.Dialog)
		if ok {
			dial = d
		}
		return true
	}
	c.dialogs.Range(findDialog)
	return
}

// Dialogs returns all character dialogs.
func (c *Character) Dialogs() (dls []*dialog.Dialog) {
	addDialog := func(k, v interface{}) bool {
		d, ok := v.(*dialog.Dialog)
		if ok {
			dls = append(dls, d)
		}
		return true
	}
	c.dialogs.Range(addDialog)
	return
}

// AddDialog adds specified dialog to character and
// sets character as dialog owner.
func (c *Character) AddDialog(d *dialog.Dialog) {
	d.SetOwner(c)
	c.dialogs.Store(d.ID(), d)
}

// Flags returns all active flags.
func (c *Character) Flags() (flags []flag.Flag) {
	addFlag := func(k, v interface{}) bool {
		f, ok := v.(flag.Flag)
		if ok {
			flags = append(flags, f)
		}
		return true
	}
	c.flags.Range(addFlag)
	return
}

// AddFlag adds specified flag.
func (c *Character) AddFlag(f flag.Flag) {
	c.flags.Store(f.ID(), f)
}

// RemoveFlag removes specified flag.
func (c *Character) RemoveFlag(f flag.Flag) {
	c.flags.Delete(f.ID())
}

// HasFlag checks if character has specified flag.
func (c *Character) HasFlag(flag flag.Flag) bool {
	_, ok := c.flags.Load(flag.ID())
	return ok
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
	t.UseAction().SetOwner(c)
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

// SetRespawn sets character respawn time in milliseconds.
func (c *Character) SetRespawn(respawn int64) {
	c.respawn = respawn
}

// Respawn returns time of character respawn in milliseconds.
func (c *Character) Respawn() int64 {
	return c.respawn
}

// Casted returns usable object currently casted by the character.
func (c *Character) Casted() useaction.Usable {
	if len(c.casted.ID) < 1 {
		return nil
	}
	// Owner.
	var owner serial.Serialer
	if c.casted.Owner.ID == c.ID() && c.casted.Owner.Serial == c.Serial() {
		owner = c
	} else {
		owner = serial.Object(c.casted.Owner.ID, c.casted.Owner.Serial)
	}
	if owner == nil {
		return nil
	}
	if c.casted.ID == owner.ID() {
		return owner.(useaction.Usable)
	}
	// Skills.
	if owner, ok := owner.(skill.User); ok {
		for _, s := range owner.Skills() {
			if s.ID() == c.casted.ID {
				return s
			}
		}
	}
	// Items.
	if owner, ok := owner.(item.Container); ok {
		for _, i := range owner.Inventory().Items() {
			if i.ID() == c.casted.ID {
				return owner.(useaction.Usable)
			}
		}
	}
	// Recipes.
	if owner, ok := owner.(craft.Crafter); ok {
		for _, r := range owner.Crafting().Recipes() {
			if r.ID() == c.casted.ID {
				return r
			}
		}
	}
	// Trainings.
	if owner, ok := owner.(training.Trainer); ok {
		for _, t := range owner.Trainings() {
			if t.ID() == c.casted.ID {
				return t
			}
		}
	}
	return nil
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
