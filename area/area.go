/*
 * area.go
 *
 * Copyright 2018-2021 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/object"
	"github.com/isangeles/flame/objects"
	"github.com/isangeles/flame/serial"
)

// Area struct represents game world area.
type Area struct {
	id       string
	chars    *sync.Map
	objects  *sync.Map
	subareas *sync.Map
}

// New creates new area.
func New() *Area {
	a := new(Area)
	a.chars = new(sync.Map)
	a.objects = new(sync.Map)
	a.subareas = new(sync.Map)
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
	a.objects.Store(o.ID()+o.Serial(), o)
}

func (a *Area) RemoveObject(o *object.Object) {
	a.objects.Delete(o.ID() + o.Serial())
}

// AddSubareas adds specified area to subareas.
func (a *Area) AddSubarea(sa *Area) {
	a.subareas.Store(sa.ID(), sa)
}

// RemoveSubareas removes specified subobject.
func (a *Area) RemoveSubarea(sa *Area) {
	a.subareas.Delete(sa.ID())
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
	addObject := func(k, v interface{}) bool {
		o, ok := v.(*object.Object)
		if ok {
			objects = append(objects, o)
		}
		return true
	}
	a.objects.Range(addObject)
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
	addArea := func(k, v interface{}) bool {
		area, ok := v.(*Area)
		if ok {
			areas = append(areas, area)
		}
		return true
	}
	a.subareas.Range(addArea)
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
func (a *Area) NearTargets(pos objects.Positioner, maxrange float64) (targets []effect.Target) {
	addTar := func(k, v interface{}) bool {
		t, ok := v.(effect.Target)
		if ok && objects.Range(t, pos) <= maxrange {
			targets = append(targets, t)
		}
		return true
	}
	a.chars.Range(addTar)
	a.objects.Range(addTar)
	return
}

// NearObjects returns all objects within specified range from specified
// XY position.
func (a *Area) NearObjects(x, y, maxrange float64) (obs []objects.Positioner) {
	addObject := func(k, v interface{}) bool {
		o, ok := v.(objects.Positioner)
		posX, posY := o.Position()
		if ok && math.Hypot(posX-x, posY-y) <= maxrange {
			obs = append(obs, o)
		}
		return true
	}
	a.chars.Range(addObject)
	a.objects.Range(addObject)
	return
}

// Apply applies specified data on the area.
func (a *Area) Apply(data res.AreaData) {
	a.id = data.ID
	// Remove characters not present anymore.
	removeChars := func(key, value interface{}) bool {
		key, _ = key.(string)
		found := false
		for _, cd := range data.Characters {
			if cd.ID+cd.Serial == key {
				found = true
				break
			}
		}
		if !found {
			a.chars.Delete(key)
		}
		return true
	}
	a.chars.Range(removeChars)
	// Characters.
	for _, areaCharData := range data.Characters {
		// Retireve char data.
		charData := res.Character(areaCharData.ID, areaCharData.Serial)
		if charData == nil {
			log.Err.Printf("area: %s: npc data not found: %s",
				a.ID(), areaCharData.ID)
			continue
		}
		ob := serial.Object(areaCharData.ID, areaCharData.Serial)
		char, ok := ob.(*character.Character)
		if ok {
			// Apply data and add to area if not present already.
			char.Apply(*charData)
			_, inArea := a.chars.Load(areaCharData.ID + areaCharData.Serial)
			if !inArea {
				a.AddCharacter(char)
			}
		} else {
			// Add new character to area.
			charData.AI = areaCharData.AI
			charData.Flags = append(charData.Flags, areaCharData.Flags...)
			char = character.New(*charData)
			a.AddCharacter(char)
		}
		// Set position.
		char.SetPosition(areaCharData.PosX, areaCharData.PosY)
		char.SetDestPoint(areaCharData.DestX, areaCharData.DestY)
		char.SetDefaultPosition(areaCharData.DefX, areaCharData.DefY)
	}
	// Objects.
	for _, areaObData := range data.Objects {
		// Retrieve object data.
		obData := res.Object(areaObData.ID, areaObData.Serial)
		if obData == nil {
			log.Err.Printf("area %s: object data not found: %s",
				a.ID(), areaObData.ID)
			continue
		}
		ob := serial.Object(areaObData.ID, areaObData.Serial)
		areaOb, ok := ob.(*object.Object)
		if ok {
			// Apply data and add to area if not present already.
			areaOb.Apply(*obData)
			_, inArea := a.objects.Load(areaObData.ID + areaObData.Serial)
			if !inArea {
				a.AddObject(areaOb)
			}
		} else {
			// Add new object to area.
			areaOb = object.New(*obData)
			a.AddObject(areaOb)
		}
		// Set position.
		areaOb.SetPosition(areaObData.PosX, areaObData.PosY)
	}
	// Subareas.
	for _, subareaData := range data.Subareas {
		v, _ := a.subareas.Load(subareaData.ID)
		subarea, _ := v.(*Area)
		if subarea == nil {
			subarea = New()
			subarea.Apply(subareaData)
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
		charData.DestX, charData.DestY = c.DestPoint()
		charData.DefX, charData.DefY = c.DefaultPosition()
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
