/*
 * inventory.go
 *
 * Copyright 2018-2023 Dariusz Sikora <ds@isangeles.dev>
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
	"github.com/isangeles/flame/rng"
	"github.com/isangeles/flame/serial"
)

// Struct for container with items.
type Inventory struct {
	items         map[string]Item
	lootItems     []Item
	tradeItems    []*TradeItem
	onItemRemoved func(i Item)
}

// Interface for objects with inventory.
type Container interface {
	serial.Serialer
	Inventory() *Inventory
}

// NewInventory creates new inventory.
func NewInventory() *Inventory {
	i := Inventory{items: make(map[string]Item)}
	return &i
}

// Update updates all items in the inventory.
func (i *Inventory) Update(delta int64) {
	for _, it := range i.Items() {
		it.Update(delta)
	}
}

// Items returns all items in inventory.
func (i *Inventory) Items() (items []Item) {
	for _, it := range i.items {
		items = append(items, it)
	}
	return
}

// Item returns item with specified ID and serial
// from the inventory or nil if no such item was
// found.
func (i *Inventory) Item(id, serial string) Item {
	return i.items[id+serial]
}

// TradeItem returns 'tradable' item with specified ID
// and serial from the inventory or nil if no such item
// was found.
func (i *Inventory) TradeItem(id, serial string) *TradeItem {
	for _, it := range i.TradeItems() {
		if it.ID() == id && it.Serial() == serial {
			return it
		}
	}
	return nil
}

// LootItem returns 'lootable' item with specified ID
// and serial from the inventory or nil if no such
// item was found.
func (i *Inventory) LootItem(id, serial string) Item {
	for _, it := range i.LootItems() {
		if it.ID() == id && it.Serial() == serial {
			return it
		}
	}
	return nil
}

// AddItems add specified item to inventory.
func (i *Inventory) AddItem(it Item) {
	i.items[it.ID()+it.Serial()] = it
}

// RemoveItem removes specified item from inventory.
func (i *Inventory) RemoveItem(it Item) {
	delete(i.items, it.ID()+it.Serial())
	if i.onItemRemoved != nil {
		i.onItemRemoved(it)
	}
}

// TradeItems returns all items for trade
// from inventory.
func (i *Inventory) TradeItems() []*TradeItem {
	return i.tradeItems
}

// AddTradeItems adds specified trade item to inventory.
func (i *Inventory) AddTradeItem(it *TradeItem) {
	i.AddItem(it)
	i.tradeItems = append(i.tradeItems, it)
}

// LootItems returns all 'lootable' items from inventory.
func (i *Inventory) LootItems() []Item {
	return i.lootItems
}

// AddLootItem adds specified 'lootable' item to
// the inventory.
func (i *Inventory) AddLootItem(it Item) {
	i.AddItem(it)
	i.lootItems = append(i.lootItems, it)
}

// Size returns current amount of items
// in inventory.
func (i *Inventory) Size() int {
	return len(i.items)
}

// SetOnItemRemoved sets function to trigger  after
// removing item from the inventory.
func (i *Inventory) SetOnItemRemovedFunc(f func(i Item)) {
	i.onItemRemoved = f
}

// Apply applies specified data on the inventory.
func (i *Inventory) Apply(data res.InventoryData) {
	// Clear removed items.
	i.tradeItems = make([]*TradeItem, 0)
	i.lootItems = make([]Item, 0)
	for _, it := range i.Items() {
		found := false
		for _, invItData := range data.Items {
			if it.ID() == invItData.ID && it.Serial() == invItData.Serial {
				found = true
				break
			}
		}
		if !found {
			delete(i.items, it.ID()+it.Serial())
		}
	}
	// Add/update items.
	for _, invItData := range data.Items {
		it := i.Item(invItData.ID, invItData.Serial)
		if it != nil {
			i.updateItem(it, invItData)
			continue
		}
		if len(invItData.Serial) > 0 {
			err := i.restoreItem(invItData)
			if err != nil {
				log.Err.Printf("Inventory: Apply: unable to restore item: %s %s: %v",
					invItData.ID, invItData.Serial, err)
			}
			continue
		}
		err := i.spawnItem(invItData)
		if err != nil {
			log.Err.Printf("Inventory: Apply: unable to spawn item: %s: %v",
				invItData.ID, err)
		}

	}
}

// Data creates data resource for inventory.
func (i *Inventory) Data() res.InventoryData {
	data := res.InventoryData{}
	for _, it := range i.Items() {
		// Build item data.
		invItemData := res.InventoryItemData{
			ID:     it.ID(),
			Serial: it.Serial(),
		}
		if it := i.TradeItem(it.ID(), it.Serial()); it != nil {
			invItemData.Trade = true
			invItemData.TradeValue = it.Price
		}
		if i.LootItem(it.ID(), it.Serial()) != nil {
			invItemData.Loot = true
		}
		data.Items = append(data.Items, invItemData)
	}
	return data
}

// Update item updates item with specified item inventory data.
func (i *Inventory) updateItem(it Item, data res.InventoryItemData) {
	if len(data.Serial) > 0 {
		it.SetSerial(data.Serial)
	}
	if data.Trade {
		ti := TradeItem{
			Item:  it,
			Price: data.TradeValue,
		}
		i.AddTradeItem(&ti)
	}
	if data.Loot {
		i.AddLootItem(it)
	}
}

// spawnItem spawns specified amount of items in the inventory.
func (i *Inventory) spawnItem(data res.InventoryItemData) error {
	if data.Random > 0 && !rng.RollChance(data.Random) {
		return nil
	}
	itData := res.Item(data.ID)
	if itData == nil {
		return fmt.Errorf("Item data not found: %s", data.ID)
	}
	if data.Amount == 0 {
		data.Amount = 1
	}
	for itemNumber := 0; itemNumber < data.Amount; itemNumber++ {
		it := New(itData)
		if it == nil {
			return fmt.Errorf("Item not created: %s", data.ID)
		}
		i.updateItem(it, data)
		i.items[it.ID()+it.Serial()] = it
	}
	return nil
}

// restoreItem restores inventory item for specified data.
func (i *Inventory) restoreItem(data res.InventoryItemData) error {
	itData := res.Item(data.ID)
	if itData == nil {
		return fmt.Errorf("Item data not found: %s", data.ID)
	}
	it := New(itData)
	if it == nil {
		return fmt.Errorf("Item not created: %s", data.ID)
	}
	i.updateItem(it, data)
	i.items[it.ID()+it.Serial()] = it
	return nil
}
