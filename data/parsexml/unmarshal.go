/*
 * unmarshal.go
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
	"strconv"
	"strings"

	"github.com/isangeles/flame/module/character"
	"github.com/isangeles/flame/module/item"
	"github.com/isangeles/flame/module/objects"
)

var (
	PositionSep = "x"
)

// UnmarshalPosition parses specified position
// attribute(X[sep]Y) to XY position.
func UnmarshalPosition(attr string) (float64, float64, error) {
	posValues := strings.Split(attr, PositionSep)
	if len(posValues) < 2 {
		return 0, 0, fmt.Errorf("invalid position attribute format")
	}
	x, err := strconv.ParseFloat(posValues[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to parse x position: %v", err)
	}
	y, err := strconv.ParseFloat(posValues[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to parse y position: %v", err)
	}
	return x, y, nil
}

// UnmarshalAttributes parses specified attributes from XML doc.
func UnmarshalAttributes(attributesAttr string) (character.Attributes, error) {
	stats := strings.Split(attributesAttr, ";")
	if len(stats) < 5 {
		return character.Attributes{},
			fmt.Errorf("unable to parse attributes text: %s", attributesAttr)
	}
	str, err := strconv.Atoi(stats[0])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("unable to parse str attribute: %s", stats[0])
	}
	con, err := strconv.Atoi(stats[1])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("unable to parse con attribute: %s", stats[1])
	}
	dex, err := strconv.Atoi(stats[2])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("unable to parse dex attribute: %s", stats[2])
	}
	inte, err := strconv.Atoi(stats[3])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("unable to parse int attribute: %s", stats[3])
	}
	wis, err := strconv.Atoi(stats[4])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("unable to parse wis attribute: %s", stats[4])
	}
	return character.Attributes{str, con, dex, inte, wis}, nil
}

// UnmarshalItemSlots parses specified slots attribute from XML doc
// to item slot types.
func UnmarshalItemSlots(slotsAttr string) ([]item.Slot, error) {
	slots := make([]item.Slot, 0)
	attrs := strings.Fields(slotsAttr)
	for _, attr := range attrs {
		switch attr {
		case item.Hand.ID():
			slots = append(slots, item.Hand)
		case item.Chest.ID():
			slots = append(slots, item.Chest)
		default: // all slots IDs must be 'parsable'
			return nil, fmt.Errorf("unable to parse slot type: %s", slotsAttr)
		}
	}
	return slots, nil
}

// UnmarshalEqSlot parses specified slot XML attribute to
// equipment slot type.
func UnmarshalEqSlot(slot string) (character.EquipmentSlotType, error) {
	switch slot {
	case "right_hand":
		return character.Hand_right, nil
	case "chest":
		return character.Chest, nil
	default:
		return -1, fmt.Errorf("unsupported eq slot type value: %s", slot)
	}
}

// UnmarshalElementType parses specified string to element type.
func UnmarshalElementType(s string) (objects.Element, error) {
	switch s {
	case "normal":
		return objects.Element_none, nil
	case "fire":
		return objects.Element_fire, nil
	case "frost":
		return objects.Element_frost, nil
	case "nature":
		return objects.Element_nature, nil
	default:
		return -1, fmt.Errorf("unsupported hit type value: %s", s)
	}
}
