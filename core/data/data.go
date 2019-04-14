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
		return fmt.Errorf("fail_to_load_effects:%v", err)
	}
	// Translate.
	for _, ed := range effectsData {
		ed.Name = lang.TextDir(mod.Conf().LangPath(), ed.ID)
	}
	res.SetEffectsData(effectsData)
	// Skills.
	skillsData, err := ImportSkillsDir(mod.Conf().SkillsPath())
	if err != nil {
		return fmt.Errorf("fail_to_load_skills:%v", err)
	}
	// Translate.
	for _, sd := range skillsData {
		sd.Name = lang.TextDir(mod.Conf().LangPath(), sd.ID)
	}
	res.SetSkillsData(skillsData)
	// Weapons.
	weaponsData, err := ImportWeaponsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("fail_to_load_weapons:%v", err)
	}
	// Translate.
	for _, wd := range weaponsData {
		wd.Name = lang.TextDir(mod.Conf().LangPath(), wd.ID)
	}
	res.SetWeaponsData(weaponsData)
	// Area objects.
	objectsData, err := ImportObjectsDir(mod.Conf().ObjectsPath())
	if err != nil {
		return fmt.Errorf("fail_to_load_area_objects:%v", err)
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
		return fmt.Errorf("fail_to_load_npc:%v", err)
	}
	res.SetCharactersData(npcData)
	return nil
}
