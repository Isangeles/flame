/*
 * inventory.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"fmt"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/rng"
)

// Struct for container with items.
type Inventory struct {
	items      map[string]Item
	tradeItems []*TradeItem
	cap        int
}

// Interface for objects with inventory.
type Container interface {
	objects.Object
	Inventory() *Inventory
}

// NewInventory creates new inventory with
// specified maximal capacity.
func NewInventory(data res.InventoryData) *Inventory {
	inv := new(Inventory)
	inv.items = make(map[string]Item)
	for _, itData := range data.Items {
		if itData.Random > 0 && !rng.RollChance(itData.Random) {
			continue
		}
		dat := res.Item(itData.ID)
		it := New(dat)
		if it == nil {
			log.Err.Printf("build inv: item: %s: fail to create item from data",
				itData.ID)
			continue
		}
		if len(itData.Serial) > 0 {
			it.SetSerial(itData.Serial)
		}
		inv.items[it.ID()+it.Serial()] = it
		if itData.Trade {
			ti := TradeItem{
				Item:  it,
				Price: itData.TradeValue,
			}
			inv.AddTradeItem(&ti)
		}
	}
	inv.cap = data.Cap
	return inv
}

// Items returns all items in inventory.
func (inv *Inventory) Items() (items []Item) {
	for _, i := range inv.items {
		items = append(items, i)
	}
	return
}

// Item returns item with specified serial ID
// from inventory or nil if no item with such serial
// ID was found in inventory.
func (inv *Inventory) Item(id, serial string) Item {
	return inv.items[id+serial]
}

// AddItems add specified item to inventory.
func (inv *Inventory) AddItem(i Item) error {
	if inv.items[i.ID()+i.Serial()] != nil {
		return nil
	}
	if len(inv.items) >= inv.Capacity() {
		return fmt.Errorf("no_inv_space")
	}
	inv.items[i.ID()+i.Serial()] = i
	return nil
}

// RemoveItem removes specified item from inventory.
func (inv *Inventory) RemoveItem(i Item) {
	delete(inv.items, i.ID()+i.Serial())
}

// TradeItems returns all items for trade
// from inventory.
func (inv *Inventory) TradeItems() []*TradeItem {
	return inv.tradeItems
}

// AddTradeItems adds specified trade item to inventory.
func (inv *Inventory) AddTradeItem(i *TradeItem) error {
	err := inv.AddItem(i)
	if err != nil {
		return err
	}
	inv.tradeItems = append(inv.tradeItems, i)
	return nil
}

// Size returns current amount of items
// in inventory.
func (inv *Inventory) Size() int {
	return len(inv.items)
}

// SetCapacity sets maximal capacity.
func (inv *Inventory) SetCapacity(c int) {
	inv.cap = c
}

// Capacity returns maximal inventory
// capacity.
func (inv *Inventory) Capacity() int {
	return inv.cap
}

// Data creates data resource for inventory.
func (inv *Inventory) Data() res.InventoryData {
	data := res.InventoryData{
		Cap: inv.Capacity(),
	}
	for _, i := range inv.TradeItems() {
		// Build trade item data.
		invItemData := res.InventoryItemData{
			ID:         i.ID(),
			Serial:     i.Serial(),
			Trade:      true,
			TradeValue: i.Price,
		}
		data.Items = append(data.Items, invItemData)
	}
	for _, i := range inv.Items() {
		// Check if item was already added as trade item.
		prs := false
		for _, id := range data.Items {
			if i.ID() == id.ID && i.Serial() == id.Serial {
				prs = true
				break
			}
		}
		if prs {
			continue
		}
		// Build item data.
		invItemData := res.InventoryItemData{
			ID:     i.ID(),
			Serial: i.Serial(),
		}
		data.Items = append(data.Items, invItemData)
	}
	return data
}
