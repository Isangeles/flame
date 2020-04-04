/*
 * character.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for character data resource.
type CharacterData struct {
	ID        string
	Serial    string
	Name      string
	AI        bool
	Level     int
	Sex       string
	Race      string
	Attitude  string
	Guild     string
	Alignment string
	PosX      float64
	PosY      float64
	DefX      float64
	DefY      float64
	HP        int
	Mana      int
	Exp       int
	Attributes AttributesData
	Inventory InventoryData
	Equipment EquipmentData
	QuestLog  QuestLogData
	Crafting  CraftingData
	Trainings TrainingsData
	Flags     []FlagData
	Effects   []ObjectEffectData
	Skills    []ObjectSkillData
	Memory    []AttitudeMemoryData
	Dialogs   []ObjectDialogData
}

// Struct for character attributes data.
type AttributesData struct {
	Str       int
	Con       int
	Dex       int
	Int       int
	Wis       int
}

// Struct for character equipement data.
type EquipmentData struct {
	Items []EquipmentItemData
}

// Struct for equipment item data
// resource.
type EquipmentItemData struct {
	ID     string
	Serial string
	Slot   string
}

// Struct for attitude memory data.
type AttitudeMemoryData struct {
	ObjectID     string
	ObjectSerial string
	Attitude     string
}

// Struct for race data.
type RaceData struct {
	ID       string
	Playable bool
}
