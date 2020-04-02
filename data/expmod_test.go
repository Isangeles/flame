/*
 * expmod_test.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/module"
	"github.com/isangeles/flame/module/area"
)

// Test for exporting module.
func TestExportModule(t *testing.T) {
	modConf := module.Config{
		ID:      "test",
		Chapter: "ch1",
	}
	mod := module.New(modConf)
	chConf := module.ChapterConfig{
		ID:        "ch1",
		StartArea: "a1",
	}
	chapter := module.NewChapter(mod, chConf)
	mod.SetChapter(chapter)
	areaData := res.AreaData{
		ID: "a1",
	}
	charData := res.CharacterData{
		ID:        "char1",
		Name:      "charName",
		AI:        true,
		Level:     2,
		Sex:       "genderMale",
		Race:      "raceHuman",
		Attitude:  "attFriendly",
		Guild:     "guildID",
		Alignment: "aliTrueNeutral",
		Str:       2,
		Con:       3,
		Dex:       4,
		Int:       5,
		Wis:       6,
	}
	res.SetCharactersData(append(res.Characters(), &charData))
	areaCharData := res.AreaCharData{ID: charData.ID}
	areaData.Characters = append(areaData.Characters, areaCharData)
	area := area.New(areaData)
	chapter.AddAreas(area)
	err := ExportModule(mod, "testexp")
	if err != nil {
		t.Errorf("Unable to export module: %v", err)
	}
}
