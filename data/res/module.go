/*
 * module.go
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

package res

// Struct for module data.
type ModuleData struct {
	ID        string              `xml:"id,attr" json:"id"`
	Config    map[string][]string `xml:"config" json:"config"`
	Chapter   ChapterData         `xml:"chapter" json:"chapter"`
	Resources ResourcesData       `xml:"resources" json:"resources"`
}

// Struct for chapter data.
type ChapterData struct {
	ID        string              `xml:"id,attr" json:"id"`
	Config    map[string][]string `xml:"config" json:"config"`
	Areas     []AreaData          `xml:"areas>area" json:"areas"`
	Resources ResourcesData       `xml:"resources" json:"resources"`
}

// Struct for module resouces data.
type ResourcesData struct {
	Characters       []CharacterData       `xml:"characters>character" json:"characters"`
	Objects          []ObjectData          `xml:"objects>object" json:"objects"`
	Effects          []EffectData          `xml:"effects>effect" json:"effects"`
	Skills           []SkillData           `xml:"skills>skill" json:"skills"`
	Armors           []ArmorData           `xml:"armors>armor" json:"armors"`
	Weapons          []WeaponData          `xml:"weapons>weapon" json:"weapons"`
	Miscs            []MiscItemData        `xml:"misc>misc" json:"miscs"`
	Dialogs          []DialogData          `xml:"dialogs>dialog" json:"dialogs"`
	Quests           []QuestData           `xml:"quests>quest" json:"quests"`
	Recipes          []RecipeData          `xml:"recipes>recipe" json:"recipes"`
	Areas            []AreaData            `xml:"areas>area" json:"areas"`
	Races            []RaceData            `xml:"races>race" json:"races"`
	Trainings        []TrainingData        `xml:"trainings>training" json:"trainings"`
	TranslationBases []TranslationBaseData `xml:"translations>base" json:"translation-base"`
}
