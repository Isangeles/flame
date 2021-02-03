/*
 * mod.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/objects"
)

// TakeModifiers handles all specified modifiers.
func (o *Object) TakeModifiers(source objects.Object, mods ...effect.Modifier) {
	for _, m := range mods {
		o.takeModifier(source, m)
	}
}

// takeModifier handles specified modifier.
func (o *Object) takeModifier(s objects.Object, m effect.Modifier) {
	switch m := m.(type) {
	case *effect.HealthMod:
		val := m.RandomValue()
		o.SetHealth(o.Health() + val)
	}
	if o.onModifierTaken != nil {
		o.onModifierTaken(m)
	}
}
