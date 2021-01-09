/*
 * data.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/module/dialog"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/flag"
	"github.com/isangeles/flame/module/skill"
	"github.com/isangeles/flame/module/training"
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
	c.SetAI(data.AI)
	c.SetGender(Gender(data.Sex))
	c.SetAttitude(Attitude(data.Attitude))
	c.SetAlignment(Alignment(data.Alignment))
	c.Attributes().Apply(data.Attributes)
	c.Inventory().Apply(data.Inventory)
	c.Equipment().Apply(data.Equipment)
	c.Journal().Apply(data.QuestLog)
	c.Crafting().Apply(data.Crafting)
	c.ChatLog().Apply(data.ChatLog)
	c.PrivateLog().Apply(data.PrivateLog)
	c.targets = data.Targets
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
	c.flags = make(map[string]flag.Flag)
	for _, fd := range data.Flags {
		f := flag.Flag(fd.ID)
		c.flags[f.ID()] = f
	}
	// Add skills.
	for _, charSkillData := range data.Skills {
		s := c.skills[charSkillData.ID]
		if s != nil {
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
	// Add dialogs.
	for _, charDialogData := range data.Dialogs {
		d := c.dialogs[charDialogData.ID]
		if d != nil {
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
		e := c.effects[charEffectData.ID]
		if e == nil {
			effectData := res.Effect(charEffectData.ID)
			if effectData == nil {
				log.Err.Printf("Character: %s: Apply: effect data not found: %s",
					c.ID(), charEffectData.ID)
				continue
			}
			e = effect.New(*effectData)
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
		ID:         c.ID(),
		Serial:     c.Serial(),
		AI:         c.AI(),
		Level:      c.Level(),
		Sex:        string(c.Gender()),
		Attitude:   string(c.Attitude()),
		Alignment:  string(c.Alignment()),
		Guild:      c.Guild().ID(),
		HP:         c.Health(),
		Mana:       c.Mana(),
		Exp:        c.Experience(),
		Attributes: c.Attributes().Data(),
		Inventory:  c.Inventory().Data(),
		Equipment:  c.Equipment().Data(),
		QuestLog:   c.Journal().Data(),
		Crafting:   c.Crafting().Data(),
		ChatLog:    c.ChatLog().Data(),
		PrivateLog: c.PrivateLog().Data(),
		Targets:    c.targets,
		Restore:    true,
	}
	if c.Race() != nil {
		data.Race = c.Race().ID()
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
	for k, e := range c.effects {
		found := false
		for _, ed := range data.Effects {
			found = e.ID() == ed.ID && e.Serial() == ed.Serial
		}
		if !found {
			delete(c.effects, k)
		}
	}
	for k, s := range c.skills {
		found := false
		for _, sd := range data.Skills {
			found = s.ID() == sd.ID
		}
		if !found {
			delete(c.skills, k)
		}
	}
	for k, d := range c.dialogs {
		found := false
		for _, dd := range data.Dialogs {
			found = d.ID() == dd.ID
		}
		if !found {
			delete(c.dialogs, k)
		}
	}
	for k, m := range c.memory {
		found := false
		for _, md := range data.Memory {
			found = m.TargetID == md.ObjectID && m.TargetSerial == md.ObjectSerial
		}
		if !found {
			delete(c.memory, k)
		}
	}
}
