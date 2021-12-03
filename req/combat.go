/*
 * combat.go
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

// Struct for combat requirement.
type Combat struct {
	combat bool
	meet   bool
}

// NewCombat creates new combat requirement.
func NewCombat(data res.CombatReqData) *Combat {
	cr := new(Combat)
	cr.combat = data.Combat
	return cr
}

// Combat checks if combat is required.
func (cr *Combat) Combat() bool {
	return cr.combat
}

// Meet checks if requirement is set as met.
func (cr *Combat) Meet() bool {
	return cr.meet
}

// SetMeet sets requirement as meet/not meet.
func (cr *Combat) SetMeet(meet bool) {
	cr.meet = meet
}

// Data returns data resource for requirement.
func (cr *Combat) Data() res.CombatReqData {
	data := res.CombatReqData{cr.combat}
	return data
}
