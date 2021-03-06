/*
 * level.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for level requirement.
type Level struct {
	min  int
	meet bool
}

// NewLevel creates new level requirement
// with specified level value.
func NewLevel(data res.LevelReqData) *Level {
	req := new(Level)
	req.min = data.Min
	return req
}

// MinLevel returns minimal required level.
func (lr *Level) MinLevel() int {
	return lr.min
}

// Meet checks whether requirements
// was set as meet.
func (lr *Level) Meet() bool {
	return lr.meet
}

// SetMeet sets requirement as meet or not meet.
func (lr *Level) SetMeet(meet bool) {
	lr.meet = meet
}

// Data returns data resource for requirement.
func (lr *Level) Data() res.LevelReqData {
	data := res.LevelReqData{
		Min: lr.MinLevel(),
	}
	return data
}
