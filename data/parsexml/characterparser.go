 /*
 * characterparser.go
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

package parsexml

import (
	"encoding/xml"
)

// Struct for XML characters base.
type Characters struct {
	XMLName    xml.Name    `xml:"characters"`
	Characters []Character `xml:"char"`
}

// Struct for XML character node.
type Character struct {
	XMLName     xml.Name         `xml:"char"`
	ID          string           `xml:"id,attr"`
	Serial      string           `xml:"serial,attr"`
	Name        string           `xml:"name,attr"`
	Gender      string           `xml:"gender,attr"`
	Race        string           `xml:"race,attr"`
	Attitude    string           `xml:"attitude,attr"`
	Alignment   string           `xml:"alignment,attr"`
	Guild       string           `xml:"guild,attr"`
	Level       int              `xml:"level,attr"`
	Attributes  Attributes       `xml:"attributes"`
	HP          int              `xml:"hp,attr"`
	Mana        int              `xml:"mana,attr"`
	Exp         int              `xml:"exp,attr"`
	PosX        float64          `xml:"position-x,attr"`
	PosY        float64          `xml:"position-y,attr"`
	DefPosX     float64          `xml:"def-position-x,attr"`
	DefPosY     float64          `xml:"def-position-y,attr"`
	AI          bool             `xml:"ai,attr"`
	Inventory   Inventory        `xml:"inventory"`
	Equipment   Equipment        `xml:"equipment"`
	Effects     []ObjectEffect   `xml:"effects>effect"`
	Skills      ObjectSkills     `xml:"skills"`
	Memory      Memory           `xml:"memory"`
	Dialogs     ObjectDialogs    `xml:"dialogs"`
	Quests      []CharacterQuest `xml:"quests>quest"`
	Flags       []Flag           `xml:"flags>flag"`
	Crafting    []ObjectRecipe   `xml:"crafting>recipe"`
	Trainings   Trainings        `xml:"trainings"`
}

// Struct for equipment XML node.
type Equipment struct {
	XMLName xml.Name        `xml:"equipment"`
	Items   []EquipmentItem `xml:"item"`
}

// Struct for equipment item XML node.
type EquipmentItem struct {
	XMLName xml.Name `xml:"item"`
	ID      string   `xml:"id,attr"`
	Serial  string   `xml:"serial,attr"`
	Slot    string   `xml:"slot"`
}

// Struct for character memory XML node.
type Memory struct {
	XMLName xml.Name       `xml:"memory"`
	Nodes   []TargetMemory `xml:"target"`
}

// Struct for target memory XML node.
type TargetMemory struct {
	XMLName  xml.Name `xml:"target"`
	ID       string   `xml:"id,attr"`
	Serial   string   `xml:"serial,attr"`
	Attitude string   `xml:"attitude,attr"`
}

// Struct for flag XML node.
type Flag struct {
	XMLName xml.Name `xml:"flag"`
	ID      string   `xml:"id,attr"`
}

// Struct for character quest XML node.
type CharacterQuest struct {
	XMLName xml.Name `xml:"quest"`
	ID      string   `xml:"id,attr"`
	Stage   string   `xml:"stage,attr"`
}

// Struct for character attributes node.
type Attributes struct {
	XMLName      xml.Name `xml:"attributes"`
	Strenght     int      `xml:"stringht,attr"`
	Constitution int      `xml:"constitution,attr"`
	Dexterity    int      `xml:"dexterity,attr"`
	Intelligence int      `xml:"inteligence,attr"`
	Wisdom       int      `xml:"wisdom,attr"`
}
