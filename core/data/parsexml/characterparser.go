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

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/log"
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
	Memory    MemoryXML        `xml:"memory"`
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
	Serial  string   `xml:"serial,attr"`
	Slot    string   `xml:"slot"`
}

// Struct for attitude memory XML node.
type MemoryXML struct {
	XMLName xml.Name            `xml:"memory"`
	Nodes   []MemoryAttitudeXML `xml:"attitude"`
}

// Struct for attitude memodry XML node.
type MemoryAttitudeXML struct {
	XMLName  xml.Name `xml:"attitude"`
	ID       string   `xml:"id,attr"`
	Serial   string   `xml:"serial,attr"`
	Attitude string   `xml:"attitude,attr"`
}

// UnmarshalCharactersBase retrieve all characters data
// from specified XML data.
func UnmarshalCharactersBase(data io.Reader) ([]*res.CharacterData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(CharactersBaseXML)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v", err)
	}
	chars := make([]*res.CharacterData, 0)
	for _, xmlChar := range xmlBase.Characters {
		char, err := buildCharacterData(&xmlChar)
		if err != nil {
			log.Err.Printf("xml:unmarshal_character:build_data_fail:%v", err)
			continue
		}
		chars = append(chars, char)
	}
	return chars, nil
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
// base string.
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
	xmlChar.Skills = *xmlObjectSkills(char.Skills()...)
	xmlChar.Memory = *xmlMemory(char.Memory())
	return xmlChar
}

// xmlEquipment parses specified character equipment to
// XML equipment node.
func xmlEquipment(eq *character.Equipment) *EquipmentXML {
	xmlEq := new(EquipmentXML)
	handRightItem := eq.HandRight().Item()
	if handRightItem != nil {
		xmlEqItem := EquipmentItemXML{
			ID:     handRightItem.ID(),
			Serial: handRightItem.Serial(),
			Slot:   MarshalEqSlot(eq.HandRight()),
		}
		xmlEq.Items = append(xmlEq.Items, xmlEqItem)
	}
	// TODO: parse all equipment slots.
	return xmlEq
}

// xmlMemory parses specified character attitude memodry to
// XML memory node.
func xmlMemory(mem []*character.AttitudeMemory) *MemoryXML {
	xmlMem := new(MemoryXML)
	for _, am := range mem {
		attAttr := marshalAttitude(am.Attitude)
		xmlAtt := MemoryAttitudeXML{
			ID:       am.Target.ID(),
			Serial:   am.Target.Serial(),
			Attitude: attAttr,
		}
		xmlMem.Nodes = append(xmlMem.Nodes, xmlAtt)
	}
	return xmlMem
}

// buildCharacterData creates character resources from specified
// XML data.
func buildCharacterData(xmlChar *CharacterXML) (*res.CharacterData, error) {
	// Basic data.
	baseData := res.CharacterBasicData{
		ID:     xmlChar.ID,
		Serial: xmlChar.Serial,
		Name:   xmlChar.Name,
		Level:  xmlChar.Level,
		Guild:  xmlChar.Guild,
	}
	data := res.CharacterData{BasicData: baseData}
	sex, err := UnmarshalGender(xmlChar.Gender)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_gender:%v", err)
	}
	data.BasicData.Sex = int(sex)
	race, err := UnmarshalRace(xmlChar.Race)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_race:%v", err)
	}
	data.BasicData.Race = int(race)
	attitude, err := UnmarshalAttitude(xmlChar.Attitude)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_attitude:%v", err)
	}
	data.BasicData.Attitude = int(attitude)
	alignment, err := UnmarshalAlignment(xmlChar.Alignment)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_alignment:%v", err)
	}
	data.BasicData.Alignment = int(alignment)
	attributes, err := UnmarshalAttributes(xmlChar.Stats)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_attributes:%v", err)
	}
	// Attributes.
	data.BasicData.Str = attributes.Str
	data.BasicData.Con = attributes.Con
	data.BasicData.Dex = attributes.Dex
	data.BasicData.Int = attributes.Int
	data.BasicData.Wis = attributes.Wis
	// Save.
	data.SavedData.PC = xmlChar.PC
	// HP, mana, exp.
	data.SavedData.HP = xmlChar.HP
	data.SavedData.Mana = xmlChar.Mana
	data.SavedData.Exp = xmlChar.Exp
	// Position.
	if xmlChar.Position != "" {
		posX, posY, err := UnmarshalPosition(xmlChar.Position)
		if err != nil {
			return nil, fmt.Errorf("fail_to_parse_position:%v", err)
		}
		data.SavedData.PosX, data.SavedData.PosY = posX, posY
	}
	// Items.
	for _, xmlIt := range xmlChar.Inventory.Items {
		itData := res.InventoryItemData{
			ID:     xmlIt.ID,
			Serial: xmlIt.Serial,
		}
		data.Items = append(data.Items, itData)
	}
	// Equipment.
	for _, xmlEqIt := range xmlChar.Equipment.Items {
		slot, err := UnmarshalEqSlot(xmlEqIt.Slot)
		if err != nil {
			log.Err.Printf("xml:build_character:%s:parse_eq_item:%s:fail_to_parse_slot:%v",
				xmlChar.ID, xmlEqIt.ID, err)
			continue
		}
		eqItData := res.EquipmentItemData{
			ID:     xmlEqIt.ID,
			Serial: xmlEqIt.Serial,
			Slot:   int(slot),
		}
		data.EqItems = append(data.EqItems, eqItData)
	}
	// Effects.
	for _, xmlEffect := range xmlChar.Effects.Nodes {
		effectData := res.ObjectEffectData{
			ID:           xmlEffect.ID,
			Serial:       xmlEffect.Serial,
			Time:         xmlEffect.Time,
			SourceID:     xmlEffect.Source.ID,
			SourceSerial: xmlEffect.Source.Serial,
		}
		data.Effects = append(data.Effects, effectData)
	}
	// Skills.
	for _, xmlSkill := range xmlChar.Skills.Nodes {
		skillData := res.ObjectSkillData{
			ID:     xmlSkill.ID,
			Serial: xmlSkill.Serial,
		}
		data.Skills = append(data.Skills, skillData)
	}
	// Memory.
	for _, xmlAtt := range xmlChar.Memory.Nodes {
		att, err := UnmarshalAttitude(xmlAtt.Attitude)
		if err != nil {
			log.Err.Printf("xml:build_character:%s:fail_to_parse_att_mem:%s",
				xmlChar.ID, err)
			continue
		}
		attData := res.AttitudeMemoryData{
			ObjectID:     xmlAtt.ID,
			ObjectSerial: xmlAtt.Serial,
			Attitude:     int(att),
		}
		data.Memory = append(data.Memory, attData)
	}
	return &data, nil
}
