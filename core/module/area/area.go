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

package area

import (
	"math"

	"github.com/isangeles/flame/core/module/effect"
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/core/module/object/area"
	"github.com/isangeles/flame/core/module/object/character"
)

// Area struct represents game world area.
type Area struct {
	id       string
	chars    map[string]*character.Character
	objects  map[string]*area.Object
	subareas map[string]*Area
}

// NewArea returns new instace of area.
func NewArea(id string) *Area {
	a := new(Area)
	a.id = id
	a.chars = make(map[string]*character.Character)
	a.objects = make(map[string]*area.Object)
	a.subareas = make(map[string]*Area)
	return a
}

// ID returns area ID.
func (a *Area) ID() string {
	return a.id
}

// AddCharacter adds specified character to area.
func (a *Area) AddCharacter(c *character.Character) {
	a.chars[c.ID()+c.Serial()] = c
	c.SetAreaID(a.ID())
}

// RemoveCharacter removes specified character from area.
func (a *Area) RemoveCharacter(c *character.Character) {
	delete(a.chars, c.ID()+c.Serial())
}

// AddObjects adds specified object to area.
func (a *Area) AddObject(o *area.Object) {
	a.objects[o.ID()+o.Serial()] = o
}

func (a *Area) RemoveObject(o *area.Object) {
	delete(a.objects, o.ID()+o.Serial())
}

// AddSubareas adds specified area to subareas.
func (a *Area) AddSubarea(sa *Area) {
	a.subareas[sa.ID()] = sa
}

// RemoveSubareas removes specified subarea.
func (a *Area) RemoveSubarea(sa *Area) {
	delete(a.subareas, sa.ID())
}

// Chracters returns list with characters in
// area(excluding subareas).
func (a *Area) Characters() (chars []*character.Character) {
	for _, c := range a.chars {
		chars = append(chars, c)
	}
	return
}

// AllCharacters returns list with all characters in
// area and subareas.
func (a *Area) AllCharacters() (chars []*character.Character) {
	chars = a.Characters()
	for _, sa := range a.Subareas() {
		chars = append(chars, sa.AllCharacters()...)
	}
	return
}

// Objects returns list with all objects in
// area(excluding subareas).
func (a *Area) Objects() (objects []*area.Object) {
	for _, o := range a.objects {
		objects = append(objects, o)
	}
	return
}

// AllObjects retuns list with all objects in
// area and subareas.
func (a *Area) AllObjects() (objects []*area.Object) {
	objects = a.Objects()
	for _, sa := range a.Subareas() {
		objects = append(objects, sa.AllObjects()...)
	}
	return
}

// Subareas returns all subareas.
func (a *Area) Subareas() (areas []*Area) {
	for _, sa := range a.subareas {
		areas = append(areas, sa)
	}
	return
}

// AllSubareas returns all subareas, including
// subareas of subareas
func (a *Area) AllSubareas() (subareas []*Area) {
	subareas = a.Subareas()
	for _, sa := range a.Subareas() {
		subareas = append(subareas, sa.AllSubareas()...)
	}
	return
}

// ContainsCharacter checks whether area
// contains specified character.
func (a *Area) ContainsCharacter(char *character.Character) bool {
	return a.chars[char.ID()+char.Serial()] != nil
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
		if math.Hypot(charX-x, charY-y) <= maxrange {
			objects = append(objects, char)
		}
	}
	// Objects.
	for _, ob := range a.objects {
		obX, obY := ob.Position()
		if math.Hypot(obX-x, obY-y) <= maxrange {
			objects = append(objects, ob)
		}
	}
	return objects
}
