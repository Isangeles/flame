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
	slotType item.Slot
	item     item.Equiper
}

// newEquipment creates new equipment for
// specified character.
func newEquipment(data res.EquipmentData, char *Character) *Equipment {
	eq := new(Equipment)
	eq.char = char
	eq.slots = append(eq.slots, newEquipmentSlot(item.Head))
	eq.slots = append(eq.slots, newEquipmentSlot(item.Neck))
	eq.slots = append(eq.slots, newEquipmentSlot(item.Chest))
	eq.slots = append(eq.slots, newEquipmentSlot(item.Hand))
	eq.slots = append(eq.slots, newEquipmentSlot(item.Hand))
	eq.slots = append(eq.slots, newEquipmentSlot(item.Finger))
	eq.slots = append(eq.slots, newEquipmentSlot(item.Finger))
	eq.slots = append(eq.slots, newEquipmentSlot(item.Legs))
	eq.slots = append(eq.slots, newEquipmentSlot(item.Feet))
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
		slot := item.Slot(itData.Slot)
		for _, s := range eq.Slots() {
			if s.Type() != slot {
				continue
			}
			s.SetItem(eqItem)
		}
	}
	return eq
}

// newEquipmentSlot creates new equipment slot for
// specified slot type.
func newEquipmentSlot(slotType item.Slot) *EquipmentSlot {
	s := new(EquipmentSlot)
	s.slotType = slotType
	return s
}

// Unequip removes specified item from all
// compatible slots.
func (eq *Equipment) Unequip(it item.Equiper) {
	for _, s := range eq.Slots() {
		if s.Item() == it {
			s.SetItem(nil)
		}
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
func (eqSlot *EquipmentSlot) Type() item.Slot {
	return eqSlot.slotType
}

// Item returns slot item or nil if slot is empty.
func (eqSlot *EquipmentSlot) Item() item.Equiper {
	return eqSlot.item
}

// SetItem sets inserts specified item to slot.
func (eqSlot *EquipmentSlot) SetItem(it item.Equiper) {
	eqSlot.item = it
}
