/*
 * attributes.go
 *
 * Copyright 2018-2025 Dariusz Sikora <ds@isangeles.dev>
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

	"github.com/isangeles/flame/data/res"
)

const (
	BaseLift   = 10
	BaseSight  = 300.0
	BaseHealth = 50
	BaseMana   = 10
)

// Attributes struct represents game character attributes: strenght,
// constitution, dexterity, wisdom, intelligence.
type Attributes struct {
	Str, Con, Dex, Wis, Int int
	MoveMod                 int64
}

// Lift returns maximal size of inventory based on
// attributes.
func (a *Attributes) Lift() int {
	return BaseLift * (1 + a.Str)
}

// Sight returns maximal sight range based on
// attributes.
func (a *Attributes) Sight() float64 {
	return BaseSight // * float64(1 + a.Wis)
}

// Health returns maximal health based on
// attributes.
func (a *Attributes) Health() int {
	return (BaseHealth * (1 + a.Con)) * (1 + a.Str/2)
}

// Mana returns maximal mana based on attributes.
func (a *Attributes) Mana() int {
	return (BaseMana * (1 + a.Int)) * (1 + a.Wis/2)
}

// Damage returns min and max damage values
// based on attributes.
func (a *Attributes) Damage() (int, int) {
	min := 1 + (10 * a.Str + a.Dex)
	max := 10 + (10 * a.Str + a.Dex)
	return min, max
}

// String returns attributes struct as string in format:
// [strengt], [constitution], [dexterity], [wisdom], [intelligence]
func (a *Attributes) String() string {
	return fmt.Sprintf("%d, %d, %d, %d, %d",
		a.Str, a.Con, a.Dex, a.Wis, a.Int)
}

// Apply applies specified data on character attributes.
func (a *Attributes) Apply(data res.AttributesData) {
	a.Str = data.Str
	a.Con = data.Con
	a.Dex = data.Dex
	a.Int = data.Int
	a.Wis = data.Wis
}

// Data returns data resource for attributes.
func (a *Attributes) Data() res.AttributesData {
	data := res.AttributesData{
		Str: a.Str,
		Con: a.Con,
		Dex: a.Dex,
		Wis: a.Wis,
		Int: a.Int,
	}
	return data
}
