/*
 * combat.go
 *
 * Copyright 2019-2022 Dariusz Sikora <ds@isangeles.dev>
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

package object

import (
	"github.com/isangeles/flame/effect"
)

// HitModifiers returns all object hit modifiers.
func (ob *Object) HitModifiers() []effect.Modifier {
	modifiers := make([]effect.Modifier, 0)
	return modifiers
}

// HitEffects returns all object hit effects.
func (ob *Object) HitEffects() []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	return effects
}

// TakeEffect handles effect casted towards object.
func (ob *Object) TakeEffect(e *effect.Effect) {
	ob.AddEffect(e)
	if ob.onEffectTaken != nil {
		ob.onEffectTaken(e)
	}
}
