/*
 * char.go
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

package data

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/character"
	"github.com/isangeles/flame/log"
)

const (
	CharsFileExt = ".characters"
)

// ImportCharacters import characters data from base file
// with specified path.
func ImportCharacters(path string) ([]res.CharacterData, error) {
	baseFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open char base file: %v", err)
	}
	defer baseFile.Close()
	buf, err := ioutil.ReadAll(baseFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.CharactersData)
	err = xml.Unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal XML: %v", err)
	}
	return data.Characters, nil
}

// ImportCharactersDir imports all characters data from
// files in directory with specified path.
func ImportCharactersDir(path string) ([]res.CharacterData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	chars := make([]res.CharacterData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), CharsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(path + "/" + finfo.Name())
		impChars, err := ImportCharacters(basePath)
		if err != nil {
			log.Err.Printf("data: import chars dir: %s: unable to parse char file: %v",
				basePath, err)
			continue
		}
		chars = append(chars, impChars...)
	}
	return chars, nil
}

// ExportCharacters saves characters to new file with specified path.
func ExportCharacters(path string, chars ...*character.Character) error {
	data := new(res.CharactersData)
	for _, c := range chars {
		data.Characters = append(data.Characters, c.Data())
	}
	// Marshal character data.
	xml, err := xml.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal characters: %v", err)
	}
	// Create character file.
	if !strings.HasSuffix(path, CharsFileExt) {
		path += CharsFileExt
	}
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create characters file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create characters file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	w := bufio.NewWriter(file)
	w.Write(xml)
	w.Flush()
	return nil
}
