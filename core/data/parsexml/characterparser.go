/*
 * characterparser.go
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

package parsexml

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/skill"
)

// Struct for XML characters base.
type CharactersBaseXML struct {
	XMLName    xml.Name       `xml:"base"`
	Characters []CharacterXML `xml:"char"`
}

// Struct for XML character node.
type CharacterXML struct {
	XMLName   xml.Name         `xml:"char"`
	ID        string           `xml:"id,attr"`
	Serial    string           `xml:"serial,attr"`
	Name      string           `xml:"name,attr"`
	Gender    string           `xml:"gender,attr"`
	Race      string           `xml:"race,attr"`
	Attitude  string           `xml:"attitude,attr"`
	Alignment string           `xml:"alignment,attr"`
	Guild     string           `xml:"guild,attr"`
	Level     int              `xml:"level,attr"`
	Stats     string           `xml:"stats,value"`
	PC        bool             `xml:"pc,attr"`
	HP        int              `xml:"hp, attr"`
	Mana      int              `xml:"mana,attr"`
	Exp       int              `xml:"exp,attr"`
	Position  string           `xml:"position,value"`
	Inventory InventoryXML     `xml:"inventory"`
	Equipment EquipmentXML     `xml:"equipment"`
	Effects   ObjectEffectsXML `xml:"effects"`
	Skills    ObjectSkillsXML  `xml:"skills"`
}

// Struct for equipment XML node.
type EquipmentXML struct {
	XMLName xml.Name           `xml:"equipment"`
	Items   []EquipmentItemXML `xml:"item"`
}

// Struct for equipment item XML node.
type EquipmentItemXML struct {
	XMLName xml.Name `xml:"item"`
	ID      string   `xml:"id,attr"`
	Slot    string   `xml:"slot"`
}

// UnmarshalCharactersBaseXML parses to XML character
// nodes.
func UnmarshalCharactersBase(data io.Reader) ([]CharacterXML, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(CharactersBaseXML)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v",
			err)
	}
	return xmlBase.Characters, nil
}

// UnmarshalCharacter parses character with specified ID from
// XML data.
func UnmarshalCharacter(data io.Reader, charID string) (CharacterXML, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlCharsBase := new(CharactersBaseXML)
	err := xml.Unmarshal(doc, xmlCharsBase)
	if err != nil {
		return CharacterXML{}, fmt.Errorf("fail_to_unmarshal_xml_data:%v",
			err)
	}
	for _, charXML := range xmlCharsBase.Characters {
		if charXML.ID == charID {
			return charXML, nil
		}
	}
	return CharacterXML{}, fmt.Errorf("char_not_found_in_xml_data:%s", charID)
}

// MarshalCharacter parses game character to XML characters
// base.
func MarshalCharacter(char *character.Character) (string, error) {
	xmlCharBase := new(CharactersBaseXML)
	xmlChar := xmlCharacter(char)
	xmlCharBase.Characters = append(xmlCharBase.Characters, *xmlChar)
	out, err := xml.Marshal(xmlCharBase)
	if err != nil {
		return "", fmt.Errorf("fail_to_marshal_char:%v", err)
	}
	return string(out[:]), nil
}

// xmlCharacter parses specified game character to
// XML character struct.
func xmlCharacter(char *character.Character) *CharacterXML {
	xmlChar := new(CharacterXML)
	xmlChar.ID = char.ID()
	xmlChar.Serial = char.Serial()
	xmlChar.Name = char.Name()
	xmlChar.Level = char.Level()
	xmlChar.Gender = marshalGender(char.Gender())
	xmlChar.Race = marshalRace(char.Race())
	xmlChar.Attitude = marshalAttitude(char.Attitude())
	xmlChar.Alignment = marshalAlignment(char.Alignment())
	xmlChar.Stats = marshalAttributes(char.Attributes())
	xmlChar.HP = char.Health()
	xmlChar.Mana = char.Mana()
	xmlChar.Exp = char.Experience()
	posX, posY := char.Position()
	xmlChar.Position = fmt.Sprintf("%fx%f", posX, posY)
	xmlChar.Inventory = *xmlInventory(char.Inventory())
	xmlChar.Equipment = *xmlEquipment(char.Equipment())
	xmlChar.Effects = *xmlObjectEffects(char.Effects()...)
	charSkills := make([]*skill.Skill, 0)
	for _, s := range char.Skills() {
		charSkills = append(charSkills, s)
	}
	xmlChar.Skills = *xmlObjectSkills(charSkills...)
	return xmlChar
}

// xmlEquipment parses specified character equipment to
// XML equipment node.
func xmlEquipment(eq *character.Equipment) *EquipmentXML {
	xmlEq := new(EquipmentXML)
	handRightItem := eq.HandRight().Item()
	if handRightItem != nil {
		xmlEqItem := EquipmentItemXML{
			ID:   handRightItem.ID() + "_" + handRightItem.Serial(),
			Slot: MarshalEqSlot(eq.HandRight()),
		}
		xmlEq.Items = append(xmlEq.Items, xmlEqItem)
	}
	// TODO: parse all equipment slots.
	return xmlEq
}
