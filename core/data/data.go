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
 * You should have received a copy of the GNU General Public License
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
	CHAR_FILE_PREFIX = ".xml"
)

// Scenario parses file on specified path
// to scenario.
func Scenario(path string) (*scenario.Scenario, error) {
	return parsexml.UnmarshalScenarioXML(path)
}

// Characters parses file on specified path to
// game characters.
func Characters(path string) ([]*character.Character, error) {
	return parsexml.UnmarshalCharactersBaseXML(path)
}

// ImportCharacters parses all characters files from directory
// with specified path.
func ImportCharacters(dirPath string) ([]*character.Character, error) {
	chars := make([]*character.Character, 0)
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return chars, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	for _, fInfo := range files {
		if !strings.HasSuffix(fInfo.Name(), CHAR_FILE_PREFIX) {
			continue
		}
		charFilePath := filepath.FromSlash(dirPath + "/" + fInfo.Name())
		impChars, err := Characters(charFilePath)
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
	xml, err := parsexml.MarshalCharacterXML(char)
	if err != nil {
		return fmt.Errorf("fail_to_export_char:%v", err)
	}

	f, err := os.Create(filepath.FromSlash(dirPath + "/" +
		char.Name()) + CHAR_FILE_PREFIX)
	if err != nil {
		return fmt.Errorf("fail_to_create_char_file:%v", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(xml)
	w.Flush()
	return nil
}
