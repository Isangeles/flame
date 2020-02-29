/*
 * data_test.go
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

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/character"
)

// Test for exporting characters to characters file.
func TestExportCharacters(t *testing.T) {
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
	chars := make([]*character.Character, 0)
	chars = append(chars, character.New(data))
	data.BasicData.ID = "char2"
	chars = append(chars, character.New(data))
	err := ExportCharacters("testchars", chars...)
	if err != nil {
		t.Errorf("Unable to export characters: %v", err)
	}
}
