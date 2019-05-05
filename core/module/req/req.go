/*
 * req.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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

// Package for requirements(e.g. weapon
// equip requirements).
package req

import (
	"github.com/isangeles/flame/core/data/res"
)

// Interface for requirements.
type Requirement interface {
	Meet() bool
	SetMeet(meet bool)
}

// NewRequirements creates new requirements from
// specified data.
func NewRequirements(data ...res.ReqData) []Requirement {
	reqs := make([]Requirement, 0)
	for _, d := range data {
		switch d := d.(type) {
		case res.LevelReqData:
			lreq := NewLevelReq(d)
			reqs = append(reqs, lreq)
		case res.GenderReqData:
			
			greq := NewGenderReq(d)
			reqs = append(reqs, greq)
		}
	}
	return reqs
}
