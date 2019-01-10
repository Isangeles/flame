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
	head        item.Equiper
	neck        item.Equiper
	chest       item.Equiper
	handRight   item.Equiper
	handLeft    item.Equiper
	fingerRight item.Equiper
	fingerLeft  item.Equiper
	legs        item.Equiper
	feets       item.Equiper
}

// newEquipment creates new equipment for
// specified character.
func newEquipment(char *Character) *Equipment {
	eq := new(Equipment)
	eq.char = char
	return eq
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
	eq.handRight = it
	return nil
}

// Items returns slice with all equiped items.
func (eq *Equipment) Items() []item.Equiper {
	its := make([]item.Equiper, 0)
	if eq.handRight != nil {
		its = append(its, eq.handRight)	
	}
	if eq.handLeft != nil {
		its = append(its, eq.handLeft)
	}
	return its
}
