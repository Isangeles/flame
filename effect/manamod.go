/*
 * manamod.go
 *
 * Copyright 2021 Dariusz Sikora <dev@isangeles.pl>
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

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/rng"
)

// Struct for mana modifier.
type ManaMod struct {
	min       int
	max       int
	lastValue int
}

// NewManaMod creates new mana modifier.
func NewManaMod(data res.ManaModData) *ManaMod {
	mm := new(ManaMod)
	mm.min = data.Min
	mm.max = data.Max
	return mm
}

// Min returns minimal value of mana modifier.
func (mm *ManaMod) Min() int {
	return mm.min
}

// Max returns maximal value of mana modifier.
func (mm *ManaMod) Max() int {
	return mm.max
}

// RandomValue returns random number from
// Min - Max range of modifier.
func (mm *ManaMod) RandomValue() int {
	mm.lastValue = rng.RollInt(mm.Min(), mm.Max())
	return mm.lastValue
}

// LastValue retuns last value generated with
// RandomValue function.
func (mm *ManaMod) LastValue() int {
	return mm.lastValue
}

// Data returns data resource for modifier.
func (mm *ManaMod) Data() res.ManaModData {
	data := res.ManaModData{
		Min: mm.Min(),
		Max: mm.Max(),
	}
	return data
}
