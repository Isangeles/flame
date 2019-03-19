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
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/item"
	"github.com/isangeles/flame/core/module/serial"
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
	var i item.Item
	// Find data resources.
	switch {
	case res.Weapon(id).ID != "":
		weaponData := res.Weapon(id)
		w := item.NewWeapon(*weaponData)
		i = w
	default:
		return nil, fmt.Errorf("item_data_not_found:%s", id)
	}
	// Assign serial.
	serial.AssignSerial(i)
	return i, nil
}

// ImportWeapons imports all XML weapons from file with specified
// path.
func ImportWeapons(basePath string) ([]*res.WeaponData, error) {
	weapons := make([]*res.WeaponData, 0)
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
	for _, weaponXML := range weaponsXML {
		w, err := buildXMLWeaponData(weaponXML)
		if err != nil {
			log.Err.Printf("imp_weapon:build_xml_data_fail:%v", err)
			continue
		}
		weapons = append(weapons, w)
	}
	return weapons, nil
}

// ImportWeaponsDir imports all weapons from files
// in specified directory.
func ImportWeaponsDir(dirPath string) ([]*res.WeaponData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	weapons := make([]*res.WeaponData, 0)
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

// buildXMLWeapon creates new weapon from specified XML data.
func buildXMLWeaponData(xmlWeapon parsexml.WeaponXML) (*res.WeaponData, error) {
	reqs := buildXMLReqs(&xmlWeapon.Reqs)
	slots, err := parsexml.UnmarshalItemSlots(xmlWeapon.Slots)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_slot_types:%v", err)
	}
	slotsID := make([]int, 0)
	for _, s := range slots {
		slotsID = append(slotsID, int(s))
	}
	dmgType, err := parsexml.UnmarshalHitType(xmlWeapon.Damage.Type)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_damage_type:%v", err)
	}
	hitEffects := make([]res.EffectData, 0)
	for _, xmlEffect := range xmlWeapon.Damage.Effects.Nodes {
		eff := res.Effect(xmlEffect.ID)
		if eff == nil {
			log.Err.Printf("data:build_xml_weapon:hit_effect_not_found:%s",
				xmlEffect.ID)
			continue
		}
		hitEffects = append(hitEffects, *eff)
	}
	w := res.WeaponData{
		ID:         xmlWeapon.ID,
		Value:      xmlWeapon.Value,
		Level:      xmlWeapon.Level,
		DMGMin:     xmlWeapon.Damage.Min,
		DMGMax:     xmlWeapon.Damage.Max,
		DMGType:    int(dmgType),
		DMGEffects: hitEffects,
		EQReqs:     reqs,
		Slots:      slotsID,
	}
	return &w, nil
}
