/*
 * weapon.go
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

// Struct for weapon requirement.
type Weapon struct {
	ranged bool
	meet   bool
}

// NewWeapon creates new weapon requirement.
func NewWeapon(data res.WeaponReqData) *Weapon {
	wr := new(Weapon)
	wr.ranged = data.Ranged
	return wr
}

// Ranged checks if ranged weapon is required.
func (wr *Weapon) Ranged() bool {
	return wr.ranged
}

// Meet checks if requirement is set as met.
func (wr *Weapon) Meet() bool {
	return wr.meet
}

// SetMeet sets requirement as meet/not meet.
func (wr *Weapon) SetMeet(meet bool) {
	wr.meet = meet
}

// Data returns data resource for requirement.
func (wr *Weapon) Data() res.WeaponReqData {
	data := res.WeaponReqData{wr.ranged}
	return data
}
