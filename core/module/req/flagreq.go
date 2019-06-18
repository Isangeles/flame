/*
 * flagreq.go
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

// Struct for flag requirement.
type FlagReq struct {
	flagID  string
	flagOff bool
	meet    bool
}

// NewFlagReq creates new flag requirement.
func NewFlagReq(data res.FlagReqData) *FlagReq {
	fr := new(FlagReq)
	fr.flagID = data.ID
	fr.flagOff = data.Off
	return fr
}

// FlagID returns ID of required flag.
func (fr *FlagReq) FlagID() string {
	return fr.flagID
}

// FlagOff checks if flag should be present
// or not.
func (fr *FlagReq) FlagOff() bool {
	return fr.flagOff
}

// Meet checks wheter requirement is set as meet.
func (fr *FlagReq) Meet() bool {
	return fr.meet
}

// SetMeet sets requirement as meet/not meet.
func (fr *FlagReq) SetMeet(meet bool) {
	fr.meet = meet
}
