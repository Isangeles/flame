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
)

// Chapter struct represents module chapter
type Chapter struct {
	id, path    string
	scensIds    []string
	startScenId string

	scenario    *scenario.Scenario
	loadedScens []*scenario.Scenario
} 

// NewChapters creates new instance of module chapter.
func NewChapter(id, path string) (*Chapter, error) {
	c := new(Chapter)
	c.id = id
	c.path = path
	err := c.loadConf()
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_config:%v", err)
	}
	startScenarioPath := filepath.FromSlash(c.ScenariosPath() + "/" +
		c.startScenId)
	s, err := data.Scenario(startScenarioPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_start_scenario:%v", err)
	}
	c.scenario = s
	return c, nil
}

// Id returns chapter ID.
func (c *Chapter) Id() string {
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
