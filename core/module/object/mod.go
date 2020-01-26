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

package object

import (
	"fmt"

	"github.com/isangeles/flame/core/data/res/lang"
	"github.com/isangeles/flame/core/module/effect"
)

// TakeModifiers handles all specified modifiers.
func (o *Object) TakeModifiers(source effect.Target, mods ...effect.Modifier) {
	for _, m := range mods {
		o.takeModifier(source, m)
	}
}

// takeModifier handles specified modifier.
func (o *Object) takeModifier(s effect.Target, m effect.Modifier) {
	switch m := m.(type) {
	case *effect.HealthMod:
		val := m.RandomValue()
		o.SetHealth(o.Health() + val)
		cmbMsg := fmt.Sprintf("%s: %s: %d", o.Name(),
			lang.Text("ob_health"), val)
		o.SendCombat(cmbMsg)
	}
}
