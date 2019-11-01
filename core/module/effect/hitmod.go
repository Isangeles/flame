/*
 * hitmod.go
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

package effect

// Struct for hit modifier.
type HitMod struct {}

// NewHitMod creates hit modifier.
func NewHitMod() *HitMod {
	hitm := new(HitMod)
	return hitm
}

// Affect takes source damage and attacks specified targets.
func (hitm *HitMod) Affect(source Target, targets ...Target) {
	if source == nil {
		return
	}
	for _, t := range targets {
		for _, e := range source.HitEffects() {
			t.TakeEffect(e)
		}
	}
}

// Undo does nothing.
func (hitm *HitMod) Undo(source Target, targets ...Target) {
}
