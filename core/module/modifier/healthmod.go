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

// Struct for health modifier.
type HealthMod struct {
	min, max int
}

// NewHealthMod creates new health modifier.
func NewHealthMod(min, max int) HealthMod {
	return HealthMod{min, max}
}

// Affect affects on specified targetable objects.
func (hm HealthMod) Affect(source Target, targets ...Target) {
	for _, t := range targets {
		val := hm.min
		t.SetHealth(val)
	}
}
