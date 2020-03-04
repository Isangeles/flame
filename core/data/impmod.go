/*
 * impmod.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"strconv"

	"github.com/isangeles/flame/core/data/parsetxt"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/area"
)

// ImportModule imports module from directory with specified path.
func ImportModule(path string) (*module.Module, error) {
	// Load module config file.
	confPath := filepath.Join(path, ".module")
	mc, err := importModuleConfig(confPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load module config: %v",
			err)
	}
	// Create module.
	m := module.New(mc)
	return m, nil
}

// LoadChapter loads module chapter with
// specified ID.
func LoadChapter(mod *module.Module, id string) error {
	// Load chapter config file.
	chapPath := filepath.Join(mod.Conf().ChaptersPath(), mod.Conf().Chapter)
	confPath := filepath.Join(chapPath, ".chapter")
	chapConf, err := importChapterConfig(confPath)
	if err != nil {
		return fmt.Errorf("unable to read chapter conf: %s: %v",
			chapPath, err)
	}
	chapConf.ID = id
	chapConf.ModulePath = mod.Conf().Path
	// Create chapter & set as current module chapter.
	startChap := module.NewChapter(mod, chapConf)
	mod.SetChapter(startChap)
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
	areaPath := filepath.Join(chap.Conf().AreasPath(), id)
	areaData, err := ImportArea(areaPath)
	if err != nil {
		return fmt.Errorf("unable to import area: %v", err)
	}
	// Build mainarea.
	mainarea := area.New(*areaData)
	// Add area to active module chapter.
	chap.AddAreas(mainarea)
	return nil
}

// exportModuleConfig imports module configuration  from
// file with specified path.
func importModuleConfig(path string) (module.Config, error) {
	conf := module.Config{Path: filepath.Dir(path)}
	// Check if mod dir exists.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return conf, fmt.Errorf("module not found: '%s': %v", path, err)
	}
	// Read conf.
	file, err := os.Open(path)
	if err != nil {
		return conf, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()
	confValues := parsetxt.UnmarshalConfig(file)
	// Set conf values.
	if len(confValues["id"]) > 0 {
		conf.ID = confValues["id"][0]
	}
	if len(confValues["chapter"]) > 0 {
		conf.Chapter = confValues["chapter"][0]
	}
	return conf, nil
}

// importChapterConfig imports chapter configuration from file with specified path,
// returns error if configuration not found or invalid.
func importChapterConfig(path string) (module.ChapterConfig, error) {
	conf := module.ChapterConfig{Path: filepath.Dir(path)}
	file, err := os.Open(path)
	if err != nil {
		return conf, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()
	confValues := parsetxt.UnmarshalConfig(file)
	if len(confValues["start-area"]) > 0 {
		conf.StartArea = confValues["start-area"][0]
	}
	if len(confValues["start-pos"]) > 1 {
		conf.StartPosX, _ = strconv.ParseFloat(confValues["start-pos"][0], 64)
		conf.StartPosY, _ = strconv.ParseFloat(confValues["start-pos"][1], 64)
	}
	return conf, nil
}
