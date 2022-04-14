/*
 * health.go
 *
 * Copyright 2022 Dariusz Sikora <dev@isangeles.pl>
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

package req

import (
	"github.com/isangeles/flame/data/res"
)

// Struct for health requirement.
type Health struct {
	less  bool
	meet  bool
	value int
}

// NewHealth creates new health requirement.
func NewHealth(data res.HealthReqData) *Health {
	h := Health{
		less:  data.Less,
		value: data.Value,
	}
	return &h
}

// Value returns required health value.
func (h *Health) Value() int {
	return h.value
}

// Less check if actual health value should be
// lesser then required value.
func (h *Health) Less() bool {
	return h.less
}

// Meet checks if requirement is set as met.
func (h *Health) Meet() bool {
	return h.meet
}

// SetMeet sets requirement as meet/not meet.
func (h *Health) SetMeet(meet bool) {
	h.meet = meet
}

// Data returns data resource for health requiremnt.
func (h *Health) Data() res.HealthReqData {
	data := res.HealthReqData{
		Less:  h.less,
		Value: h.value,
	}
	return data
}
