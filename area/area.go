/*
 * area.go
 *
 * Copyright 2018-2023 Dariusz Sikora <ds@isangeles.dev>
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
	"time"

	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/serial"
)

// List with names of the map layers on which the
// area objects can move.
var PassableMapLayers = []string{"ground"}

// Area struct represents game world area.
type Area struct {
	id       string
	Time     time.Time
	weather  *Weather
	areaMap  Map
	respawn  *Respawn
	objects  *sync.Map
	subareas *sync.Map
}

// Interface for area objects.
type Object interface {
	effect.Target
	Update(d int64)
	Live() bool
	Respawn() int64
	AreaID() string
	SetAreaID(s string)
	SightRange() float64
	Moving() bool
	DestPoint() (float64, float64)
	Interrupt()
	BaseMoveCooldown() int64
	MoveCooldown() int64
	SetMoveCooldown(c int64)
}

// New creates new area.
func New(data res.AreaData) *Area {
	a := new(Area)
	a.objects = new(sync.Map)
	a.subareas = new(sync.Map)
	a.weather = newWeather(a)
	a.respawn = newRespawn(a)
	a.Apply(data)
	return a
}

// Update updates area.
func (a *Area) Update(delta int64) {
	a.Time = a.Time.Add(time.Duration(delta) * time.Millisecond)
	a.Weather().update()
	for _, o := range a.Objects() {
		o.Update(delta)
		// Move to dest point.
		if o.Moving() {
			x, y := o.DestPoint()
			a.moveObject(o, x, y)
		}
	}
	for _, sa := range a.Subareas() {
		sa.Update(delta)
	}
	a.respawn.Update()
}

// ID returns area ID.
func (a *Area) ID() string {
	return a.id
}

// Map returns area map.
func (a *Area) Map() Map {
	return a.areaMap
}

// AddObjects adds specified object to area.
func (a *Area) AddObject(o Object) {
	a.objects.Store(o.ID()+o.Serial(), o)
	o.SetAreaID(a.ID())
}

// RemoveObject removes specified object from area.
func (a *Area) RemoveObject(o Object) {
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

// Objects returns list with all objects in
// area(excluding subareas).
func (a *Area) Objects() (objects []Object) {
	addObject := func(k, v interface{}) bool {
		o, ok := v.(Object)
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
func (a *Area) AllObjects() (objects []Object) {
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

// Weather retuns area weather.
func (a *Area) Weather() *Weather {
	return a.weather
}

// NearObjects returns all objects within specified range from specified
// XY position.
func (a *Area) NearObjects(x, y, maxrange float64) (obs []Object) {
	addObject := func(k, v interface{}) bool {
		o, ok := v.(Object)
		posX, posY := o.Position()
		if ok && math.Hypot(posX-x, posY-y) <= maxrange {
			obs = append(obs, o)
		}
		return true
	}
	a.objects.Range(addObject)
	return
}

// SightRangeObjects retuns all objects that have specified XY position
// in their sight range.
func (a *Area) SightRangeObjects(x, y float64) (obs []Object) {
	addObject := func(k, v interface{}) bool {
		ob, ok := v.(Object)
		if !ok {
			return true
		}
		obX, obY := ob.Position()
		if math.Hypot(obX-x, obY-y) <= ob.SightRange() {
			obs = append(obs, ob)
		}
		return true
	}
	a.objects.Range(addObject)
	return
}

// Apply applies specified data on the area.
func (a *Area) Apply(data res.AreaData) {
	a.id = data.ID
	a.Time, _ = time.Parse(time.Kitchen, data.Time)
	a.weather.Conditions = Conditions(data.Weather)
	if data.Map != nil {
		a.areaMap = newMap(data.Map)
	}
	a.respawn.Apply(data.Respawn)
	// Remove objects not present anymore.
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
			a.objects.Delete(key)
		}
		return true
	}
	a.objects.Range(removeChars)
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
			_, inArea := a.objects.Load(areaCharData.ID + areaCharData.Serial)
			if !inArea {
				a.AddObject(char)
			}
		} else {
			// Add new character to area.
			charData.Flags = append(charData.Flags, areaCharData.Flags...)
			char = character.New(*charData)
			a.AddObject(char)
		}
		char.SetRespawn(areaCharData.Respawn)
		// Set position.
		if areaCharData.InitX > 0 && areaCharData.InitY > 0 {
			char.SetPosition(areaCharData.InitX, areaCharData.InitY)
			char.SetDestPoint(areaCharData.InitX, areaCharData.InitY)
			char.SetDefaultPosition(areaCharData.InitX, areaCharData.InitY)
			continue
		}
		char.SetPosition(areaCharData.PosX, areaCharData.PosY)
		char.SetDestPoint(areaCharData.DestX, areaCharData.DestY)
		char.SetDefaultPosition(areaCharData.DefX, areaCharData.DefY)
	}
	// Subareas.
	for _, subareaData := range data.Subareas {
		v, _ := a.subareas.Load(subareaData.ID)
		subarea, _ := v.(*Area)
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
		ID:      a.ID(),
		Time:    a.Time.Format(time.Kitchen),
		Respawn: a.respawn.Data(),
		Map:     a.areaMap.Data(),
	}
	for _, o := range a.Objects() {
		c, ok := o.(*character.Character)
		if !ok {
			continue
		}
		charData := res.AreaCharData{
			ID:     c.ID(),
			Serial: c.Serial(),
		}
		charData.PosX, charData.PosY = c.Position()
		charData.DestX, charData.DestY = c.DestPoint()
		charData.DefX, charData.DefY = c.DefaultPosition()
		charData.Respawn = c.Respawn()
		data.Characters = append(data.Characters, charData)
	}
	for _, sa := range a.Subareas() {
		data.Subareas = append(data.Subareas, sa.Data())
	}
	return data
}

// moveObject moves object towards speicifed
// XY position.
func (a *Area) moveObject(ob Object, x, y float64) {
	ob.Interrupt()
	obX, obY := ob.Position()
	if obX < x && a.passable(obX+1, obY) {
		obX += 1
	}
	if obX > x && a.passable(obX-1, obY) {
		obX -= 1
	}
	if obY < y && a.passable(obX, obY+1) {
		obY += 1
	}
	if obY > y && a.passable(obX, obY-1) {
		obY -= 1
	}
	ob.SetPosition(obX, obY)
	ob.SetMoveCooldown(ob.BaseMoveCooldown())
}

// passable checks if specified XY position is passable
// i.e. is withing on of passable layers of the area map.
func (a *Area) passable(x, y float64) bool {
	layer := a.Map().PositionLayer(x, y)
	if len(layer.Name()) < 1 {
		return true
	}
	for _, l := range PassableMapLayers {
		if l == layer.Name() {
			return true
		}
	}
	return false
}
