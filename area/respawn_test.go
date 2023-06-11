/*
 * respawn_test.go
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

package area

import (
	"testing"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/character"
)

// TestAreaRespawn tests respawn for area.
func TestAreaRespawn(t *testing.T) {
	// Create area.
	res.Characters = append(res.Characters, charData)
	ob := character.New(charData)
	ob.SetRespawn(1000)
	area := New(areaData)
	area.AddObject(ob)
	// Test.
	ob.SetHealth(0)
	area.Update(1)
	area.Update(1000)
	if len(area.Objects()) != 1 {
		t.Errorf("Invalid number of area objects: %d != 1", len(area.Objects()))
	}
	respOb := area.Objects()[0]
	if respOb == ob {
		t.Errorf("Object was not respawned")
	}
}

