/*
 * item.go
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

package item

import (
	"github.com/isangeles/flame/core/module/req"
)

// Interface for items.
type Item interface {
	ID() string
	Serial() string
	SerialID() string
	SetSerial(serial string)
	Name() string
	Value() int
	Level() int
}

// Interface for 'equipable' items.
type Equiper interface {
	EquipReqs() []req.Requirement
	Slots()     []Slot
}

// Type for slot type occupated by item.
type Slot int

const (
	None Slot = iota
	Head
	Neck
	Chest
	Hand
	Finger
	Legs
	Feets
)
