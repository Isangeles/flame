/*
 * area.go
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

// Package with game area struct.
package area

import (
	"math"
	"sync"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/module/character"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/object"
	"github.com/isangeles/flame/module/objects"
)

// Area struct represents game world area.
type Area struct {
	id       string
	chars    *sync.Map
	objects  map[string]*object.Object
	subareas map[string]*Area
}

// New creates new area.
func New(data res.AreaData) *Area {
	a := new(Area)
	a.chars = new(sync.Map)
	a.objects = make(map[string]*object.Object)
	a.subareas = make(map[string]*Area)
	a.Apply(data)
	return a
}

// Update updates area.
func (a *Area) Update(delta int64) {
	for _, c := range a.Characters() {
		c.Update(delta)
	}
	for _, o := range a.Objects() {
		o.Update(delta)
	}
	for _, sa := range a.Subareas() {
		sa.Update(delta)
	}
}

// ID returns area ID.
func (a *Area) ID() string {
	return a.id
}

// AddCharacter adds specified character to object.
func (a *Area) AddCharacter(c *character.Character) {
	a.chars.Store(c.ID()+c.Serial(), c)
	c.SetAreaID(a.ID())
}

// RemoveCharacter removes specified character from object.
func (a *Area) RemoveCharacter(c *character.Character) {
	a.chars.Delete(c.ID() + c.Serial())
}

// AddObjects adds specified object to object.
func (a *Area) AddObject(o *object.Object) {
	a.objects[o.ID()+o.Serial()] = o
}

func (a *Area) RemoveObject(o *object.Object) {
	delete(a.objects, o.ID()+o.Serial())
}

// AddSubareas adds specified area to subareas.
func (a *Area) AddSubarea(sa *Area) {
	a.subareas[sa.ID()] = sa
}

// RemoveSubareas removes specified subobject.
func (a *Area) RemoveSubarea(sa *Area) {
	delete(a.subareas, sa.ID())
}

// Chracters returns list with characters in
// area(excluding subareas).
func (a *Area) Characters() (chars []*character.Character) {
	addChar := func(k, v interface{}) bool {
		c, ok := v.(*character.Character)
		if ok {
			chars = append(chars, c)
		}
		return true
	}
	a.chars.Range(addChar)
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
func (a *Area) Objects() (objects []*object.Object) {
	for _, o := range a.objects {
		objects = append(objects, o)
	}
	return
}

// AllObjects retuns list with all objects in
// area and subareas.
func (a *Area) AllObjects() (objects []*object.Object) {
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

// NearTargets returns all targets near specified position.
func (a *Area) NearTargets(pos objects.Positioner, maxrange float64) []effect.Target {
	targets := make([]effect.Target, 0)
	// Characters.
	addChar := func(k, v interface{}) bool {
		t, ok := v.(*character.Character)
		if ok && objects.Range(t, pos) <= maxrange {
			targets = append(targets, t)
		}
		return true
	}
	a.chars.Range(addChar)
	// Objects.
	for _, ob := range a.objects {
		if objects.Range(ob, pos) <= maxrange {
			targets = append(targets, ob)
		}
	}
	return targets
}

// NearObjects returns all objects within specified range from specified
// XY position.
func (a *Area) NearObjects(x, y, maxrange float64) []objects.Positioner {
	objects := make([]objects.Positioner, 0)
	// Characters.
	addChar := func(k, v interface{}) bool {
		c, ok := v.(*character.Character)
		charX, charY := c.Position()
		if ok && math.Hypot(charX-x, charY-y) <= maxrange {
			objects = append(objects, c)
		}
		return true
	}
	a.chars.Range(addChar)
	// Objects.
	for _, ob := range a.objects {
		obX, obY := ob.Position()
		if math.Hypot(obX-x, obY-y) <= maxrange {
			objects = append(objects, ob)
		}
	}
	return objects
}

// Apply applies specified data on the area.
func (a *Area) Apply(data res.AreaData) {
	a.id = data.ID
	// Characters.
	for _, areaCharData := range data.Characters {
		ob, _ := a.chars.Load(areaCharData.ID+areaCharData.Serial)
		char, _ := ob.(*character.Character)
		if char == nil {
			// Retireve char data.
			charData := res.Character(areaCharData.ID, areaCharData.Serial)
			if charData == nil {
				log.Err.Printf("area: %s: npc data not found: %s",
					a.ID(), areaCharData.ID)
				continue
			}
			charData.AI = areaCharData.AI
			char = character.New(*charData)
			// Add to area.
			a.AddCharacter(char)
		}
		// Set position.
		char.SetPosition(areaCharData.PosX, areaCharData.PosY)
		char.SetDefaultPosition(areaCharData.PosX, areaCharData.PosY)
	}
	// Objects.
	for _, areaObData := range data.Objects {
		ob := a.objects[areaObData.ID+areaObData.Serial]
		if ob == nil {
			// Retrieve object data.
			obData := res.Object(areaObData.ID, areaObData.Serial)
			if obData == nil {
				log.Err.Printf("area %s: object data not found: %s",
					a.ID(), areaObData.ID)
				continue
			}
			ob = object.New(*obData)
			// Add to area.
			a.AddObject(ob)
		}
		// Set position.
		ob.SetPosition(areaObData.PosX, areaObData.PosY)
	}
	// Subareas.
	for _, subareaData := range data.Subareas {
		subarea := a.subareas[subareaData.ID]
		if subarea == nil {
			subarea = New(subareaData)
			a.AddSubarea(subarea)
		}
		subarea.Apply(subareaData)
	}
}

// Data returns area data resource.
func (a *Area) Data() res.AreaData {
	data := res.AreaData{
		ID: a.ID(),
	}
	for _, c := range a.Characters() {
		charData := res.AreaCharData{
			ID:     c.ID(),
			Serial: c.Serial(),
			AI:     c.AI(),
		}
		charData.PosX, charData.PosY = c.Position()
		data.Characters = append(data.Characters, charData)
	}
	for _, o := range a.Objects() {
		obData := res.AreaObjectData{
			ID:     o.ID(),
			Serial: o.Serial(),
		}
		obData.PosX, obData.PosY = o.Position()
		data.Objects = append(data.Objects, obData)
	}
	for _, sa := range a.Subareas() {
		data.Subareas = append(data.Subareas, sa.Data())
	}
	return data
}
