/*
 * race.go
 *
 * Copyright 2020-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"io"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

// ImportRaces imports all reces from file with specified path.
func ImportRaces(path string) ([]res.RaceData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open data file: %v", err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.RacesData)
	err = xml.Unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal XML data: %v", err)
	}
	return data.Races, nil
}

// ImportRacesDir imports all races from data files from
// directory with specified path.
func ImportRacesDir(path string) ([]res.RaceData, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	races := make([]res.RaceData, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.Join(path, file.Name())
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

// ExportRaces exports races to data file under specified path.
func ExportRaces(path string, races ...res.RaceData) error {
	data := new(res.RacesData)
	for _, r := range races {
		data.Races = append(data.Races, r)
	}
	// Marshal races data.
	xml, err := xml.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal races: %v", err)
	}
	// Create races file.
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create races file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create races file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	writer := bufio.NewWriter(file)
	writer.Write(xml)
	writer.Flush()
	return nil
}
