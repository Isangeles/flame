/*
 * mod.go
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

package data

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/scenario"
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
	chapPath := filepath.FromSlash(mod.ChaptersPath() +
		"/" + mod.Conf().Chapters[0])
	chapConf, err := chapterConf(chapPath)
	if err != nil {
		return fmt.Errorf("fail_to_read_chapter_conf:%s:%v",
			chapPath, err)
	}
	chapConf.ID = id
	// Create chapter & set as current module
	// chapter.
	startChap := module.NewChapter(mod, chapConf)
	err = mod.SetChapter(startChap) // move to start chapter
	if err != nil {
		return fmt.Errorf("fail_to_set_mod_chapter:%v",
			err)
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
	scenPath := filepath.FromSlash(chap.ScenariosPath() + "/" +
		id)
	docScen, err := os.Open(scenPath)
	if err != nil {
		return fmt.Errorf("fail_to_open_scenario_file:%v", err)
	}
	defer docScen.Close()
	npcsBasePath := filepath.FromSlash(chap.NPCPath() + "/npc" + CHARS_FILE_EXT)
	docNPCs, err := os.Open(npcsBasePath)
	if err != nil {
		return fmt.Errorf("fail_to_open_characters_base_file:%v", err)
	}
	defer docNPCs.Close()
	npcsLangPath := filepath.FromSlash(chap.LangPath() + "/npc" + text.LANG_FILE_EXT)
	// Unmarshal scenario file.
	xmlScen, err := parsexml.UnmarshalScenario(docScen)
	if err != nil {
		return fmt.Errorf("fail_to_parse_scenario_file:%v", err)
	}
	// Build scenario mainarea.
	mainarea := scenario.NewArea(xmlScen.Mainarea.ID)
	for _, xmlAreaChar := range xmlScen.Mainarea.NPCs.Characters {
		// Unmarshal area NPC.
		charXML, err := parsexml.UnmarshalCharacter(docNPCs, xmlAreaChar.ID)
		if err != nil {
			log.Err.Printf("data_scenario_unmarshal_npc:%s:fail:%v",
				xmlAreaChar.ID, err)
			continue
		}
		// Build area NPC.
		char, err := buildXMLCharacter(&charXML)
		if err != nil {
			log.Err.Printf("data_scenario_build_npc:%s:fail:%v",
				xmlAreaChar.ID, err)
		}
		x, y, err := parsexml.UnmarshalPosition(xmlAreaChar.Position)
		if err != nil {
			log.Err.Printf("data_scenario_spawn_npc:%s:unmarshal_position_fail:%v",
				xmlAreaChar.ID, err)
			continue
		}
		char.SetPosition(x, y)
		name := text.ReadDisplayText(npcsLangPath, char.ID())
		char.SetName(name[0])
		err = mod.AssignSerial(char)
		if err != nil {
			log.Err.Printf("data_scenario_spawn_npc:%s:fail_to_assign_serial:%v",
				xmlAreaChar.ID, err)
			continue
		}
		mainarea.AddCharacter(char)
	}
	subareas := make([]*scenario.Area, 0)
	for _, xmlArea := range xmlScen.Subareas {
		area := scenario.NewArea(xmlArea.ID)
		// TODO: area NPCs.
		subareas = append(subareas, area)
	}
	scen := scenario.NewScenario(xmlScen.ID, mainarea, subareas)
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
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return module.ModConf{}, fmt.Errorf("module_not_found:'%s':%v",
			path, err)
	}
	modConfPath := filepath.FromSlash(path + "/mod.conf")
	confInts, err := text.ReadConfigInt(modConfPath, "new_char_attrs_min",
		"new_char_attrs_max")
	if err != nil {
		return module.ModConf{}, fmt.Errorf("fail_to_retrieve_int_values:%s",
			err)
	}
	confValues, err := text.ReadConfigValue(modConfPath, "name", "chapters")
	if err != nil {
		return module.ModConf{}, fmt.Errorf("fail_to_retrieve_values:%s",
			err)
	}
	chapters := strings.Split(confValues[1], ";")
	if len(chapters) < 1 {
		return module.ModConf{}, fmt.Errorf("no_chapters_specified")
	}
	conf := module.ModConf{
		Name:            confValues[0],
		Path:            path,
		Lang:            lang,
		NewcharAttrsMin: confInts[0],
		NewcharAttrsMax: confInts[1],
		Chapters:        chapters,
	}
	return conf, nil
}

// chapterConf loads chapter configuration file,
// returns error if configuration not found or corrupted.
func chapterConf(chapterPath string) (module.ChapterConf, error) {
	confPath := filepath.FromSlash(chapterPath + "/chapter.conf")
	confValues, err := text.ReadConfigValue(confPath, "start_scenario")
	if err != nil {
		return module.ChapterConf{}, fmt.Errorf("fail_to_read_conf_values:%v",
			err)
	}
	conf := module.ChapterConf{
		Path:chapterPath,
		StartScenID:confValues[0],
	}
	return conf, nil
}
