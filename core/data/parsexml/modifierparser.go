/*
 * modifierparser.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

package parsexml

import (
	"encoding/xml"
	
	"github.com/isangeles/flame/core/data/res"
)

// Struct for modifiers XML node.
type ModifiersXML struct {
	XMLName    xml.Name       `xml:"modifiers"`
	HealthMods []HealthModXML `xml:"healthMod"`
	HitMods    []HitModXML    `xml:"hitMod"`
}

// Struct for health modifier XML node.
type HealthModXML struct {
	XMLName  xml.Name `xml:"healthMod"`
	MinValue int      `xml:"min,attr"`
	MaxValue int      `xml:"max,attr"`
}

// Struct for hit modifier XML node.
type HitModXML struct {
	XMLName xml.Name `xml:"hitMod"`
}

// buildModifiers creates modifiers from specified XML data.
func buildModifiers(xmlModifiers *ModifiersXML) (mods []res.ModifierData) {
	// Health modifiers.
	for _, xmlMod := range xmlModifiers.HealthMods {
		mod := res.HealthModData{xmlMod.MinValue, xmlMod.MaxValue}
		mods = append(mods, mod)
	}
	// Hit modifiers.
	for _ = range xmlModifiers.HitMods {
		mod := res.HitModData{}
		mods = append(mods, mod)
	}
	return
}
