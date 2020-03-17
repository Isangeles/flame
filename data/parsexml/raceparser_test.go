/*
 * raceparser_test.go
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

package parsexml

import (
	"strings"
	"testing"
)

// Test for unmarshaling races data.
func TestUnmarshalRaces(t *testing.T) {
	xmlRaces := `<races>
	<race id="rHuman"></race>
	<race id="rElf"></race>
	<race id="rDwarf"></race>
	<race id="rGnome"></race>
</races>`
	data, err := UnmarshalRaces(strings.NewReader(xmlRaces))
	if err != nil {
		t.Errorf("Unable to unmarshal XML string: %v", err)
		return
	}
	if len(data) != 4 {
		t.Errorf("Unmarshaled data is invalid: len: %d != 4", len(data))
		return
	}
}
