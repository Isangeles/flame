/*
 * transferitemmod.go
 *
 * Copyright 2025 Dariusz Sikora <ds@isangeles.dev>
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

// Type for transfer item modifier.
type TransferItemMod AddItemMod

// NewTransferItemMod creates new remove item modifier.
func NewTransferItemMod(data res.TransferItemModData) *TransferItemMod {
	im := TransferItemMod{
		itemID: data.ItemID,
		amount: data.Amount,
	}
	return &im
}

// ItemID returns ID of the item to remove.
func (tim *TransferItemMod) ItemID() string {
	return tim.itemID
}

// Amount returns number of items to remove.
func (tim *TransferItemMod) Amount() int {
	return tim.amount
}

// Data creates data resource for modifier.
func (tim *TransferItemMod) Data() res.TransferItemModData {
	data := res.TransferItemModData{
		ItemID: tim.itemID,
		Amount: tim.amount,
	}
	return data
}
