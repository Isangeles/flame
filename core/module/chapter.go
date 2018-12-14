/*
 * chapter.go
 * 
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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
	//"strings"

	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/object/character"
)

// Chapter struct represents module chapter
type Chapter struct {
	id, path    string
	scensIDs    []string
	startScenId string

	mod         *Module
	scenario    *scenario.Scenario
	loadedScens []*scenario.Scenario
	npcs        []*character.Character
} 

// NewChapters creates new instance of module chapter.
func NewChapter(mod *Module, id string) (*Chapter, error) {
	c := new(Chapter)
	c.id = id
	c.path = mod.ChaptersPath()
	c.mod = mod
	err := c.loadConf()
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_config:%v", err)
	}
	startScenarioPath := filepath.FromSlash(c.ScenariosPath() + "/" +
		c.startScenId)
	s, err := data.Scenario(startScenarioPath, c.NPCPath(), c.LangPath())
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_start_scenario:%v", err)
	}
	c.loadedScens = append(c.loadedScens, s)
	c.scenario = s
	c.generateSerials()
	return c, nil
}

// ID returns chapter ID.
func (c *Chapter) ID() string {
	return c.id
}

// FullPath returns path to chapter directory.
func (c *Chapter) FullPath() string {
	return filepath.FromSlash(c.path + "/" + c.id)
}

// ConfPath returns path to chapter configuration file.
func (c *Chapter) ConfPath() string {
	return filepath.FromSlash(c.FullPath() + "/chapter.conf")
}

// ScenariosPath returns path to chapter
// scenarios directory.
func (c *Chapter) ScenariosPath() string {
	return filepath.FromSlash(c.FullPath() + "/area/scenarios")
}

// LangPath returns path to chapter
// lang directory.
func (c *Chapter) LangPath() string {
	return filepath.FromSlash(c.FullPath() + "/lang" +
		c.mod.LangID())
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

// Scneario returns current chapter scenario.
func (c *Chapter) Scenario() *scenario.Scenario {
	return c.scenario
}

// ChangeScenario changes current scenario to scenario
// with specified ID.
func (c *Chapter) ChangeScenario(scenID string) error {
	for _, s := range c.loadedScens {
		if s.ID() == scenID {
			c.scenario = s
			return nil
		}
	}
	for _, sID := range c.scensIDs {
		if sID == scenID {
			scenPath := filepath.FromSlash(c.ScenariosPath() + "/" +
				sID)
			s, err := data.Scenario(scenPath, c.NPCPath(), c.LangPath())
			if err != nil {
				return fmt.Errorf("fail_to_retrieve_scenario:%v", err)
			}
			c.scenario = s
			c.loadedScens = append(c.loadedScens, s)
			c.generateSerials()
			return nil
		}
	}
	return fmt.Errorf("scenario_not_found:%s", scenID)
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

// GenerateSerial sets unique serial value for specified
// object with serial ID.
func (c *Chapter) GenerateCharacterSerial(char *character.Character) {
	chars := c.CharactersWithID(char.ID())
	objects := make([]Serializer, 0)
	for _, c := range chars {
		objects = append(objects, c)
	}
	serial := uniqueSerial(objects)
	// Assing serial value to char.
	char.SetSerial(serial)
}

// generateSerials generates unique serial values
// for all chapter objects without serial value.
func (c *Chapter) generateSerials() {
	// Characters.
	for _, char := range c.Characters() {
		if char.HasSerial() { // assumes assigned serial uniqueness
			continue
		}
		c.GenerateCharacterSerial(char)
	}
}

// loadConf loads configuration file for this chapter,
// returns error if configuration not found or corrupted.
func (c *Chapter) loadConf() error {
	confValues, err := text.ReadConfigValue(c.ConfPath(), "start_scenario")
	if err != nil {
		return fmt.Errorf("fail_to_read_conf_values:%v", err)
	}
	c.startScenId = confValues[0]
	return nil
}
