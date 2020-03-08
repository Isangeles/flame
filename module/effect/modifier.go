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
type Modifier interface {
}

// NewModifiers creatas modifiers for specified data.
func NewModifiers(data ...res.ModifierData) (mods []Modifier) {
	for _, md := range data {
		switch md := md.(type) {
		case res.HealthModData:
			hpMod := NewHealthMod(md)
			mods = append(mods, hpMod)
		case res.FlagModData:
			flagMod := NewFlagMod(md)
			mods = append(mods, flagMod)
		case res.QuestModData:
			questMod := NewQuestMod(md)
			mods = append(mods, questMod)
		case res.AreaModData:
			areaMod := NewAreaMod(md)
			mods = append(mods, areaMod)
		}
	}
	return
}
