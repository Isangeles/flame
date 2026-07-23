/*
 * visibility.go
 *
 * Copyright 2026 Dariusz Sikora <ds@isangeles.dev>
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

// Struct for visibility requirement.
type Visibility struct {
	less   bool
	charge bool
	meet   bool
	value  int
}

// NewVisibility creates new visibility requirement.
func NewVisibility(data res.ValueReqData) *Visibility {
	v := Visibility{
		less:   data.Less,
		charge: data.Charge,
		value:  data.Value,
	}
	return &v
}

// Value returns required health value.
func (v *Visibility) Value() int {
	return v.value
}

// Less check if actual health value should be
// lesser then required value.
func (v *Visibility) Less() bool {
	return v.less
}

// Charge checks if required health value should
// be taken from object health pool after
// requirement check.
func (v *Visibility) Charge() bool {
	return v.charge
}

// Meet checks if requirement is set as met.
func (v *Visibility) Meet() bool {
	return v.meet
}

// SetMeet sets requirement as meet/not meet.
func (v *Visibility) SetMeet(meet bool) {
	v.meet = meet
}

// Data returns data resource for health requiremnt.
func (v *Visibility) Data() res.ValueReqData {
	data := res.ValueReqData{
		Less:   v.less,
		Charge: v.charge,
		Value:  v.value,
	}
	return data
}
