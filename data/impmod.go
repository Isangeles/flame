/*
 * impmod.go
 *
 * Copyright 2018-2023 Dariusz Sikora <dev@isangeles.pl>
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
	"io"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/text"
	"github.com/isangeles/flame/log"
)

const (
	ModuleFileExt     = ".mod"
	ModuleConfigFile  = ".module"
	ChapterConfigFile = ".chapter"
)

// ImportModule imports module from module file with specified path.
func ImportModule(path string) (res.ModuleData, error) {
	data := res.ModuleData{}
	file, err := os.Open(path)
	if err != nil {
		return data, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return data, fmt.Errorf("unable to read file: %v", err)
	}
	err = unmarshal(buf, &data)
	if err != nil {
		return data, fmt.Errorf("unable to unmarshal JSON data: %v", err)
	}
	return data, nil
}

// ImportModuleDir imports module from directory with specified path.
func ImportModuleDir(path string) (data res.ModuleData, err error) {
	// Load module config file.
	file, err := os.Open(filepath.Join(path, ModuleConfigFile))
	if err != nil {
		return data, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()
	data.Config, err = text.UnmarshalConfig(file)
	if err != nil {
		return data, fmt.Errorf("unable to unmarshal config file: %v", err)
	}
	data.Config["id"] = []string{filepath.Base(path)}
	data.Config["path"] = []string{path}
	// Characters.
	data.Resources.Characters, err = ImportCharactersDir(filepath.Join(path, "characters"))
	if err != nil {
		log.Err.Printf("Import module: unable to import characters: %v", err)
	}
	// Races.
	data.Resources.Races, err = ImportRacesDir(filepath.Join(path, "characters/races"))
	if err != nil {
		log.Err.Printf("Import module: unable to imports races: %v", err)
	}
	// Skills.
	data.Resources.Skills, err = ImportSkillsDir(filepath.Join(path, "skills"))
	if err != nil {
		log.Err.Printf("Import module: unable to import skills: %v", err)
	}
	// Effects.
	data.Resources.Effects, err = ImportEffectsDir(filepath.Join(path, "effects"))
	if err != nil {
		log.Err.Printf("Import module: unable to import effects: %v", err)
	}
	// Armors.
	data.Resources.Armors, err = ImportArmorsDir(filepath.Join(path, "items/armors"))
	if err != nil {
		log.Err.Printf("Import module: unable to import armors: %v", err)
	}
	// Weapons.
	data.Resources.Weapons, err = ImportWeaponsDir(filepath.Join(path, "items/weapons"))
	if err != nil {
		log.Err.Printf("Import module: unable to import weapons: %v", err)
	}
	// Miscs.
	data.Resources.Miscs, err = ImportMiscItemsDir(filepath.Join(path, "items/misc"))
	if err != nil {
		log.Err.Printf("Import module: unable to import misc items: %v", err)
	}
	// Recipes.
	data.Resources.Recipes, err = ImportRecipesDir(filepath.Join(path, "recipes"))
	if err != nil {
		log.Err.Printf("Import module: unable to import recipes: %v", err)
	}
	// Trainings.
	data.Resources.Trainings, err = ImportTrainingsDir(filepath.Join(path, "trainings"))
	if err != nil {
		log.Err.Printf("Import module: unable to import trainings: %v", err)
	}
	// Translations.
	data.Resources.TranslationBases, err = ImportLangDirs(filepath.Join(path, "lang"))
	if err != nil {
		log.Err.Printf("Import module: unable to import translations: %v", err)
	}
	// Chapter.
	if len(data.Config["chapter"]) < 1 {
		return data, fmt.Errorf("no chapter set: %v", err)
	}
	chapterPath := filepath.Join(path, "chapters", data.Config["chapter"][0])
	data.Chapter, err = ImportChapterDir(chapterPath)
	if err != nil {
		return data, fmt.Errorf("unable to import chapter: %v", err)
	}
	return data, nil
}

// ImportChapterDir imports chapter from directory with specified path.
func ImportChapterDir(path string) (data res.ChapterData, err error) {
	// Load module config file.
	file, err := os.Open(filepath.Join(path, ChapterConfigFile))
	if err != nil {
		return data, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()
	data.Config, err = text.UnmarshalConfig(file)
	if err != nil {
		return data, fmt.Errorf("unable to unmarshal config file: %v", err)
	}
	data.Config["id"] = []string{filepath.Base(path)}
	data.Config["path"] = []string{path}
	// Characters.
	data.Resources.Characters, err = ImportCharactersDir(filepath.Join(path, "characters"))
	if err != nil {
		log.Err.Printf("Import chapter: unable to import characters: %v", err)
	}
	// Quests.
	data.Resources.Quests, err = ImportQuestsDir(filepath.Join(path, "quests"))
	if err != nil {
		log.Err.Printf("Import chapter: unable to import quests: %v", err)
	}
	// Dialogs.
	data.Resources.Dialogs, err = ImportDialogsDir(filepath.Join(path, "dialogs"))
	if err != nil {
		log.Err.Printf("Import chapter: unable to import dialogs: %v", err)
	}
	// Areas.
	data.Resources.Areas, err = ImportAreasDir(filepath.Join(path, "areas"))
	if err != nil {
		log.Err.Printf("Import chapter: unable to import areas: %v", err)
	}
	// Translations.
	data.Resources.TranslationBases, err = ImportLangDirs(filepath.Join(path, "lang"))
	if err != nil {
		log.Err.Printf("Import chapter: unable to import translations: %v", err)
	}
	return data, nil
}
