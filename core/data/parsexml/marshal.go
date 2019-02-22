/*
 * marshal.go
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

package parsexml

import (
	"fmt"

	"github.com/isangeles/flame/core/module/object/character"
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
	case character.Lawful_good:
		return "lawful_good"
	case character.Neutral_good:
		return "neutral_good"
	case character.Chaotic_good:
		return "chaotic_good"
	case character.Lawful_neutral:
		return "lawful_neutral"
	case character.True_neutral:
		return "true_neutral"
	case character.Chaotic_neutral:
		return "chaotic_neutral"
	case character.Lawful_evil:
		return "lawful_evil"
	case character.Neutral_evil:
		return "neutral_evil"
	case character.Chaotic_evil:
		return "chaotic_evil"
	default:
		return "unknown"
	}
}

// marshalAttributes parses attributes to XML node value.
func marshalAttributes(attrs character.Attributes) string {
	return fmt.Sprintf("%d;%d;%d;%d;%d;", attrs.Str,
		attrs.Con, attrs.Dex, attrs.Wis, attrs.Int)
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
		return "right hand"
	case character.Hand_left:
		return "left hand"
	case character.Finger_right:
		return "finger right"
	case character.Finger_left:
		return "finger left"
	case character.Legs:
		return "legs"		
	case character.Feet:
		return "feet"
	default:
		return "unknow"
	}
}
