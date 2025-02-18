/*
 * additemmod.go
 *
 * Copyright 2020-2025 Dariusz Sikora <ds@isangeles.dev>
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

// Struct for add item modifier.
type AddItemMod struct {
	itemID string
	amount int
}

// NewAddItemMod creates new add item modifier.
func NewAddItemMod(data res.AddItemModData) *AddItemMod {
	aim := AddItemMod{
		itemID: data.ItemID,
		amount: data.Amount,
	}
	if aim.amount < 1 {
		aim.amount = 1
	}
	return &aim
}

// ItemID returns ID of the item to add.
func (aim *AddItemMod) ItemID() string {
	return aim.itemID
}

// Amount returns number of items to add.
func (aim *AddItemMod) Amount() int {
	return aim.amount
}

// Data creates data resource for modifier.
func (aim *AddItemMod) Data() res.AddItemModData {
	data := res.AddItemModData{
		ItemID: aim.ItemID(),
		Amount: aim.Amount(),
	}
	return data
}
