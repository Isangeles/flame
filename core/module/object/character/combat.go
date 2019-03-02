/*
 * combat.go
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

	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/object/skill"
	"github.com/isangeles/flame/core/rng"
)

// Hit creates character hit.
func (c *Character) Hit() effect.Hit {
	return effect.Hit{
		Source: c,
		Type:   effect.Hit_normal,
		HP:     rng.RollInt(c.Damage()),
	}
}

// Hit handles specified hit.
func (c *Character) TakeHit(hit effect.Hit) {
	// TODO: handle resists.
	c.SetHealth(c.Health() + hit.HP)
	msg := fmt.Sprintf("%s:%s:%d", c.Name(), lang.Text("ui", "ob_health"), hit.HP)
	select {
	case c.combatlog <- msg:
	default:
	}
}

// TakeEffects adds specified effects
func (c *Character) TakeEffect(e *effect.Effect) {
	// TODO: handle resists.
	c.AddEffect(e)
	msg := fmt.Sprintf("%s:%s:%s", c.Name(), lang.Text("ui", "ob_effect"), e.Name())
	select {
	case c.combatlog <- msg:
	default:
	}
}

// UseSkill uses specified skill on current target.
func (c *Character) UseSkill(s *skill.Skill) {
	for _, charSkill := range c.Skills() {
		if charSkill == s {
			err := s.Cast(c, c.Targets()[0])
			if err != nil {
				c.combatlog <- fmt.Sprintf("%s:%s:%v", c.Name(), s.Name(), err)
			}
			return
		}
	}
	msg := fmt.Sprintf("%s:%s:skill_not_known", c.Name(), s.Name())
	select {
	case c.combatlog <- msg:
	default:
	}
}
