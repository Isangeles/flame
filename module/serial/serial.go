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
// This package also stores all registered objects.
package serial

import (
	"fmt"
	"sync"
)

var (
	objects *sync.Map
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
	Reset()
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
	obs, _ := objects.Load(s.ID())
	serialers, _ := obs.([]Serialer)
	// Check whether this is first object with such ID.
	if len(serialers) < 1 {
		if len(s.Serial()) < 1 {
			s.SetSerial("0")
		}
		objects.Store(s.ID(), append(serialers, s))
		return
	}
	// Check whether object already has unique serial value.
	if s.Serial() != "" && unique(serialers, s.Serial()) {
		objects.Store(s.ID(), append(serialers, s))
		return
	}
	// Generate & assign unique serial.
	s.SetSerial(uniqueSerial(serialers))
	// Save to base.
	objects.Store(s.ID(), append(serialers, s))
	return
}

// Object returns object with specified ID and
// serial value or nil if no such object was
// found among registered serial objects.
func Object(id, serial string) Serialer {
	obs, _ := objects.Load(id)
	serialers, _ := obs.([]Serialer)
	for _, s := range serialers {
		if s.ID() == id && s.Serial() == serial {
			return s
		}
	}
	return nil
}

// Reset removes all registered objects from
// base.
func Reset() {
	objects = new(sync.Map)
}

// uinqueSerial generates unique serial value accross
// specified group of objects with serial value.
func uniqueSerial(group []Serialer) string {
	// Choose unique serial value.
	serial := len(group)
	for i := serial; !unique(group, fmt.Sprintf("%d", i)); i++ {
		serial = i
	}
	return fmt.Sprintf("%d", serial)
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
