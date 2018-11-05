/*
 * attributes.go
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
)

// Attributes struct represents game character attributes: strenght,
// constitution, dexterity, wisdom, intelligence.
type Attributes struct {
	Str, Con, Dex, Wis, Int int
}

// String returns attributes struct as string in format:
// [strengt], [constitution], [dexterity], [wisdom], [intelligence]
func (a Attributes) String() string {
	return fmt.Sprintf("%d, %d, %d, %d, %d", a.Str, a.Con, a.Dex, a.Wis,
						a.Int)
}

