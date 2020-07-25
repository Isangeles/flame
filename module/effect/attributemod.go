/*
 * attributemod.go
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

package effect

import (
	"github.com/isangeles/flame/data/res"
)

// Struct for attribute modifier.
type AttributeMod struct {
	strength     int
	constitution int
	dexterity    int
	intelligence int
	wisdom       int
}

// NewAttributeMod creates new attribute modifier.
func NewAttributeMod(data res.AttributeModData) *AttributeMod {
	am := AttributeMod{
		strength:     data.Str,
		constitution: data.Con,
		dexterity:    data.Dex,
		wisdom:       data.Wis,
		intelligence: data.Int,
	}
	return &am
}

// Strength returns modifier value for strength attribute.
func (am *AttributeMod) Strength() int {
	return am.strength
}

// Constitution returns modifier value for constitution attribute.
func (am *AttributeMod) Constitution() int {
	return am.constitution
}

// Dexterity returns modifier value for dexterity attribute.
func (am *AttributeMod) Dexterity() int {
	return am.dexterity
}

// Intelligence returns modifier value for intelligence attribute.
func (am *AttributeMod) Intelligence() int {
	return am.intelligence
}

// Wisdom returns modifer value for wisom attribute.
func (am *AttributeMod) Wisdom() int {
	return am.wisdom
}

// Data creates data resource for modifier.
func (am *AttributeMod) Data() res.AttributeModData {
	data := res.AttributeModData{
		Str: am.Strength(),
		Con: am.Constitution(),
		Dex: am.Dexterity(),
		Int: am.Intelligence(),
		Wis: am.Wisdom(),
	}
	return data
}
