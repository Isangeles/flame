/*
 * chapter.go
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

package module

import (	
	"github.com/isangeles/flame/core/module/object/area"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/serial"
)

// Chapter struct represents module chapter.
type Chapter struct {
	conf        ChapterConf
	mod         *Module
	loadedAreas map[string]*scenario.Area
	onAreaAdded func(s *scenario.Area)
}

// NewChapters creates new instance of module chapter.
func NewChapter(mod *Module, conf ChapterConf) *Chapter {
	c := new(Chapter)
	c.mod = mod
	c.conf = conf
	c.conf.Lang = c.mod.Conf().Lang
	c.loadedAreas = make(map[string]*scenario.Area)
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

// Area returns active(loaded) area with specified ID,
// or nil if area with such ID was not found.
func (c *Chapter) Area(areaID string) *scenario.Area {
	return c.loadedAreas[areaID]
}

// Areas returns all active(loaded) areas.
func (c *Chapter) Areas() (areas []*scenario.Area) {
	for _, a := range c.loadedAreas {
		areas = append(areas, a)
	}
	return
}

// SetAreas sets specified areas as loaded areas.
func (c *Chapter) AddAreas(areas ...*scenario.Area) {
	for _, a := range areas {
		c.loadedAreas[a.ID()] = a
		if c.onAreaAdded != nil {
			c.onAreaAdded(a)
		}
	}
}

// Conf returns chapter configuration.
func (c *Chapter) Conf() ChapterConf {
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
// loaded scenarios.
func (c *Chapter) AreaObjects() (objects []*area.Object) {
	for _, a := range c.loadedAreas {
		for _, o := range a.AllObjects() {
			objects = append(objects, o)
		}
	}
	return
}

// CharactersWithID returns all existing characters with
// specified ID.
func (c *Chapter) CharactersWithID(id string) (chars []*character.Character) {
	for _, a := range c.loadedAreas {
		for _, c := range a.AllCharacters() {
			if c.ID() == id {
				chars = append(chars, c)
			}
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
func (c *Chapter) AreaObject(id, serial string) *area.Object {
	for _, a := range c.loadedAreas {
		for _, o := range a.AllObjects() {
			if o.ID() == id && o.Serial() == serial {
				return o
			}
		}
	}
	return nil
}

// CharacterArea returns area where specified character
// is present, or nil if no such area was found.
func (c *Chapter) CharacterArea(char *character.Character) *scenario.Area {
	for _, a := range c.loadedAreas {
		for _, c := range a.Characters() {
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
func (c *Chapter) SetOnAreaAddedFunc(f func(s *scenario.Area)) {
	c.onAreaAdded = f
}

// generateSerials generates unique serial values
// for all chapter objects without serial value.
// TODO: remove this(?).
func (c *Chapter) generateSerials() {
	// Characters.
	for _, char := range c.Characters() {
		if char.Serial() == "" {
			serial.AssignSerial(char)
		}
		for _, i := range char.Inventory().Items() {
			if i.Serial() != "" {
				continue
			}
			serial.AssignSerial(i)
		}
		for _, e := range char.Effects() {
			if e.Serial() != "" {
				continue
			}
			serial.AssignSerial(e)
		}
		for _, s := range char.Skills() {
			if s.Serial() != "" {
				continue
			}
			serial.AssignSerial(s)
		}
	}
}
