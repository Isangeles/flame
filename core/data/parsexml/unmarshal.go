/*
 * unmarshal.go
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

package parsexml

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/isangeles/flame/core/module/character"
	"github.com/isangeles/flame/core/module/item"
	"github.com/isangeles/flame/core/module/objects"
	"github.com/isangeles/flame/core/module/skill"
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
		return 0, 0, fmt.Errorf("fail to parse x position: %v", err)
	}
	y, err := strconv.ParseFloat(posValues[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("fail to parse y position: %v", err)
	}
	return x, y, nil
}

// UnmarshalGender parses specified gender XML attribute,
func UnmarshalGender(genderAttr string) (character.Gender, error) {
	switch genderAttr {
	case "male":
		return character.Male, nil
	case "female":
		return character.Female, nil
	case "unknwon", "*":
		return character.UnknownGender, nil
	default:
		return -1, fmt.Errorf("unsupported gender value: %s", genderAttr)
	}
}

// UnmarshalRace parses specified race XML attribute.
func UnmarshalRace(raceAttr string) (character.Race, error) {
	switch raceAttr {
	case "human":
		return character.Human, nil
	case "elf":
		return character.Elf, nil
	case "dwarf":
		return character.Dwarf, nil
	case "gnome":
		return character.Gnome, nil
	case "wolf":
		return character.Wolf, nil
	case "goblin":
		return character.Goblin, nil
	case "unknown", "*":
		return character.UnknownRace, nil
	default:
		return -1, fmt.Errorf("unsupported race value: %s", raceAttr)
	}
}

// UnmarshalAttitude parses specified attitude XML attribute.
func UnmarshalAttitude(attitudeAttr string) (character.Attitude, error) {
	switch attitudeAttr {
	case "friendly":
		return character.Friendly, nil
	case "neutral":
		return character.Neutral, nil
	case "hostile":
		return character.Hostile, nil
	default:
		return -1, fmt.Errorf("unsupported attitude value: %s", attitudeAttr)
	}
}

// UnmarshalAlignment parses specified alignemnt XML attribute.
func UnmarshalAlignment(aliAttr string) (character.Alignment, error) {
	switch aliAttr {
	case "lawful_good":
		return character.Lawful_good, nil
	case "neutral_good":
		return character.Neutral_good, nil
	case "chaotic_good":
		return character.Chaotic_good, nil
	case "lawful_neutral":
		return character.Lawful_neutral, nil
	case "chaotic_neutral":
		return character.Chaotic_neutral, nil
	case "true_neutral":
		return character.True_neutral, nil
	case "lawful_evil":
		return character.Lawful_evil, nil
	case "neutral_evil":
		return character.Neutral_evil, nil
	case "chaotic_evil":
		return character.Chaotic_evil, nil
	default:
		return -1, fmt.Errorf("unsupported alignment value: %s", aliAttr)
	}
}

// UnmarshalAttributes parses specified attributes from XML doc.
func UnmarshalAttributes(attributesAttr string) (character.Attributes, error) {
	stats := strings.Split(attributesAttr, ";")
	if len(stats) < 5 {
		return character.Attributes{},
			fmt.Errorf("fail to parse attributes text: %s", attributesAttr)
	}
	str, err := strconv.Atoi(stats[0])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("fail to parse str attribute: %s", stats[0])
	}
	con, err := strconv.Atoi(stats[1])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("fail to parse con attribute: %s", stats[1])
	}
	dex, err := strconv.Atoi(stats[2])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("fail to parse dex attribute: %s", stats[2])
	}
	inte, err := strconv.Atoi(stats[3])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("fail to parse int attribute: %s", stats[3])
	}
	wis, err := strconv.Atoi(stats[4])
	if err != nil {
		return character.Attributes{},
			fmt.Errorf("fail to parse wis attribute: %s", stats[4])
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
			return nil, fmt.Errorf("fail to parse slot type: %s", slotsAttr)
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

// UnmarshalSkillRange parses specified skill range XML attribute
// to skill range type.
func UnmarshalSkillRange(rg string) (skill.Range, error) {
	switch rg {
	case "touch":
		return skill.Range_touch, nil
	case "close":
		return skill.Range_close, nil
	case "far":
		return skill.Range_far, nil
	case "huge":
		return skill.Range_huge, nil
	default:
		return -1, fmt.Errorf("unsupported skill range type value: s", rg)
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
