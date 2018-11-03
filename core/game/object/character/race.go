/*
 * race.go
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

// Type for characters races
type Race int

const (
	HUMAN Race = iota
	ELF
	DWARF
	GNOME
	WOLF
	GOBLIN
	UNKNOWN
	//...
)

// Id returns race ID.
func (r Race) Id() string {
	switch r {
	case HUMAN:
		return "race_human"
	case ELF:
		return "race_elf"
	case DWARF:
		return "race_dwarf"
	case GNOME:
		return "race_gnome"
	case WOLF:
		return "race_wolf"
	case GOBLIN:
		return "race_goblin"
	default:
		return "race_unknown"
	}
}
