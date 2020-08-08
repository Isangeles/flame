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
)

// Data creates data resource struct for character.
func (c *Character) Data() res.CharacterData {
	data := res.CharacterData{
		ID:         c.ID(),
		Serial:     c.Serial(),
		Name:       c.Name(),
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
		CombatLog:  c.CombatLog().Data(),
		PrivateLog: c.PrivateLog().Data(),
		Restore:    true,
	}
	if c.Race() != nil {
		data.Race = c.Race().ID()
	}
	data.PosX, data.PosY = c.Position()
	data.DefX, data.DefY = c.DefaultPosition()
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
