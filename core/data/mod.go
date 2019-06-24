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
	"strings"

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/serial"
	"github.com/isangeles/flame/log"
)

const (
	SCENARIO_FILE_EXT = ".scenario"
)

// Module creates new module from specified path.
func Module(path, langID string) (*module.Module, error) {
	// Load module config file.
	mc, err := modConf(path, langID)
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_module_config:%v",
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
		return fmt.Errorf("fail_to_read_chapter_conf:%s:%v",
			chapPath, err)
	}
	chapConf.ID = id
	chapConf.ModulePath = mod.Conf().Path
	// Create chapter & set as current module chapter.
	startChap := module.NewChapter(mod, chapConf)
	err = mod.SetChapter(startChap) // move to start chapter
	if err != nil {
		return fmt.Errorf("fail_to_set_mod_chapter:%v", err)
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
		return fmt.Errorf("fail_to_open_scenario_file:%v", err)
	}
	defer docScen.Close()
	// Unmarshal scenario file.
	scenData, err := parsexml.UnmarshalScenario(docScen)
	if err != nil {
		return fmt.Errorf("fail_to_parse_scenario_file:%v", err)
	}
	// Build mainarea.
	var mainarea *scenario.Area
	subareas := make([]*scenario.Area, 0)
	for _, areaData := range scenData.Areas {
		area := scenario.NewArea(areaData.ID)
		// NPCs.
		for _, areaChar := range areaData.NPCS {
			// Retireve char data.
			charData := res.Character(areaChar.ID)
			if charData == nil {
				log.Err.Printf("data:unmarshal_scenario:%s:area:%s:npc_data_not_found:%s",
					scenData.ID, areaData.ID, areaChar.ID)
				continue
			}
			char := buildCharacter(mod, charData)
			// Set serial & position.
			serial.AssignSerial(char)
			char.SetPosition(areaChar.PosX, areaChar.PosY)
			char.SetDefaultPosition(areaChar.PosX, areaChar.PosY)
			// Char to area.
			area.AddCharacter(char)
		}
		// Objects.
		for _, areaObject := range areaData.Objects {
			// Retrieve object data.
			object, err := Object(mod, areaObject.ID)
			if err != nil {
				log.Err.Printf("data:unmarshal_scenario:area:%s:%s:%v",
					scenData.ID, areaData.ID, err)
				continue
			}
			// Set serial & position.
			serial.AssignSerial(object)
			object.SetPosition(areaObject.PosX, areaObject.PosY)
			// Object to area.
			area.AddObject(object)
		}
		if areaData.Main {
			mainarea = area
			continue
		}
		subareas = append(subareas, area)
	}
	scen := scenario.NewScenario(scenData.ID, mainarea, subareas)
	// Add scenario to active module chapter.
	err = chap.AddScenario(scen)
	if err != nil {
		return fmt.Errorf("fail_to_add_scenario_to_chapter:%v",
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
		return conf, fmt.Errorf("module_not_found:'%s':%v", path, err)
	}
	modConfPath := filepath.FromSlash(path + "/mod.conf")
	// Read conf.
	confValues, err := text.ReadValue(modConfPath, "id", "start-chapter",
		"char-skills", "char-items")
	if err != nil {
		return conf, fmt.Errorf("fail_to_retrieve_values:%s", err)
	}
	charSkills := strings.Split(confValues["char-skills"], ";")
	charItems := strings.Split(confValues["char-items"], ";")
	// Set conf values.
	conf.ID = confValues["id"]
	conf.StartChapter = confValues["start-chapter"]
	for _, sid := range charSkills {
		if len(sid) < 1 {
			continue
		}
		conf.CharSkills = append(conf.CharSkills, sid)
	}
	for _, iid := range charItems {
		if len(iid) < 1 {
			continue
		}
		conf.CharItems = append(conf.CharItems, iid)
	}
	return conf, nil
}

// chapterConf loads chapter configuration file,
// returns error if configuration not found or corrupted.
func chapterConf(chapterPath string) (module.ChapterConf, error) {
	confPath := filepath.FromSlash(chapterPath + "/chapter.conf")
	confValues, err := text.ReadValue(confPath, "start-scenario")
	if err != nil {
		return module.ChapterConf{}, fmt.Errorf("fail_to_read_conf_values:%v",
			err)
	}
	conf := module.ChapterConf{
		Path:        chapterPath,
		StartScenID: confValues["start-scenario"],
	}
	return conf, nil
}
