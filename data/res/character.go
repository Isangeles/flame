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
	XMLName    xml.Name        `xml:"characters" json:"-"`
	Characters []CharacterData `xml:"char" json:"characters"`
}

// Struct for character data resource.
type CharacterData struct {
	XMLName    xml.Name             `xml:"char" json:"-"`
	ID         string               `xml:"id,attr" json:"id"`
	Serial     string               `xml:"serial,attr" json:"serial"`
	Name       string               `xml:"name,attr" json:"name"`
	AI         bool                 `xml:"ai,attr" json:"ai"`
	Level      int                  `xml:"level,attr" json:"level"`
	Sex        string               `xml:"gender,attr" json:"sex"`
	Race       string               `xml:"race,attr" json:"race"`
	Attitude   string               `xml:"attitude,attr" json:"attitude"`
	Guild      string               `xml:"guild,attr" json:"guild"`
	Alignment  string               `xml:"alignment,attr" json:"alignment"`
	PosX       float64              `xml:"position-x,attr" json:"pos-x"`
	PosY       float64              `xml:"position-y,attr" json:"pos-y"`
	DefX       float64              `xml:"def-position-x,attr" json:"def-pos-x"`
	DefY       float64              `xml:"def-position-y,attr" json:"def-pos-y"`
	HP         int                  `xml:"hp,attr" json:"hp"`
	Mana       int                  `xml:"mana,attr" json:"mana"`
	Exp        int                  `xml:"exp,attr" json:"exp"`
	Restore    bool                 `xml:"restore,attr" json:"resore"`
	Attributes AttributesData       `xml:"attributes" json:"attributes"`
	Inventory  InventoryData        `xml:"inventory" json:"inventory"`
	Equipment  EquipmentData        `xml:"equipment" json:"equipment"`
	QuestLog   QuestLogData         `xml:"quests" json:"quests"`
	Crafting   CraftingData         `xml:"crafting" json:"crafting"`
	Trainings  TrainingsData        `xml:"trainings" json:"trainings"`
	Flags      []FlagData           `xml:"flags>flag" json:"flags"`
	Effects    []ObjectEffectData   `xml:"effects>effect" json:"effects"`
	Skills     []ObjectSkillData    `xml:"skills>skill" json:"skills"`
	Memory     []AttitudeMemoryData `xml:"memory>target" json:"memory"`
	Dialogs    []ObjectDialogData   `xml:"dialogs>dialog" json:"dialogs"`
}

// Struct for character attributes data.
type AttributesData struct {
	Str int `xml:"strenght.attr" json:"str"`
	Con int `xml:"constitution,attr" json:"con"`
	Dex int `xml:"dexterity,attr" json:"dex"`
	Int int `xml:"inteligence,attr" json:"int"`
	Wis int `xml:"wisdom,attr" json:"wis"`
}

// Struct for character equipement data.
type EquipmentData struct {
	Items []EquipmentItemData `xml:"item" json:"items"`
}

// Struct for equipment item data
// resource.
type EquipmentItemData struct {
	ID     string `xml:"id,attr" json:"id"`
	Serial string `xml:"serial,attr" json:"serial"`
	Slot   string `xml:"slot,attr" json:"slot"`
}

// Struct for attitude memory data.
type AttitudeMemoryData struct {
	ObjectID     string `xml:"id,attr" json:"id"`
	ObjectSerial string `xml:"serial,attr" json:"serial"`
	Attitude     string `xml:"attitude,attr" json:"attitude"`
}

// Struct for race data.
type RaceData struct {
	ID       string `xml:"id,attr" json:"id"`
	Playable bool   `xml:"playable,attr" json:"playable"`
}
