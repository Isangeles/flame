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
	Slots     []string
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

// Struct for miscellaneous items.
type MiscItemData struct {
	ID       string
	Value    int
	Level    int
	Loot     bool
	Currency bool
}

// Struct for inventory data.
type InventoryData struct {
	Cap   int
	Items []InventoryItemData
}

// Struct for inventory item data
// resource.
type InventoryItemData struct {
	ID         string
	Serial     string
	Trade      bool
	TradeValue int
	Random     float64
}
