/*
 * mod_test.go
 *
 * Copyright 2023 Dariusz Sikora <ds@isangeles.dev>
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

package character

import (
	"testing"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/effect"
)

var (
	areaModData          = res.AreaModData{"testArea", 10, 10}
	changeChapterModData = res.ChangeChapterModData{"testChapter"}
)

// TestTakeModifiersArea tests handling of area
// modifier.
func TestTakeModifiersArea(t *testing.T) {
	ob := New(charData)
	mod := effect.NewAreaMod(areaModData)
	ob.TakeModifiers(nil, mod)
	if ob.AreaID() != mod.AreaID() {
		t.Errorf("Invalid area ID: '%s' != '%s'",
			ob.AreaID(), mod.AreaID())
	}
	obPosX, obPosY := ob.Position()
	entryPosX, entryPosY := mod.EntryPosition()
	if obPosX != entryPosX || obPosY != entryPosY {
		t.Errorf("Invalid character position: %fx%f != %fx%f",
			obPosX, obPosY, entryPosX, entryPosY)
	}
}

// TestTakeModifiersChapterChange tests handling of
// chapter change modifier.
func TestTakeModifiersChapterChange(t *testing.T) {
	ob := New(charData)
	mod := effect.NewChangeChapterMod(changeChapterModData)
	ob.TakeModifiers(nil, mod)
	if ob.ChapterID() != mod.ChapterID() {
		t.Errorf("Invalid chapter ID: '%s' != '%s'",
			ob.ChapterID(), mod.ChapterID())
	}
}
