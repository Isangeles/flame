/*
 * area.go
 *
 * Copyright 2022-2023 Dariusz Sikora <ds@isangeles.dev>
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

package area

import (
	"testing"

	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/data/res"
)

var (
	areaData = res.AreaData{ID: "area"}
	charData = res.CharacterData{ID: "char", Level: 1}
)

// TestNearObjects tests function for retrieving near objects.
func TestNearObjects(t *testing.T) {
	// Create objects & area.
	char1 := character.New(charData)
	char1.SetPosition(30, 50)
	char2 := character.New(charData)
	char2.SetPosition(10, 15)
	char3 := character.New(charData)
	char3.SetPosition(10, 10)
	area := New(areaData)
	area.AddObject(char1)
	area.AddObject(char2)
	area.AddObject(char3)
	// Test.
	objects := area.NearObjects(0, 0, 20)
	if len(objects) != 2 {
		t.Errorf("Invalid number of objects returned: %d", len(objects))
	}
	if containsObject(char1.ID(), char1.Serial(), objects...) {
		t.Errorf("Object should not be among returned objects: %s %s",
			char1.ID(), char1.Serial())
	}
	if !containsObject(char2.ID(), char2.Serial(), objects...) {
		t.Errorf("Object should be among returned objects: %s %s",
			char2.ID(), char2.Serial())
	}
	if !containsObject(char3.ID(), char3.Serial(), objects...) {
		t.Errorf("Object should be among returned objects: %s %s",
			char3.ID(), char3.Serial())
	}
}

// TestSightRangeObjects tests function for retrieving
// objects with specified XY position in range.
func TestSightRangeObjects(t *testing.T) {
	// Create objects & area.
	char1 := character.New(charData)
	char1.SetPosition(0, 0)
	char2 := character.New(charData)
	char2.SetPosition(10, 15)
	char3 := character.New(charData)
	char3.SetPosition(30, 50)
	area := New(areaData)
	area.AddObject(char1)
	area.AddObject(char2)
	area.AddObject(char3)
	// Test
	objects := area.SightRangeObjects(220, 220)
	if len(objects) != 2 {
		t.Errorf("Invalid number of objects returned: %d", len(objects))
	}
	if containsObject(char1.ID(), char1.Serial(), objects...) {
		t.Errorf("Object should not be among returned objects: %s %s",
			char1.ID(), char1.Serial())
	}
	if !containsObject(char2.ID(), char2.Serial(), objects...) {
		t.Errorf("Object should be among returned objects: %s %s",
			char2.ID(), char2.Serial())
	}
	if !containsObject(char3.ID(), char3.Serial(), objects...) {
		t.Errorf("Object should be among returned objects: %s %s",
			char3.ID(), char3.Serial())
	}
}

// TestCharacterMove tests moving objects to their
// destination points.
func TestCharacterMove(t *testing.T) {
	// Creates object & area.
	ob := character.New(charData)
	area := New(areaData)
	area.AddObject(ob)
	// Test.
	x, y := ob.Position()
	if x != 0.0 || y != 0.0 {
		t.Errorf("Invalid starting position: %f %f", x, y)
	}
	ob.SetDestPoint(2, 2)
	area.Update(1)
	area.Update(1000)
	x, y = ob.Position()
	if x != 2 || y != 2 {
		t.Errorf("Object not moved to destination point: %f %f != 2 2",
			x, y)
	}
}

// containsObject checks if object with specified ID and serial
func containsObject(id, serial string, obs ...Object) bool {
	for _, ob := range obs {
		if ob.ID() == id && ob.Serial() == serial {
			return true
		}
	}
	return false
}
