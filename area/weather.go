/*
 * weather.go
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

package area

import (
	"time"

	"github.com/isangeles/flame/rng"
)

// Struct for area weather.
type Weather struct {
	conditions Conditions
	lastChange time.Time
}

// Type for area weather conditions.
type Conditions string

const (
	Sunny           Conditions = Conditions("weatherSunny")
	Rain                       = Conditions("weatherRain")
	conditionsTimer            = 60.0 // minutes
)

// Conditions retuns weather conditions.
func (w *Weather) Conditions() Conditions {
	return w.conditions
}

// update updates weather.
func (w *Weather) update(areaTime time.Time) {
	if w.lastChange.IsZero() {
		w.lastChange = areaTime
		return
	}
	if areaTime.Sub(w.lastChange).Minutes() < conditionsTimer {
		return
	}
	roll := rng.RollInt(1, 4)
	switch roll {
	case 1, 2:
		w.conditions = Sunny
	case 3:
		w.conditions = Rain
	}
	w.lastChange = areaTime
}
