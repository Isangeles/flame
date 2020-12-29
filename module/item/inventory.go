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
	lootItems  []Item
	tradeItems []*TradeItem
	cap        int
}

// Interface for objects with inventory.
type Container interface {
	objects.Object
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
func (i *Inventory) AddItem(it Item) error {
	if i.items[it.ID()+it.Serial()] != nil {
		return nil
	}
	if len(i.items) >= i.Capacity() {
		return fmt.Errorf("no_inv_space")
	}
	i.items[it.ID()+it.Serial()] = it
	return nil
}

// RemoveItem removes specified item from inventory.
func (i *Inventory) RemoveItem(it Item) {
	delete(i.items, it.ID()+it.Serial())
}

// TradeItems returns all items for trade
// from inventory.
func (i *Inventory) TradeItems() []*TradeItem {
	return i.tradeItems
}

// AddTradeItems adds specified trade item to inventory.
func (i *Inventory) AddTradeItem(it *TradeItem) error {
	err := i.AddItem(it)
	if err != nil {
		return err
	}
	i.tradeItems = append(i.tradeItems, it)
	return nil
}

// LootItems returns all 'lootable' items from inventory.
func (i *Inventory) LootItems() []Item {
	return i.lootItems
}

// AddLootItem adds specified 'lootable' item to
// the inventory.
func (i *Inventory) AddLootItem(it Item) error {
	err := i.AddItem(it)
	if err != nil {
		return err
	}
	i.lootItems = append(i.lootItems, it)
	return nil
}

// Size returns current amount of items
// in inventory.
func (i *Inventory) Size() int {
	return len(i.items)
}

// SetCapacity sets maximal capacity.
func (i *Inventory) SetCapacity(c int) {
	i.cap = c
}

// Capacity returns maximal inventory
// capacity.
func (i *Inventory) Capacity() int {
	return i.cap
}

// Apply applies specified data on the inventory.
func (i *Inventory) Apply(data res.InventoryData) {
	for _, invItData := range data.Items {
		it := i.Item(invItData.ID, invItData.Serial)
		if it == nil {
			if invItData.Random > 0 && !rng.RollChance(invItData.Random) {
				continue
			}
			itData := res.Item(invItData.ID)
			if itData == nil {
				log.Err.Printf("Inventory: Apply: item: %s: data not found", invItData.ID)
				continue
			}
			it = New(itData)
			if it == nil {
				log.Err.Printf("Inventory: Apply: item: %s: unable to create item from data",
					invItData.ID)
				continue
			}
			if len(invItData.Serial) > 0 {
				it.SetSerial(invItData.Serial)
			}
			i.items[it.ID()+it.Serial()] = it
		}
		if invItData.Trade {
			ti := TradeItem{
				Item:  it,
				Price: invItData.TradeValue,
			}
			err := i.AddTradeItem(&ti)
			if err != nil {
				log.Err.Printf("Inventory: Apply: item: %s: unable to add trade item: %v",
					invItData.ID, err)
			}
		}
		if invItData.Loot {
			err := i.AddLootItem(it)
			if err != nil {
				log.Err.Printf("Inventory: Apply: item: %s: unable to add loot item: %v",
					invItData.ID, err)
			}
		}
	}
	i.cap = data.Cap
}

// Data creates data resource for inventory.
func (i *Inventory) Data() res.InventoryData {
	data := res.InventoryData{
		Cap: i.Capacity(),
	}
	for _, it := range i.Items() {
		// Build item data.
		invItemData := res.InventoryItemData{
			ID:         it.ID(),
			Serial:     it.Serial(),
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
