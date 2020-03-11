/*
 * alignment.go
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

package character

// Type for character alignment.
type Alignment int

const (
	LawfulGood Alignment = iota
	NeutralGood
	ChaoticGood
	LawfulNeutral
	TrueNeutral
	ChaoticNeutral
	LawfulEvil
	NeutralEvil
	ChaoticEvil
)

// String returns text representation of
// alignement.
func (a Alignment) String() string {
	return a.ID()
}

// Id returns alignemnt ID.
func (a Alignment) ID() string {
	switch a {
	case LawfulGood:
		return "ali_law_good"
	case NeutralGood:
		return "ali_neu_good"
	case ChaoticGood:
		return "ali_cha_good"
	case LawfulNeutral:
		return "ali_law_neutral"
	case TrueNeutral:
		return "ali_tru_neutral"
	case ChaoticNeutral:
		return "ali_cha_neutral"
	case LawfulEvil:
		return "ali_law_evil"
	case NeutralEvil:
		return "ali_neu_evil"
	case ChaoticEvil:
		return "alsi_cha_evil"
	default:
		return "ali_unknown"
	}
}
