/*
 * characterparser.go
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

package parse

import (
	"encoding/xml"
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/log"
)

// Struct for XML characters base.
type XMLCharactersBase struct {
	XMLName    xml.Name       `xml:"base"`
	Characters []XMLCharacter `xml:"char"`
}

// Struct for XML character node. 
type XMLCharacter struct {
	XMLName   xml.Name `xml:"char"`
	Id        string   `xml:"id,attr"`
	Name      string   `xml:"name,attr"`
	Gender    string   `xml:"gender,attr"`
	Race      string   `xml:"race,attr"`
	Attitude  string   `xml:"attitude,attr"`
	Alignment string   `xml:"alignment,attr"`
	Guild     string   `xml:"guild,attr"`
	Level     string   `xml:"level,attr"`
	Stats     string   `xml:"stats,value"`
}

// UnmarshalCharactersBaseXML parses characters base from XML file
// in specified path.
func UnmarshalCharactersBaseXML(xmlPath string) ([]*character.Character, error) {
	chars := make([]*character.Character, 0)
	doc, err := os.Open(xmlPath)
	if err != nil {
		return chars, fmt.Errorf("fail_to_find_chars_base_file:%v", err)
	}
	defer doc.Close()

	data, _ := ioutil.ReadAll(doc)
	xmlCharsBase := new(XMLCharactersBase)
	err = xml.Unmarshal(data, xmlCharsBase)
	if err != nil {
		return chars, fmt.Errorf("xml:%s:fail_to_unmarshal_xml:%v",
			xmlPath, err)
	}
	for _, charXML := range xmlCharsBase.Characters {
		id := charXML.Id
		name := charXML.Name
		level, err := strconv.Atoi(charXML.Level)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_level:%v",
			xmlPath, err)
			continue
		}
		sex, err := attrToGender(charXML.Gender)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_gender:%v",
			xmlPath, err)
			continue
		}
		race, err := attrToRace(charXML.Race)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_race:%v",
			xmlPath, err)
			continue
		}
		attitude, err := attrToAttitude(charXML.Attitude)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_attitude:%v",
			xmlPath, err)
			continue
		}
		guild := character.NewGuild(charXML.Guild) // TODO: search and assign guild
		attributes, err := unmarshalAttributes(charXML.Stats)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_attributes:%v",
			xmlPath, err)
			continue
		}
		alignment, err := attrToAlignment(charXML.Alignment)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_alignment:%v",
			xmlPath, err)
			continue
		}
		char := character.NewCharacter(id, name, level, sex, race,
			attitude, guild, attributes, alignment)
		chars = append(chars, char)
	}
	return chars, nil
}

// MarshalCharacter parses game character to XML node
// representation.
func MarshalCharacterXML(char *character.Character) ([]byte, error) {
	xmlCharBase := new(XMLCharactersBase)
	xmlChar := new(XMLCharacter)
	xmlChar.Id = char.ID()
	xmlChar.Name = char.Name()
	xmlChar.Level = fmt.Sprintf("%d", char.Level())
	xmlChar.Gender = genderToAttr(char.Gender())
	xmlChar.Race = raceToAttr(char.Race())
	xmlChar.Attitude = attitudeToAttr(char.Attitude())
	xmlChar.Alignment = alignmentToAttr(char.Alignment())
	xmlChar.Stats = marshalAttributes(char.Attributes())
	xmlCharBase.Characters = append(xmlCharBase.Characters, *xmlChar)
	out, err := xml.Marshal(xmlCharBase)
	if err != nil {
		return []byte{}, fmt.Errorf("fail_to_marshal_char:%v", err)
	}
	return out, nil
}

// attrToGender parses specified gender XML attribute,
func attrToGender(genderAttr string) (character.Gender, error) {
	switch genderAttr {
	case "male":
		return character.Male, nil
	case "female":
		return character.Female, nil
	default:
		return -1, fmt.Errorf("fail to parse gender:%s", genderAttr)
	}
}

// attrToRace parses specified race XML attribute.
func attrToRace(raceAttr string) (character.Race, error) {
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

// attrToAttitude parses specified attitude XML attribute.
func attrToAttitude(attitudeAttr string) (character.Attitude, error) {
	switch attitudeAttr {
	case "friendly":
		return character.Friendly, nil
	case "neutral":
		return character.Neutral, nil
	case "hostile":
		return character.Hostile, nil
	default:
		return -1, fmt.Errorf("fail to parse attitude:%s", attitudeAttr)
	}
}

// attrToAlignment parses specified alignemnt XML attribute.
func attrToAlignment(aliAttr string) (character.Alignment, error) {
	switch aliAttr {
	case "lawful_good":
		return character.Lawful_good, nil
	case "neutral_good":
		return character.Neutral_good, nil
	case "chaotic_good":
		return character.Chaotic_good, nil
	case "lawful_neutral":
		return character.Lawful_neutral, nil
	case "true_neutral":
		return character.True_neutral, nil
	case "lawful_evil":
		return character.Lawful_evil, nil
	case "neutral_evil":
		return character.Neutral_evil, nil
	case "chaotic_evil":
		return character.Chaotic_evil, nil
	default:
		return -1, fmt.Errorf("fail to parse alignment:%s", aliAttr)
	}
}

// genderToAttr parses specified gender to gender XML
// attribute value.
func genderToAttr(sex character.Gender) string {
	switch sex {
	case character.Male:
		return "male"
	case character.Female:
		return "female"
	default:
		return "male"
	}
}

// raceToAttr parses specified race to race XML
// attribute value.
func raceToAttr(race character.Race) string {
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

// attitudeToAttr parses specified attitude to attitude XML
// attribute value.
func attitudeToAttr(attitude character.Attitude) string {
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

// alingmentToAttr parses specified alignment to alignment XML
// attribute value.
func alignmentToAttr(alignment character.Alignment) string {
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

// unmarshalAttributes parses specified attributes from XML doc.
func unmarshalAttributes(attributesAttr string) (character.Attributes, error) {
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

// marshalAttributes parses attributes to XML node value.
func marshalAttributes(attrs character.Attributes) string {
	return fmt.Sprintf("%d;%d;%d;%d;%d;", attrs.Str,
		attrs.Con, attrs.Dex, attrs.Wis, attrs.Int)
}


