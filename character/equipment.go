/*
 * equipment.go
 *
 * Copyright 2018-2022 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/item"
	"github.com/isangeles/flame/log"
)

// Struct for character equipment.
type Equipment struct {
	char  *Character
	slots []*EquipmentSlot
}

// Struct for equipment slots.
type EquipmentSlot struct {
	id       int
	slotType item.Slot
	item     item.Equiper
}

// newEquipment creates new equipment for
// specified character.
func newEquipment(char *Character) *Equipment {
	eq := new(Equipment)
	eq.char = char
	eq.slots = append(eq.slots, eq.newEquipmentSlot(item.Head))
	eq.slots = append(eq.slots, eq.newEquipmentSlot(item.Neck))
	eq.slots = append(eq.slots, eq.newEquipmentSlot(item.Chest))
	eq.slots = append(eq.slots, eq.newEquipmentSlot(item.Hand))
	eq.slots = append(eq.slots, eq.newEquipmentSlot(item.Hand))
	eq.slots = append(eq.slots, eq.newEquipmentSlot(item.Finger))
	eq.slots = append(eq.slots, eq.newEquipmentSlot(item.Finger))
	eq.slots = append(eq.slots, eq.newEquipmentSlot(item.Legs))
	eq.slots = append(eq.slots, eq.newEquipmentSlot(item.Feet))
	return eq
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
		if s.Item() == nil {
			continue
		}
		uniqueItems[s.Item().ID()+s.Item().Serial()] = s.Item()
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

// Apply applies specified data on the equipment.
func (eq *Equipment) Apply(data res.EquipmentData) {
	// Remove unequiped items.
	for _, s := range eq.Slots() {
		if s.Item() == nil {
			continue
		}
		found := false
		for _, id := range data.Items {
			slotType := item.Slot(id.Slot)
			if id.ID == s.Item().ID() && id.Serial == s.Item().Serial() &&
				slotType == s.Type() && id.SlotID == s.ID() {
				found = true
				break
			}
		}
		if !found {
			s.SetItem(nil)
		}
	}
	// Equip items.
	for _, itData := range data.Items {
		it := inventoryItem(eq.char.Inventory(), itData.ID, itData.Serial)
		if it == nil {
			log.Err.Printf("character: %s %s: eq: unable to retrieve item from inventory: %s %s",
				eq.char.ID(), eq.char.Serial(), itData.ID, itData.Serial)
			continue
		}
		eqItem, ok := it.(item.Equiper)
		if !ok {
			log.Err.Printf("character: %s %s: eq: not eqipable item: %s %s",
				eq.char.ID(), eq.char.Serial(), it.ID(), it.Serial())
			continue
		}
		// Equip.
		if !eq.char.MeetReqs(eqItem.EquipReqs()...) {
			log.Err.Printf("character: %s %s: eq: equip reqs not meet: %s %s",
				eq.char.ID(), eq.char.Serial(), it.ID(), it.Serial())
			continue
		}
		slot := item.Slot(itData.Slot)
		for _, s := range eq.Slots() {
			if s.Type() != slot || s.ID() != itData.SlotID {
				continue
			}
			s.SetItem(eqItem)
		}
	}
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
			SlotID: s.ID(),
		}
		data.Items = append(data.Items, eqItemData)
	}
	return data
}

// ID returns slot ID.
func (eqSlot *EquipmentSlot) ID() int {
	return eqSlot.id
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

// newEquipmentSlot creates new equipment slot for
// specified slot type.
func (eq *Equipment) newEquipmentSlot(slotType item.Slot) *EquipmentSlot {
	s := new(EquipmentSlot)
	s.slotType = slotType
	// Count existing slots with specified type and set unique ID for new slot.
	slots := make([]*EquipmentSlot, 0)
	for _, s := range eq.Slots() {
		if s.Type() == slotType {
			slots = append(slots, s)
		}
	}
	s.id = len(slots)
	return s
}

// inventoryItem returns an item with specified ID and serial from
// the inventory or any item with specified ID if serial is empty.
func inventoryItem(inv *item.Inventory, id string, serial string) item.Item {
	if len(serial) > 0 {
		return inv.Item(id, serial)
	}
	for _, it := range inv.Items() {
		if it.ID() == id {
			return it
		}
	}
	return nil
}
