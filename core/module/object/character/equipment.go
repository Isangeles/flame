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
	char      *Character
	handRight item.Equiper
	handLeft  item.Equiper
}

// newEquipment creates new equipment for
// specified character.
func newEquipment(char *Character) *Equipment {
	eq := new(Equipment)
	eq.char = char
	return eq
}

// EquipHandRight equips specified 'equipable' item,
// returns error if equip fail(e.q. equip reqs not meet).
func (eq *Equipment) EquipHandRight(item item.Equiper) error {
	if eq.char.MeetReqs(item.EquipReqs()) {
		eq.handRight = item
		return nil
	}
	return fmt.Errorf("reqs_not_meet")
}
