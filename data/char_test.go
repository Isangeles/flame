/*
 * char_test.go
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

package data

import (
	"testing"

	"github.com/isangeles/flame/data/res"
)

// Test for exporting characters to characters file.
func TestExportCharacters(t *testing.T) {
	data := res.CharacterData{
		ID:        "char1",
		AI:        true,
		Level:     2,
		Sex:       "genderMale",
		Race:      "raceHuman",
		Attitude:  "attFriendly",
		Guild:     "guildID",
		Alignment: "aliTrueNeutral",
	}
	data.Attributes = res.AttributesData{
		Str:       2,
		Con:       3,
		Dex:       4,
		Int:       5,
		Wis:       6,
	}
	chars := make([]res.CharacterData, 0)
	chars = append(chars, data)
	data.ID = "char2"
	chars = append(chars, data)
	err := ExportCharacters("testchars", chars...)
	if err != nil {
		t.Errorf("Unable to export characters: %v", err)
	}
}
