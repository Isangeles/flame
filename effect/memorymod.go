/*
 * memorymod.go
 *
 * Copyright 2021 Dariusz Sikora <dev@isangeles.pl>
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

package effect

import (
	"github.com/isangeles/flame/data/res"
)

// Struct for memory modifier.
type MemoryMod struct {
	attitude string
}

// NewMemoryModifer creates new memory modifer.
func NewMemoryMod(data res.MemoryModData) *MemoryMod {
	mm := MemoryMod{data.Attitude}
	return &mm
}

// Attitude returns ID of character attitude to set.
func (mm *MemoryMod) Attitude() string {
	return mm.attitude
}

// Data returns data resource for modifier.
func (mm *MemoryMod) Data() res.MemoryModData {
	return res.MemoryModData{mm.attitude}
}
