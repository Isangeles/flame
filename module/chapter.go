/*
 * chapter.go
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

package module

import (
	"strconv"
	
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/area"
	"github.com/isangeles/flame/module/character"
	"github.com/isangeles/flame/module/object"
	"github.com/isangeles/flame/module/objects"
)

// Chapter struct represents module chapter.
type Chapter struct {
	Res         res.ResourcesData
	conf        ChapterConfig
	mod         *Module
	loadedAreas map[string]*area.Area
	onAreaAdded func(s *area.Area)
}

// NewChapter creates new instance of module chapter
// from specified data, adds chapter resources to
// resources base in res package.
func NewChapter(mod *Module, data res.ChapterData) *Chapter {
	c := new(Chapter)
	c.mod = mod
	c.conf = ChapterConfig{}
	if len(data.Config["id"]) > 0 {
		c.conf.ID = data.Config["id"][0]
	}
	if len(data.Config["path"]) > 0 {
		c.conf.Path = data.Config["path"][0]
	}
	if len(data.Config["start-area"]) > 0 {
		c.conf.StartArea = data.Config["start-area"][0]
	}
	if len(data.Config["start-pos"]) > 1 {
		c.conf.StartPosX, _ = strconv.ParseFloat(data.Config["start-pos"][0], 64)
		c.conf.StartPosY, _ = strconv.ParseFloat(data.Config["start-pos"][1], 64)
	}
	if len(data.Config["start-attrs"]) > 0 {
		c.conf.StartAttrs, _ = strconv.Atoi(data.Config["start-attrs"][0])
	}
	c.conf.StartItems = data.Config["start-items"]
	c.conf.StartSkills = data.Config["start-skills"]
	c.Res = data.Resources
	res.Add(c.Res)
	c.loadedAreas = make(map[string]*area.Area)
	return c
}

// ID returns chapter ID.
func (c *Chapter) ID() string {
	return c.conf.ID
}

// Module returns chapter module.
func (c *Chapter) Module() *Module {
	return c.mod
}

// Area returns area with specified ID,
// or nil if area with such ID was not found.
// Loads area if area with specified ID was not
// requested before.
func (c *Chapter) Area(areaID string) *area.Area {
	a := c.loadedAreas[areaID]
	if a != nil {
		return a
	}
	areaData, ok := res.Areas[areaID]
	if !ok {
		return nil
	}
	a = area.New(areaData)
	c.AddAreas(a)
	return a
}

// Areas returns all active(loaded) areas.
func (c *Chapter) Areas() (areas []*area.Area) {
	for _, a := range c.loadedAreas {
		areas = append(areas, a)
	}
	return
}

// AddAreas adds specified areas to loaded
// areas list.
func (c *Chapter) AddAreas(areas ...*area.Area) {
	for _, a := range areas {
		c.loadedAreas[a.ID()] = a
		if c.onAreaAdded != nil {
			c.onAreaAdded(a)
		}
	}
}

// Conf returns chapter configuration.
func (c *Chapter) Conf() ChapterConfig {
	return c.conf
}

// Characters returns list with all existing(loaded)
// characters in chapter.
func (c *Chapter) Characters() (chars []*character.Character) {
	for _, a := range c.loadedAreas {
		for _, c := range a.AllCharacters() {
			chars = append(chars, c)
		}
	}
	return
}

// Objects returns list with all area objects from all
// loaded areas.
func (c *Chapter) AreaObjects() (objects []*object.Object) {
	for _, a := range c.loadedAreas {
		for _, o := range a.AllObjects() {
			objects = append(objects, o)
		}
	}
	return
}

// Character returns existing game character with specified
// serial ID or nil if no character with specified ID exists.
func (c *Chapter) Character(id, serial string) *character.Character {
	for _, a := range c.loadedAreas {
		for _, c := range a.AllCharacters() {
			if c.ID() == id && c.Serial() == serial {
				return c
			}
		}
	}
	return nil
}

// AreaObject retruns area object with specified ID and serial
// or nil if no object was found.
func (c *Chapter) AreaObject(id, serial string) *object.Object {
	for _, a := range c.loadedAreas {
		for _, o := range a.AllObjects() {
			if o.ID() == id && o.Serial() == serial {
				return o
			}
		}
	}
	return nil
}

// Object returns game object with specified ID and serial
// or nil if no such object was found.
func (c *Chapter) Object(id, serial string) objects.Object {
	char := c.Character(id, serial)
	if char != nil {
		return char
	}
	ob := c.AreaObject(id, serial)
	if ob != nil {
		return ob
	}
	return nil
}

// CharacterArea returns area where specified character
// is present, or nil if no such area was found.
func (c *Chapter) CharacterArea(char *character.Character) *area.Area {
	for _, a := range c.loadedAreas {
		for _, c := range a.AllCharacters() {
			if c.SerialID() == char.SerialID() {
				return a
			}
			for _, a := range a.AllSubareas() {
				for _, c := range a.Characters() {
					if c.SerialID() == char.SerialID() {
						return a
					}
				}
			}
		}

	}
	return nil
}

// SetOnAreaAddedFunc sets function triggered after adding
// new area to chapter.
func (c *Chapter) SetOnAreaAddedFunc(f func(s *area.Area)) {
	c.onAreaAdded = f
}

// Data creates data resource for chapter.
func (c *Chapter) Data() res.ChapterData {
	data := res.ChapterData{ID: c.Conf().ID}
	data.Config = make(map[string][]string)
	data.Config["id"] = []string{c.Conf().ID}
	data.Config["path"] = []string{c.Conf().Path}
	for _, a := range c.Areas() {
		data.Areas = append(data.Areas, a.Data())
	}
	return data
}
