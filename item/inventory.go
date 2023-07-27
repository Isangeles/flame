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
	"sync"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/rng"
	"github.com/isangeles/flame/serial"
)

// Struct for container with items.
type Inventory struct {
	items         *sync.Map
	onItemRemoved func(i Item)
}

// Struct for inventory items.
type InventoryItem struct {
	Item
	Price int
	Trade bool
	Loot  bool
}

// Interface for objects with inventory.
type Container interface {
	serial.Serialer
	Inventory() *Inventory
}

// NewInventory creates new inventory.
func NewInventory() *Inventory {
	i := Inventory{items: new(sync.Map)}
	return &i
}

// Update updates all items in the inventory.
func (i *Inventory) Update(delta int64) {
	for _, it := range i.Items() {
		it.Update(delta)
	}
}

// Items returns all items in inventory.
func (i *Inventory) Items() (items []*InventoryItem) {
	addItem := func(k, v interface{}) bool {
		it, ok := v.(*InventoryItem)
		if ok {
			items = append(items, it)
		}
		return true
	}
	i.items.Range(addItem)
	return
}

// Item returns item with specified ID and serial
// from the inventory or nil if no such item was
// found.
func (i *Inventory) Item(id, serial string) *InventoryItem {
	val, _ := i.items.Load(id + serial)
	if val == nil {
		return nil
	}
	return val.(*InventoryItem)
}

// AddItems add specified item to inventory.
// Item will be marked as tradeable and lootable inside the inventory.
// Trade value will be set as the same as item value.
func (i *Inventory) AddItem(it Item) {
	invIt := InventoryItem{it, it.Value(), true, true}
	i.items.Store(it.ID()+it.Serial(), &invIt)
}

// RemoveItem removes specified item from inventory.
func (i *Inventory) RemoveItem(it Item) {
	i.items.Delete(it.ID() + it.Serial())
	if i.onItemRemoved != nil {
		i.onItemRemoved(it)
	}
}

// Size returns current amount of items
// in inventory.
func (i *Inventory) Size() int {
	return len(i.Items())
}

// SetOnItemRemoved sets function to trigger  after
// removing item from the inventory.
func (i *Inventory) SetOnItemRemovedFunc(f func(i Item)) {
	i.onItemRemoved = f
}

// Apply applies specified data on the inventory.
func (i *Inventory) Apply(data res.InventoryData) {
	// Clear removed items.
	for _, it := range i.Items() {
		found := false
		for _, invItData := range data.Items {
			if it.ID() == invItData.ID && it.Serial() == invItData.Serial {
				found = true
				break
			}
		}
		if !found {
			i.items.Delete(it.ID() + it.Serial())
		}
	}
	// Add/update items.
	for _, invItData := range data.Items {
		it := i.Item(invItData.ID, invItData.Serial)
		if it != nil {
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
			ID:         it.ID(),
			Serial:     it.Serial(),
			TradeValue: it.Price,
			NoTrade:    !it.Trade,
			NoLoot:     !it.Loot,
		}
		data.Items = append(data.Items, invItemData)
	}
	return data
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
		invIt := InventoryItem{it, data.TradeValue, !data.NoTrade, !data.NoLoot}
		i.items.Store(it.ID()+it.Serial(), &invIt)
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
	it.SetSerial(data.Serial)
	invIt := InventoryItem{it, data.TradeValue, !data.NoTrade, !data.NoLoot}
	i.items.Store(it.ID()+it.Serial(), &invIt)
	return nil
}
