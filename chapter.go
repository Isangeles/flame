/*
 * chapter.go
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

package flame

import (
	"fmt"
	"strconv"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/area"
	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/object"
	"github.com/isangeles/flame/objects"
	"github.com/isangeles/flame/log"
)

// Chapter struct represents module chapter.
type Chapter struct {
	res         *res.ResourcesData
	conf        *ChapterConfig
	mod         *Module
	areas       map[string]*area.Area
	onAreaAdded func(s *area.Area)
}

// NewChapter creates new module chapter.
func NewChapter(mod *Module, data res.ChapterData) *Chapter {
	c := new(Chapter)
	c.mod = mod
	c.conf = new(ChapterConfig)
	c.areas = make(map[string]*area.Area)
	c.Apply(data)
	return c
}

// Update updates chapter.
func (c *Chapter) Update(delta int64) {
	for _, a := range c.areas {
		a.Update(delta)
	}
	c.updateObjectsArea()
}

// ID returns chapter ID.
func (c *Chapter) ID() string {
	return c.conf.ID
}

// Module returns chapter module.
func (c *Chapter) Module() *Module {
	return c.mod
}

// Resources chapter resources.
func (c *Chapter) Resources() *res.ResourcesData {
	return c.res
}

// Area returns area with specified ID,
// or nil if area with such ID was not found.
// Loads area if area with specified ID was not
// requested before.
func (c *Chapter) Area(areaID string) *area.Area {
	return c.areas[areaID]
}

// Areas returns all active(loaded) areas.
func (c *Chapter) Areas() (areas []*area.Area) {
	for _, a := range c.areas {
		areas = append(areas, a)
	}
	return
}

// AddAreas adds specified areas to loaded
// areas list.
func (c *Chapter) AddAreas(areas ...*area.Area) {
	for _, a := range areas {
		c.areas[a.ID()] = a
		if c.onAreaAdded != nil {
			c.onAreaAdded(a)
		}
	}
}

// Conf returns chapter configuration.
func (c *Chapter) Conf() *ChapterConfig {
	return c.conf
}

// Characters returns list with all existing(loaded)
// characters in chapter.
func (c *Chapter) Characters() (chars []*character.Character) {
	for _, a := range c.areas {
		for _, c := range a.AllCharacters() {
			chars = append(chars, c)
		}
	}
	return
}

// Objects returns list with all area objects from all
// loaded areas.
func (c *Chapter) AreaObjects() (objects []*object.Object) {
	for _, a := range c.areas {
		for _, o := range a.AllObjects() {
			objects = append(objects, o)
		}
	}
	return
}

// Character returns existing game character with specified
// serial ID or nil if no character with specified ID exists.
func (c *Chapter) Character(id, serial string) *character.Character {
	for _, a := range c.areas {
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
	for _, a := range c.areas {
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
	for _, a := range c.areas {
		for _, c := range a.Characters() {
			if c.ID() == char.ID() && c.Serial() == char.Serial() {
				return a
			}
		}
		for _, a := range a.AllSubareas() {
			for _, c := range a.Characters() {
				if c.ID() == char.ID() && c.Serial() == char.Serial() {
					return a
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

// Apply applies specified data on the chapter.
// Also, adds chapter resources to resources
// base in res package.
func (c *Chapter) Apply(data res.ChapterData) {
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
	if len(data.Config["start-level"]) > 0 {
		c.conf.StartLevel, _ = strconv.Atoi(data.Config["start-level"][0])
	}
	c.conf.StartItems = data.Config["start-items"]
	c.conf.StartSkills = data.Config["start-skills"]
	c.res = &data.Resources
	res.Add(*c.res)
	for _, ad := range data.Resources.Areas {
		a := c.Area(ad.ID)
		if a == nil {
			a = area.New()
			a.Apply(ad)
			c.AddAreas(a)
			continue
		}
		a.Apply(ad)
	}
}

// Data creates data resource for chapter.
func (c *Chapter) Data() res.ChapterData {
	data := res.ChapterData{ID: c.Conf().ID}
	data.Config = make(map[string][]string)
	data.Config["id"] = []string{c.Conf().ID}
	data.Config["path"] = []string{c.Conf().Path}
	data.Config["start-area"] = []string{c.Conf().StartArea}
	data.Config["start-pos"] = []string{fmt.Sprintf("%f", c.Conf().StartPosX),
		fmt.Sprintf("%f", c.Conf().StartPosY)}
	data.Config["start-items"] = c.Conf().StartItems
	data.Config["start-skills"] = c.Conf().StartSkills
	data.Config["start-attrs"] = []string{fmt.Sprintf("%d", c.Conf().StartAttrs)}
	data.Config["start-level"] = []string{fmt.Sprintf("%d", c.Conf().StartLevel)}
	for _, a := range c.Areas() {
		data.Areas = append(data.Areas, a.Data())
	}
	data.Resources = *c.res
	// Remove old characters and objects from resources, besides basic ones.
	data.Resources.Characters = make([]res.CharacterData, 0)
	for _, c := range c.Resources().Characters {
		if len(c.Serial) < 1 {
			data.Resources.Characters = append(data.Resources.Characters, c)
		}
	}
	data.Resources.Objects = make([]res.ObjectData, 0)
	for _, o := range c.Resources().Objects {
		if len(o.Serial) < 1 {
			data.Resources.Objects = append(data.Resources.Objects, o)
		}
	}
	for _, c := range c.Characters() {
		data.Resources.Characters = append(data.Resources.Characters, c.Data())
	}
	for _, o := range c.AreaObjects() {
		data.Resources.Objects = append(data.Resources.Objects, o.Data())
	}
	data.Resources.Areas = make([]res.AreaData, 0)
	for _, a := range c.Areas() {
		data.Resources.Areas = append(data.Resources.Areas, a.Data())
	}
	return data
}

// updateObjectsArea checks and moves game objects to
// proper areas, if needed.
func (c *Chapter) updateObjectsArea() {
	for _, char := range c.Characters() {
		currentArea := c.CharacterArea(char)
		if currentArea != nil && currentArea.ID() == char.AreaID() {
			continue
		}
		var newArea *area.Area
		// Search for area in current chapter.
		for _, a := range c.Areas() {
			if a.ID() == char.AreaID() {
				newArea = a
				break
			}
			for _, sa := range a.AllSubareas() {
				if sa.ID() == char.AreaID() {
					newArea = sa
					break
				}
			}
		}
		if newArea == nil {
			// Search for area data in res package.
			areaData := res.Area(char.AreaID())
			if areaData == nil {
				log.Err.Printf("area update: %s %s: area not found: %s\n",
					char.ID(), char.Serial(), char.AreaID())
				char.SetAreaID(currentArea.ID())
				return
			}
			newArea = area.New()
			newArea.Apply(*areaData)
			c.AddAreas(newArea)
		}
		newArea.AddCharacter(char)
		currentArea.RemoveCharacter(char)
	}
}
