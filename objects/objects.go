/*
 * objects.go
 *
 * Copyright 2019-2022 Dariusz Sikora <ds@isangeles.dev>
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

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/serial"
)

// Interface for all object with
// position on game world map.
type Positioner interface {
	serial.Serialer
	SetPosition(x, y float64)
	Position() (x, y float64)
}

// Interface for objects with
// health points.
type Killable interface {
	serial.Serialer
	SetHealth(v int)
	Health() int
	MaxHealth() int
	Live() bool
}

// Interface for objects with kill records.
type Killer interface {
	serial.Serialer
	AddKill(k res.KillData)
	Kills() []res.KillData
}

// Interfece for objects with
// experience points.
type Experiencer interface {
	Killable
	SetExperience(v int)
	Experience() int
	MaxExperience() int
	Level() int
}

// Interface for objects with
// mana points.
type Magician interface {
	Killable
	SetMana(v int)
	Mana() int
	MaxMana() int
}

// Interface for objects with log
// channels.
type Logger interface {
	serial.Serialer
	ChatLog() *Log
}

// Interface for area objects.
type AreaObject interface {
	Positioner
	AreaID() string
	SetAreaID(s string)
	SightRange() float64
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
func Equals(ob1, ob2 serial.Serialer) bool {
	return ob1.ID() + ob1.Serial() == ob2.ID() + ob2.Serial()
}

// Range returns range between two objects.
func Range(ob1, ob2 Positioner) float64 {
	x1, y1 := ob1.Position()
	x2, y2 := ob2.Position()
	return math.Hypot(x1-x2, y1-y2)
}
