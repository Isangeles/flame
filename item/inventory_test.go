/*
 * inventory_test.go
 *
 * Copyright 2023 Dariusz Sikora <ds@isangeles.dev>
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

package item

import (
	"testing"

	"github.com/isangeles/flame/data/res"
)

var (
	invItemData = res.InventoryItemData{ID: "item", Serial: "0", TradeValue: 10}
	invData     = res.InventoryData{Items: []res.InventoryItemData{invItemData}}
)

// TestInventoryApply tests inventory data apply function.
func TestInventoryApply(t *testing.T) {
	// Add test item to resources base.
	res.Miscs = append(res.Miscs, res.MiscItemData{ID: "item"})
	// Create inventory.
	inv := NewInventory()
	invData.Items = append(invData.Items, res.InventoryItemData{ID: "item", Amount: 2})
	// Test.
	inv.Apply(invData)
	if len(inv.Items()) != 3 {
		t.Errorf("Invalid amount of items: %d != 3", len(inv.Items()))
	}
	item := inv.Item(invItemData.ID, invItemData.Serial)
	if item == nil {
		t.Errorf("No item in inventory")
	}
	if !item.Loot {
		t.Errorf("No loot item in inventory")
	}
	if !item.Trade {
		t.Errorf("No trade item in inventory")
	}
	if item.Price != 10 {
		t.Errorf("Invalid trade value: %d != 10", item.Price)
	}
}
