/*
 * mana.go
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

// Struct for mana requirement.
type Mana struct {
	less  bool
	meet  bool
	value int
}

// NewMana creates new mana requirement.
func NewMana(data res.ManaReqData) *Mana {
	m := Mana{
		less:  data.Less,
		value: data.Value,
	}
	return &m
}

// Value returns required mana value.
func (m *Mana) Value() int {
	return m.value
}

// Less check if actual mana value should be
// lesser then required value.
func (m *Mana) Less() bool {
	return m.less
}

// Meet checks if requirement is set as met.
func (m *Mana) Meet() bool {
	return m.meet
}

// SetMeet sets requirement as meet/not meet.
func (m *Mana) SetMeet(meet bool) {
	m.meet = meet
}

// Data returns data resource for mana requiremnt.
func (m *Mana) Data() res.ManaReqData {
	data := res.ManaReqData{
		Less:  m.less,
		Value: m.value,
	}
	return data
}
