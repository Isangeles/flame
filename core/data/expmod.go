/*
 * expmod.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/core/data/parsetxt"
	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/area"
	"github.com/isangeles/flame/core/module/character"
)

// ExportModule exports module to new a directory under specified path.
func ExportModule(mod *module.Module, path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("unable to create module dir: %v", err)
	}
	confPath := filepath.Join(path, "mod.conf")
	err = exportModuleConfig(mod.Conf(), confPath)
	if err != nil {
		return fmt.Errorf("unable to create module config file: %v", err)
	}
	chapterPath := filepath.Join(path, "chapters", mod.Chapter().ID())
	err = exportChapter(mod.Chapter(), chapterPath)
	if err != nil {
		return fmt.Errorf("unable to export module chapter: %v", err)
	}
	return nil
}

// exportChapter exports chapter to a new directory under specified path.
func exportChapter(chapter *module.Chapter, path string) error {
	// Dir.
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("unable to create chapter dir: %v", err)
	}
	// Config.
	confPath := filepath.Join(path, "chapter.conf")
	err = exportChapterConfig(chapter.Conf(), confPath)
	if err != nil {
		return fmt.Errorf("unable to create chapter config file: %v", err)
	}
	// Areas.
	areasPath := filepath.Join(path, "areas")
	for _, a := range chapter.Areas() {
		areaPath := filepath.Join(areasPath, a.ID())
		err = exportArea(a, areaPath)
		if err != nil {
			return fmt.Errorf("unable to export area: %s: %v", a.ID(), err)
		}
	}
	// Characters.
	charsPath := filepath.Join(path, "characters")
	err = exportCharacters(chapter.Characters(), charsPath)
	if err != nil {
		return fmt.Errorf("unable to export characters: %v", err)
	}
	return nil
}

// exportArea exports area to a new file under specified
// directory.
func exportArea(area *area.Area, path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("unable to create area dir: %v", err)
	}
	areaData := area.Data()
	xmlArea, err := parsexml.MarshalArea(&areaData)
	if err != nil {
		return fmt.Errorf("unable to marshal area data: %v", err)
	}
	areaFilePath := filepath.Join(path, "main.area")
	areaFile, err := os.Create(areaFilePath)
	if err != nil {
		return fmt.Errorf("unable to create area file: %v", err)
	}
	defer areaFile.Close()
	w := bufio.NewWriter(areaFile)
	w.WriteString(xmlArea)
	w.Flush()
	return nil
}

// exportCharacters exports character to a new directory under
// specified path.
func exportCharacters(chars []*character.Character, path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("unable to create characters dir: %v", err)
	}
	xmlChars, err := parsexml.MarshalCharacters(chars...)
	if err != nil {
		return fmt.Errorf("unable to marshal characters: %v", err)
	}
	charsFilePath := filepath.Join(path, "main.characters")
	charsFile, err := os.Create(charsFilePath)
	if err != nil {
		return fmt.Errorf("unable to create characters file: %v", err)
	}
	defer charsFile.Close()
	w := bufio.NewWriter(charsFile)
	w.WriteString(xmlChars)
	w.Flush()
	return nil
}

// exportModuleConfig exports module config to a new
// file under specified path.
func exportModuleConfig(conf module.Config, path string) error {
	confValues := make(map[string][]string)
	confValues["id"] = []string{conf.ID}
	confValues["path"] = []string{conf.Path}
	confValues["lang"] = []string{conf.Lang}
	confValues["chapter"] = []string{conf.Chapter}
	config := parsetxt.MarshalConfig(confValues)
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create config file: %v", err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString(config)
	w.Flush()
	return nil
}

// exportChapterConfig exports chapter config to a new
// file under specified path.
func exportChapterConfig(conf module.ChapterConfig, path string) error {
	confValues := make(map[string][]string)
	confValues["id"] = []string{conf.ID}
	confValues["start-area"] = []string{conf.StartArea}
	config := parsetxt.MarshalConfig(confValues)
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create config file: %v", err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString(config)
	w.Flush()
	return nil
}
