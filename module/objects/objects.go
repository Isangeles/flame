/*
 * objects.go
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

// Package with utils for module objects.
package objects

import (
	"math"
)

// Interface for game objects.
type Object interface {
	ID() string
	Serial() string
}

// Interface for all object with
// position on game world map.
type Positioner interface {
	Object
	SetPosition(x, y float64)
	Position() (x, y float64)
}

// Interface for objects with
// health points.
type Killable interface {
	Object
	SetHealth(v int)
	SetMaxHealth(v int)
	Health() int
	MaxHealth() int
	Live() bool
}

// Interfece for objects with
// experience points.
type Experiencer interface {
	Killable
	SetExperience(v int)
	SetMaxExperience(v int)
	Experience() int
	MaxExperience() int
	Level() int
}

// Interface for objects with
// mana points.
type Magician interface {
	Killable
	SetMana(v int)
	SetMaxMana(v int)
	Mana() int
	MaxMana() int
}

// Interface for objects with log
// channels.
type Logger interface {
	Object
	Name() string
	CombatLog() *Log
	ChatLog() *Log
	PrivateLog() *Log
}

// Interface for area objects.
type AreaObject interface {
	Positioner
	AreaID() string
	SetAreaID(s string)
}

// Struct for object effect
// types.
type Element string

const (
	ElementNone Element = Element("elementNone")
	ElementFire = Element("elementFire")
	ElementFrost = Element("elementFrost")
	ElementNature = Element("elementNature")
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
