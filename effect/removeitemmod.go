/*
 * removeitemmod.go
 *
 * Copyright 2021-2025 Dariusz Sikora <ds@isangeles.dev>
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

// Type for remove item modifier.
type RemoveItemMod AddItemMod

// NewRemoveItemMod creates new remove item modifier.
func NewRemoveItemMod(data res.RemoveItemModData) *RemoveItemMod {
	rim := RemoveItemMod{
		itemID: data.ItemID,
		amount: data.Amount,
	}
	if rim.amount < 1 {
		rim.amount = 1
	}
	return &rim
}

// ItemID returns ID of the item to remove.
func (rim *RemoveItemMod) ItemID() string {
	return rim.itemID
}

// Amount returns number of items to remove.
func (rim *RemoveItemMod) Amount() int {
	return rim.amount
}

// Data creates data resource for modifier.
func (rim *RemoveItemMod) Data() res.RemoveItemModData {
	data := res.RemoveItemModData{
		ItemID: rim.itemID,
		Amount: rim.amount,
	}
	return data
}
