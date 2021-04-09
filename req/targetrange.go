/*
 * req.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for target range requirement.
type TargetRange struct {
	minRange float64
	meet     bool
}

// NewTargetRange creates target range requirement.
func NewTargetRange(data res.TargetRangeReqData) *TargetRange {
	tr := new(TargetRange)
	tr.minRange = data.MinRange
	return tr
}

// MinRange returns value of minimal
// required range to target.
func (tr *TargetRange) MinRange() float64 {
	return tr.minRange
}

// Meet checks if requirement is meet.
func (tr *TargetRange) Meet() bool {
	return tr.meet
}

// SetMeet sets requirement as meet/not meet.
func (tr *TargetRange) SetMeet(meet bool) {
	tr.meet = meet
}

// Data returns data resource for requirement.
func (tr *TargetRange) Data() res.TargetRangeReqData {
	data := res.TargetRangeReqData{
		MinRange: tr.MinRange(),
	}
	return data
}

