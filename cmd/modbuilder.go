/*
 * modbuilder.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
)

// NewModule creates new module directory
// in data/modules with all one chapter and
// one empty start area.
func NewModule(name string) error {
	path := filepath.FromSlash("data/modules/" + name)
	// Mod dir.
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_module_dir:%v", err)
	}
	// Sub-dirs.
	err = os.MkdirAll(filepath.FromSlash(path+"/characters"), 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_characters_dir:%v", err)
	}
	err = os.MkdirAll(filepath.FromSlash(path+"/items"), 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_items_dir:%v", err)
	}
	err = os.MkdirAll(filepath.FromSlash(path+"/effects"), 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_effects_dir:%v", err)
	}
	err = os.MkdirAll(filepath.FromSlash(path+"/skills"), 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_skills_dir:%v", err)
	}
	err = os.MkdirAll(filepath.FromSlash(path+"/objects"), 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_objects_dir:%v", err)
	}
	err = os.MkdirAll(filepath.FromSlash(path+"/lang"), 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_lang_dir:%v", err)
	}
	// Mod conf.
	confPath := filepath.FromSlash(path + "/mod.conf")
	confFile, err := os.Create(confPath)
	if err != nil {
		return fmt.Errorf("fail_to_create_module_conf_file:%v", err)
	}
	defer confFile.Close()
	confFormat := "id:%s;\nstart-chapter:%s;\nchar-skills:%s;\nchar-items:%s;\n"
	conf := fmt.Sprintf(confFormat, name, "prologue", "", "")
	w := bufio.NewWriter(confFile)
	w.WriteString(conf)
	w.Flush()
	// Start chapter.
	err = createChapter(path+"/chapters/prologue", "prologue")
	if err != nil {
		return fmt.Errorf("fail_to_create_chapter:%v", err)
	}
	return nil
}

// createChapter creates new chapter
// directory.
func createChapter(path, id string) error {
	// Dir.
	err := os.MkdirAll(filepath.FromSlash(path), 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_dir:%v", err)
	}
	// Sub-dirs.
	err = os.MkdirAll(filepath.FromSlash(path+"/npc"), 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_npc_dir:%v", err)
	}
	err = os.MkdirAll(filepath.FromSlash(path+"/lang"), 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_lang_dir:%v", err)
	}
	// Conf.
	confPath := filepath.FromSlash(path + "/chapter.conf")
	confFile, err := os.Create(confPath)
	if err != nil {
		return fmt.Errorf("fail_to_create_conf_file:%v", err)
	}
	defer confFile.Close()
	conf := fmt.Sprintf("start-scenario:%s;\nscenarios:%s;\n",
		"area1.scenario", "")
	w := bufio.NewWriter(confFile)
	w.WriteString(conf)
	w.Flush()
	// Start area.
	scensPath := filepath.FromSlash(path + "/area/scenarios")
	err = os.MkdirAll(scensPath, 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_scenarios_dir:%v", err)
	}
	sd := res.ModuleScenarioData{ID: "area1"}
	adMain := res.ModuleAreaData{ID: "area1_main", Main: true}
	sd.Areas = append(sd.Areas, adMain)
	xmlScen, err := parsexml.MarshalScenario(&sd)
	if err != nil {
		return fmt.Errorf("fail_to_marshal_start_scenario:%v", err)
	}
	scenPath := filepath.FromSlash(scensPath + "/area1" + data.SCENARIO_FILE_EXT)
	scenFile, err := os.Create(scenPath)
	if err != nil {
		return fmt.Errorf("fail_to_create_start_scenario_file:%v", err)
	}
	defer scenFile.Close()
	w = bufio.NewWriter(scenFile)
	w.WriteString(xmlScen)
	w.Flush()
	return nil
}
