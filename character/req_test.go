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
	healthReqData = res.HealthReqData{10, false, true}
	manaReqData   = res.ManaReqData{10, false, true}
	itemReqData   = res.ItemReqData{"item1", 1, true}
	combatReqData = res.CombatReqData{true}
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

// TestMeetReqsCombat test meet requirement check function
// for combat requirement.
func TestMeetReqsCombat(t *testing.T) {
	// Meet.
	char := New(charData)
	hostileCharData := charData
	hostileCharData.Attitude = string(Hostile)
	hostileChar := New(hostileCharData)
	char.SetTarget(hostileChar)
	combatReq := req.NewCombat(combatReqData)
	if !char.MeetReqs(combatReq) {
		t.Errorf("Requirement should be meet: character in combat: %v", char.Fighting())
	}
	// Not Meet.
	char.SetTarget(nil)
	if char.MeetReqs(combatReq) {
		t.Errorf("Requirement should not be meet: character in combat: %v", char.Fighting())
	}
}

// TestChargeReqs tests charge requirements function.
func TestChargeReqs(t *testing.T) {
	// Handle mixed reqs(chargeable and non chargeable)
	char := New(charData)
	reqs := make([]req.Requirement, 3)
	reqs = append(reqs, req.NewMana(manaReqData))
	reqs = append(reqs, req.NewItem(itemReqData))
	reqs = append(reqs, req.NewCombat(combatReqData))
	char.ChargeReqs(reqs...)
}

// TestChargeReqsMana tests charge requrements function
// for mana requirement.
func TestChargeReqsMana(t *testing.T) {
	// Charge.
	char := New(charData)
	char.SetMana(15)
	manaReq := req.NewMana(manaReqData)
	char.ChargeReqs(manaReq)
	if char.Mana() != 5 {
		t.Errorf("Invalid mana value after charge: %d != 5", char.Mana())
	}
	// No charge.
	char.SetMana(15)
	manaReqData.Charge = false
	manaReq = req.NewMana(manaReqData)
	char.ChargeReqs(manaReq)
	if char.Mana() != 15 {
		t.Errorf("Mana value should not change: %d != 15", char.Mana())
	}
}

// TestChargeReqsHealth tests charge requrements function
// for health requirement.
func TestChargeReqsHealth(t *testing.T) {
	// Charge.
	char := New(charData)
	char.SetHealth(15)
	healthReq := req.NewHealth(healthReqData)
	char.ChargeReqs(healthReq)
	if char.Health() != 5 {
		t.Errorf("Invalid health value after charge: %d != 5", char.Health())
	}
	// No charge.
	char.SetHealth(15)
	healthReqData.Charge = false
	healthReq = req.NewHealth(healthReqData)
	char.ChargeReqs(healthReq)
	if char.Health() != 15 {
		t.Errorf("Health value should not change: %d != 15", char.Health())
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
