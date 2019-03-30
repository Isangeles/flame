/*
 * object.go
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

package object

import (
	"math"
)

// Interface for game objects.
type Object interface {
	ID() string
	Serial() string
}

// Interface for useable gmae
// objects.
type UseObject interface {
	ID() string
	Serial() string
	Activate()
}

// Interface for all object with
// position on game world map.
type Positioner interface {
	Position() (x, y float64)
}

// Struct for object effect
// types.
type Element int

const (
	Element_none Element = iota
	Element_fire 
	Element_frost
	Element_nature
)

// Equals checks whether two specified objects
// represents the same game object.
func Equals(ob1, ob2 Object) bool {
	return ob1.ID() + ob1.Serial() == ob2.ID() + ob2.Serial()
}

// Range returns range between two objects.
func Range(ob1, ob2 Positioner) float64 {
	x1, y1 := ob1.Position()
	x2, y2 := ob2.Position()
	return math.Hypot(x1-x2, y1-y2)
}
