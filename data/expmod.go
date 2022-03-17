/*
 * expmod.go
 *
 * Copyright 2020-2022 Dariusz Sikora <dev@isangeles.pl>
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
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/text"
)

// ExportModule exports module data to the single file.
func ExportModule(path string, data res.ModuleData) error {
	json, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal module data: %v", err)
	}
	if !strings.HasSuffix(path, ModuleFileExt) {
		path += ModuleFileExt
	}
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("unable to create module file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create module file: %v", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(json)
	writer.Flush()
	return nil
}

// ExportModuleDir exports module data to new a directory under specified path.
func ExportModuleDir(path string, data res.ModuleData) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("unable to create module dir: %v", err)
	}
	// Config.
	confPath := filepath.Join(path, ".module")
	err = exportConfig(confPath, data.Config)
	if err != nil {
		return fmt.Errorf("unable to create config file: %v", err)
	}
	// Characters.
	charsPath := filepath.Join(path, "characters", "main")
	err = ExportCharacters(charsPath, data.Resources.Characters...)
	if err != nil {
		return fmt.Errorf("unable to export characters: %v", err)
	}
	// Races.
	racesPath := filepath.Join(path, "characters/races/main")
	err = ExportRaces(racesPath, data.Resources.Races...)
	if err != nil {
		return fmt.Errorf("unable to export races: %v", err)
	}
	// Objects.
	objectsPath := filepath.Join(path, "objects", "main")
	err = ExportObjects(objectsPath, data.Resources.Objects...)
	if err != nil {
		return fmt.Errorf("unable to export objects: %v", err)
	}
	// Skills.
	skillsPath := filepath.Join(path, "skills", "main")
	err = ExportSkills(skillsPath, data.Resources.Skills...)
	if err != nil {
		return fmt.Errorf("unable to export skills: %v", err)
	}
	// Effects.
	effectsPath := filepath.Join(path, "effects", "main")
	err = ExportEffects(effectsPath, data.Resources.Effects...)
	if err != nil {
		return fmt.Errorf("unable to export effects: %v", err)
	}
	// Armors.
	armorsPath := filepath.Join(path, "items/armors/main")
	err = ExportArmors(armorsPath, data.Resources.Armors...)
	if err != nil {
		return fmt.Errorf("unable to export armors: %v", err)
	}
	// Weapons.
	weaponsPath := filepath.Join(path, "items/weapons/main")
	err = ExportWeapons(weaponsPath, data.Resources.Weapons...)
	if err != nil {
		return fmt.Errorf("unable to export weapons: %v", err)
	}
	// Miscs.
	miscPath := filepath.Join(path, "items/misc/main")
	err = ExportMiscItems(miscPath, data.Resources.Miscs...)
	if err != nil {
		return fmt.Errorf("unable to export misc items: %v", err)
	}
	// Recipes.
	recipesPath := filepath.Join(path, "recipes", "main")
	err = ExportRecipes(recipesPath, data.Resources.Recipes...)
	if err != nil {
		return fmt.Errorf("unable to export recipes: %v", err)
	}
	// Trainings.
	trainingsPath := filepath.Join(path, "trainings", "main")
	err = ExportTrainings(trainingsPath, data.Resources.Trainings...)
	if err != nil {
		return fmt.Errorf("unable to export trainings: %v", err)
	}
	// Translations.
	langPath := filepath.Join(path, "lang")
	err = ExportLangDirs(langPath, data.Resources.TranslationBases...)
	if err != nil {
		return fmt.Errorf("unable to export translations: %v", err)
	}
	// Chapters.
	chapterPath := filepath.Join(path, "chapters", data.Chapter.ID)
	err = exportChapterDir(chapterPath, data.Chapter)
	if err != nil {
		return fmt.Errorf("unable to export chapter: %v", err)
	}
	return nil
}

// exportChapterDir exports chapter to a new directory under specified path.
func exportChapterDir(path string, data res.ChapterData) error {
	// Dir.
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("unable to create chapter dir: %v", err)
	}
	// Config.
	confPath := filepath.Join(path, ".chapter")
	err = exportConfig(confPath, data.Config)
	if err != nil {
		return fmt.Errorf("unable to create config file: %v", err)
	}
	// Characters.
	charsPath := filepath.Join(path, "characters", "main")
	err = ExportCharacters(charsPath, data.Resources.Characters...)
	if err != nil {
		return fmt.Errorf("unable to export characters: %v", err)
	}
	// Objects.
	objectsPath := filepath.Join(path, "objects", "main")
	err = ExportObjects(objectsPath, data.Resources.Objects...)
	if err != nil {
		return fmt.Errorf("unable to export objects: %v", err)
	}
	// Quests.
	questsPath := filepath.Join(path, "quests", "main")
	err = ExportQuests(questsPath, data.Resources.Quests...)
	if err != nil {
		return fmt.Errorf("unable to export quests: %v", err)
	}
	// Dialogs.
	dialogsPath := filepath.Join(path, "dialogs", "main")
	err = ExportDialogs(dialogsPath, data.Resources.Dialogs...)
	if err != nil {
		return fmt.Errorf("unable to export dialogs: %v", err)
	}
	// Areas.
	areasPath := filepath.Join(path, "areas")
	for _, a := range data.Areas {
		areaPath := filepath.Join(areasPath, a.ID, "main")
		err = ExportArea(areaPath, a)
		if err != nil {
			return fmt.Errorf("unable to export area: %s: %v", a.ID, err)
		}
	}
	// Translations.
	langPath := filepath.Join(path, "lang")
	err = ExportLangDirs(langPath, data.Resources.TranslationBases...)
	if err != nil {
		return fmt.Errorf("unable to export translations: %v", err)
	}
	return nil
}

// exportConfig exports config values to a config file
// new under specified path.
func exportConfig(path string, values map[string][]string) error {
	config := text.MarshalConfig(values)
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
