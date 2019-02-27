/*
 * char.go
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

package data

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/item"
	"github.com/isangeles/flame/log"
)

const (
	CHARS_FILE_EXT = ".characters"
)

// Character parses specified characters base and creates
// game character.
func Character(mod *module.Module, charID string) (*character.Character, error) {
	dataChar := res.Character(charID)
	if dataChar.ID == "" {
		return nil, fmt.Errorf("character_data_not_found:%%s", charID)
	}
	char := character.New(dataChar)
	// Inventory.
	for _, invItData := range dataChar.Items {
		it, err := Item(mod, invItData.ID)
		if it == nil {
			log.Err.Printf("data:character:%s:fail_to_retrieve_inv_item:%v",
				char.ID(), err)
			continue
		}
		it.SetSerial(invItData.Serial)
		char.Inventory().AddItem(it)
	}
	// Equipment.
	for _, eqItData := range dataChar.EqItems {
		it := char.Inventory().Item(eqItData.ID)
		if it == nil {
			log.Err.Printf("data:character:%s:eq:fail_to_retrieve_eq_item_from_inv:%s",
				char.ID(), eqItData.ID)
			continue
		}
		eqItem, ok := it.(item.Equiper)
		if !ok {
			log.Err.Printf("data:character:%s:eq:not_eqipable_item:%s",
				char.ID(), it.ID())
			continue
		}
		switch character.EquipmentSlotType(eqItData.Slot) {
		case character.Hand_right:
			err := char.Equipment().EquipHandRight(eqItem)
			if err != nil {
				log.Err.Printf("data_build_character:%s:eq:fail_to_equip_item:%v",
					char.ID(), err)
			}
		default:
			log.Err.Printf("data:character:%s:unknown_equipment_slot:%s",
				char.ID(), eqItData.Slot)
		}
	}
	return char, nil
}

// ImportCharactersData import characters data from base file
// with specified path.
func ImportCharactersData(basePath string) ([]res.CharacterData, error) {
	baseFile, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_char_base_file:%v", err)
	}
	defer baseFile.Close()
	xmlChars, err := parsexml.UnmarshalCharactersBase(baseFile)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_chars_base:%v", err)
	}
	chars := make([]res.CharacterData, 0)
	for _, xmlChar := range xmlChars {
		char, err := buildXMLCharacterData(&xmlChar)
		if err != nil {
			log.Err.Printf("data:import_char:%s:fail_to_build_xml_data:%v",
				xmlChar.ID, err)
			continue
		}
		chars = append(chars, char)
	}
	return chars, nil
}

// ImportCharactersDir imports all characters files from directory
// with specified path.
func ImportCharactersDataDir(dirPath string) ([]res.CharacterData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	chars := make([]res.CharacterData, 0)
	for _, fInfo := range files {
		if !strings.HasSuffix(fInfo.Name(), CHARS_FILE_EXT) {
			continue
		}
		charFilePath := filepath.FromSlash(dirPath + "/" + fInfo.Name())
		impChars, err := ImportCharactersData(charFilePath)
		if err != nil {
			log.Err.Printf("data_char_import:%s:fail_to_parse_char_file:%v",
				charFilePath, err)
			continue
		}
		for _, c := range impChars {
			chars = append(chars, c)
		}
	}
	return chars, nil
}

// ImportCharacters imports characters from base file with
// specified path.
func ImportCharacters(mod *module.Module, path string) ([]*character.Character, error) {
	charFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_char_base_file:%v", err)
	}
	defer charFile.Close()
	charsXML, err := parsexml.UnmarshalCharactersBase(charFile)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_chars_base:%v", err)
	}
	chars := make([]*character.Character, 0)
	for _, charXML := range charsXML {
		char, err := buildXMLCharacter(mod, &charXML)
		if err != nil {
			log.Err.Printf("data_import_chars:%s:fail_to_build_char:%s:%v",
				path, charXML.ID, err)
			continue
		}
		chars = append(chars, char)
	}
	return chars, nil
}

// ImportCharactersDir imports all characters files from directory
// with specified path.
func ImportCharactersDir(mod *module.Module, dirPath string) ([]*character.Character, error) {
	chars := make([]*character.Character, 0)
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return chars, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	for _, fInfo := range files {
		if !strings.HasSuffix(fInfo.Name(), CHARS_FILE_EXT) {
			continue
		}
		charFilePath := filepath.FromSlash(dirPath + "/" + fInfo.Name())
		impChars, err := ImportCharacters(mod, charFilePath)
		if err != nil {
			log.Err.Printf("data_char_import:%s:fail_to_parse_char_file:%v",
				charFilePath, err)
			continue
		}
		for _, c := range impChars {
			chars = append(chars, c)
		}
	}
	return chars, nil
}

// ExportCharacter saves specified character to
// [Module]/characters directory.
func ExportCharacter(char *character.Character, dirPath string) error {
	// Parse character data.
	xml, err := parsexml.MarshalCharacter(char)
	if err != nil {
		return fmt.Errorf("fail_to_export_char:%v", err)
	}
	// Create character file.
	f, err := os.Create(filepath.FromSlash(dirPath+"/"+
		strings.ToLower(char.Name())) + CHARS_FILE_EXT)
	if err != nil {
		return fmt.Errorf("fail_to_create_char_file:%v", err)
	}
	defer f.Close()
	// Write data to file.
	w := bufio.NewWriter(f)
	w.WriteString(xml)
	w.Flush()
	return nil
}

// buildXMLCharacter creates new game character from XML
// character data.
func buildXMLCharacter(mod *module.Module,
	charXML *parsexml.CharacterXML) (*character.Character, error) {
	id := charXML.ID
	name := charXML.Name
	level := charXML.Level
	sex, err := parsexml.UnmarshalGender(charXML.Gender)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_char_gender:%v",
			err)
	}
	race, err := parsexml.UnmarshalRace(charXML.Race)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_char_race:%v",
			err)
	}
	attitude, err := parsexml.UnmarshalAttitude(charXML.Attitude)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_char_attitude:%v",
			err)
	}
	guild := character.NewGuild(charXML.Guild) // TODO: search and assign guild
	attributes, err := parsexml.UnmarshalAttributes(charXML.Stats)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_char_attributes:%v",
			err)
	}
	alignment, err := parsexml.UnmarshalAlignment(charXML.Alignment)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_char_alignment:%v",
			err)
	}
	char := character.NewCharacter(id, name, level, sex, race,
		attitude, guild, attributes, alignment)
	// Inventory.
	for _, xmlInvItem := range charXML.Inventory.Items {
		it, err := Item(mod, xmlInvItem.ID)
		if err != nil {
			log.Err.Printf("data_build_character:%s:inv_item:%v",
				char.ID(), err)
			continue
		}
		it.SetSerial(xmlInvItem.Serial)
		err = char.Inventory().AddItem(it)
		if err != nil {
			log.Err.Printf("data_build_character:%s::add_item:%v",
				char.ID(), err)
		}
	}
	// Equipment.
	for _, xmlEqItem := range charXML.Equipment.Items {
		it := char.Inventory().Item(xmlEqItem.ID)
		if it == nil {
			log.Err.Printf("data_build_character:%s:eq:fail_to_retrieve_eq_item_from_inv:%v",
				char.SerialID(), err)
			continue
		}
		eqItem, ok := it.(item.Equiper)
		if !ok {
			log.Err.Printf("data_build_character:%s:eq:not_eqipable_item:%s",
				char.ID(), it.ID())
			continue
		}
		switch xmlEqItem.Slot {
		case parsexml.MarshalEqSlot(char.Equipment().HandRight()):
			err := char.Equipment().EquipHandRight(eqItem)
			if err != nil {
				log.Err.Printf("data_build_character:%s:eq:fail_to_equip_item:%v",
					char.ID(), err)
			}
		default:
			log.Err.Printf("data_build_character:%s:unknown_equipment_slot:%s",
				char.ID(), xmlEqItem.Slot)
		}
	}
	return char, nil
}

// buildXMLCharacterData creates character resources from specified
// XML data.
func buildXMLCharacterData(xmlChar *parsexml.CharacterXML) (res.CharacterData, error) {
	data := res.CharacterData{
		ID: xmlChar.ID,
		Name: xmlChar.Name,
		Level: xmlChar.Level,
		Guild: xmlChar.Guild,
	}
	sex, err := parsexml.UnmarshalGender(xmlChar.Gender)
	if err != nil {
		return data, fmt.Errorf("fail_to_parse_gender:%v", err)
	}
	data.Sex = int(sex)
	race, err := parsexml.UnmarshalRace(xmlChar.Race)
	if err != nil {
		return data, fmt.Errorf("fail_to_parse_race:%v", err)
	}
	data.Race = int(race)
	attitude, err := parsexml.UnmarshalAttitude(xmlChar.Attitude)
	if err != nil {
		return data, fmt.Errorf("fail_to_parse_attitude:%v", err)
	}
	data.Attitude = int(attitude)
	alignment, err := parsexml.UnmarshalAlignment(xmlChar.Alignment)
	if err != nil {
		return data, fmt.Errorf("fail_to_parse_alignment:%v", err)
	}
	data.Alignment = int(alignment)
	attributes, err := parsexml.UnmarshalAttributes(xmlChar.Stats)
	if err != nil {
		return data, fmt.Errorf("fail_to_parse_attributes:%v", err)
	}
	data.Str = attributes.Str
	data.Con = attributes.Con
	data.Dex = attributes.Dex
	data.Int = attributes.Int
	data.Wis = attributes.Wis
	for _, xmlInvIt := range xmlChar.Inventory.Items {
		invItData := res.InventoryItemData{
			ID: xmlInvIt.ID,
			Serial: xmlInvIt.Serial,
		}
		data.Items = append(data.Items, invItData)
	}
	for _, xmlEqIt := range xmlChar.Equipment.Items {
		slot, err := parsexml.UnmarshalEqSlot(xmlEqIt.Slot)
		if err != nil {
			log.Err.Printf("data:build_xml_character:%s:parse_eq_item:%s:fail_to_parse_slot:%v",
				xmlChar.ID, xmlEqIt.ID, err)
			continue
		}
		eqItData := res.EquipmentItemData{
			ID: xmlEqIt.ID,
		        Slot: int(slot),
		}
		data.EqItems = append(data.EqItems, eqItData)
	}
	return data, nil
}
