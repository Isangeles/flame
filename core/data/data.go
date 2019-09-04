/*
 * data.go
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

// data package provides connection with external data
// files like items base, savegames, etc.
package data

import (
	"fmt"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module"
)

// LoadModuleData loads module data(items, skills, etc.)
// for specified module.
func LoadModuleData(mod *module.Module) error {
	// Effects.
	effectsData, err := ImportEffectsDir(mod.Conf().EffectsPath())
	if err != nil {
		return fmt.Errorf("fail to load effects: %v", err)
	}
	for _, ed := range effectsData {
		ed.Name = lang.TextDir(mod.Conf().LangPath(), ed.ID)
	}
	res.SetEffectsData(effectsData)
	// Skills.
	skillsData, err := ImportSkillsDir(mod.Conf().SkillsPath())
	if err != nil {
		return fmt.Errorf("fail to load skills: %v", err)
	}
	for _, sd := range skillsData {
		sd.Name = lang.TextDir(mod.Conf().LangPath(), sd.ID)
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
	for _, od := range objectsData {
		od.BasicData.Name = lang.TextDir(mod.Conf().LangPath(), od.BasicData.ID)
	}
	res.SetObjectsData(objectsData)
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
	for _, npcd := range npcData {
		npcd.BasicData.Name = lang.TextDir(chapter.Conf().LangPath(), npcd.BasicData.ID)
	}
	res.SetCharactersData(npcData)
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
	return nil
}
