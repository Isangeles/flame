/*
 * effect.go
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

// Struct for effect requirement.
type Effect struct {
	id   string
	meet bool
}

// NewEffect creates new effect requirement.
func NewEffect(data res.EffectReqData) *Effect {
	er := new(Effect)
	er.id = data.ID
	return er
}

// ID returns required effect ID.
func (er *Effect) ID() string {
	return er.id
}

// Meet checks if requirement is set as met.
func (er *Effect) Meet() bool {
	return er.meet
}

// SetMeet sets requirement as meet/not meet.
func (er *Effect) SetMeet(meet bool) {
	er.meet = meet
}

// Data returns data resource for requirement.
func (er *Effect) Data() res.EffectReqData {
	data := res.EffectReqData{er.id}
	return data
}
