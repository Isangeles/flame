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

var (
	miscItemData = res.MiscItemData{ID: "testItem"}
)

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

// TestTakeModifiersTransferItem tests handling of
// transfer item modifier.
func TestTakeModifiersTransferItem(t *testing.T) {
	ob1 := New(charData)
	for i := 0; i < 3; i++ {
		it := item.NewMisc(miscItemData)
		ob1.Inventory().AddItem(it)
	}
	ob2 := New(charData)
	mod := effect.NewTransferItemMod(res.TransferItemModData{"testItem", 2})
	ob1.TakeModifiers(ob2, mod)
	itemsCount := 0
	for _, i := range ob2.Inventory().Items() {
		if i.ID() == mod.ItemID() {
			itemsCount ++
		}
	}
	if itemsCount != 2 {
		t.Errorf("Invalid number of items in source inventory: %d",
			itemsCount)
	}
	if len(ob1.Inventory().Items()) != 1 {
		t.Errorf("Invalid number od items in target inventory: %d",
			itemsCount)
	}
}

// TestTakeModifiersAddSkill tests handling of add
// skill modifier.
func TestTakeModifiersAddSkill(t *testing.T) {
	ob := New(charData)
	res.Skills = append(res.Skills, skillData)
	mod := effect.NewAddSkillMod(res.AddSkillModData{"skill"})
	ob.TakeModifiers(nil, mod)
	found := false
	for _, s := range ob.Skills() {
		if s.ID() == mod.SkillID() {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Skill not added")
	}
}
