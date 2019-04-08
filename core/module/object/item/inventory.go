/*
 * inventory.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/core/module/object"
)

// Struct for container with items.
type Inventory struct {
	items []Item
	cap   int
}

// NewInventory creates new inventory with
// specified maximal capacity.
func NewInventory(cap int) *Inventory {
	inv := new(Inventory)
	inv.cap = cap
	return inv
}

// Items returns all items in inventory.
func (inv *Inventory) Items() []Item {
	return inv.items
}

// Item returns item with specified serial ID
// from inventory or nil if no item with such serial
// ID was found in inventory.
func (inv *Inventory) Item(id, serial string) Item {
	for _, i := range inv.Items() {
		if i.ID()+i.Serial() == id+serial  {	
			return i
		}
	}
	return nil
}

// AddItems add specified item to inventory.
func (inv *Inventory) AddItem(i Item) error {
	if len(inv.items) >= inv.Capacity() {
		return fmt.Errorf("no_inv_space")
	}
	inv.items = append(inv.items, i)
	return nil
}

// RemoveItem removes specified item from inventory.
func (inv *Inventory) RemoveItem(i Item) {
	for _, it := range inv.items {
		if !object.Equals(it, i) {
			continue
		}
		//delete(inv.items, id)
	}
}

// Size returns current amount of items
// in inventory.
func (inv *Inventory) Size() int {
	return len(inv.items)
}

// Capacity returns maximal inventory
// capacity.
func (inv *Inventory) Capacity() int {
	return inv.cap
}
