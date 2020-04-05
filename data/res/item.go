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

// Empty interface for item data
// structs.
type ItemData interface{}

// Struct for armor resource data.
type ArmorData struct {
	ID        string
	Value     int
	Level     int
	Armor     int
	EQEffects []EffectData
	EQReqs    ReqsData
	Slots     []ItemSlotData
	Loot      bool
}

// Struct for weapon resource data.
type WeaponData struct {
	ID         string
	Value      int
	Level      int
	DMGMin     int
	DMGMax     int
	DMGType    string
	DMGEffects []EffectData
	EQReqs     ReqsData
	Slots      []string
	Loot       bool
}

// Struct for miscellaneous items data.
type MiscItemData struct {
	ID       string
	Value    int
	Level    int
	Loot     bool
	Currency bool
}

// Struct for item slot data.
type ItemSlotData struct {
	ID string
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
