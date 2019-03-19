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

	"github.com/isangeles/flame/config"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/object/item"
	"github.com/isangeles/flame/core/module/object/skill"
	"github.com/isangeles/flame/core/rng"
)

// Damage retruns min and max damage value,
// including weapons, active effects, etc.
func (c *Character) Damage() (int, int) {
	min, max := c.Attributes().Damage()
	if it := c.Equipment().HandRight().Item(); it != nil {
		if w, ok := it.(*item.Weapon); ok {
			dmgMin, dmgMax := w.Damage()
			min += dmgMin
			max += dmgMax
		}
	}
	return min, max
}

// DamgaeType returns type of damage caused by
// character.
func (c *Character) DamageType() effect.HitType {
	rightHandItem := c.Equipment().HandRight().Item()
	if rightHandItem != nil {
		if w, ok := rightHandItem.(*item.Weapon); ok {
			w.DamageType()
		}
	}
	return effect.Hit_normal
}

// DamageEffects returns character damage effects.
func (c *Character) DamageEffects() []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	rightHandItem := c.Equipment().HandRight().Item()
	if rightHandItem != nil {
		if w, ok := rightHandItem.(*item.Weapon); ok {
			effects = append(effects, c.buildEffects(w.DamageEffects())...)
		}
	}
	return effects
}

// Hit creates character hit.
func (c *Character) Hit() effect.Hit {
	return effect.Hit{
		Source:  c,
		Type:    c.DamageType(),
		HP:      -(rng.RollInt(c.Damage())),
		Effects: c.DamageEffects(),
	}
}

// UseSkill uses specified skill on current target.
func (c *Character) UseSkill(s *skill.Skill) {
	if c.Casting() || c.Moving() {
		msg := fmt.Sprintf("%s:%s:%s", c.Name(), s.Name(),
			lang.Text("ui", "cant_do_right_now"))
		c.sendCmb(msg)
		return
	}
	charSkill := c.skills[s.ID()+s.Serial()]
	if charSkill == s {
		tar := c.Targets()[0]
		err := s.Cast(c, tar)
		if err != nil {
			// Move to target if is too far.
			if fmt.Sprintf("%v", err) == skill.RANGE_ERR {
				c.SetDestPoint(tar.Position())
			}
			c.combatlog <- fmt.Sprintf("%s:%s:%v", c.Name(), s.Name(), err)
		}
		return
	}
	msg := fmt.Sprintf("%s:%s:%s", c.Name(), s.Name(),
		lang.Text("ui", "skill_not_known"))
	c.sendCmb(msg)
}

// Hit handles specified hit.
func (c *Character) TakeHit(hit effect.Hit) {
	// TODO: handle resists.
	c.SetHealth(c.Health() + hit.HP)
	for _, e := range hit.Effects {
		c.TakeEffect(e)
	}
	msg := fmt.Sprintf("%s:%s:%d", c.Name(), lang.Text("ui", "ob_health"), hit.HP)
	c.sendCmb(msg)
}

// TakeEffects adds specified effects
func (c *Character) TakeEffect(e *effect.Effect) {
	// TODO: handle resists.
	c.AddEffect(e)
	msg := fmt.Sprintf("%s:%s:%s", c.Name(), lang.Text("ui", "ob_effect"), e.Name())
	if config.Debug() {
		msg = fmt.Sprintf("%s(%s_%s)", msg, e.ID(), e.Serial())
	}
	c.sendCmb(msg)
}
