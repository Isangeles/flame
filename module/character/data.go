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
	"github.com/isangeles/flame/module/train"
)

// Data creates data resource struct for character.
func (c *Character) Data() res.CharacterData {
	data := res.CharacterData{
		ID:        c.ID(),
		Serial:    c.Serial(),
		Name:      c.Name(),
		AI:        c.AI(),
		Level:     c.Level(),
		Sex:       string(c.Gender()),
		Attitude:  string(c.Attitude()),
		Alignment: string(c.Alignment()),
		Guild:     c.Guild().ID(),
		HP:        c.Health(),
		Mana:      c.Mana(),
		Exp:       c.Experience(),
		Inventory: c.Inventory().Data(),
		Equipment: c.Equipment().Data(),
		QuestLog:  c.Journal().Data(),
		Crafting:  c.Crafting().Data(),
		Trainings: train.TrainingsData(c.Trainings()...),
	}
	if c.Race() != nil {
		data.Race = c.Race().ID()
	}
	data.PosX, data.PosY = c.Position()
	data.DefX, data.DefY = c.DefaultPosition()
	data.Attributes = res.AttributesData{
		Str: c.Attributes().Str,
		Con: c.Attributes().Con,
		Dex: c.Attributes().Dex,
		Wis: c.Attributes().Wis,
		Int: c.Attributes().Int,
	}
	for _, f := range c.Flags() {
		data.Flags = append(data.Flags, res.FlagData{string(f)})
	}
	for _, e := range c.Effects() {
		effData := res.ObjectEffectData{
			ID:     e.ID(),
			Serial: e.Serial(),
			Time:   e.Time(),
		}
		if e.Source() != nil {
			effData.SourceID = e.Source().ID()
			effData.SourceSerial = e.Source().Serial()
		}
		data.Effects = append(data.Effects, effData)
	}
	for _, s := range c.Skills() {
		skillData := res.ObjectSkillData{
			ID:       s.ID(),
			Serial:   s.Serial(),
			Cooldown: s.Cooldown(),
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
		data.Dialogs = append(data.Dialogs, dialogData)
	}
	return data
}
