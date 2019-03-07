/*
 * serial.go
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

package serial

import (
	"fmt"
)

var (
	base map[string][]Serialer
)

// Interface for all game objects with
// unique serial ID.
type Serialer interface {
	ID() string
	Serial() string
	SetSerial(serial string)
}

// On init.
func init() {
	base = make(map[string][]Serialer)
}

// AssignSerial assigns and registers unique serial value to
// specified object among all previously registered objects with
// same ID, only registers serial if object already has uniqe
// serial value.
// Assigns new serial if specified object has serial
// value but its not unique among previously
// registered objects.
func AssignSerial(s Serialer) {
	// Get all objects with same ID.
	serialers := base[s.ID()]
	// Check whether this is first object with such ID.
	if serialers == nil {
		if s.Serial() == "" {
			s.SetSerial("0")
		}
		regSerial(s)
		return
	}
	// Check whether object has unique serial value already.
	if s.Serial() != "" && unique(serialers, s.Serial()) {
		regSerial(s)
		return
	}
	// Generate & assign unique serial.
	s.SetSerial(UniqueSerial(serialers))
	// Save to base.
	base[s.ID()] = append(serialers, s)
	return
}

// FullSerial retruns full serial ID string
// with specified ID and serial value.
func FullSerial(id, serial string) string {
	return fmt.Sprintf("%s_%s", id, serial)
}

// UinqueSerial generates unique serial value accross
// specified group of objects with serial value.
func UniqueSerial(group []Serialer) string {
	// Choose serial value unique accross chars.
	serial := fmt.Sprintf("%d", len(group))
	// Ensure serial value uniqueness.
	for i := len(group); !unique(group, serial); i++ {
		serial = fmt.Sprintf("%d", i)
	}
	return serial
}

// serialUnique checks whether specified serial value
// is unique across specified objects with serial value.
func unique(group []Serialer, serial string) bool {
	for _, ob := range group {
		if ob.Serial() == serial {
			return false
		}
	}
	return true
}

// regSerial registers specified object
// in base with object with serial values.
func regSerial(s Serialer) {
	serialers := base[s.ID()]
	if serialers == nil {
		serialers = make([]Serialer, 0)
	}
	base[s.ID()] = append(serialers, s)
}
