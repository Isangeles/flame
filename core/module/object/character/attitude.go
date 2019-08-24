/*
 * attitude.go
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
 
package character

import (
	"github.com/isangeles/flame/core/module/object"
)

// Interface for object with attitude.
type Attituder interface {
	Attitude() Attitude
	AttitudeFor(o object.Object) Attitude
}

// Type for character attitude.
type Attitude int

const (
	Friendly Attitude = iota
	Neutral
	Hostile
)

// ID returns attitude ID.
func (a Attitude) ID() string {
	switch a {
	case Friendly:
		return "att_friendly"
	case Neutral:
		return "att_neutral"
	case Hostile:
		return "att_hostile"
	default:
		return "att_unknown"
	}
}
