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
		sex, err := parseGenderAttr(charXML.Gender)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_gender:%v",
			xmlPath, err)
			continue
		}
		race, err := parseRaceAttr(charXML.Race)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_race:%v",
			xmlPath, err)
			continue
		}
		attitude, err := parseAttitudeAttr(charXML.Attitude)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_attitude:%v",
			xmlPath, err)
			continue
		}
		guild := character.NewGuild(charXML.Guild) // TODO: search and assign guild
		attributes, err := parseAttributesAttr(charXML.Stats)
		if err != nil {
			log.Err.Printf("xml:%s:parse:fail_to_parse_char_attributes:%v",
			xmlPath, err)
			continue
		}
		alignment, err := parseAlignmentAttr(charXML.Alignment)
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
	xmlChar.Id = char.Id()
	xmlChar.Name = char.Name()
	xmlChar.Level = fmt.Sprintf("%d", char.Level())
	xmlCharBase.Characters = append(xmlCharBase.Characters, *xmlChar)
	out, err := xml.Marshal(xmlCharBase)
	if err != nil {
		return []byte{}, fmt.Errorf("fail_to_marshal_char:%v", err)
	}
	return out, nil
}

// parseGenderAttr parses specified gender XML attribute,
func parseGenderAttr(genderAttr string) (character.Gender, error) {
	switch genderAttr {
	case "male":
		return character.MALE, nil
	case "female":
		return character.FEMALE, nil
	default:
		return -1, fmt.Errorf("fail to parse gender:%s", genderAttr)
	}
}

// parseRaceAttr parses specified race XML attribute.
func parseRaceAttr(raceAttr string) (character.Race, error) {
	switch raceAttr {
	case "human":
		return character.HUMAN, nil
	// TODO: handle all races.
	default:
		return -1, fmt.Errorf("fail to parse race:%s", raceAttr)
	}
}

// parseAttitideAttr parses specified attitude XML attribute.
func parseAttitudeAttr(attitudeAttr string) (character.Attitude, error) {
	switch attitudeAttr {
	case "friendly":
		return character.Friendly, nil
	default:
		return -1, fmt.Errorf("fail to parse attitude:%s", attitudeAttr)
	}
}

// parse AttributesAttr parses specified attributes XML attribute.
func parseAttributesAttr(attributesAttr string) (character.Attributes, error) {
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

// parseAlignmentAttr parses specified alignemnt XML attribute.
func parseAlignmentAttr(aliAttr string) (character.Alignment, error) {
	switch aliAttr {
	case "lawful_good":
		return character.Lawful_good, nil
	default:
		return -1, fmt.Errorf("fail to parse alignment:%s", aliAttr)
	}
}


