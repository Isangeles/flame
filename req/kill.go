/*
 * kill.go
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

// Struct for kill requirement.
type Kill struct {
	id     string
	amount int
	meet   bool
}

// NewKill creates new kill requirement.
func NewKill(data res.KillReqData) *Kill {
	kr := Kill{
		id:     data.ID,
		amount: data.Amount,
	}
	return &kr
}

// ID returns ID of the object to kill.
func (kr *Kill) ID() string {
	return kr.id
}

// Amount returns required amount of objects to kill.
func (kr *Kill) Amount() int {
	return kr.amount
}

// Meet checks if requirement is meet.
func (kr *Kill) Meet() bool {
	return kr.meet
}

// Meet sets requirement as meet/not meet.
func (kr *Kill) SetMeet(meet bool) {
	kr.meet = meet
}

// Data returns data resource for requirement.
func (kr *Kill) Data() res.KillReqData {
	data := res.KillReqData{
		ID:     kr.id,
		Amount: kr.amount,
	}
	return data
}
