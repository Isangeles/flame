/*
 * gender.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for gender requirement.
type Gender struct {
	gender string
	meet   bool
}

// NewGender creates new gender requirement.
func NewGender(data res.GenderReqData) *Gender {
	gr := new(Gender)
	gr.gender = data.Gender
	return gr
}

// Type returns ID of required gender type.
func (gr *Gender) Gender() string {
	return gr.gender
}

// Meet checks wheter requirement is set as meet.
func (gr *Gender) Meet() bool {
	return gr.meet 
}

// SetMeet sets requirement as meet/not meet.
func (gr *Gender) SetMeet(meet bool) {
	gr.meet = meet
}

// Data returns data resource for requirement.
func (gr *Gender) Data() res.GenderReqData {
	data := res.GenderReqData{
		Gender: gr.Gender(),
	}
	return data
}
