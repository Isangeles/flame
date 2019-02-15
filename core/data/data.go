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
	"fmt"

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/module"
)

var (
	weaponsData map[string]*parsexml.WeaponNodeXML
	effectsData map[string]*parsexml.EffectNodeXML
)

// LoadModuleData loads data(items, skills, etc.) from
// specified module.
func LoadModuleData(mod *module.Module) error {
	// Weapons.
	weaponsData = make(map[string]*parsexml.WeaponNodeXML)
	xmlWeapons, err := ImportWeaponsDir(mod.Conf().ItemsPath())
	if err != nil {
		return fmt.Errorf("fail_to_load_weapons:%v", err)
	}
	for _, xmlWeapon := range xmlWeapons {
		weaponsData[xmlWeapon.ID] = xmlWeapon
	}
	// Effects.
	effectsData = make(map[string]*parsexml.EffectNodeXML)
	xmlEffects, err := ImportEffectsDir(mod.Conf().EffectsPath())
	if err != nil {
		return fmt.Errorf("fail_to_load_effects:%v", err)
	}
	for _, xmlEffect := range xmlEffects {
		effectsData[xmlEffect.ID] = xmlEffect
	}
	return nil
}
