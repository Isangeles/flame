/*
 * character_test.go
 *
 * Copyright 2022 Dariusz Sikora <dev@isangeles.pl>
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
)

var charData = res.CharacterData{ID: "char", Level: 1, Attributes: res.AttributesData{5, 5, 5, 5, 5}}

// TestLive tests live check function.
func TestLive(t *testing.T) {
	char := New(charData)
	if !char.Live() {
		t.Errorf("Character should be live: %v != true", char.Live())
	}
	char.SetHealth(0)
	char.Update(1)
	if char.Live() {
		t.Errorf("Character should be dead: %v != false", char.Live())
	}
}
