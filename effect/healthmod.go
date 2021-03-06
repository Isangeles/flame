/*
 * healthmod.go
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

package effect

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/rng"
)

// Struct for health modifier.
type HealthMod struct {
	min       int
	max       int
	lastValue int
}

// NewHealthMod creates new health modifier.
func NewHealthMod(data res.HealthModData) *HealthMod {
	hm := new(HealthMod)
	hm.min = data.Min
	hm.max = data.Max
	return hm
}

// Min returns minimal value of health modifier.
func (hm *HealthMod) Min() int {
	return hm.min
}

// Max returns maximal vallue of healtg modifier.
func (hm *HealthMod) Max() int {
	return hm.max
}

// RandomValue returns random number from
// Min - Max range of modifier.
func (hm *HealthMod) RandomValue() int {
	hm.lastValue = rng.RollInt(hm.Min(), hm.Max())
	return hm.lastValue
}

// LastValue returns last value generated with
// RandomValue function.
func (hm *HealthMod) LastValue() int {
	return hm.lastValue
}

// Data creates data resource for modifier.
func (hm *HealthMod) Data() res.HealthModData {
	data := res.HealthModData{
		Min: hm.Min(),
		Max: hm.Max(),
	}
	return data
}
