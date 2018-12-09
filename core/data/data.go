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
`1 * You should have received a copy of the GNU General Public License
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
	"strings"
	
	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/log"
)

const (
	CHAR_FILE_EXT = ".characters"
)

// Scenario parses file on specified path
// to chapter scenario.
func Scenario(scenPath, npcsDirPath string) (*scenario.Scenario, error) {
	docScen, err := os.Open(scenPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_scenario_file:%v", err)
	}
	defer docScen.Close()
	npcsBasePath := filepath.FromSlash(npcsDirPath + "/npc" + CHAR_FILE_EXT)
	docNPCs, err := os.Open(npcsBasePath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_characters_base_file:%", err)
	}
	defer docNPCs.Close()
	
	xmlScen, err := parsexml.UnmarshalScenario(docScen)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_scenario_file:%v", err)
	}
	mainarea := scenario.NewArea(xmlScen.Mainarea.ID)
	for _, xmlAreaChar := range xmlScen.Mainarea.NPCs.Characters {
		char, err := parsexml.UnmarshalCharacter(docNPCs, xmlAreaChar.ID)
		if err != nil {
			log.Err.Printf("fail_to_create_scenario_npc:%s:%v",
				xmlAreaChar.ID, err)
			continue
		}
		// TODO: set NPC position.
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
	char, err := parsexml.UnmarshalCharacter(doc, charID)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_character:%v", err)
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
	return parsexml.UnmarshalCharactersBase(charFile)
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
		if !strings.HasSuffix(fInfo.Name(), CHAR_FILE_EXT) {
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
		char.Name()) + CHAR_FILE_EXT)
	if err != nil {
		return fmt.Errorf("fail_to_create_char_file:%v", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(xml)
	w.Flush()
	return nil
}
