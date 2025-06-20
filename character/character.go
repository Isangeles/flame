/*
 * character.go
 *
 * Copyright 2018-2025 Dariusz Sikora <ds@isangeles.dev>
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
	"sync"

	"github.com/isangeles/flame/craft"
	"github.com/isangeles/flame/data/res"
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
	id, serial      string
	level           int
	hp              int
	mana            int
	exp             int
	live            bool
	agony           bool
	openLoot        bool
	sex             Gender
	race            Race
	attitude        Attitude
	alignment       Alignment
	guild           Guild
	attributes      *Attributes
	resilience      objects.Resilience
	action          *useaction.UseAction
	posX, posY      float64
	destX, destY    float64
	defX, defY      float64
	useCooldown     int64 // millis
	moveCooldown    int64 // millis
	respawn         int64
	areaID          string
	chapterID       string
	inventory       *item.Inventory
	equipment       *Equipment
	journal         *quest.Journal
	crafting        *craft.Crafting
	targets         []res.SerialObjectData
	kills           []res.KillData
	effects         *sync.Map
	skills          *sync.Map
	memory          *sync.Map
	dialogs         *sync.Map
	startedDialogs  *sync.Map
	flags           *sync.Map
	trainings       []*training.TrainerTraining
	casted          res.CastedObjectData
	chatLog         *objects.Log
	onModifierTaken func(m effect.Modifier)
}

const (
	baseExp               = 1000
	useCD                 = 2000 // millis
	moveCD                = 15   // millis
	startedDialogIDFormat = "%s_%s#%s" // [dialog ID]_[object ID]#[object serial]
)

// New creates new character from specified data.
func New(data res.CharacterData) *Character {
	c := Character{
		attributes:     new(Attributes),
		inventory:      item.NewInventory(),
		effects:        new(sync.Map),
		skills:         new(sync.Map),
		memory:         new(sync.Map),
		dialogs:        new(sync.Map),
		startedDialogs: new(sync.Map),
		flags:          new(sync.Map),
		chatLog:        objects.NewLog(),
		race:           NewRace(res.RaceData{}),
	}
	c.equipment = newEquipment(&c)
	c.journal = quest.NewJournal(&c)
	c.crafting = craft.NewCrafting(&c)
	c.Inventory().SetOnItemRemovedFunc(c.removeItem)
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
	// Cooldowns.
	if c.useCooldown > 0 {
		c.useCooldown -= delta
	}
	if c.moveCooldown > 0 {
		c.moveCooldown -= delta
	}
	// Check experience value.
	if c.Experience() >= c.MaxExperience() {
		c.levelup()
	}
	// Journal && inventory.
	c.Journal().Update(delta)
	c.Inventory().Update(delta)
	// Dialogs.
	c.startedDialogs.Range(c.removeFinishedDialog)
	// Skills.
	for _, s := range c.Skills() {
		s.Update(delta)
	}
	// Effects.
	for _, e := range c.Effects() {
		e.Update(delta)
		// Remove expired effects.
		if e.Time() <= 0 && !e.Infinite() {
			c.effects.Delete(e.ID() + e.Serial())
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
	// Use action.
	if c.UseAction() != nil {
		c.UseAction().Update(delta)
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

// OpenLoot checks whether it should be possible
// to loot the character at any time.
func (c *Character) OpenLoot() bool {
	return c.openLoot
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
func (c *Character) Race() Race {
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

// AttitudeFor returns attitude for specified object.
// For dead objects returns neutral attitude.
// For memorized targets returns memorized attitude.
// For not memorized objects from the same guild returns
// friendly attitude.
// For not memorized hostile objects returns hostile attitude.
// For not memorized neutral objects returns neutral attitude.
// For other cases returns character attitude.
func (c *Character) AttitudeFor(o serial.Serialer) Attitude {
	char, ok := o.(*Character)
	if ok && !char.Live() {
		return Neutral
	}
	ob, _ := c.memory.Load(o.ID() + o.Serial())
	mem, ok := ob.(*TargetMemory)
	if ok {
		return mem.Attitude
	}
	if char == nil {
		return c.Attitude()
	}
	if len(char.Guild().ID()) > 0 && char.Guild().ID() == c.Guild().ID() {
		return Friendly
	}
	if char.Attitude() == Neutral {
		return Neutral
	}
	if char.Attitude() == Hostile {
		return Hostile
	}
	return c.Attitude()
}

// SetGuild sets character guild.
func (c *Character) SetGuild(guild Guild) {
	c.guild = guild
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
// position.
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
	// Update ownerships.
	if c.UseAction() != nil {
		c.UseAction().SetOwner(c)
	}
	for _, s := range c.Skills() {
		if s.UseAction() != nil {
			s.UseAction().SetOwner(c)
		}
	}
	for _, r := range c.Crafting().Recipes() {
		r.UseAction().SetOwner(c)
	}
	for _, t := range c.Trainings() {
		t.UseAction().SetOwner(c)
	}
}

// Effects returns character all effects.
func (c *Character) Effects() (effects []*effect.Effect) {
	addEffect := func(k, v interface{}) bool {
		e, ok := v.(*effect.Effect)
		if ok {
			effects = append(effects, e)
		}
		return true
	}
	c.effects.Range(addEffect)
	return
}

// AddEffect add specified effect to character effects.
func (c *Character) AddEffect(e *effect.Effect) {
	e.SetTarget(c)
	c.effects.Store(e.ID()+e.Serial(), e)
}

// RemoveEffect removes effect from character.
func (c *Character) RemoveEffect(e *effect.Effect) {
	c.effects.Delete(e.ID() + e.Serial())
}

// Skills return all character skills.
func (c *Character) Skills() (skills []*skill.Skill) {
	addSkill := func(k, v interface{}) bool {
		s, ok := v.(*skill.Skill)
		if ok {
			skills = append(skills, s)
		}
		return true
	}
	c.skills.Range(addSkill)
	return
}

// AddSkill adds specified skill to characters
// skills.
func (c *Character) AddSkill(s *skill.Skill) {
	s.SetOwner(c)
	c.skills.Store(s.ID(), s)
}

// RemoveSkill removes specified skill.
func (c *Character) RemoveSkill(s *skill.Skill) {
	c.skills.Delete(s.ID())
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

// Action returns character use action.
func (c *Character) UseAction() *useaction.UseAction {
	return c.action
}

// Moving checks whether character is moving.
func (c *Character) Moving() bool {
	return c.Live() && c.MoveCooldown() < 1 && (c.posX != c.destX || c.posY != c.destY)
}

// Fighting checks if character is in combat.
func (c *Character) Fighting() bool {
	if len(c.Targets()) < 1 {
		return false
	}
	tar := c.Targets()[0]
	if c.AttitudeFor(tar) != Hostile {
		return false
	}
	return objects.Range(c, tar) <= c.SightRange()
}

// ChatLog returns character speech log channel.
func (c *Character) ChatLog() *objects.Log {
	return c.chatLog
}

// SetOnModifierTakenFunc sets function triggered after receiving new
// modifier.
// The event function will be called for every recieved modifier
// after it was handled by the character.
func (c *Character) SetOnModifierTakenFunc(f func(m effect.Modifier)) {
	c.onModifierTaken = f
}

// Interrupt stops any acction(like skill
// casting) performed by character.
func (c *Character) Interrupt() {
	c.casted.ID = ""
}

// Cooldown returns character cooldown in milliseconds.
func (c *Character) Cooldown() int64 {
	return c.useCooldown
}

// BaseMoveCooldown returns character base movement
// cooldown in milliseconds.
func (c *Character) BaseMoveCooldown() int64 {
	return moveCD - int64((1+c.Attributes().Dex)/4)
}

// MoveCooldown returns character movement cooldown in
// milliseconds.
func (c *Character) MoveCooldown() int64 {
	return c.moveCooldown
}

// SetMoveCooldown sets character movement cooldown to
// the specified value(in milliseconds).
func (c *Character) SetMoveCooldown(cooldown int64) {
	c.moveCooldown = cooldown
}

// Dialog returns dialog for specified character.
// If there is already dialog in-progress started by specified character,
// then this dialog will be retruned, otherwise the new dialog will be returned.
func (c *Character) Dialog(ob dialog.Talker) (dial *dialog.Dialog) {
	findDialog := func(k, v interface{}) bool {
		d, ok := v.(*dialog.Dialog)
		if ok && k == fmt.Sprintf(startedDialogIDFormat, d.ID(), ob.ID(), ob.Serial()) {
			dial = d
		}
		return true

	}
	c.startedDialogs.Range(findDialog)
	if dial != nil {
		return
	}
	var dialogData res.DialogData
	findDialog = func(k, v interface{}) bool {
		d, ok := v.(res.DialogData)
		// TODO: find proper dialog for specified character.
		if ok {
			dialogData = d
		}
		return true
	}
	c.dialogs.Range(findDialog)
	if len(dialogData.ID) < 1 {
		return
	}
	dial = dialog.New(dialogData)
	dial.SetOwner(c)
	dial.SetTarget(ob)
	id := fmt.Sprintf(startedDialogIDFormat, dial.ID(), ob.ID(), ob.Serial())
	c.startedDialogs.Store(id, dial)
	return
}

// StartedDialogs returns all character dialogs that are currently in-progress.
func (c *Character) StartedDialogs() (dialogs []*dialog.Dialog) {
	addDialog := func(k, v interface{}) bool {
		d, ok := v.(*dialog.Dialog)
		if ok {
			dialogs = append(dialogs, d)
		}
		return true
	}
	c.startedDialogs.Range(addDialog)
	return
}

// Dialogs returns all character dialogs.
func (c *Character) Dialogs() (dialogs []res.DialogData) {
	addDialog := func(k, v interface{}) bool {
		d, ok := v.(res.DialogData)
		if ok {
			dialogs = append(dialogs, d)
		}
		return true
	}
	c.dialogs.Range(addDialog)
	return
}

// AddDialog adds specified dialog to character and
// sets character as dialog owner.
func (c *Character) AddDialog(d res.DialogData) {
	c.dialogs.Store(d.ID, d)
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

// ChapterID returns ID of the chapter that this
// character should be in.
func (c *Character) ChapterID() string {
	return c.chapterID
}

// SetChapterID sets specified chapter ID as character
// chapter ID.
func (c *Character) SetChapterID(chapterID string) {
	c.chapterID = chapterID
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

// removeItem removes specific item from usage.
func (c *Character) removeItem(it item.Item) {
	if eqIt, ok := it.(item.Equiper); ok {
		c.Equipment().Unequip(eqIt)
	}
}

// removeFinishedDialog removes specified key-value pair from the started dialogs
// map if it contains finished dialog or dialog without the target.
func (c *Character) removeFinishedDialog(id, value interface{}) bool {
	dialog, ok := value.(*dialog.Dialog)
	if ok && (dialog.Finished() || dialog.Target() == nil) {
		c.startedDialogs.Delete(id)
	}
	return true
}
