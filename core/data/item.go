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
	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/item"
	"github.com/isangeles/flame/log"
)

const (
	WEAPONS_FILE_EXT = ".weapons"
)

// Item creates new instance of item with specified ID
// for specified module, returns error if item data with such ID
// was not found or module failed to assign serial value for
// item.
func Item(mod *module.Module, id string) (item.Item, error) {
	switch {
	case weaponsData[id] != nil:
		return weapon(mod, id)
	default:
		return nil, fmt.Errorf("item_not_found:%s",
			id)
	}
}

// ImportWeapons imports all XML weapons from file with specified
// path.
func ImportWeapons(basePath string) ([]*parsexml.WeaponNodeXML, error) {
	weapons := make([]*parsexml.WeaponNodeXML, 0)
	doc, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_weapons_base_file:%v",
			err)
	}
	defer doc.Close()
	weaponsXML, err := parsexml.UnmarshalWeaponsBase(doc)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_weapons:%v",
			err)
	}
	for _, w := range weaponsXML {
		weapons = append(weapons, &w)
	}
	return weapons, nil
}

// ImportWeaponsDir imports all weapons from files
// in specified directory.
func ImportWeaponsDir(dirPath string) ([]*parsexml.WeaponNodeXML, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	weapons := make([]*parsexml.WeaponNodeXML, 0)
	for _, fInfo := range files {
		if !strings.HasSuffix(fInfo.Name(), WEAPONS_FILE_EXT) {
			continue
		}
		baseFilePath := filepath.FromSlash(dirPath + "/" + fInfo.Name())
		weaps, err := ImportWeapons(baseFilePath)
		if err != nil {
			log.Err.Printf("data_weapons_import:%s:fail_to_load_weapons_file:%v",
				baseFilePath, err)
			continue
		}
		for _, w := range weaps {
			weapons = append(weapons, w)
		}
	}
	return weapons, nil
}

// weapon creates new instance of weapon with specified ID
// for specified module, returns error if weapon with such ID
// was not found or module failed to assign serial value for
// weapon.
func weapon(mod *module.Module, id string) (*item.Weapon, error) {
	xmlWeapon := weaponsData[id]
	if xmlWeapon == nil {
		return nil, fmt.Errorf("weapon_not_found:%s", id)
	}
	itemsLangPath := filepath.FromSlash(mod.Conf().LangPath() + "/items" +
		text.LANG_FILE_EXT)
	w, err := buildXMLWeapon(xmlWeapon)
	if err != nil {
		return nil, fmt.Errorf("fail_to_build_weapon:%v", err)
	}
	name := text.ReadDisplayText(itemsLangPath, w.ID())
	w.SetName(name[0])
	err = mod.AssignSerial(w)
	if err != nil {
		return nil, fmt.Errorf("fail_to_assign_item_serial:%v",
			err)
	}
	return w, nil
}

// buildXMLWeapon creates new weapon from specified XML data.
func buildXMLWeapon(xmlWeapon *parsexml.WeaponNodeXML) (*item.Weapon, error) {
	reqs := buildXMLReqs(&xmlWeapon.Reqs)
	slots, err := parsexml.UnmarshalItemSlots(xmlWeapon.Slots)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_slot_types:%v",
			err)
	}
	w := item.NewWeapon(xmlWeapon.ID, xmlWeapon.Value, xmlWeapon.Level,
		xmlWeapon.Damage.Min, xmlWeapon.Damage.Max, reqs, slots)
	return w, nil
}
