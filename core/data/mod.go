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

	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/area"
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
	m := module.New(mc)
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
	areaPath := filepath.FromSlash(fmt.Sprintf("%s/%s",
		chap.Conf().AreasPath(), id))
	areaData, err := ImportArea(areaPath)
	if err != nil {
		return fmt.Errorf("fail to import area: %v", err)
	}
	// Build mainarea.
	mainarea := area.New(*areaData)
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
