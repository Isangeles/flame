/*
 * area.go
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

package scenario

import (
	"math"
	
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/core/module/object/area"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/effect"
)

// Area struct represents game world area.
type Area struct {
	id      string
	chars   []*character.Character
	objects []*area.Object
}

// NewArea returns new instace of area.
func NewArea(id string) (*Area) {
	a := new(Area)
	a.id = id
	return a
}

// ID returns area ID.
func (a *Area) ID() string {
	return a.id
}

// AddCharacter adds specified character to area.
func (a *Area) AddCharacter(c *character.Character) {
	a.chars = append(a.chars, c)
}

// AddObjects adds specified object to area.
func (a *Area) AddObject(o *area.Object) {
	a.objects = append(a.objects, o)
}

// Chracters returns list with characters in area.
func (a *Area) Characters() []*character.Character {
	return a.chars
}

// Objects returns list with all object in area.
func (a *Area) Objects() []*area.Object {
	return a.objects
}

// ContainsCharacter checks whether area
// contains specified character.
func (a *Area) ContainsCharacter(char *character.Character) bool {
	for _, c := range a.chars {
		if char == c {
			return true
		}
	}
	return false
}

// NearTargets returns all targets near specified position.
func (a *Area) NearTargets(pos object.Positioner, maxrange float64) []effect.Target {
	objects := make([]effect.Target, 0)
	// Characters.
	for _, char := range a.chars {
		if object.Range(char, pos) <= maxrange {
			objects = append(objects, char)
		}
	}
	// Objects.
	for _, ob := range a.objects {
		if object.Range(ob, pos) <= maxrange {
			objects = append(objects, ob)
		}
	}
	return objects
}

// NearObjects returns all objects within specified range from specified
// XY position.
func (a *Area) NearObjects(x, y, maxrange float64) []object.Positioner {
	objects := make([]object.Positioner, 0)
	// Characters.
	for _, char := range a.chars {
		charX, charY := char.Position()
		if math.Hypot(charX - x, charY - y) <= maxrange {
			objects = append(objects, char)
		}
	}
	// Objects.
	for _, ob := range a.objects {
		obX, obY := ob.Position()
		if math.Hypot(obX - x, obY - y) <= maxrange {
			objects = append(objects, ob)
		}
	}
	return objects
}
