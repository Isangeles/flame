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

	"github.com/isangeles/flame/data/text"
	"github.com/isangeles/flame/data/res"
)

const (
	ModuleConfigFile = ".module"
	ChapterConfigFile = ".chapter"
)

// ImportModule imports module from directory with specified path.
func ImportModule(path string) (data res.ModuleData, err error) {
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
		return data, fmt.Errorf("unable to import characters: %v", err)
	}
	// Races.
	data.Resources.Races, err = ImportRacesDir(filepath.Join(path, "characters/races"))
	if err != nil {
		return data, fmt.Errorf("unable to imports races: %v", err)
	}
	// Objects.
	data.Resources.Objects, err = ImportObjectsDir(filepath.Join(path, "objects"))
	if err != nil {
		return data, fmt.Errorf("unable to import objects: %v", err)
	}
	// Skills.
	data.Resources.Skills, err = ImportSkillsDir(filepath.Join(path, "skills"))
	if err != nil {
		return data, fmt.Errorf("unable to import skills: %v", err)
	}
	// Effects.
	data.Resources.Effects, err = ImportEffectsDir(filepath.Join(path, "effects"))
	if err != nil {
		return data, fmt.Errorf("unable to import effects: %v", err)
	}
	// Armors.
	data.Resources.Armors, err = ImportArmorsDir(filepath.Join(path, "items"))
	if err != nil {
		return data, fmt.Errorf("unable to import armors: %v", err)
	}
	// Weapons.
	data.Resources.Weapons, err = ImportWeaponsDir(filepath.Join(path, "items"))
	if err != nil {
		return data, fmt.Errorf("unable to import weapons: %v", err)
	}
	// Miscs.
	data.Resources.Miscs, err = ImportMiscItemsDir(filepath.Join(path, "items"))
	if err != nil {
		return data, fmt.Errorf("unable to import misc items: %v", err)
	}
	// Recipes.
	data.Resources.Recipes, err = ImportRecipesDir(filepath.Join(path, "recipes"))
	if err != nil {
		return data, fmt.Errorf("unable to import recipes: %v", err)
	}
	// Trainings.
	data.Resources.Trainings, err = ImportTrainingsDir(filepath.Join(path, "trainings"))
	if err != nil {
		return data, fmt.Errorf("unable to import trainings: %v", err)
	}
	// Chapter.
	if len(data.Config["chapter"]) < 1 {
		return data, fmt.Errorf("no chapter set: %v", err)
	}
	chapterPath := filepath.Join(path, "chapters", data.Config["chapter"][0])
	data.Chapter, err = ImportChapter(chapterPath)
	if err != nil {
		return data, fmt.Errorf("unable to import chapter: %v", err)
	}
	return data, nil
}

// ImportChapter imports chapter from directory with specified path.
func ImportChapter(path string) (data res.ChapterData, err error) {
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
		return data, fmt.Errorf("unable to import characters: %v", err)
	}
	// Objects.
	data.Resources.Objects, err = ImportObjectsDir(filepath.Join(path, "objects"))
	if err != nil {
		return data, fmt.Errorf("unable to import objects: %v", err)
	}
	// Quests.
	data.Resources.Quests, err = ImportQuestsDir(filepath.Join(path, "quests"))
	if err != nil {
		return data, fmt.Errorf("unable to import quests: %v", err)
	}
	// Dialogs.
	data.Resources.Dialogs, err = ImportDialogsDir(filepath.Join(path, "dialogs"))
	if err != nil {
		return data, fmt.Errorf("unable to import dialogs: %v", err)
	}
	// Areas.
	data.Resources.Areas, err = ImportAreasDir(filepath.Join(path, "areas"))
	if err != nil {
		return data, fmt.Errorf("unable to import areas: %v", err)
	}
	return data, nil
}
