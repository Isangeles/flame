/*
 * manapercent.go
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

// Struct for mana requirement.
type ManaPercent struct {
	value int
	less  bool
	meet  bool
}

// NewManaPercent creates new mana percent requirement.
func NewManaPercent(data res.ManaPercentReqData) *ManaPercent {
	mpr := ManaPercent{
		value: data.Value,
		less:  data.Less,
	}
	return &mpr
}

// Value retuns required mana percentage value.
func (mpr *ManaPercent) Value() int {
	return mpr.value
}

// Less checks if actual mana percentage should be
// less then requirement mana percent.
func (mpr *ManaPercent) Less() bool {
	return mpr.less
}

// Meet checks if requirement is set as met.
func (mpr *ManaPercent) Meet() bool {
	return mpr.meet
}

// SetMeet sets requirements as meet/not meet.
func (mpr *ManaPercent) SetMeet(meet bool) {
	mpr.meet = meet
}

// Data returns data resource for requirement.
func (mpr *ManaPercent) Data() res.ManaPercentReqData {
	data := res.ManaPercentReqData{
		Value: mpr.value,
		Less:  mpr.less,
	}
	return data
}
