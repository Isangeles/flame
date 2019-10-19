/*
 * mod.go
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

package data

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/serial"
	"github.com/isangeles/flame/log"
)

const (
	ScenarioFileExt = ".scenario"
)

// Module creates new module from specified path.
func Module(path, langID string) (*module.Module, error) {
	// Load module config file.
	mc, err := modConf(path, langID)
	if err != nil {
		return nil, fmt.Errorf("fail to load module config: %v",
			err)
	}
	// Create module.
	m := module.NewModule(mc)
	return m, nil
}

// LoadChapter loads module chapter with
// specified ID.
func LoadChapter(mod *module.Module, id string) error {
	// Load chapter config file.
	chapPath := filepath.FromSlash(mod.Conf().ChaptersPath() +
		"/" + mod.Conf().StartChapter)
	chapConf, err := chapterConf(chapPath)
	if err != nil {
		return fmt.Errorf("fail to read chapter conf: %s: %v",
			chapPath, err)
	}
	chapConf.ID = id
	chapConf.ModulePath = mod.Conf().Path
	// Create chapter & set as current module chapter.
	startChap := module.NewChapter(mod, chapConf)
	err = mod.SetChapter(startChap) // move to start chapter
	if err != nil {
		return fmt.Errorf("fail to set mod chapter: %v", err)
	}
	return nil
}

// LoadScenario loads scenario with specified
// ID for current module chapter.
func LoadScenario(mod *module.Module, id string) error {
	// Check whether mod has active chapter.
	chap := mod.Chapter()
	if chap == nil {
		return fmt.Errorf("no module chapter set")
	}
	// Load files.
	scenPath := filepath.FromSlash(chap.Conf().ScenariosPath() +
		"/" + id)
	docScen, err := os.Open(scenPath)
	if err != nil {
		return fmt.Errorf("fail to open scenario file: %v", err)
	}
	defer docScen.Close()
	// Unmarshal scenario file.
	scenData, err := parsexml.UnmarshalScenario(docScen)
	if err != nil {
		return fmt.Errorf("fail to parse scenario file: %v", err)
	}
	// Build mainarea.
	mainarea := buildArea(mod, scenData.Area)
	scen := scenario.NewScenario(scenData.ID, mainarea)
	// Add scenario to active module chapter.
	err = chap.AddScenario(scen)
	if err != nil {
		return fmt.Errorf("fail to add scenario to chapter: %v",
			err)
	}
	return nil
}

// modConf loads module configuration file
// from specified path.
func modConf(path, lang string) (module.ModConf, error) {
	conf := module.ModConf{Path: path, Lang: lang}
	// Check if mod dir exists.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return conf, fmt.Errorf("module not found: '%s': %v", path, err)
	}
	modConfPath := filepath.FromSlash(path + "/mod.conf")
	// Read conf.
	confValues, err := text.ReadValue(modConfPath, "id", "start-chapter")
	if err != nil {
		return conf, fmt.Errorf("fail to retrieve values: %s", err)
	}
	// Set conf values.
	conf.ID = confValues["id"]
	conf.StartChapter = confValues["start-chapter"]
	return conf, nil
}

// chapterConf loads chapter configuration file,
// returns error if configuration not found or corrupted.
func chapterConf(chapterPath string) (module.ChapterConf, error) {
	confPath := filepath.FromSlash(chapterPath + "/chapter.conf")
	confValues, err := text.ReadValue(confPath, "start-scenario")
	if err != nil {
		return module.ChapterConf{}, fmt.Errorf("fail to read conf values: %v",
			err)
	}
	conf := module.ChapterConf{
		Path:        chapterPath,
		StartScenID: confValues["start-scenario"],
	}
	return conf, nil
}

// buildArea creates area from specified data.
func buildArea(mod *module.Module, data res.ModuleAreaData) *scenario.Area {
	area := scenario.NewArea(data.ID)
	// NPCs.
	for _, areaCharData := range data.NPCS {
		// Retireve char data.
		charData := res.Character(areaCharData.ID)
		if charData == nil {
			log.Err.Printf("data: build area: %s: npc data not found: %s",
				data.ID, areaCharData.ID)
			continue
		}
		char := buildCharacter(mod, charData)
		// Set serial & position.
		serial.AssignSerial(char)
		char.SetPosition(areaCharData.PosX, areaCharData.PosY)
		char.SetDefaultPosition(areaCharData.PosX, areaCharData.PosY)
		// Char to area.
		area.AddCharacter(char)
	}
	// Objects.
	for _, areaObData := range data.Objects {
		// Retrieve object data.
		object, err := Object(mod, areaObData.ID)
		if err != nil {
			log.Err.Printf("data: build area %s: object: %s: %v",
				data.ID, areaObData.ID, err)
			continue
		}
		// Set serial & position.
		serial.AssignSerial(object)
		object.SetPosition(areaObData.PosX, areaObData.PosY)
		// Object to area.
		area.AddObject(object)
	}
	for _, subareaData := range data.Subareas {
		subarea := buildArea(mod, subareaData)
		area.AddSubarea(subarea)
	}
	return area
}
