/*
 * serial.go
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

package module

import (
	"fmt"
)

// Interface for all game objects with unique
// serial ID.
type Serializer interface {
	ID() string
	Serial() string
	SerialID() string
	SetSerial(serial string)
	HasSerial() bool
}

// FullSerial retruns full serial ID string
// with specified ID and serial value.
func FullSerial(id, serial string) string {
	return fmt.Sprintf("%s_%s", id, serial)
}

// uinqueSerial generates unique serial value accross
// specified group of objects with serial value.
func uniqueSerial(group []Serializer) string {
	// Choose serial value unique accross chars.
	serial := fmt.Sprintf("%d", len(group))
	// Ensure serial value uniqueness.
	for i := len(group); !isSerialUnique(group, serial); i++ {
		serial = fmt.Sprintf("%d", i)
	}
	return serial
}

// isSerialUnique checks whether specified serial value
// is unique across specified objects with serial value.
func isSerialUnique(group []Serializer, serial string) bool {
	for _, ob := range group {
		if ob.Serial() == serial {
			return false
		}
	}
	return true
}
