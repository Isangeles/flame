/*
 * effectmod.go
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

package modifier

import (
	"github.com/isangeles/flame/core/module/object"
)

// Struct for effect modifier.
type EffectMod struct {
	effectID     string
	activeSerial string
}

// NewEffectMod creates new effect modifier.
func NewEffectMod(effectID string) EffectMod {
	em := EffectMod{effectID: effectID}
	return em
}

// Affect puts modifier effect on specified targets.
func (em EffectMod) Affect(source object.Object, targets ...object.Object) {
}

// Undo removes modifier effect from specified targets.
func (em EffectMod) Undo(source object.Object, targets ...object.Object) {
}
