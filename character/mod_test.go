/*
 * mod_test.go
 *
 * Copyright 2023-2025 Dariusz Sikora <ds@isangeles.dev>
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

package character

import (
	"testing"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/item"
)

var miscItemData = res.MiscItemData{ID: "testItem"}

// TestTakeModifiersArea tests handling of area
// modifier.
func TestTakeModifiersArea(t *testing.T) {
	ob := New(charData)
	mod := effect.NewAreaMod(res.AreaModData{"testArea", 10, 10})
	ob.TakeModifiers(nil, mod)
	if ob.AreaID() != mod.AreaID() {
		t.Errorf("Invalid area ID: '%s' != '%s'",
			ob.AreaID(), mod.AreaID())
	}
	obPosX, obPosY := ob.Position()
	entryPosX, entryPosY := mod.EnterPosition()
	if obPosX != entryPosX || obPosY != entryPosY {
		t.Errorf("Invalid character position: %fx%f != %fx%f",
			obPosX, obPosY, entryPosX, entryPosY)
	}
}

// TestTakeModifiersChapter tests handling of chapter
// modifier.
func TestTakeModifiersChapter(t *testing.T) {
	ob := New(charData)
	mod := effect.NewChapterMod(res.ChapterModData{"testChapter"})
	ob.TakeModifiers(nil, mod)
	if ob.ChapterID() != mod.ChapterID() {
		t.Errorf("Invalid chapter ID: '%s' != '%s'",
			ob.ChapterID(), mod.ChapterID())
	}
}

// TestTakeModifiersAddItem tests handling of add item
// modifier.
func TestTakeModifiersAddItem(t *testing.T) {
	ob := New(charData)
	res.Miscs = append(res.Miscs, miscItemData)
	mod := effect.NewAddItemMod(res.AddItemModData{"testItem", 2})
	ob.TakeModifiers(nil, mod)
	itemsCount := 0
	for _, i := range ob.Inventory().Items() {
		if i.ID() == mod.ItemID() {
			itemsCount ++
		}
	}
	if itemsCount != 2 {
		t.Errorf("Invalid number of items after taking modifier: %d",
			itemsCount)
	}
}

// TestTakeModifiersRemoveItem tests handling of remove
// item modifier.
func TestTakeModifiersRemoveItem(t *testing.T) {
	ob := New(charData)
	for i := 0; i < 3; i++ {
		it := item.NewMisc(miscItemData)
		ob.Inventory().AddItem(it)
	}
	mod := effect.NewRemoveItemMod(res.RemoveItemModData{"testItem", 2})
	ob.TakeModifiers(nil, mod)
	if ob.Inventory().Size() != 1 {
		t.Errorf("Invalid amout of items after taking modifier: %d",
			ob.Inventory().Size())
	}
}
