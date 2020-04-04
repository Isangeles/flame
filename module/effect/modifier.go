/*
 * modifier.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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
type Modifier interface {}

// NewModifiers creatas modifiers for specified data.
func NewModifiers(data res.ModifiersData) (mods []Modifier) {
	for _, md := range data.HealthMods {
		hpMod := NewHealthMod(md)
		mods = append(mods, hpMod)
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
	return
}

// ModifiersData creates data resource for modifiers.
func ModifiersData(mods ...Modifier) (data res.ModifiersData) {
	for _, m := range mods {
		switch m := m.(type) {
		case *HealthMod:
			data.HealthMods = append(data.HealthMods, m.Data())
		case *FlagMod:
			data.FlagMods = append(data.FlagMods, m.Data())
		case *QuestMod:
			data.QuestMods = append(data.QuestMods, m.Data())
		case *AreaMod:
			data.AreaMods = append(data.AreaMods, m.Data())
		}
	}
	return
}
