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
	
	"github.com/isangeles/flame/core/data/parse"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/object/character"
)

const (
	CHAR_FILE_PREFIX = ".xml"
)

// Scenario parses file on specified path
// to scenario.
func Scenario(path string) (*scenario.Scenario, error) {
	return parse.UnmarshalScenarioXML(path)
}

// Characters parses file on specified path to
// game characters.
func Characters(path string) (*[]*character.Character, error) {
	return parse.UnmarshalCharactersBaseXML(path)
}

// ExportCharacter saves specified character to
// [Module]/characters directory.
func ExportCharacter(char *character.Character, dirPath string) error {
	out, err := parse.MarshalCharacterXML(char)
	if err != nil {
		return fmt.Errorf("fail_to_export_char:%v", err)
	}

	f, err := os.Create(filepath.FromSlash(dirPath + "/" +
		char.Id()) + CHAR_FILE_PREFIX)
	if err != nil {
		return fmt.Errorf("fail_to_create_char_file:%v", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.Write(out)
	w.Flush()
	return nil
}
