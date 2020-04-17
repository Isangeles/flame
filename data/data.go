/*
 * data.go
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

// data package provides connection with external data
// files like items base, savegames, etc.
package data

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/data/parsetxt"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module"
	"github.com/isangeles/flame/module/area"
)

// LoadChapter loads chapter with  specified ID
// for specified module.
func LoadChapter(mod *module.Module, id string) error {
	// Load chapter config file.
	chapPath := filepath.Join(mod.Conf().ChaptersPath(), mod.Conf().Chapter)
	confPath := filepath.Join(chapPath, ".chapter")
	file, err := os.Open(confPath)
	if err != nil {
		return fmt.Errorf("unable to read chapter conf: %s: %v",
			chapPath, err)
	}
	defer file.Close()
	conf := parsetxt.UnmarshalConfig(file)
	conf["id"] = []string{id}
	conf["path"] = []string{chapPath}
	data := res.ChapterData{
		Config: conf,
	}
	// Create chapter & set as current module chapter.
	chapter := module.NewChapter(mod, data)
	// Load chapter data.
	err = loadChapterData(chapter)
	if err != nil {
		return fmt.Errorf("unable to load chapter data: %v", err)
	}
	mod.SetChapter(chapter)
	setModuleResources(mod)
	return nil
}

// LoadArea loads area with specified ID for current
// chapter of specified module.
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
	mainarea := area.New(areaData)
	// Add area to active module chapter.
	chap.AddAreas(mainarea)
	return nil
}

// LoadModuleLang loads translation data for specified
// language for specified module.
func LoadModuleLang(mod *module.Module, lang string) error {
	modLangPath := filepath.Join(mod.Conf().LangPath(), lang)
	err := LoadTranslationData(modLangPath)
	if err != nil {
		return fmt.Errorf("unable to load module translation data: %v", err)
	}
	if mod.Chapter() == nil {
		return nil
	}
	chapterLangPath := filepath.Join(mod.Chapter().Conf().LangPath(), lang)
	err = LoadTranslationData(chapterLangPath)
	if err != nil {
		return fmt.Errorf("unable to load chapter translation data: %v", err)
	}
	return nil
}

// LoadTranslationData loads all lang files from
// from directory with specified path.
func LoadTranslationData(path string) error {
	// Translation.
	langData, err := ImportLangDir(path)
	if err != nil {
		return fmt.Errorf("unable to import lang dir: %v", err)
	}
	resData := res.Translations()
	for _, td := range langData {
		resData = append(resData, td)
	}
	res.SetTranslationData(resData)
	return nil
}

// loadModuleData loads module data(items, skills, etc.)
// for specified module.
func loadModuleData(mod *module.Module) error {
	// Effects.
	effectsData, err := ImportEffectsDir(mod.Conf().EffectsPath())
	if err != nil {
		return fmt.Errorf("unable to load effects: %v", err)
	}
	mod.Res.Effects = effectsData
	// Skills.
	skillsData, err := ImportSkillsDir(mod.Conf().SkillsPath())
	if err != nil {
		return fmt.Errorf("unable to load skills: %v", err)
	}
	mod.Res.Skills = skillsData
	// Armors.
	armorsData, err := ImportArmorsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("unable to load armros: %v", err)
	}
	mod.Res.Armors = armorsData
	// Weapons.
	weaponsData, err := ImportWeaponsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("unable to load weapons: %v", err)
	}
	mod.Res.Weapons = weaponsData
	// Misc items.
	miscItemsData, err := ImportMiscItemsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("unable to load misc items: %v", err)
	}
	mod.Res.Miscs = miscItemsData
	// Recipes.
	recipesData, err := ImportRecipesDir(mod.Conf().RecipesPath())
	if err != nil {
		return fmt.Errorf("unable to load recipes: %v", err)
	}
	mod.Res.Recipes = recipesData
	// Area objects.
	objectsData, err := ImportObjectsDir(mod.Conf().ObjectsPath())
	if err != nil {
		return fmt.Errorf("unable to load area objects: %v", err)
	}
	mod.Res.Objects = objectsData
	// Races.
	racesData, err := ImportRacesDir(mod.Conf().RacesPath())
	if err != nil {
		return fmt.Errorf("unable to load races: %v", err)
	}
	mod.Res.Races = racesData
	setModuleResources(mod)
	return nil
}

// loadChapterData loads chapter data(NPC, quests, etc.)
// for specified chapter.
func loadChapterData(chapter *module.Chapter) error {
	// Characters.
	charactersData, err := ImportCharactersDataDir(chapter.Conf().CharactersPath())
	if err != nil {
		return fmt.Errorf("unable to import characters: %v", err)
	}
	chapter.Res.Characters = charactersData
	// Area objects.
	objectsData, err := ImportObjectsDir(chapter.Conf().ObjectsPath())
	if err != nil {
		return fmt.Errorf("unable to import object: %v", err)
	}
	chapter.Res.Objects = objectsData
	// Dialogs.
	dialogsData, err := ImportDialogsDir(chapter.Conf().DialogsPath())
	if err != nil {
		return fmt.Errorf("unable to import dialogs: %v", err)
	}
	chapter.Res.Dialogs = dialogsData
	// Quests.
	questsData, err := ImportQuestsDir(chapter.Conf().QuestsPath())
	if err != nil {
		return fmt.Errorf("unable to import quests: %v", err)
	}
	chapter.Res.Quests = questsData
	// Areas.
	areasData, err := ImportAreasDir(chapter.Conf().AreasPath())
	if err != nil {
		return fmt.Errorf("unable to import areas: %v", err)
	}
	chapter.Res.Areas = areasData
	return nil
}

// setModuleResources sets resources from specified module
func setModuleResources(mod *module.Module) {
	chars := mod.Res.Characters
	objects := mod.Res.Objects
	effects := mod.Res.Effects
	skills := mod.Res.Skills
	armors := mod.Res.Armors
	weapons := mod.Res.Weapons
	miscs := mod.Res.Miscs
	dialogs := mod.Res.Dialogs
	quests := mod.Res.Quests
	recipes := mod.Res.Recipes
	areas := mod.Res.Areas
	translations := append(res.Translations(), mod.Res.Translations...)
	if mod.Chapter() != nil {
		chars = append(chars, mod.Chapter().Res.Characters...)
		objects = append(objects, mod.Chapter().Res.Objects...)
		effects = append(effects, mod.Chapter().Res.Effects...)
		skills = append(skills, mod.Chapter().Res.Skills...)
		armors = append(armors, mod.Chapter().Res.Armors...)
		weapons = append(weapons, mod.Chapter().Res.Weapons...)
		miscs = append(miscs, mod.Chapter().Res.Miscs...)
		dialogs = append(dialogs, mod.Chapter().Res.Dialogs...)
		quests = append(quests, mod.Chapter().Res.Quests...)
		recipes = append(recipes, mod.Chapter().Res.Recipes...)
		areas = append(areas, mod.Chapter().Res.Areas...)
		translations = append(translations, mod.Chapter().Res.Translations...)
	}
	res.SetCharactersData(chars)
	res.SetObjectsData(objects)
	res.SetEffectsData(effects)
	res.SetSkillsData(skills)
	res.SetArmorsData(armors)
	res.SetWeaponsData(weapons)
	res.SetMiscItemsData(miscs)
	res.SetDialogsData(dialogs)
	res.SetQuestsData(quests)
	res.SetRecipesData(recipes)
	res.SetAreasData(areas)
	res.SetTranslationData(translations)
}
