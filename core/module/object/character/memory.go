/*
 * memory.go
 * 
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/core/module/effect"
)

// Struct for saved data about target.
type TargetMemory struct {
	Target   effect.Target
	Attitude Attitude
	Talked   bool
	Killed   bool
}

// Memory returns character tergets memory.
func (c *Character) Memory() (mem []*TargetMemory) {
	for _, am := range c.memory {
		mem = append(mem, am)
	}
	return
}

// MemorizeTarget saves specified target memory.
func (c *Character) MemorizeTarget(mem *TargetMemory) {
	c.memory[mem.Target.ID() + mem.Target.Serial()] = mem
}
