/*
 * alignment.go
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

// Type for character alignment.
type Alignment int

const (
	Lawful_good Alignment = iota
	Neutral_good
	Chaotic_good
	Lawful_neutral
	True_neutral
	Chaotic_neutral
	Lawful_evil
	Neutral_evil
	Chaotic_evil
)

// String returns text representation of
// alignement.
func (a Alignment) String() string {
	return a.ID()
}

// Id returns alignemnt ID.
func (a Alignment) ID() string {
	switch a {
	case Lawful_good:
		return "ali_law_good"
	case Neutral_good:
		return "ali_neu_good"
	case Chaotic_good:
		return "ali_cha_good"
	case Lawful_neutral:
		return "ali_law_neutral"
	case True_neutral:
		return "ali_tru_neutral"
	case Chaotic_neutral:
		return "ali_cha_neutral"
	case Lawful_evil:
		return "ali_law_evil"
	case Neutral_evil:
		return "ali_neu_evil"
	case Chaotic_evil:
		return "alsi_cha_evil"
	default:
		return "ali_unknown"
	}
}
