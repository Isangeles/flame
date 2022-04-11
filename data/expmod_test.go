/*
 * expmod_test.go
 *
 * Copyright 2020-2022 Dariusz Sikora <dev@isangeles.pl>
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
	"path/filepath"
	"testing"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/area"
	"github.com/isangeles/flame/data/res"
)

// Test for exporting module to the single file.
func TestExportModule(t *testing.T) {
	modData := res.ModuleData{
		Config:  make(map[string][]string),
		Chapter: res.ChapterData{Config: make(map[string][]string)},
	}
	modData.Config["id"] = []string{"test"}
	modData.Config["chapter"] = []string{"ch1"}
	modData.Chapter.Config["id"] = []string{"ch1"}
	modData.Chapter.Config["start-area"] = []string{"a1"}
	mod := flame.NewModule(modData)
	area := area.New()
	areaData := res.AreaData{
		ID: "a1",
	}
	charData := res.CharacterData{
		ID:        "char1",
		Level:     2,
		Sex:       "genderMale",
		Race:      "raceHuman",
		Attitude:  "attFriendly",
		Guild:     "guildID",
		Alignment: "aliTrueNeutral",
	}
	charData.Attributes = res.AttributesData{
		Str: 2,
		Con: 3,
		Dex: 4,
		Int: 5,
		Wis: 6,
	}
	res.Characters = append(res.Characters, charData)
	areaCharData := res.AreaCharData{ID: charData.ID}
	areaData.Characters = append(areaData.Characters, areaCharData)
	area.Apply(areaData)
	mod.Chapter().AddAreas(area)
	path := filepath.Join(t.TempDir(), "testexp")
	err := ExportModule(path, mod.Data())
	if err != nil {
		t.Errorf("Unable to export module file: %v", err)
	}
}

// Test for exporting module into the directory.
func TestExportModuleDir(t *testing.T) {
	modData := res.ModuleData{
		Config:  make(map[string][]string),
		Chapter: res.ChapterData{Config: make(map[string][]string)},
	}
	modData.Config["id"] = []string{"test"}
	modData.Config["chapter"] = []string{"ch1"}
	modData.Chapter.Config["id"] = []string{"ch1"}
	modData.Chapter.Config["start-area"] = []string{"a1"}
	mod := flame.NewModule(modData)
	area := area.New()
	areaData := res.AreaData{
		ID: "a1",
	}
	charData := res.CharacterData{
		ID:        "char1",
		Level:     2,
		Sex:       "genderMale",
		Race:      "raceHuman",
		Attitude:  "attFriendly",
		Guild:     "guildID",
		Alignment: "aliTrueNeutral",
	}
	charData.Attributes = res.AttributesData{
		Str: 2,
		Con: 3,
		Dex: 4,
		Int: 5,
		Wis: 6,
	}
	res.Characters = append(res.Characters, charData)
	areaCharData := res.AreaCharData{ID: charData.ID}
	areaData.Characters = append(areaData.Characters, areaCharData)
	area.Apply(areaData)
	mod.Chapter().AddAreas(area)
	err := ExportModuleDir(t.TempDir(), mod.Data())
	if err != nil {
		t.Errorf("Unable to export module: %v", err)
	}
}
