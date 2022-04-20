/*
 * req_test.go
 *
 * Copyright 2022 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/item"
	"github.com/isangeles/flame/req"
)

var (
	charData      = res.CharacterData{ID: "char", Attributes: res.AttributesData{5, 5, 5, 5, 5}}
	healthReqData = res.HealthReqData{10, false}
	manaReqData   = res.ManaReqData{10, false}
	itemReqData   = res.ItemReqData{"item1", 1, true}
)


// TestMeetReqsItem tests meet requiremet check function
// for item requirement.
func TestMeetReqsItem(t *testing.T) {
	// Meet
	char := New(charData)
	char.Update(1)
	item := item.NewMisc(res.MiscItemData{ID: "item1"})
	err := char.Inventory().AddItem(item)
	if err != nil {
		t.Fatalf("Unable to add item to the inventory: %v", err)
	}
	itemReq := req.NewItem(itemReqData)
	if !char.MeetReqs(itemReq) {
		t.Errorf("Requirement should be meet: %s not in inventory", itemReq.ItemID())
	}
	// Not meet.
	char.Inventory().RemoveItem(item)
	if char.MeetReqs(itemReq) {
		t.Errorf("Requirement should not be meet: %s in inventory", itemReq.ItemID())
	}
}

// TestMeetReqsHealth tests meet requirement check function
// for health requirement.
func TestMeetReqsHealth(t *testing.T) {
	// Meet
	char := New(charData)
	char.SetHealth(15)
	healthReq := req.NewHealth(healthReqData)
	if !char.MeetReqs(healthReq) {
		t.Errorf("Requirement should be meet: required health: %d, character health: %d",
			healthReq.Value(), char.Health())
	}
	// Not meet.
	char.SetHealth(5)
	healthReq = req.NewHealth(healthReqData)
	if char.MeetReqs(healthReq) {
		t.Errorf("Requirement should not be meet: required health: %d, character health: %d",
			healthReq.Value(), char.Health())
	}
}

// TestMeetReqsManaMeet tests meet requirement check function
// for mana requirement.
func TestMeetReqsMana(t *testing.T) {
	// Meet.
	char := New(charData)
	char.SetMana(15)
	manaReq := req.NewMana(manaReqData)
	if !char.MeetReqs(manaReq) {
		t.Errorf("Requirement should be meet: required mana: %d, character mana: %d",
			manaReq.Value(), char.Mana())
	}
	// Not meet.
	char.SetMana(5)
	manaReq = req.NewMana(manaReqData)
	if char.MeetReqs(manaReq) {
		t.Errorf("Requirement should not be meet: required mana: %d, character mana: %d",
			manaReq.Value(), char.Mana())
	}
}

// TestChargeReqsItem tests charge requirements function
// for item requirement.
func TestChargeReqsItem(t *testing.T) {
	// Charge.
	char := New(charData)
	char.Update(1)
	item := item.NewMisc(res.MiscItemData{ID: "item1"})
	err := char.Inventory().AddItem(item)
	if err != nil {
		t.Fatalf("Unable to add item to the inventory: %v", err)
	}
	itemReq := req.NewItem(itemReqData)
	char.ChargeReqs(itemReq)
	if char.Inventory().Item(item.ID(), item.Serial()) != nil {
		t.Errorf("Required item should be removed from the inventory")
	}
	// No charge.
	err = char.Inventory().AddItem(item)
	if err != nil {
		t.Fatalf("Unable to add item to the inventory: %v", err)
	}
	itemReqData.Charge = false
	itemReq = req.NewItem(itemReqData)
	char.ChargeReqs(itemReq)
	if char.Inventory().Item(item.ID(), item.Serial()) == nil {
		t.Errorf("Required item should not be removed from the inventory")
	}
}
