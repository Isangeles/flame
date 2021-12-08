/*
 * modifier.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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

package effect

import (
	"github.com/isangeles/flame/data/res"
)

// Interface for object modifiers.
type Modifier interface{}

// NewModifiers creatas modifiers for specified data.
func NewModifiers(data res.ModifiersData) (mods []Modifier) {
	for _, md := range data.HealthMods {
		hpMod := NewHealthMod(md)
		mods = append(mods, hpMod)
	}
	for _, md := range data.ManaMods {
		manaMod := NewManaMod(md)
		mods = append(mods, manaMod)
	}
	for _, md := range data.FlagMods {
		flagMod := NewFlagMod(md)
		mods = append(mods, flagMod)
	}
	for _, md := range data.QuestMods {
		questMod := NewQuestMod(md)
		mods = append(mods, questMod)
	}
	for _, md := range data.AreaMods {
		areaMod := NewAreaMod(md)
		mods = append(mods, areaMod)
	}
	for _, md := range data.AddItemMods {
		addItemMod := NewAddItemMod(md)
		mods = append(mods, addItemMod)
	}
	for _, md := range data.RemoveItemMods {
		removeItemMod := NewRemoveItemMod(md)
		mods = append(mods, removeItemMod)
	}
	for _, md := range data.AddSkillMods {
		addSkillMod := NewAddSkillMod(md)
		mods = append(mods, addSkillMod)
	}
	for _, md := range data.AttributeMods {
		attributeMod := NewAttributeMod(md)
		mods = append(mods, attributeMod)
	}
	for _, md := range data.MemoryMods {
		memoryMod := NewMemoryMod(md)
		mods = append(mods, memoryMod)
	}
	return
}

// ModifiersData creates data resource for modifiers.
func ModifiersData(mods ...Modifier) (data res.ModifiersData) {
	for _, m := range mods {
		switch m := m.(type) {
		case *HealthMod:
			data.HealthMods = append(data.HealthMods, m.Data())
		case *ManaMod:
			data.ManaMods = append(data.ManaMods, m.Data())
		case *FlagMod:
			data.FlagMods = append(data.FlagMods, m.Data())
		case *QuestMod:
			data.QuestMods = append(data.QuestMods, m.Data())
		case *AreaMod:
			data.AreaMods = append(data.AreaMods, m.Data())
		case *AddItemMod:
			data.AddItemMods = append(data.AddItemMods, m.Data())
		case *AddSkillMod:
			data.AddSkillMods = append(data.AddSkillMods, m.Data())
		case *RemoveItemMod:
			data.RemoveItemMods = append(data.RemoveItemMods, m.Data())
		case *AttributeMod:
			data.AttributeMods = append(data.AttributeMods, m.Data())
		case *MemoryMod:
			data.MemoryMods = append(data.MemoryMods, m.Data())
		}
	}
	return
}
