/*
 * data.go
 *
 * Copyright 2020-2022 Dariusz Sikora <ds@isangeles.dev>
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

package character

import (
	"sync"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/dialog"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/flag"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/skill"
	"github.com/isangeles/flame/training"
	"github.com/isangeles/flame/useaction"
)

// Apply applies specified data on the character.
func (c *Character) Apply(data res.CharacterData) {
	c.id = data.ID
	c.level = data.Level
	c.SetSerial(data.Serial)
	c.SetExperience(data.Exp)
	c.SetPosition(data.PosX, data.PosY)
	c.SetDefaultPosition(data.DefX, data.DefY)
	c.SetDestPoint(data.DestX, data.DestY)
	c.SetGender(Gender(data.Sex))
	c.SetAttitude(Attitude(data.Attitude))
	c.SetAlignment(Alignment(data.Alignment))
	c.Attributes().Apply(data.Attributes)
	c.Inventory().Apply(data.Inventory)
	c.Equipment().Apply(data.Equipment)
	c.Journal().Apply(data.QuestLog)
	c.Crafting().Apply(data.Crafting)
	c.ChatLog().Apply(data.ChatLog)
	c.SetAreaID(data.Area)
	c.SetGuild(NewGuild(data.Guild))
	c.action = useaction.New(data.Action)
	c.useCooldown = data.UseCooldown
	c.moveCooldown = data.MoveCooldown
	c.casted = data.Casted
	c.targets = data.Targets
	c.kills = data.Kills
	if data.Restore {
		c.SetHealth(data.HP)
		c.SetMana(data.Mana)
	} else {
		c.SetHealth(c.MaxHealth())
		c.SetMana(c.MaxMana())
	}
	// Set Race.
	raceData := res.Race(data.Race)
	if raceData != nil && (c.Race() == nil || c.Race().ID() != raceData.ID) {
		c.race = NewRace(*raceData)
	}
	// Clear old data.
	c.clearOldObjects(data)
	// Add flags.
	c.flags = new(sync.Map)
	for _, fd := range data.Flags {
		f := flag.Flag(fd.ID)
		c.flags.Store(f.ID(), f)
	}
	// Add skills.
	for _, charSkillData := range data.Skills {
		ob, _ := c.skills.Load(charSkillData.ID)
		s, ok := ob.(*skill.Skill)
		if ok {
			continue
		}
		skillData := res.Skill(charSkillData.ID)
		if skillData == nil {
			log.Err.Printf("Character: %s: Apply: skill data not found: %v",
				c.ID(), charSkillData.ID)
			continue
		}
		s = skill.New(*skillData)
		s.UseAction().SetCooldown(charSkillData.Cooldown)
		c.AddSkill(s)
	}
	// Add racial skills.
	for _, raceSkillData := range c.Race().Skills() {
		ob, _ := c.skills.Load(raceSkillData.ID)
		s, ok := ob.(*skill.Skill)
		if ok {
			continue
		}
		skillData := res.Skill(raceSkillData.ID)
		if skillData == nil {
			log.Err.Printf("Character: %s: Apply: race skill data not found: %v",
				c.ID(), raceSkillData.ID)
			continue
		}
		s = skill.New(*skillData)
		c.AddSkill(s)
	}
	// Add dialogs.
	for _, charDialogData := range data.Dialogs {
		ob, _ := c.dialogs.Load(charDialogData.ID)
		d, ok := ob.(*dialog.Dialog)
		if ok {
			continue
		}
		dialogData := res.Dialog(charDialogData.ID)
		if dialogData == nil {
			log.Err.Printf("Character: %s: Apply: dialog data not found: %s",
				c.ID(), charDialogData.ID)
			continue
		}
		d = dialog.New(*dialogData)
		for _, s := range d.Stages() {
			if s.ID() == charDialogData.Stage {
				d.SetStage(s)
			}
		}
		c.AddDialog(d)
	}
	// Effects.
	for _, charEffectData := range data.Effects {
		ob, _ := c.effects.Load(charEffectData.ID + charEffectData.Serial)
		e, ok := ob.(*effect.Effect)
		if ok {
			continue
		}
		effectData := res.Effect(charEffectData.ID)
		if effectData == nil {
			log.Err.Printf("Character: %s: Apply: effect data not found: %s",
				c.ID(), charEffectData.ID)
			continue
		}
		e = effect.New(*effectData)
		e.SetSerial(charEffectData.Serial)
		e.SetTime(charEffectData.Time)
		e.SetSource(charEffectData.SourceID, charEffectData.SourceSerial)
		c.AddEffect(e)
	}
	// Trainings.
	for _, charTrainingData := range data.Trainings {
		hasTraining := false
		for _, t := range c.Trainings() {
			if t.ID() == charTrainingData.ID {
				hasTraining = true
				break
			}
		}
		if hasTraining {
			continue
		}
		trainingData := res.Training(charTrainingData.ID)
		if trainingData == nil {
			log.Err.Printf("Character: %s: Apply: training data not found: %s",
				c.ID(), charTrainingData.ID)
			continue
		}
		t := training.New(*trainingData)
		trainerTraining := training.NewTrainerTraining(t, charTrainingData)
		c.trainings = append(c.trainings, trainerTraining)
	}
	// Memory.
	for _, memData := range data.Memory {
		att := Attitude(memData.Attitude)
		mem := TargetMemory{
			TargetID:     memData.ObjectID,
			TargetSerial: memData.ObjectSerial,
			Attitude:     att,
		}
		c.MemorizeTarget(&mem)
	}
}

// Data creates data resource struct for character.
func (c *Character) Data() res.CharacterData {
	data := res.CharacterData{
		ID:           c.ID(),
		Serial:       c.Serial(),
		Level:        c.Level(),
		Sex:          string(c.Gender()),
		Attitude:     string(c.Attitude()),
		Alignment:    string(c.Alignment()),
		Guild:        c.Guild().ID(),
		HP:           c.Health(),
		Mana:         c.Mana(),
		Exp:          c.Experience(),
		Attributes:   c.Attributes().Data(),
		Inventory:    c.Inventory().Data(),
		Equipment:    c.Equipment().Data(),
		QuestLog:     c.Journal().Data(),
		Crafting:     c.Crafting().Data(),
		ChatLog:      c.ChatLog().Data(),
		Casted:       c.casted,
		Targets:      c.targets,
		Kills:        c.kills,
		Restore:      true,
		Area:         c.AreaID(),
		UseCooldown:  c.useCooldown,
		MoveCooldown: c.moveCooldown,
	}
	if c.Race() != nil {
		data.Race = c.Race().ID()
	}
	if c.UseAction() != nil {
		data.Action = c.UseAction().Data()
	}
	data.PosX, data.PosY = c.Position()
	data.DefX, data.DefY = c.DefaultPosition()
	data.DestX, data.DestY = c.DestPoint()
	for _, f := range c.Flags() {
		data.Flags = append(data.Flags, res.FlagData{string(f)})
	}
	for _, e := range c.Effects() {
		effData := res.ObjectEffectData{
			ID:     e.ID(),
			Serial: e.Serial(),
			Time:   e.Time(),
		}
		effData.SourceID, effData.SourceSerial = e.Source()
		data.Effects = append(data.Effects, effData)
	}
	for _, s := range c.Skills() {
		skillData := res.ObjectSkillData{
			ID:       s.ID(),
			Cooldown: s.UseAction().Cooldown(),
		}
		data.Skills = append(data.Skills, skillData)
	}
	for _, m := range c.Memory() {
		memData := res.AttitudeMemoryData{
			ObjectID:     m.TargetID,
			ObjectSerial: m.TargetSerial,
			Attitude:     string(m.Attitude),
		}
		data.Memory = append(data.Memory, memData)
	}
	for _, d := range c.Dialogs() {
		dialogData := res.ObjectDialogData{
			ID: d.ID(),
		}
		if d.Stage() != nil {
			dialogData.Stage = d.Stage().ID()
		}
		data.Dialogs = append(data.Dialogs, dialogData)
	}
	for _, t := range c.Trainings() {
		data.Trainings = append(data.Trainings, t.Data())
	}
	return data
}

// clearOldData clears all effects, skills, dialogs, and memory
// targets not present in specified data.
func (c *Character) clearOldObjects(data res.CharacterData) {
	for _, e := range c.Effects() {
		found := false
		for _, ed := range data.Effects {
			found = e.ID() == ed.ID && e.Serial() == ed.Serial
		}
		if !found {
			c.RemoveEffect(e)
		}
	}
	for _, s := range c.Skills() {
		found := false
		for _, sd := range data.Skills {
			found = s.ID() == sd.ID
		}
		if !found {
			c.RemoveSkill(s)
		}
	}
	for _, d := range c.Dialogs() {
		found := false
		for _, dd := range data.Dialogs {
			found = d.ID() == dd.ID
		}
		if !found {
			c.dialogs.Delete(d.ID())
		}
	}
	for _, m := range c.Memory() {
		found := false
		for _, md := range data.Memory {
			found = m.TargetID == md.ObjectID && m.TargetSerial == md.ObjectSerial
		}
		if !found {
			c.memory.Delete(m.TargetID + m.TargetSerial)
		}
	}
}
