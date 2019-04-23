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
	"fmt"

	"github.com/isangeles/flame/core/module/object/area"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/serial"
)

// Chapter struct represents module chapter
type Chapter struct {
	conf            ChapterConf
	mod             *Module
	loadedScens     []*scenario.Scenario
	onScenarioAdded func(s *scenario.Scenario)
}

// NewChapters creates new instance of module chapter.
func NewChapter(mod *Module, conf ChapterConf) *Chapter {
	c := new(Chapter)
	c.mod = mod
	c.conf = conf
	c.conf.Lang = c.mod.Conf().Lang
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

// Scenario returns active(loaded) scenario with specified ID,
// or error if no such scenario was found.
func (c *Chapter) Scenario(scenID string) (*scenario.Scenario, error) {
	for _, s := range c.loadedScens {
		return s, nil
	}
	return nil, fmt.Errorf("loaded_scenario_not_found:%s", scenID)
}

// Scenarios returns all active(loaded) scenarios.
func (c *Chapter) Scenarios() []*scenario.Scenario {
	return c.loadedScens
}

// ClearScenarios removes all loaded scenarios
// from chapter.
func (c *Chapter) ClearScenarios() {
	c.loadedScens = make([]*scenario.Scenario, 0)
}

// AddScenario add specified scenario to loaded
// scenarios list.
func (c *Chapter) AddScenario(scen *scenario.Scenario) error {
	for _, s := range c.loadedScens { // check if scenario is already added
		if s.ID() == scen.ID() { // prevent scenrios duplication
			return fmt.Errorf("scenario_already_added:%s", scen.ID())
		}
	}
	c.loadedScens = append(c.loadedScens, scen)
	c.generateSerials() // generate serials for all new objects
	return nil
}

// Conf returns chapter configuration.
func (c *Chapter) Conf() ChapterConf {
	return c.conf
}

// Characters returns list with all existing(loaded)
// characters in chapter.
func (c *Chapter) Characters() (chars []*character.Character) {
	for _, s := range c.loadedScens {
		for _, a := range s.Areas() {
			for _, c := range a.Characters() {
				chars = append(chars, c)
			}
		}
	}
	return
}

// Objects returns list with all area objects from all
// loaded scenarios.
func (c *Chapter) AreaObjects() (objects []*area.Object) {
	for _, s := range c.loadedScens {
		for _, a := range s.Areas() {
			for _, o := range a.Objects() {
				objects = append(objects, o)
			}
		}
	}
	return
}

// CharactersWithID returns all existing characters with
// specified ID.
func (c *Chapter) CharactersWithID(id string) (chars []*character.Character) {
	for _, s := range c.loadedScens {
		for _, a := range s.Areas() {
			for _, c := range a.Characters() {
				if c.ID() == id {
					chars = append(chars, c)
				}
			}
		}
	}
	return
}

// Character returns existing game character with specified
// serial ID or nil if no character with specified ID exists.
func (c *Chapter) Character(serialID string) *character.Character {
	for _, s := range c.loadedScens {
		for _, a := range s.Areas() {
			for _, c := range a.Characters() {
				if c.SerialID() == serialID {
					return c
				}
			}
		}
	}
	return nil
}

// AreaObject retruns area object with specified ID and serial
// or nil if no object was found.
func (c *Chapter) AreaObject(id, serial string) *area.Object {
	for _, s := range c.loadedScens {
		for _, a := range s.Areas() {
			for _, o := range a.Objects() {
				if o.ID() == id && o.Serial() == serial {
					return o
				}
			}
		}
	}
	return nil
}

// CharacterArea returns area where specified character
// is present, or error if there is no such area.
func (c *Chapter) CharacterArea(char *character.Character) (*scenario.Area, error) {
	for _, s := range c.loadedScens {
		for _, a := range s.Areas() {
			for _, c := range a.Characters() {
				if c.SerialID() == char.SerialID() {
					return a, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("character not found in any active scenario")
}

// SetOnScenarioAddedFunc sets function triggered after adding
// new scenario to chapter.
func (c *Chapter) SetOnScenarioAddedFunc(f func(s *scenario.Scenario)) {
	c.onScenarioAdded = f
}

// generateSerials generates unique serial values
// for all chapter objects without serial value.
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
