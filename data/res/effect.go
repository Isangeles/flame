/*
 * effect.go
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

package res

import (
	"encoding/xml"
)

// Struct for effects data.
type EffectsData struct {
	XMLName xml.Name     `xml:"effects" json:"-"`
	Effects []EffectData `xml:"effect" json:"effects"`
}

// Struct for effect data resource.
type EffectData struct {
	XMLName   xml.Name      `xml:"effect" json:"-"`
	ID        string        `xml:"id,attr" json:"id"`
	Duration  int64         `xml:"duration,attr" json:"duration"`
	MeleeHit  bool          `xml:"melee-hit,attr" json:"melee-hit"`
	Modifiers ModifiersData `xml:"modifiers" json:"modifiers"`
}

// Struct for modifiers data resource.
type ModifiersData struct {
	XMLName     xml.Name         `xml:"modifiers" json:"-"`
	HealthMods  []HealthModData  `xml:"health-mod" json:"health-mods"`
	FlagMods    []FlagModData    `xml:"flag-mod" json:"flag-mods"`
	QuestMods   []QuestModData   `xml:"quest-mod" json:"quest-mods"`
	AreaMods    []AreaModData    `xml:"area-mod" json:"area-mods"`
	AddItemMods []AddItemModData `xml:"add-item-mod" json:"add-item-mods"`
}

// Struct for health modifier
// data.
type HealthModData struct {
	Min int `xml:"min,attr" json:"min"`
	Max int `xml:"max,attr" json:"max"`
}

// Struct for flag modifier
// data.
type FlagModData struct {
	ID string `xml:"id,attr" json:"id"`
	On bool   `xml:"disable,attr" json:"on"`
}

// Struct for quest modifier
// data.
type QuestModData struct {
	ID string `xml:"start,attr" json:"id"`
}

// Struct for area modifier
// data.
type AreaModData struct {
	ID     string  `xml:"id,attr" json:"id"`
	EnterX float64 `xml:"enter-pos-x,attr" json:"enter-pos-x"`
	EnterY float64 `xml:"enter-pos-y,attr" json:"enter-pos-y"`
}

// Struct for add item modifier data.
type AddItemModData struct {
	ItemID string `xml:"item-id,attr" json:"item-id"`
	Amount int    `xml:"amount,attr" json:"amount"`
}
