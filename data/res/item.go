/*
 * item.go
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

// Empty interface for item data
// structs.
type ItemData interface{}

// Struct for armors data.
type ArmorsData struct {
	XMLName xml.Name    `xml:"armors" json:"-"`
	Armors  []ArmorData `xml:"armor" json:"armors"`
}

// Struct for armor resource data.
type ArmorData struct {
	ID        string         `xml:"id,attr" json:"id"`
	Value     int            `xml:"value,attr" json:"value"`
	Level     int            `xml:"level,attr" json:"level"`
	Armor     int            `xml:"armor,attr" json:"armor"`
	Loot      bool           `xml:"loot,attr" json:"loot"`
	EQEffects []EffectData   `xml:"eq>effects>effect" json:"eq-effects"`
	EQReqs    ReqsData       `xml:"eq>reqs" json:"eq-reqs"`
	Slots     []ItemSlotData `xml:"slots>slot" json:"slots"`
}

// Struct for weapons data.
type WeaponsData struct {
	XMLName xml.Name     `xml:"weapons" json:"-"`
	Weapons []WeaponData `xml:"weapon" json:"weapons"`
}

// Struct for weapon resource data.
type WeaponData struct {
	ID     string         `xml:"id,attr" json:"id"`
	Value  int            `xml:"value,attr" json:"value"`
	Level  int            `xml:"level,attr" json:"level"`
	Damage DamageData     `xml:"damage" json:"damage"`
	EQReqs ReqsData       `xml:"reqs" json:"eq-reqs"`
	Loot   bool           `xml:"loot,attr" json:"loot"`
	Slots  []ItemSlotData `xml:"slots>slot" json:"slots"`
}

// Struct for damage data.
type DamageData struct {
	Type    string             `xml:"type,attr" json:"type"`
	Min     int                `xml:"min,attr" json:"min"`
	Max     int                `xml:"max,attr" json:"max"`
	Effects []ObjectEffectData `xml:"effects>effect" json:"effects"`
}

// Struct for misc items data.
type MiscItemsData struct {
	XMLName xml.Name       `xml:"miscs" json:"-"`
	Miscs   []MiscItemData `xml:"misc" json:"miscs"`
}

// Struct for miscellaneous items data.
type MiscItemData struct {
	ID         string        `xml:"id,attr" json:"id"`
	Value      int           `xml:"value,attr" json:"value"`
	Level      int           `xml:"level,attr" json:"level"`
	Loot       bool          `xml:"loot,attr" json:"loot"`
	Currency   bool          `xml:"currency,attr" json:"currency"`
	Consumable bool          `xml:"consumable,attr" json:"consumable"`
	UseAction  UseActionData `xml:"use" json:"use"`
}

// Struct for item slot data.
type ItemSlotData struct {
	ID string `xml:"id,attr" json:"id"`
}

// Struct for inventory data.
type InventoryData struct {
	Cap   int                 `xml:"cap,attr" json:"cap"`
	Items []InventoryItemData `xml:"item" json:"items"`
}

// Struct for inventory item data
// resource.
type InventoryItemData struct {
	ID         string  `xml:"id,attr" json:"id"`
	Serial     string  `xml:"serial,attr" json:"serial"`
	Trade      bool    `xml:"trade,attr" json:"trade"`
	TradeValue int     `xml:"trade-value,attr" json:"trade-value"`
	Random     float64 `xml:"random,attr" json:"random"`
}
