/*
 * flag.go
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

// Struct for flag requirement.
type Flag struct {
	flagID  string
	flagOff bool
	meet    bool
}

// NewFlag creates new flag requirement.
func NewFlag(data res.FlagReqData) *Flag {
	fr := new(Flag)
	fr.flagID = data.ID
	fr.flagOff = data.Off
	return fr
}

// FlagID returns ID of required flag.
func (fr *Flag) FlagID() string {
	return fr.flagID
}

// FlagOff checks if flag should be present
// or not.
func (fr *Flag) FlagOff() bool {
	return fr.flagOff
}

// Meet checks wheter requirement is set as meet.
func (fr *Flag) Meet() bool {
	return fr.meet
}

// SetMeet sets requirement as meet/not meet.
func (fr *Flag) SetMeet(meet bool) {
	fr.meet = meet
}

// Data returns data resource for requirement.
func (cr *Flag) Data() res.FlagReqData {
	data := res.FlagReqData{
		ID: cr.FlagID(),
		Off: cr.FlagOff(),
	}
	return data
}
