/*
 * equipment.go
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

package character

import (
	"fmt"

	"github.com/isangeles/flame/core/module/object/item"
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
type EquipmentSlotType int

const (
	Head EquipmentSlotType = iota
	Neck
	Chest
	Hand_right
	Hand_left
	Finger_right
	Finger_left
	Legs
	Feet
)

// newEquipment creates new equipment for
// specified character.
func newEquipment(char *Character) *Equipment {
	eq := new(Equipment)
	eq.char = char
	eq.slots = append(eq.slots, newEquipmentSlot(Head))
	eq.slots = append(eq.slots, newEquipmentSlot(Neck))
	eq.slots = append(eq.slots, newEquipmentSlot(Chest))
	eq.slots = append(eq.slots, newEquipmentSlot(Hand_right))
	eq.slots = append(eq.slots, newEquipmentSlot(Hand_left))
	eq.slots = append(eq.slots, newEquipmentSlot(Finger_right))
	eq.slots = append(eq.slots, newEquipmentSlot(Finger_left))
	eq.slots = append(eq.slots, newEquipmentSlot(Legs))
	eq.slots = append(eq.slots, newEquipmentSlot(Feet))
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
		return fmt.Errorf("reqs_not_meet")
	}
	for _, s := range it.Slots() {
		switch s {
		case item.Hand:
			for _, s := range eq.Slots() {
				if s.Type() == Hand_right {
					s.SetItem(it)
				}
			}
		}
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

// ID returns slot ID.
func (eqSlot *EquipmentSlot) ID() string {
	switch eqSlot.Type() {
	case Head:
		return "eq_slot_head"
	case Neck:
		return "eq_slot_neck"
	case Chest:
		return "eq_slot_chest"
	case Hand_right:
		return "eq_slot_hand_right"
	case Hand_left:
		return "eq_slot_hand_left"
	case Finger_right:
		return "eq_slot_finger_right"
	case Finger_left:
		return "eq_slot_finger_left"
	case Legs:
		return "eq_slot_legs"
	case Feet:
		return "eq_slot_feet"
	}
	return "eq_slot_unknown"
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
	case Hand_right, Hand_left:
		return itSlot == item.Hand
	case Finger_right, Finger_left:
		return itSlot == item.Finger
	case Legs:
		return itSlot == item.Legs
	case Feet:
		return itSlot == item.Feet
	}
	return false
}
