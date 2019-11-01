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

package area

import (
	"fmt"

	"github.com/isangeles/flame/config"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/effect"
)

// HitEffects returns all object hit effects.
func (ob *Object) HitEffects() []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	return effects
}

// TakeEffect handles effect casted towards object.
func (ob *Object) TakeEffect(e *effect.Effect) {
	ob.AddEffect(e)
	msg := fmt.Sprintf("%s:%s:%s", ob.Name(), lang.Text("ui", "ob_effect"), e.Name())
	if config.Debug() { // add effect serial ID to combat message
		msg = fmt.Sprintf("%s(%s_%s)", msg, e.ID(), e.Serial())
	}
	ob.SendCombat(msg)
}
