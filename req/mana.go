/*
 * mana.go
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

// Struct for mana requirement.
type Mana struct {
	percent int
	less    bool
	meet    bool
}

// NewMana creates new mana requirement.
func NewMana(data res.ManaReqData) *Mana {
	mr := Mana{
		percent: data.Percent,
		less:    data.Less,
	}
	return &mr
}

// Percent retuns required mana percentage value.
func (mr *Mana) Percent() int {
	return mr.percent
}

// Less checks if actual mana percentage should be
// less then requirement mana value.
func (mr *Mana) Less() bool {
	return mr.less
}

// Meet checks if requirement is set as met.
func (mr *Mana) Meet() bool {
	return mr.meet
}

// SetMeet sets requirements as meet/not meet.
func (mr *Mana) SetMeet(meet bool) {
	mr.meet = meet
}

// Data returns data resource for requirement.
func (mr *Mana) Data() res.ManaReqData {
	data := res.ManaReqData{
		Percent: mr.percent,
		Less:    mr.less,
	}
	return data
}
