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
	"path/filepath"

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
	chapConf, err := importChapterConfig(confPath)
	if err != nil {
		return fmt.Errorf("unable to read chapter conf: %s: %v",
			chapPath, err)
	}
	chapConf.ID = id
	chapConf.ModulePath = mod.Conf().Path
	// Create chapter & set as current module chapter.
	chapter := module.NewChapter(mod, chapConf)
	// Load chapter data.
	err = loadChapterData(chapter)
	if err != nil {
		return fmt.Errorf("unable to load chapter data: %v", err)
	}
	mod.SetChapter(chapter)
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
	mainarea := area.New(*areaData)
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
	res.SetEffectsData(effectsData)
	// Skills.
	skillsData, err := ImportSkillsDir(mod.Conf().SkillsPath())
	if err != nil {
		return fmt.Errorf("unable to load skills: %v", err)
	}
	res.SetSkillsData(skillsData)
	// Armors.
	armorsData, err := ImportArmorsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("unable to load armros: %v", err)
	}
	res.SetArmorsData(armorsData)
	// Weapons.
	weaponsData, err := ImportWeaponsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("unable to load weapons: %v", err)
	}
	res.SetWeaponsData(weaponsData)
	// Misc items.
	miscItemsData, err := ImportMiscItemsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("unable to load misc items: %v", err)
	}
	res.SetMiscItemsData(miscItemsData)
	// Recipes.
	recipesData, err := ImportRecipesDir(mod.Conf().RecipesPath())
	if err != nil {
		return fmt.Errorf("unable to load recipes: %v", err)
	}
	res.SetRecipesData(recipesData)
	// Area objects.
	objectsData, err := ImportObjectsDir(mod.Conf().ObjectsPath())
	if err != nil {
		return fmt.Errorf("unable to load area objects: %v", err)
	}
	res.SetObjectsData(objectsData)
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
	res.SetCharactersData(charactersData)
	// Area objects.
	objectsData, err := ImportObjectsDir(chapter.Conf().ObjectsPath())
	if err != nil {
		return fmt.Errorf("unable to import object: %v", err)
	}
	res.AddObjectData(objectsData...) // adding to global module objects
	// Dialogs.
	dialogsData, err := ImportDialogsDir(chapter.Conf().DialogsPath())
	if err != nil {
		return fmt.Errorf("unable to import dialogs: %v", err)
	}
	res.SetDialogsData(dialogsData)
	// Quests.
	questsData, err := ImportQuestsDir(chapter.Conf().QuestsPath())
	if err != nil {
		return fmt.Errorf("unable to import quests: %v", err)
	}
	res.SetQuestsData(questsData)
	// Areas.
	areasData, err := ImportAreasDir(chapter.Conf().AreasPath())
	if err != nil {
		return fmt.Errorf("unable to import areas: %v", err)
	}
	res.SetAreasData(areasData)
	return nil
}
