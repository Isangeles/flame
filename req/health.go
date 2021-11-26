/*
 * health.go
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

package req

import (
	"github.com/isangeles/flame/data/res"
)

// Struct for health requirement.
type Health struct {
	full bool
	meet bool
}

// NewHealth creates new health requirement.
func NewHealth(data res.HealthReqData) *Health {
	hr := Health{full: data.Full}
	return &hr
}

// Full checks if maximal amount of health is required.
func (hr *Health) Full() bool {
	return hr.full
}

// Meet checks if requirement is set as met.
func (hr *Health) Meet() bool {
	return hr.meet
}

// SetMeet sets requirements as meet/not meet.
func (hr *Health) SetMeet(meet bool) {
	hr.meet = meet
}

// Data returns data resource for requirement.
func (hr *Health) Data() res.HealthReqData {
	data := res.HealthReqData{Full: hr.full}
	return data
}
