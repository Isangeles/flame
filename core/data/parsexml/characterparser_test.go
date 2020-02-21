/*
 * characterparser_test.go
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
	"strings"
	"testing"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/character"
)

// Test for unmarshaling characers data.
func TestUnmarshalCharacters(t *testing.T) {
	xmlChars := `<characters>
        <char id="char"
                gender="male"
                race="human"
                attitude="neutral"
                alignment="chaotic_good"
                guild=""
                level="1">
        <attributes strenght="4"
                constitution="4"
                dexterity="3"
                intelligence="1"
                wisdom="0"/>
        <dialogs>
                <dialog id="dEugene1"/>
        </dialogs>
        <skills>
                <skill id="melee1"/>
        </skills>
        <inventory>
                <item id="imWater1"
                        trade="true"
                        trade-value="5"/>
                <item id="iwIronSword"/>
        </inventory>
        <trainings>
                <attrs-train str="1">
                        <reqs>
                                <currency-req amount="100"
                                        charge="true"/>
                        </reqs>
                </attrs-train>
        </trainings>
        </char>
</characters>`
	data, err := UnmarshalCharacters(strings.NewReader(xmlChars))
	if err != nil {
		t.Errorf("Unable to unmarshal XML string: %v", err)
		return
	}
	if len(data) != 1 {
		t.Errorf("Unmarshaled data is inavalid: len: %d != 1", len(data))
	}
	char := data[0]
	if char.BasicData.ID != "char" {
		t.Errorf("Unmarshaled data is inavalid: ID: %s != char", char.BasicData.ID)	
	}
}

// Test for marshaling character data.
func TestMarshalCharacter(t *testing.T) {
	var data res.CharacterData
	data.BasicData = res.CharacterBasicData{
		ID:        "char",
		Name:      "charName",
		AI:        true,
		Level:     2,
		Sex:       1,
		Race:      1,
		Attitude:  1,
		Guild:     "guildID",
		Alignment: 1,
		Str:       2,
		Con:       3,
		Dex:       4,
		Int:       5,
		Wis:       6,
	}
	char := character.New(data)
	xmlChar, err := MarshalCharacter(char)
	if err != nil {
		t.Errorf("Unable to marshal character: %v", err)
	}
	if !strings.Contains(xmlChar, "id=\"char\"") {
		t.Errorf("Marshaled data is invalid: ID: %s", xmlChar)
	}
	if !strings.Contains(xmlChar, "race=\"elf\"") {
		t.Errorf("Marshaled data is invalid: race: %s", xmlChar)
	}
	if !strings.Contains(xmlChar, "gender=\"female\"") {
		t.Errorf("Marshaled data is invalid: gender: %s", xmlChar)
	}
	if !strings.Contains(xmlChar, "attitude=\"neutral\"") {
		t.Errorf("Marshaled data is invalid: attitude: %s", xmlChar)
	}
	if !strings.Contains(xmlChar, "alignment=\"neutral_good\"") {
		t.Errorf("Marshaled data is invalid: alignment: %s", xmlChar)
	}
}

// Test for marshaling characters data.
func TestMarshalCharacters(t *testing.T) {
	var data res.CharacterData
	data.BasicData = res.CharacterBasicData{
		ID:        "char1",
		Name:      "charName",
		AI:        true,
		Level:     2,
		Sex:       1,
		Race:      1,
		Attitude:  1,
		Guild:     "guildID",
		Alignment: 1,
		Str:       2,
		Con:       3,
		Dex:       4,
		Int:       5,
		Wis:       6,
	}
	char1 := character.New(data)
	data.BasicData.ID = "char2"
	char2 := character.New(data)
	xmlChars, err := MarshalCharacters(char1, char2)
	if err != nil {
		t.Errorf("Unable to marshal character: %v", err)
	}
	if !strings.Contains(xmlChars, "id=\"char1\"") {
		t.Errorf("Marshaled data is invalid: char1 ID: %s", xmlChars)
	}
	if !strings.Contains(xmlChars, "id=\"char2\"") {
		t.Errorf("Marshaled data is invalid: char2 ID: %s", xmlChars)
	}
}
