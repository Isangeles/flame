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
	"github.com/isangeles/flame/core/module/area"
	"github.com/isangeles/flame/core/module/serial"
	"github.com/isangeles/flame/log"
)

const (
	AreaFileExt = ".area"
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

// LoadArea loads area with specified
// ID for current module chapter.
func LoadArea(mod *module.Module, id string) error {
	// Check whether mod has active chapter.
	chap := mod.Chapter()
	if chap == nil {
		return fmt.Errorf("no module chapter set")
	}
	// Load files.
	areaPath := filepath.FromSlash(fmt.Sprintf("%s/%s/main%s",
		chap.Conf().AreasPath(), id, AreaFileExt))
	docArea, err := os.Open(areaPath)
	if err != nil {
		return fmt.Errorf("fail to open area file: %v", err)
	}
	defer docArea.Close()
	// Unmarshal area file.
	areaData, err := parsexml.UnmarshalArea(docArea)
	if err != nil {
		return fmt.Errorf("fail to parse area data: %v", err)
	}
	// Build mainarea.
	mainarea := buildArea(mod, *areaData)
	// Add area to active module chapter.
	chap.AddAreas(mainarea)
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
	confValues, err := text.ReadValue(confPath, "start-area")
	if err != nil {
		return module.ChapterConf{}, fmt.Errorf("fail to read conf values: %v",
			err)
	}
	conf := module.ChapterConf{
		Path:        chapterPath,
		StartAreaID: confValues["start-area"],
	}
	return conf, nil
}

// buildArea creates area from specified data.
func buildArea(mod *module.Module, data res.ModuleAreaData) *area.Area {
	area := area.NewArea(data.ID)
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
	// Subareas.
	for _, subareaData := range data.Subareas {
		subarea := buildArea(mod, subareaData)
		area.AddSubarea(subarea)
	}
	return area
}
