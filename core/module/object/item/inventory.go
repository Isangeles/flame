/*
 * inventory.go
 *
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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
)

// Struct for container with items.
type Inventory struct {
	items   []Item
	maxSize int
}

// NewInventory creates new inventory with
// specified maximal capacity.
func NewInventory(size int) *Inventory {
	inv := new(Inventory)
	inv.maxSize = size
	return inv
}

// Items returns all items in inventory.
func (inv *Inventory) Items() []Item {
	return inv.items
}

// Item returns item with specified serial ID
// from inventory or nil if no item with such serial
// ID was found in inventory.
func (inv *Inventory) Item(serialID string) Item {
	for _, i := range inv.Items() {
		if i.SerialID() == serialID {
			return i
		}
	}
	return nil
}

// AddItems add specified item to inventory.
func (inv *Inventory) AddItem(i Item) error {
	if len(inv.items) >= inv.maxSize {
		return fmt.Errorf("no_inv_space")
	}
	inv.items = append(inv.items, i)
	return nil
}

// Size returns current amount of items
// in inventory.
func (inv *Inventory) Size() int {
	return len(inv.items)
}
