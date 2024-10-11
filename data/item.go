/*
 * item.go
 *
 * Copyright 2018-2024 Dariusz Sikora <ds@isangeles.dev>
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
	"io"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

// ImportArmors imports all JSON armors from file with specified path.
func ImportArmors(path string) ([]res.ArmorData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open data file: %v", err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.ArmorsData)
	err = unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal json data: %v", err)
	}
	return data.Armors, nil
}

// ImportArmorsDir imports all armors data from files
func ImportArmorsDir(dirPath string) ([]res.ArmorData, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	armors := make([]res.ArmorData, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.FromSlash(dirPath + "/" + file.Name())
		impArmors, err := ImportArmors(filePath)
		if err != nil {
			log.Err.Printf("data armors import: %s: unable to import base: %v",
				filePath, err)
			continue
		}
		armors = append(armors, impArmors...)
	}
	return armors, nil
}

// ImportWeapons imports all JSON weapons from file with specified
// path.
func ImportWeapons(path string) ([]res.WeaponData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open data file: %v", err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.WeaponsData)
	err = unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal json data: %v", err)
	}
	return data.Weapons, nil
}

// ImportWeaponsDir imports all weapons from files
// in specified directory.
func ImportWeaponsDir(dirPath string) ([]res.WeaponData, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	weapons := make([]res.WeaponData, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.FromSlash(dirPath + "/" + file.Name())
		impWeapons, err := ImportWeapons(filePath)
		if err != nil {
			log.Err.Printf("data weapons import: %s: unable to import base: %v",
				filePath, err)
			continue
		}
		weapons = append(weapons, impWeapons...)
	}
	return weapons, nil
}

// ImportMiscItems imports all JSON miscellaneous items from file
// with specified path.
func ImportMiscItems(path string) ([]res.MiscItemData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open data file: %v", err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.MiscItemsData)
	err = unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal json data: %v", err)
	}
	return data.Miscs, nil
}

// ImportMiscItemsDir imports all miscellaneous items from files
// in specified directory.
func ImportMiscItemsDir(dirPath string) ([]res.MiscItemData, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	miscs := make([]res.MiscItemData, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.FromSlash(dirPath + "/" + file.Name())
		impMiscs, err := ImportMiscItems(filePath)
		if err != nil {
			log.Err.Printf("data misc items import: %s: unable to import base: %v",
				filePath, err)
			continue
		}
		miscs = append(miscs, impMiscs...)
	}
	return miscs, nil
}

// ExportArmors exports armors to the data file under specified path.
func ExportArmors(path string, armors ...res.ArmorData) error {
	data := new(res.ArmorsData)
	for _, a := range armors {
		data.Armors = append(data.Armors, a)
	}
	// Marshal armors data.
	json, err := marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal armors: %v", err)
	}
	// Create armors file.
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create armors file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create armors file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	writer := bufio.NewWriter(file)
	writer.Write(json)
	writer.Flush()
	return nil
}

// ExportWeapons exports weapons to the data file under specified path.
func ExportWeapons(path string, weapons ...res.WeaponData) error {
	data := new(res.WeaponsData)
	for _, w := range weapons {
		data.Weapons = append(data.Weapons, w)
	}
	// Marshal weapons data.
	json, err := marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal weapons: %v", err)
	}
	// Create weapons file.
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create weapons file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create weapons file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	writer := bufio.NewWriter(file)
	writer.Write(json)
	writer.Flush()
	return nil
}

// ExportMiscItems exports misc items to the data file under specified path.
func ExportMiscItems(path string, miscs ...res.MiscItemData) error {
	data := new(res.MiscItemsData)
	for _, m := range miscs {
		data.Miscs = append(data.Miscs, m)
	}
	// Marshal misc items data.
	json, err := marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal misc items: %v", err)
	}
	// Create misc items file.
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create misc items file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create misc items file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	writer := bufio.NewWriter(file)
	writer.Write(json)
	writer.Flush()
	return nil
}
