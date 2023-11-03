/*
 * mod.go
 *
 * Copyright 2019-2023 Dariusz Sikora <ds@isangeles.dev>
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
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/item"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/objects"
	"github.com/isangeles/flame/quest"
	"github.com/isangeles/flame/serial"
	"github.com/isangeles/flame/skill"
)

// TakeModifiers handles all specified modifiers.
// Source can be nil.
func (c *Character) TakeModifiers(source serial.Serialer, mods ...effect.Modifier) {
	for _, m := range mods {
		c.takeModifier(source, m)
	}
}

// takeModifier handles specified modifier.
// Source can be nil.
func (c *Character) takeModifier(s serial.Serialer, m effect.Modifier) {
	switch m := m.(type) {
	case *effect.AreaMod:
		c.SetAreaID(m.AreaID())
		x, y := m.EnterPosition()
		c.SetPosition(x, y)
	case *effect.FlagMod:
		if m.Off() {
			c.RemoveFlag(m.Flag())
			break
		}
		c.AddFlag(m.Flag())
	case *effect.HealthMod:
		lived := c.Live()
		val := m.RandomValue()
		c.SetHealth(c.Health() + val)
		if s == nil {
			break
		}
		if s, ok := s.(objects.Killer); ok && lived && !c.Live() {
			kill := res.KillData{c.ID(), c.Serial(), 100 * c.Level()}
			s.AddKill(kill)
		}
	case *effect.ManaMod:
		val := m.RandomValue()
		c.SetMana(c.Mana() + val)
	case *effect.QuestMod:
		data := res.Quest(m.QuestID())
		if data == nil {
			log.Err.Printf("char: %s %s: quest mod: data not found: %s", c.ID(),
				c.Serial(), m.QuestID())
			break
		}
		q := quest.New(*data)
		c.Journal().AddQuest(q)
	case *effect.AddItemMod:
		data := res.Item(m.ItemID())
		if data == nil {
			log.Err.Printf("char: %s %s: add item mod: data not found: %s", c.ID(),
				c.Serial(), m.ItemID())
			break
		}
		for i := 0; i < m.Amount(); i++ {
			i := item.New(data)
			c.Inventory().AddItem(i)
		}
	case *effect.RemoveItemMod:
		for _, it := range c.Inventory().Items() {
			c.Inventory().RemoveItem(it)
		}
	case *effect.AddSkillMod:
		data := res.Skill(m.SkillID())
		if data == nil {
			log.Err.Printf("char: %s %s: add skill mod: data not found: %s", c.ID(),
				c.Serial(), m.SkillID())
			break
		}
		s := skill.New(*data)
		c.AddSkill(s)
	case *effect.AttributeMod:
		c.Attributes().Str += m.Strength()
		c.Attributes().Con += m.Constitution()
		c.Attributes().Dex += m.Dexterity()
		c.Attributes().Int += m.Intelligence()
		c.Attributes().Wis += m.Wisdom()
	case *effect.MemoryMod:
		tar := TargetMemory{
			TargetID:     s.ID(),
			TargetSerial: s.Serial(),
			Attitude:     Attitude(m.Attitude()),
		}
		c.MemorizeTarget(&tar)
	}
	for _, event := range c.onModifierEvents {
		event(m)
	}
}
