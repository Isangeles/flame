/*
 * healthpercent.go
 *
 * Copyright 2021-2022 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for health percent requirement.
type HealthPercent struct {
	value int
	less  bool
	meet  bool
}

// NewHealth creates new health requirement.
func NewHealthPercent(data res.HealthPercentReqData) *HealthPercent {
	hpr := HealthPercent{
		value: data.Value,
		less:  data.Less,
	}
	return &hpr
}

// Percent retuns required health percentage value.
func (hpr *HealthPercent) Value() int {
	return hpr.value
}

// Less checks if actual health percentage should be
// less then requirement health value.
func (hpr *HealthPercent) Less() bool {
	return hpr.less
}

// Meet checks if requirement is set as met.
func (hpr *HealthPercent) Meet() bool {
	return hpr.meet
}

// SetMeet sets requirements as meet/not meet.
func (hpr *HealthPercent) SetMeet(meet bool) {
	hpr.meet = meet
}

// Data returns data resource for requirement.
func (hpr *HealthPercent) Data() res.HealthPercentReqData {
	data := res.HealthPercentReqData{
		Value: hpr.value,
		Less:  hpr.less,
	}
	return data
}
