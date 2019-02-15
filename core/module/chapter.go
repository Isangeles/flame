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
	"path/filepath"
	
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/object/character"
)

// Chapter struct represents module chapter
type Chapter struct {
	conf        ChapterConf
	mod         *Module
	loadedScens []*scenario.Scenario
	npcs        []*character.Character
} 

// NewChapters creates new instance of module chapter.
func NewChapter(mod *Module, conf ChapterConf) *Chapter {
	c := new(Chapter)
	c.mod = mod
	c.conf = conf
	return c
}

// ID returns chapter ID.
func (c *Chapter) ID() string {
	return c.conf.ID
}

// FullPath returns path to chapter directory.
func (c *Chapter) FullPath() string {
	return filepath.FromSlash(c.conf.Path)
}

// ScenariosPath returns path to chapter
// scenarios directory.
func (c *Chapter) ScenariosPath() string {
	return filepath.FromSlash(c.FullPath() +
		"/area/scenarios")
}

// LangPath returns path to chapter
// lang directory.
func (c *Chapter) LangPath() string {
	return filepath.FromSlash(c.FullPath() + "/lang" +
		"/" + c.mod.LangID())
}

// NPCPath returns path to chapter NPCs
// directory.
func (c *Chapter) NPCPath() string {
	return filepath.FromSlash(c.FullPath() + "/npc")
}

// AreasPath returns path to chapter
// areas directory.
func (c *Chapter) AreasPath() string {
	return filepath.FromSlash(c.FullPath() + "/area")
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

// generateSerials generates unique serial values
// for all chapter objects without serial value.
func (c *Chapter) generateSerials() {
	// Characters.
	for _, char := range c.Characters() {
		if char.Serial() == "" { // assumes assigned serial uniqueness
			c.Module().AssignSerial(char)
		}
		for _, i := range char.Inventory().Items() {
			if i.Serial() == "" {
				c.Module().AssignSerial(i)
			}
		}
		for _, e := range char.Effects() {
			if e.Serial() == "" {
				c.Module().AssignSerial(e)
			}
		}
	}
}
