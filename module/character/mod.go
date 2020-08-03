/*
 * mod.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"fmt"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/item"
	"github.com/isangeles/flame/module/skill"
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/quest"
)

// TakeModifiers handles all specified modifiers.
// Source can be nil.
func (c *Character) TakeModifiers(source objects.Object, mods ...effect.Modifier) {
	for _, m := range mods {
		c.takeModifier(source, m)
	}
}

// takeModifier handles specified modifier.
// Source can be nil.
func (c *Character) takeModifier(s objects.Object, m effect.Modifier) {
	switch m := m.(type) {
	case *effect.AreaMod:
		c.SetAreaID(m.AreaID())
		x, y := m.EnterPosition()
		c.SetPosition(x, y)
	case *effect.FlagMod:
		c.AddFlag(m.Flag())
	case *effect.HealthMod:
		lived := c.Live()
		val := m.RandomValue()
		c.SetHealth(c.Health() + val)
		if c.onHealthMod != nil {
			c.onHealthMod(val)
		}
		cmbMsg := fmt.Sprintf("%s: %s: %d", c.Name(),
			lang.Text("ob_health"), val)
		c.CombatLog().Add(cmbMsg)
		if s == nil {
			break
		}
		if s, ok := s.(objects.Experiencer); ok && lived && !c.Live() {
			exp := 100 * c.Level()
			s.SetExperience(s.Experience() + exp)
		}
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
	}
}
