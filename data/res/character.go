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

import (
	"encoding/xml"
)

// Struct for characters data resource.
type CharactersData struct {
	XMLName    xml.Name        `xml:"characters"`
	Characters []CharacterData `xml:"char"`
}

// Struct for character data resource.
type CharacterData struct {
	XMLName    xml.Name             `xml:"char"`
	ID         string               `xml:"id,attr"`
	Serial     string               `xml:"serial,attr"`
	Name       string               `xml:"name,attr"`
	AI         bool                 `xml:"ai,attr"`
	Level      int                  `xml:"level,attr"`
	Sex        string               `xml:"gender,attr"`
	Race       string               `xml:"race,attr"`
	Attitude   string               `xml:"attitude,attr"`
	Guild      string               `xml:"guild,attr"`
	Alignment  string               `xml:"alignment,attr"`
	PosX       float64              `xml:"position-x,attr"`
	PosY       float64              `xml:"position-y,attr"`
	DefX       float64              `xml:"def-position-x,attr"`
	DefY       float64              `xml:"def-position-y,attr"`
	HP         int                  `xml:"hp,attr"`
	Mana       int                  `xml:"mana,attr"`
	Exp        int                  `xml:"exp,attr"`
	Attributes AttributesData       `xml:"attributes"`
	Inventory  InventoryData        `xml:"inventory"`
	Equipment  EquipmentData        `xml:"equipment"`
	QuestLog   QuestLogData         `xml:"quests"`
	Crafting   CraftingData         `xml:"crafting"`
	Trainings  TrainingsData        `xml:"trainings"`
	Flags      []FlagData           `xml:"flags>flag"`
	Effects    []ObjectEffectData   `xml:"effects>effect"`
	Skills     []ObjectSkillData    `xml:"skills>skill"`
	Memory     []AttitudeMemoryData `xml:"memory>target"`
	Dialogs    []ObjectDialogData   `xml:"dialogs>dialog"`
}

// Struct for character attributes data.
type AttributesData struct {
	Str int `xml:"strenght.attr"`
	Con int `xml:"constitution,attr"`
	Dex int `xml:"dexterity,attr"`
	Int int `xml:"inteligence,attr"`
	Wis int `xml:"wisdom,attr"`
}

// Struct for character equipement data.
type EquipmentData struct {
	Items []EquipmentItemData `xml:"item"`
}

// Struct for equipment item data
// resource.
type EquipmentItemData struct {
	ID     string `xml:"id,attr"`
	Serial string `xml:"serial,attr"`
	Slot   string `xml:"slot,attr"`
}

// Struct for attitude memory data.
type AttitudeMemoryData struct {
	ObjectID     string `xml:"id,attr"`
	ObjectSerial string `xml:"serial,attr"`
	Attitude     string `xml:"attitude,attr"`
}

// Struct for race data.
type RaceData struct {
	ID       string `xml:"id,attr"`
	Playable bool   `xml:"playable,attr"`
}
