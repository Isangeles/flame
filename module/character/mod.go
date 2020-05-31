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
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/quest"
	"github.com/isangeles/flame/log"
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
		c.SendCombat(cmbMsg)
		if s == nil {
			break
		}
		if s, ok := s.(objects.Experiencer); ok && lived && !c.Live() {
			exp := 100 * c.Level()
			s.SetExperience(s.Experience() + exp)
		}
	case *effect.QuestMod:
		data, ok := res.Quests[m.QuestID()]
		if !ok {
			log.Err.Printf("char: %s#%s: quest mod: data not found:%s", m.QuestID())
			break
		}
		q := quest.New(data)
		c.Journal().AddQuest(q)
	}
}
