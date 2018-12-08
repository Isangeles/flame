/*
 * area.go
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

package scenario

import (
	"github.com/isangeles/flame/core/module/object/character"
)

// Area struct represents game world area.
type Area struct {
	id    string
	chars []*character.Character
}

// NewArea returns new instace of area.
func NewArea(id string) (*Area) {
	a := new(Area)
	a.id = id
	return a
}

// Id returns area ID.
func (a *Area) ID() string {
	return a.id
}

// AddCharacter adds specified character to area.
func (a *Area) AddCharacter(c *character.Character) {
	a.chars = append(a.chars, c)
}

// Chracters returns list with characters in area.
func (a *Area) Characters() []*character.Character {
	return a.chars
}

// ContainsCharacter checks whether area
// contains specified character.
func (a *Area) ContainsCharacter(char *character.Character) bool {
	for _, c := range a.chars {
		if char == c {
			return true
		}
	}
	return false
}
