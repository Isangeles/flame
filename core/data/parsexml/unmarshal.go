/*
 * unmarshal.go
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
	"strings"
	"strconv"

	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/item"
)

var (
	POSITION_SEP = "x"
)

// UnmarshalPosition parses specified position
// attribute(X[sep]Y) to XY position.
func UnmarshalPosition(attr string) (float64, float64, error) {
	posValues := strings.Split(attr, POSITION_SEP)
	if len(posValues) < 2 {
		return 0, 0, fmt.Errorf("invalid position attribute format")
	}
	x, err := strconv.ParseFloat(posValues[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("fail_to_parse_x_position:%v",
			err)
	}
	y, err := strconv.ParseFloat(posValues[1], 64)
	if err != nil {
		return 0, 0,fmt.Errorf("fail_to_parse_y_position:%v",
			err)
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
	default:
		return -1, fmt.Errorf("fail_to_parse_gender:%s", genderAttr)
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
	default:
		return character.Race_unknown, nil//fmt.Errorf("fail to parse race:%s", raceAttr)
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
		return -1, fmt.Errorf("fail_to_parse_attitude:%s", attitudeAttr)
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
		return -1, fmt.Errorf("fail_to_parse_alignment:%s", aliAttr)
	}
}

// UnmarshalAttributes parses specified attributes from XML doc.
func UnmarshalAttributes(attributesAttr string) (character.Attributes, error) {
	stats := strings.Split(attributesAttr, ";")
	if len(stats) < 5 {
		return character.Attributes{},
		fmt.Errorf("fail to parse attributes text:%s", attributesAttr)
	}
	str, err := strconv.Atoi(stats[0])
	if err != nil {
		return character.Attributes{},
		fmt.Errorf("fail to parse str attribute:%s", stats[0])
	}
	con, err := strconv.Atoi(stats[1])
	if err != nil {
		return character.Attributes{},
		fmt.Errorf("fail to parse con attribute:%s", stats[1])
	}
	dex, err := strconv.Atoi(stats[2])
	if err != nil {
		return character.Attributes{},
		fmt.Errorf("fail to parse dex attribute:%s", stats[2])
	}
	inte, err := strconv.Atoi(stats[3])
	if err != nil {
		return character.Attributes{},
		fmt.Errorf("fail to parse int attribute:%s", stats[3])
	}
	wis, err := strconv.Atoi(stats[4])
	if err != nil {
		return character.Attributes{},
		fmt.Errorf("fail to parse wis attribute:%s", stats[4])
	}
	return character.Attributes{str, con, dex, inte, wis}, nil
}

// UnmarshalItemSlots parses specified slots attribute from XML doc
// to item slot types.
func UnmarshalItemSlots(slotsAttr string) ([]item.Slot, error) {
	slots := make([]item.Slot, 0)
	attrs := strings.Split(slotsAttr, " ")
	for _, attr := range attrs {
		switch attr {
		case "one_hand":
			slots = append(slots, item.Hand)
		default: // all slots IDs must be 'parsable'
			return nil, fmt.Errorf("fail_to_parse_slot_type:%s", slotsAttr)
		}
	}
	return slots, nil
}
