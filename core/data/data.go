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

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module"
)

// LoadTranslationData loads all lang files from
// from directory with specified path.
func LoadTranslationData(path string) error {	
	// Translation.
	langData, err := ImportLangDir(path)
	if err != nil {
		return fmt.Errorf("fail to import lang dir: %v", err)
	}
	resData := res.Translations()
	for _, td := range langData {
		resData = append(resData, td)
	}
	res.SetTranslationData(resData)
	return nil
}

// LoadModuleData loads module data(items, skills, etc.)
// for specified module.
func LoadModuleData(mod *module.Module) error {
	// Effects.
	effectsData, err := ImportEffectsDir(mod.Conf().EffectsPath())
	if err != nil {
		return fmt.Errorf("fail to load effects: %v", err)
	}
	res.SetEffectsData(effectsData)
	// Skills.
	skillsData, err := ImportSkillsDir(mod.Conf().SkillsPath())
	if err != nil {
		return fmt.Errorf("fail to load skills: %v", err)
	}
	res.SetSkillsData(skillsData)
	// Armors.
	armorsData, err := ImportArmorsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("fail to load armros: %v", err)
	}
	res.SetArmorsData(armorsData)
	// Weapons.
	weaponsData, err := ImportWeaponsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("fail to load weapons: %v", err)
	}
	res.SetWeaponsData(weaponsData)
	// Misc items.
	miscItemsData, err := ImportMiscItemsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("fail to load misc items: %v", err)
	}
	res.SetMiscItemsData(miscItemsData)
	// Recipes.
	recipesData, err := ImportRecipesDir(mod.Conf().RecipesPath())
	if err != nil {
		return fmt.Errorf("fail to load recipes: %v", err)
	}
	res.SetRecipesData(recipesData)
	// Area objects.
	objectsData, err := ImportObjectsDir(mod.Conf().ObjectsPath())
	if err != nil {
		return fmt.Errorf("fail to load area objects: %v", err)
	}
	res.SetObjectsData(objectsData)
	// Translation.
	err = LoadTranslationData(mod.Conf().LangPath())
	if err != nil {
		return fmt.Errorf("fail to load translation data: %v", err)
	}
	return nil
}

// LoadChapterData loads chapter data(NPC, quests, etc.)
// for specified chapter.
func LoadChapterData(chapter *module.Chapter) error {
	// NPC.
	npcData, err := ImportCharactersDataDir(chapter.Conf().NPCPath())
	if err != nil {
		return fmt.Errorf("fail to import npc: %v", err)
	}
	res.SetCharactersData(npcData)
	// Area objects.
	objectsData, err := ImportObjectsDir(chapter.Conf().ObjectsPath())
	if err != nil {
		return fmt.Errorf("fail to import object: %v", err)
	}
	res.AddObjectData(objectsData...) // adding to global module objects
	// Dialogs.
	dialogsData, err := ImportDialogsDir(chapter.Conf().DialogsPath())
	if err != nil {
		return fmt.Errorf("fail to import dialogs: %v", err)
	}
	res.SetDialogsData(dialogsData)
	// Quests.
	questsData, err := ImportQuestsDir(chapter.Conf().QuestsPath())
	if err != nil {
		return fmt.Errorf("fail to import quests: %v", err)
	}
	res.SetQuestsData(questsData)
	// Areas.
	areasData, err := ImportAreasDir(chapter.Conf().AreasPath())
	if err != nil {
		return fmt.Errorf("fail to import areas: %v", err)
	}
	res.SetAreasData(areasData)
	// Translation.
	err = LoadTranslationData(chapter.Conf().LangPath())
	if err != nil {
		return fmt.Errorf("fail to import translation data: %v", err)
	}
	return nil
}
