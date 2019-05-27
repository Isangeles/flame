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

// Character creates module character with specified ID.
func Character(mod *module.Module, charID string) (*character.Character, error) {
	// Get data.
	data := res.Character(charID)
	if data == nil {
		return nil, fmt.Errorf("character_data_not_found:%s", charID)
	}
	// Full build character(with skills, itmes, etc.).
	char := buildCharacter(mod, data)
	// Add skills & items from mod config.
	for _, sid := range mod.Conf().CharSkills {
		s, err := Skill(sid)
		if err != nil {
			log.Err.Printf("fail_to_retireve_conf_char_skill:%v", err)
			continue
		}
		char.AddSkill(s)
	}
	for _, iid := range mod.Conf().CharItems {
		i, err := Item(iid)
		if err != nil {
			log.Err.Printf("fail_to_retireve_conf_char_item:%v", err)
			continue
		}
		char.Inventory().AddItem(i)
	}
	return char, nil
}

// ImportCharactersData import characters data from base file
// with specified path.
func ImportCharactersData(path string) ([]*res.CharacterData, error) {
	baseFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_char_base_file:%v", err)
	}
	defer baseFile.Close()
	chars, err := parsexml.UnmarshalCharactersBase(baseFile)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_chars_base:%v", err)
	}
	return chars, nil
}

// ImportCharactersDataDir imports all characters data from
// files in directory with specified path.
func ImportCharactersDataDir(path string) ([]*res.CharacterData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	chars := make([]*res.CharacterData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), CHARS_FILE_EXT) {
			continue
		}
		basePath := filepath.FromSlash(path + "/" + finfo.Name())
		impChars, err := ImportCharactersData(basePath)
		if err != nil {
			log.Err.Printf("data:import_chars_dir:%s:fail_to_parse_char_file:%v",
				basePath, err)
			continue
		}
		chars = append(chars, impChars...)
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
	charsData, err := parsexml.UnmarshalCharactersBase(charFile)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_chars_base:%v", err)
	}
	chars := make([]*character.Character, 0)
	for _, charData := range charsData {
		char := buildCharacter(mod, charData)
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
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_chrs_dir:%v", err)
	}
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

// buildCharacter builds new character from specified data(with items and equipment).
func buildCharacter(mod *module.Module, data *res.CharacterData) *character.Character {
	char := character.New(data.BasicData)
	// Inventory.
	for _, invItData := range data.Items {
		it, err := Item(invItData.ID)
		if err != nil {
			log.Err.Printf("data:character:%s:fail_to_retrieve_inv_item:%v",
				char.ID(), err)
			continue
		}
		it.SetSerial(invItData.Serial)
		char.Inventory().AddItem(it)
	}
	// Equipment.
	for _, eqItData := range data.EqItems {
		it := char.Inventory().Item(eqItData.ID, eqItData.Serial)
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
	// Skills.
	for _, skillData := range data.Skills {
		skill, err := Skill(skillData.ID)
		if err != nil {
			log.Err.Printf("data:build_character:%s:fail_to_retrieve_skill:%v",
				char.ID(), err)
			continue
		}
		if len(skillData.Serial) > 0 {
			skill.SetSerial(skillData.Serial)
		}
		skill.SetCooldown(skillData.Cooldown)
		char.AddSkill(skill)
	}
	// Dialogs.
	for _, dialogData := range data.Dialogs {
		dialog, err := Dialog(dialogData.ID)
		if err != nil {
			log.Err.Printf("data:build_character:%s:fail_to_retrieve_dialog:%v",
				char.ID(), err)
			continue
		}
		char.AddDialog(dialog)
	}
	// Quests.
	for _, questData := range data.Quests {
		quest, err := Quest(questData.ID)
		if err != nil {
			log.Err.Printf("data:build_character:%s:fail_to_retrieve_quest:%v",
				char.ID(), err)
			continue
		}
		for _, s := range quest.Stages() {
			if s.ID() == questData.Stage {
				quest.SetActiveStage(s)
			}
		}
		if quest.ActiveStage() == nil {
			log.Err.Printf("data:build_character:%s:quest:%s:fail_to_set_active_stage:%v",
				err)
		}
		char.Journal().AddQuest(quest)
	}
	return char
}
