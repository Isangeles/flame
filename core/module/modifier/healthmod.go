/*
 * healthmod.go
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
	"math/rand"
	"time"

	"github.com/isangeles/flame/core/module/object"
)

// Struct for health modifier.
type HealthMod struct {
	min, max int
	rng  *rand.Rand
}

// NewHealthMod creates new health modifier.
func NewHealthMod(min, max int) HealthMod {
	hm := HealthMod{min: min, max: max}
	rngSrc := rand.NewSource(time.Now().UnixNano())
	hm.rng = rand.New(rngSrc)
	return hm
}

// Affect modifies targets health points.
func (hm HealthMod) Affect(source object.Target, targets ...object.Target) {
	for _, t := range targets {
		val := hm.rollValue()
		hit := object.Hit{
			Source: source,
			Type: object.Hit_normal,
			Damage: val,
		}
		t.TakeHit(hit)
	}
}

// Undo undos health modification on specified targets.
func (hm HealthMod) Undo(source object.Target, targets ...object.Target) {
}

// rollValue returns random value
// from min - max range.
func (hm HealthMod) rollValue() int {
	return hm.min + hm.rng.Intn(hm.max - hm.min)
}
