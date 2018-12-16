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

	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/module"
)

// LoadModConf loads module configuration file
// from specified path.
func LoadModConf(path, lang string) (module.ModConf, error) {
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

// ChapterConf loads chapter configuration file,
// returns error if configuration not found or corrupted.
func LoadChapterConf(chapterPath string) (module.ChapterConf, error) {
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
