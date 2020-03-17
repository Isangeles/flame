/*
 * race.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/data/parsexml"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

const (
	RacesFileExt = ".races"
)

// ImportRaces imports all reces from file with specified path.
func ImportRaces(path string) ([]*res.RaceData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open races data file: %v", err)
	}
	defer file.Close()
	races, err := parsexml.UnmarshalRaces(file)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal races data file: %v", err)
	}
	return races, nil
}

// ImportRacesDir imports all races from data files from
// directory with specified path.
func ImportRacesDir(path string) ([]*res.RaceData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	races := make([]*res.RaceData, 0)
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), RacesFileExt) {
			continue
		}
		filePath := filepath.Join(path, f.Name())
		impRaces, err := ImportRaces(filePath)
		if err != nil {
			log.Err.Printf("data: import races dir: %s: unable to import file: %v",
				filePath, err)
			continue
		}
		races = append(races, impRaces...)
	}
	return races, nil
}
