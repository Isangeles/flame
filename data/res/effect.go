/*
 * effect.go
 *
 * Copyright 2019-2025 Dariusz Sikora <ds@isangeles.dev>
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
	XMLName           xml.Name      `xml:"effect" json:"-"`
	ID                string        `xml:"id,attr" json:"id"`
	Duration          int64         `xml:"duration,attr" json:"duration"`
	MeleeHit          bool          `xml:"melee-hit,attr" json:"melee-hit"`
	Infinite          bool          `xml:"infinite,attr" json:"infinite"`
	Hostile           bool          `xml:"hostile,attr" json:"hostile"`
	Modifiers         ModifiersData `xml:"modifiers" json:"modifiers"`
	OverTimeModifiers ModifiersData `xml:"over-time-modifiers" json:"over-time-modifiers"`
}

// Struct for modifiers data resource.
type ModifiersData struct {
	HealthMods       []HealthModData       `xml:"health-mod" json:"health-mods"`
	ManaMods         []ManaModData         `xml:"mana-mod" json:"mana-mods"`
	FlagMods         []FlagModData         `xml:"flag-mod" json:"flag-mods"`
	QuestMods        []QuestModData        `xml:"quest-mod" json:"quest-mods"`
	AreaMods         []AreaModData         `xml:"area-mod" json:"area-mods"`
	ChapterMods      []ChapterModData      `xml:"chapter-mod" json:"chapter-mods"`
	AddItemMods      []AddItemModData      `xml:"add-item-mod" json:"add-item-mods"`
	AddSkillMods     []AddSkillModData     `xml:"add-skill-mod" json:"add-skill-mods"`
	RemoveItemMods   []RemoveItemModData   `xml:"remove-item-mod" json:"remove-item-mods"`
	TransferItemMods []TransferItemModData `xml:"transfer-item-mod" json:"transfer-item-mods"`
	AttributeMods    []AttributeModData    `xml:"attribute-mod" json:"attribute-mods"`
	MemoryMods       []MemoryModData       `xml:"memory-mod" json:"memory-mods"`
	MoveSpeedMods    []MoveSpeedModData    `xml:"move-speed-mod" json:"move-speed-mods"`
}

// Struct for health modifier data.
type HealthModData struct {
	Min int `xml:"min,attr" json:"min"`
	Max int `xml:"max,attr" json:"max"`
}

// Struct for mana modifier data.
type ManaModData struct {
	Min int `xml:"min,attr" json:"min"`
	Max int `xml:"max,attr" json:"max"`
}

// Struct for flag modifier data.
type FlagModData struct {
	ID  string `xml:"id,attr" json:"id"`
	Off bool   `xml:"off,attr" json:"off"`
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

// Struct for chapter modifier.
type ChapterModData struct {
	ID string `xml:"id,attr" json:"id"`
}

// Struct for add item modifier data.
type AddItemModData struct {
	ItemID string `xml:"item-id,attr" json:"item-id"`
	Amount int    `xml:"amount,attr" json:"amount"`
}

// Type for remove item modifier data.
type RemoveItemModData AddItemModData

// Type for transfer item modifier data.
type TransferItemModData AddItemModData

// Struct for add skill modifier data.
type AddSkillModData struct {
	SkillID string `xml:"skill-id,attr" json:"skill-id"`
}

// Type for attribute modifier data.
type AttributeModData AttributesData

// Struct for memory modifier data.
type MemoryModData struct {
	Attitude string `xml:"attitude,attr" json:"attitude"`
}

// Struct for movement speed modifier data.
type MoveSpeedModData struct {
	Value int64 `xml:"value,attr" json:"value"`
}
