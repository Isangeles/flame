/*
 * attributes.go
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

package character

import (
	"fmt"
)

const (
	base_lift = 10
	base_sight = 300.0
	base_health = 100
	base_mana = 10
)

// Attributes struct represents game character attributes: strenght,
// constitution, dexterity, wisdom, intelligence.
type Attributes struct {
	Str, Con, Dex, Wis, Int int
}

// Lift returns maximal size of inventory based on
// attributes.
func (a Attributes) Lift() int {
	return base_lift * a.Str
}

// Sight returns maximal sight range based on
// attributes.
func (a Attributes) Sight() float64 {
	return base_sight * float64(a.Wis)
}

// Health returns maximal health based on
// attributes.
func (a Attributes) Health() int {
	return (base_health * a.Con) * a.Str/2
}

// Mana returns maximal mana based on attributes.
func (a Attributes) Mana() int {
	return (base_mana * a.Int) * a.Wis/2
}

// String returns attributes struct as string in format:
// [strengt], [constitution], [dexterity], [wisdom], [intelligence]
func (a Attributes) String() string {
	return fmt.Sprintf("%d, %d, %d, %d, %d",
		a.Str, a.Con, a.Dex, a.Wis, a.Int)
}

