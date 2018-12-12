/*
 * data.go
 * 
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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
`* You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 * 
 * 
 */

// data package provides connection with external data files like items
// base, savegames, etc.
package data

import (
	"bufio"
	"fmt"
	"path/filepath"
	"os"
	"io/ioutil"
	"strconv"
	"strings"
	
	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/log"
)

const (
	CHARS_FILE_EXT = ".characters"
)

// Scenario parses file on specified path
// to chapter scenario.
func Scenario(scenPath, npcsDirPath string) (*scenario.Scenario, error) {
	docScen, err := os.Open(scenPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_scenario_file:%v", err)
	}
	defer docScen.Close()
	npcsBasePath := filepath.FromSlash(npcsDirPath + "/npc" + CHARS_FILE_EXT)
	docNPCs, err := os.Open(npcsBasePath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_characters_base_file:%v", err)
	}
	defer docNPCs.Close()
	
	xmlScen, err := parsexml.UnmarshalScenario(docScen)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_scenario_file:%v", err)
	}
	mainarea := scenario.NewArea(xmlScen.Mainarea.ID)
	for _, xmlAreaChar := range xmlScen.Mainarea.NPCs.Characters {
		charXML, err := parsexml.UnmarshalCharacter(docNPCs, xmlAreaChar.ID)
		if err != nil {
			log.Err.Printf("data_scenario_unmarshal_npc:%s:fail:%v",
				xmlAreaChar.ID, err)
			continue
		}
		char, err := buildXMLCharacter(&charXML)
		if err != nil {
			log.Err.Printf("data_scenario_build_npc:%s:fail:%v",
				xmlAreaChar.ID, err)
		}
		x, y, err := parsexml.UnmarshalPosition(xmlAreaChar.Position)
		if err != nil {
			log.Err.Printf("data_scenario_spawn_npc:%s:unmarshal_position_fail:%v",
				xmlAreaChar.ID, err)
			continue
		}
		char.SetPosition(x, y)
		mainarea.AddCharacter(char)
	}
	subareas := make([]*scenario.Area, 0)
	for _, xmlArea := range xmlScen.Subareas {
		area := scenario.NewArea(xmlArea.ID)
		// TODO: area NPCs.
		subareas = append(subareas, area)
	}
	scen := scenario.NewScenario(xmlScen.ID, mainarea, subareas)
	return scen, nil
}

// Character parses specified characters base and creates
// game character.
func Character(basePath, charID string) (*character.Character, error) {
	doc, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_characters_base_file:%v", err)
	}
	defer doc.Close()
	charXML, err := parsexml.UnmarshalCharacter(doc, charID)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_character:%v", err)
	}
	char, err := buildXMLCharacter(&charXML)
	if err != nil {
		return nil, fmt.Errorf("fail_to_build_character_from_xml:%v", err)
	}
	return char, nil
}

// ImportCharacters imports char file with specified path.
func ImportCharacters(path string) ([]*character.Character, error) {
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
		char, err := buildXMLCharacter(&charXML)
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
func ImportCharactersDir(dirPath string) ([]*character.Character, error) {
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
		impChars, err := ImportCharacters(charFilePath)
		if err != nil {
			log.Err.Printf("data_char_import:fail_to_parse_char_file:%v",
				err)
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
	xml, err := parsexml.MarshalCharacter(char)
	if err != nil {
		return fmt.Errorf("fail_to_export_char:%v", err)
	}

	f, err := os.Create(filepath.FromSlash(dirPath + "/" +
		char.Name()) + CHARS_FILE_EXT)
	if err != nil {
		return fmt.Errorf("fail_to_create_char_file:%v", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(xml)
	w.Flush()
	return nil
}

// buildXMLCharacter creates new game character from XML
// character data.
func buildXMLCharacter(charXML *parsexml.CharacterXML) (*character.Character, error) {
	id := charXML.ID
	name := charXML.Name
	level, err := strconv.Atoi(charXML.Level)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_char_level:%v",
			err)
	}
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
	return char, nil
}
