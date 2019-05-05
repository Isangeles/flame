/*
 * genderreq.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/core/data/res"
)

// Struct for gender requirement.
type GenderReq struct {
	genType int
	meet    bool
}

// NewGenderReq creates new gender requirement.
func NewGenderReq(data res.GenderReqData) *GenderReq {
	gr := new(GenderReq)
	gr.genType = data.Type
	return gr
}

// Type returns ID of required gender type.
func (gr *GenderReq) Type() int {
	return gr.genType
}

// Meet checks wheter requirement is set as meet.
func (gr *GenderReq) Meet() bool {
	return gr.meet 
}

// SetMeet sets requirement as meet/not meet.
func (gr *GenderReq) SetMeet(meet bool) {
	gr.meet = meet
}