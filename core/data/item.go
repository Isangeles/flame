/*
 * item.go
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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/item"
	"github.com/isangeles/flame/core/module/serial"
	"github.com/isangeles/flame/log"
)

const (
	ArmorsFileExt    = ".armors"
	WeaponsFileExt   = ".weapons"
	MiscItemsFileExt = ".misc"
)

// Item creates new instance of item with specified ID
// for specified module, returns error if item data with such ID
// was not found or module failed to assign serial value for
// item.
func Item(id string) (item.Item, error) {
	var i item.Item
	// Find data resources.
	switch {
	case res.Armor(id) != nil:
		armorData := res.Armor(id)
		a := item.NewArmor(*armorData)
		i = a
	case res.Weapon(id) != nil:
		weaponData := res.Weapon(id)
		w := item.NewWeapon(*weaponData)
		i = w
	case res.MiscItem(id) != nil:
		miscData := res.MiscItem(id)
		m := item.NewMisc(*miscData)
		i = m
	default:
		return nil, fmt.Errorf("item data not found: %s", id)
	}
	// Assign serial.
	serial.AssignSerial(i)
	return i, nil
}

// ImportArmors imports all XML armors from file with specified path.
func ImportArmors(basePath string) ([]*res.ArmorData, error) {
	file, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open base file: %v", err)
	}
	defer file.Close()
	armors, err := parsexml.UnmarshalArmors(file)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal xml data: %v", err)
	}
	return armors, nil
}

// ImportArmorsDir imports all armors data from files 
func ImportArmorsDir(dirPath string) ([]*res.ArmorData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	armors := make([]*res.ArmorData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), ArmorsFileExt) {
			continue
		}
		baseFilePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		impArmors, err := ImportArmors(baseFilePath)
		if err != nil {
			log.Err.Printf("data armors import: %s: fail to import base: %v",
				baseFilePath, err)
			continue
		}
		armors = append(armors, impArmors...)
	}
	return armors, nil
}

// ImportWeapons imports all XML weapons from file with specified
// path.
func ImportWeapons(basePath string) ([]*res.WeaponData, error) {
	file, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open base file: %v", err)
	}
	defer file.Close()
	weapons, err := parsexml.UnmarshalWeapons(file)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal xml data: %v", err)
	}
	return weapons, nil
}

// ImportWeaponsDir imports all weapons from files
// in specified directory.
func ImportWeaponsDir(dirPath string) ([]*res.WeaponData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	weapons := make([]*res.WeaponData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), WeaponsFileExt) {
			continue
		}
		baseFilePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		impWeapons, err := ImportWeapons(baseFilePath)
		if err != nil {
			log.Err.Printf("data weapons import: %s: fail to import base: %v",
				baseFilePath, err)
			continue
		}
		weapons = append(weapons, impWeapons...)
	}
	return weapons, nil
}

// ImportMiscItems imports all XML miscellaneous items from file
// with specified path.
func ImportMiscItems(basePath string) ([]*res.MiscItemData, error) {
	file, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open base file: %v", err)
	}
	defer file.Close()
	miscs, err := parsexml.UnmarshalMiscItems(file)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal xml data: %v", err)
	}
	return miscs, nil
}

// ImportMiscItemsDir imports all miscellaneous items from files
// in specified directory.
func ImportMiscItemsDir(dirPath string) ([]*res.MiscItemData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	miscs := make([]*res.MiscItemData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), MiscItemsFileExt) {
			continue
		}
		baseFilePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		impMiscs, err := ImportMiscItems(baseFilePath)
		if err != nil {
			log.Err.Printf("data misc items import: %s: fail to import base: %v",
				baseFilePath, err)
			continue
		}
		miscs = append(miscs, impMiscs...)
	}
	return miscs, nil
}
