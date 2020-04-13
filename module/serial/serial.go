/*
 * serial.go
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

// Package for generating unique serial values
// for game objects.
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

// Register assigns and registers unique serial value to
// specified object among all previously registered objects with
// same ID, or only registers serial if object already has unique
// serial value.
// Assigns new serial if specified object has serial
// value but its not unique among previously
// registered objects.
func Register(s Serialer) {
	// Get all objects with same ID.
	serialers := base[s.ID()]
	// Check whether this is first object with such ID.
	if len(serialers) < 1 {
		if len(s.Serial()) < 1 {
			s.SetSerial("0")
		}
		regSerial(s)
		return
	}
	// Check whether object already has unique serial value.
	if s.Serial() != "" && unique(serialers, s.Serial()) {
		regSerial(s)
		return
	}
	// Generate & assign unique serial.
	s.SetSerial(uniqueSerial(serialers))
	// Save to base.
	base[s.ID()] = append(serialers, s)
	return
}

// Object returns object with specified ID and
// serial value or nil if no such object was
// found among registered serial objects.
func Object(id, serial string) Serialer {
	objects := base[id]
	for _, o := range objects {
		if o.Serial() == serial {
			return o
		}
	}
	return nil
}

// Reset removes all registered objects from
// base.
func Reset() {
	base = make(map[string][]Serialer)
}

// uinqueSerial generates unique serial value accross
// specified group of objects with serial value.
func uniqueSerial(group []Serialer) string {
	// Choose unique serial value.
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
