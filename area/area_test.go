/*
 * area.go
 *
 * Copyright 2022 Dariusz Sikora <ds@isangeles.dev>
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

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/character"
)

var charData = res.CharacterData{ID: "char"}

// TestNearObjects function for retrieving near objects.
func TestNearObjects(t *testing.T) {
	// Create objects & area.
	char1 := character.New(charData)
	char1.SetPosition(10, 10)
	char2 := character.New(charData)
	char2.SetPosition(10, 15)
	char3 := character.New(charData)
	char3.SetPosition(30, 50)
	area := New()
	area.AddCharacter(char1)
	area.AddCharacter(char2)
	area.AddCharacter(char3)
	// Test.
	objects := area.NearObjects(0, 0, 20)
	if len(objects) != 2 {
		t.Errorf("Invalid number of objects returned: %d", len(objects))
	}
}

