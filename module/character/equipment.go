/*
 * equipment.go
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

package character

import (
	"fmt"
	
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/item"
	"github.com/isangeles/flame/log"
)

// Struct for character equipment.
type Equipment struct {
	char  *Character
	slots []*EquipmentSlot
}

// Struct for equipment slots.
type EquipmentSlot struct {
	sType EquipmentSlotType
	item  item.Equiper
}

// Type for equipment slot type.
type EquipmentSlotType string

const (
	Head EquipmentSlotType = EquipmentSlotType("eqSlotHead")
	Neck = EquipmentSlotType("eqSlotNeck")
	Chest = EquipmentSlotType("eqSlotChest")
	HandRight = EquipmentSlotType("eqSlotHandRight")
	HandLeft = EquipmentSlotType("eqSlotHandLeft")
	FingerRight = EquipmentSlotType("eqSlotFingerRight")
	FingerLeft = EquipmentSlotType("eqSlotFingerLeft")
	Legs = EquipmentSlotType("eqSlotLegs")
	Feet = EquipmentSlotType("eqSlotFeet")
)

// newEquipment creates new equipment for
// specified character.
func newEquipment(data res.EquipmentData, char *Character) *Equipment {
	eq := new(Equipment)
	eq.char = char
	eq.slots = append(eq.slots, newEquipmentSlot(Head))
	eq.slots = append(eq.slots, newEquipmentSlot(Neck))
	eq.slots = append(eq.slots, newEquipmentSlot(Chest))
	eq.slots = append(eq.slots, newEquipmentSlot(HandRight))
	eq.slots = append(eq.slots, newEquipmentSlot(HandLeft))
	eq.slots = append(eq.slots, newEquipmentSlot(FingerRight))
	eq.slots = append(eq.slots, newEquipmentSlot(FingerLeft))
	eq.slots = append(eq.slots, newEquipmentSlot(Legs))
	eq.slots = append(eq.slots, newEquipmentSlot(Feet))
	for _, itData := range data.Items {
		it := char.Inventory().Item(itData.ID, itData.Serial)
		if it == nil {
			log.Err.Printf("character: %s: eq: fail to retrieve eq item from inv: %s",
				char.ID(), itData.ID)
			continue
		}
		eqItem, ok := it.(item.Equiper)
		if !ok {
			log.Err.Printf("character: %s: eq: not eqipable item: %s",
				char.ID(), it.ID())
			continue
		}
		// Equip.
		if !char.MeetReqs(eqItem.EquipReqs()...) {
			continue
		}
		st := EquipmentSlotType(itData.Slot)
		for _, s := range eq.Slots() {
			if s.Type() != st {
				continue
			}
			s.SetItem(eqItem)
		}
	}
	return eq
}

// newEquipmentSlot creates new equipment slot for
// specified slot type.
func newEquipmentSlot(sType EquipmentSlotType) *EquipmentSlot {
	s := new(EquipmentSlot)
	s.sType = sType
	return s
}

// Equip add specified equipable item to all
// compatible slots.
func (eq *Equipment) Equip(it item.Equiper) error {
	if !eq.char.MeetReqs(it.EquipReqs()...) {
		return fmt.Errorf("reqs not meet")
	}
	for _, s := range it.Slots() {
		switch s {
		case item.Hand:
			for _, s := range eq.Slots() {
				if s.Type() == HandRight {
					s.SetItem(it)
				}
			}
		case item.Chest:
			for _, s := range eq.Slots() {
				if s.Type() == Chest {
					s.SetItem(it)
				}
			}
		}
	}
	if !eq.Equiped(it) {
		return fmt.Errorf("no compatible slots")
	}
	return nil
}

// Unequip removes specified item from all
// compatible slots.
func (eq *Equipment) Unequip(it item.Equiper) {
	for _, s := range eq.Slots() {
		if s.Item() != it {
			continue
		}
		s.SetItem(nil)
	}
}

// Items returns slice with all equiped items.
func (eq *Equipment) Items() (items []item.Equiper) {
	uniqueItems := make(map[string]item.Equiper)
	for _, s := range eq.Slots() {
		it := s.Item()
		if it == nil {
			continue
		}
		uniqueItems[it.ID()+it.Serial()] = it
	}
	for _, it := range uniqueItems {
		items = append(items, it)
	}
	return
}

// Equiped checks whether specified item is
// equiped.
func (eq *Equipment) Equiped(item item.Equiper) bool {
	for _, i := range eq.Items() {
		if i.ID() == item.ID() && i.Serial() == item.Serial() {
			return true
		}
	}
	return false
}

// Slots retuns all equipment slots.
func (eq *Equipment) Slots() []*EquipmentSlot {
	return eq.slots
}

// Data creates data resource for equipment.
func (eq *Equipment) Data() res.EquipmentData {
	data := res.EquipmentData{}
	for _, s := range eq.Slots() {
		if s.Item() == nil {
			continue
		}
		eqItemData := res.EquipmentItemData{
			ID:     s.Item().ID(),
			Serial: s.Item().Serial(),
			Slot:   string(s.Type()),
		}
		data.Items = append(data.Items, eqItemData)
	}
	return data
}

// Type returns slot type.
func (eqSlot *EquipmentSlot) Type() EquipmentSlotType {
	return eqSlot.sType
}

// Item returns slot item or nil if slot is empty.
func (eqSlot *EquipmentSlot) Item() item.Equiper {
	return eqSlot.item
}

// SetItem sets inserts specified item to slot.
func (eqSlot *EquipmentSlot) SetItem(it item.Equiper) {
	eqSlot.item = it
}

// compact checks whether equipment slot is compatible with
// specified item slot.
func (eqSlotType EquipmentSlotType) compact(itSlot item.Slot) bool {
	switch eqSlotType {
	case Head:
		return itSlot == item.Head
	case Neck:
		return itSlot == item.Neck
	case Chest:
		return itSlot == item.Chest
	case HandRight, HandLeft:
		return itSlot == item.Hand
	case FingerRight, FingerLeft:
		return itSlot == item.Finger
	case Legs:
		return itSlot == item.Legs
	case Feet:
		return itSlot == item.Feet
	}
	return false
}
