/*
 * mod.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/effect"
	"github.com/isangeles/flame/core/module/quest"
	"github.com/isangeles/flame/log"
)

// TakeModifiers handles all specified modifiers.
func (c *Character) TakeModifiers(source effect.Target, mods ...effect.Modifier) {
	for _, m := range mods {
		c.takeModifier(source, m)
	}
}

// takeModifier handles specified modifier.
func (c *Character) takeModifier(s effect.Target, m effect.Modifier) {
	switch m := m.(type) {
	case *effect.AreaMod:
		c.SetAreaID(m.AreaID())
		x, y := m.EnterPosition()
		c.SetPosition(x, y)
	case *effect.FlagMod:
		c.AddFlag(m.Flag())
	case *effect.HealthMod:
		val := m.RandomValue()
		c.SetHealth(c.Health() + val)
		if c.onDamageTaken != nil {
			c.onDamageTaken(val)
		}
		cmbMsg := fmt.Sprintf("%s:%s:%d", c.Name(),
			lang.Text("ui", "ob_health"), val)
		c.SendCombat(cmbMsg)
	case *effect.HitMod:
		for _, e := range s.HitEffects() {
			c.TakeEffect(e)
		}
	case *effect.QuestMod:
		data := res.Quest(m.QuestID())
		if data == nil {
			log.Err.Printf("char: %s#%s: quest mod: data not found:%s", m.QuestID())
			return
		}
		q := quest.New(*data)
		c.Journal().AddQuest(q)
	}
}
