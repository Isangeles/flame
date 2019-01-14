/*
 * equipment.go
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

package character

import (
	"fmt"
	
	"github.com/isangeles/flame/core/module/object/item"
)

// Struct for character equipment.
type Equipment struct {
	char        *Character
	head        *EquipmentSlot
	neck        *EquipmentSlot
	chest       *EquipmentSlot
	handRight   *EquipmentSlot
	handLeft    *EquipmentSlot
	fingerRight *EquipmentSlot
	fingerLeft  *EquipmentSlot
	legs        *EquipmentSlot
	feet        *EquipmentSlot
}

// Struct for equipment slots.
type EquipmentSlot struct {
	stype EquipmentSlotType
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
	eq.head = newEquipmentSlot(Head)
	eq.neck = newEquipmentSlot(Neck)
	eq.chest = newEquipmentSlot(Chest)
	eq.handRight = newEquipmentSlot(Hand_right)
	eq.handLeft = newEquipmentSlot(Hand_left)
	eq.fingerRight = newEquipmentSlot(Finger_right)
	eq.fingerLeft = newEquipmentSlot(Finger_left)
	eq.legs = newEquipmentSlot(Legs)
	eq.feet = newEquipmentSlot(Feet)
	return eq
}

func newEquipmentSlot(stype EquipmentSlotType) *EquipmentSlot {
	s := new(EquipmentSlot)
	s.stype = stype
	return s
}

// EquipHandRight assigns specified 'equipable' item to right hand slot,
// returns error if equip fail(e.q. equip reqs aren't meet).
func (eq *Equipment) EquipHandRight(it item.Equiper) error {
	if !eq.char.MeetReqs(it.EquipReqs()) {
		return fmt.Errorf("reqs_not_meet")
	}
	if len(it.Slots()) != 1 || it.Slots()[0] != item.Hand {
		return fmt.Errorf("slot_not_match")
	}
	eq.handRight.item = it
	return nil
}

// Items returns slice with all equiped items.
func (eq *Equipment) Items() []item.Equiper {
	its := make([]item.Equiper, 0)
	if eq.handRight.item != nil {
		its = append(its, eq.handRight.item)	
	}
	if eq.handLeft.item != nil {
		its = append(its, eq.handLeft.item)
	}
	return its
}

// HandRight returns item from right hand slot.
func (eq *Equipment) HandRight() item.Equiper {
	return eq.handRight.item
}

// Type returns slot type.
func (eqSlot *EquipmentSlot) Type() EquipmentSlotType {
	return eqSlot.stype
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
	case Hand_right: case Hand_left:
		return itSlot == item.Hand
	case Finger_right: case Finger_left:
		return itSlot == item.Finger
	case Legs:
		return itSlot == item.Legs
	case Feet:
		return itSlot == item.Feet
	}
	return false
}
