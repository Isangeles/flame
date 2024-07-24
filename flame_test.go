/*
 * flame_test.go
 *
 * Copyright 2023-2024 Dariusz Sikora <ds@isangeles.dev>
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

package flame

import (
	"testing"

	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/data/res"
)

var (
	charData    = res.CharacterData{ID: "char"}
	areaData    = res.AreaData{ID: "area"}
	resData     = res.ResourcesData{Areas: []res.AreaData{areaData}}
	chapterData = res.ChapterData{ID: "chapter", Resources: resData}
	modData     = res.ModuleData{ID: "module", Chapter: chapterData}
)

// TestChapterChangeEvent tests adding and triggering chapter
// change event.
func TestChapterChangeEvent(t *testing.T) {
	// Create test objects
	mod := NewModule(modData)
	area := mod.Chapter().Area("area")
	if area == nil {
		t.Fatalf("Test area not found")
	}
	ob := character.New(charData)
	area.AddObject(ob)
	// Test
	evTriggered := false
	ev := func(evOb *character.Character) {
		evTriggered = true
		if evOb != ob {
			t.Errorf("Event triggered with invalid object: %v",
				ob)
		}
	}
	mod.AddChangeChapterEvent(ev)
	ob.SetChapterID("nextChapter")
	mod.Update(1)
	if !evTriggered {
		t.Errorf("Event was not triggered")
	}
}
