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
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/rng"
	"github.com/isangeles/flame/log"
)

// Hit creates character hit.
func (c *Character) Hit() object.Hit {
	return object.Hit{
		Source: c,
		Type:   object.Hit_normal,
		HP:     rng.RollInt(c.Damage()),
	}
}

// Hit handles specified hit.
func (c *Character) TakeHit(hit object.Hit) {
	// TODO: handle resists.
	c.SetHealth(c.Health() + hit.HP)
	log.Cmb.Printf("%s:%s:%d", c.Name(), lang.Text("ui", "ob_health"), hit.HP)
}

// TakeEffects adds specified effects
func (c *Character) TakeEffect(e *effect.Effect) {
	// TODO: handle resists.
	c.AddEffect(e)
	log.Cmb.Printf("%s:%s:%s", c.Name(), lang.Text("ui", "ob_effect"), e.Name())
}
