/*
 * marshal.go
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

package parsexml

import (
	"fmt"

	"github.com/isangeles/flame/module/character"
)

// marshalGender parses specified gender to gender XML
// attribute value.
func marshalGender(sex character.Gender) string {
	switch sex {
	case character.Male:
		return "male"
	case character.Female:
		return "female"
	default:
		return "male"
	}
}

// marshalRace parses specified race to race XML
// attribute value.
func marshalRace(race character.Race) string {
	switch race {
	case character.Human:
		return "human"
	case character.Elf:
		return "elf"
	case character.Dwarf:
		return "dwarf"
	case character.Gnome:
		return "gnome"
	case character.Wolf:
		return "wolf"
	case character.Goblin:
		return "goblin"
	default:
		return "unknown"
	}
}

// marshalAttitude parses specified attitude to attitude XML
// attribute value.
func marshalAttitude(attitude character.Attitude) string {
	switch attitude {
	case character.Friendly:
		return "friendly"
	case character.Neutral:
		return "neutral"
	case character.Hostile:
		return "hostile"
	default:
		return "unknown"
	}
}

// marshalAlignment parses specified alignment to alignment XML
// attribute value.
func marshalAlignment(alignment character.Alignment) string {
	switch alignment {
	case character.LawfulGood:
		return "lawful_good"
	case character.NeutralGood:
		return "neutral_good"
	case character.ChaoticGood:
		return "chaotic_good"
	case character.LawfulNeutral:
		return "lawful_neutral"
	case character.TrueNeutral:
		return "true_neutral"
	case character.ChaoticNeutral:
		return "chaotic_neutral"
	case character.LawfulEvil:
		return "lawful_evil"
	case character.NeutralEvil:
		return "neutral_evil"
	case character.ChaoticEvil:
		return "chaotic_evil"
	default:
		return "unknown"
	}
}

// marshalEqSlot parses specified equipment slot to XML
// attribute value.
func MarshalEqSlot(eqSlot *character.EquipmentSlot) string {
	switch eqSlot.Type() {
	case character.Head:
		return "head"
	case character.Neck:
		return "neck"
	case character.Chest:
		return "chest"
	case character.Hand_right:
		return "right_hand"
	case character.Hand_left:
		return "left_hand"
	case character.Finger_right:
		return "finger_right"
	case character.Finger_left:
		return "finger_left"
	case character.Legs:
		return "legs"		
	case character.Feet:
		return "feet"
	default:
		return "unknow"
	}
}

// MarshalPosition parses specified XY position
// to string.
func MarshalPosition(x, y float64) string {
	return fmt.Sprintf("%fx%f", x, y)
}
