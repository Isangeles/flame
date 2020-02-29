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

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/area"
)

// Test for exporting module.
func TestExportModule(t *testing.T) {
	modConf := module.Config{
		ID:      "test",
		Chapter: "ch1",
	}
	mod := module.New(modConf)
	chConf := module.ChapterConfig{
		ID:          "ch1",
		StartAreaID: "a1",
	}
	chapter := module.NewChapter(mod, chConf)
	mod.SetChapter(chapter)
	areaData := res.AreaData{
		ID: "a1",
	}
	area := area.New(areaData)
	chapter.AddAreas(area)
	err := ExportModule(mod, "testexp")
	if err != nil {
		t.Errorf("Unable to export module: %v", err)
	}
}
