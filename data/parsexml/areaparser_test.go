/*
 * areaparser_test.go
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
	"testing"
	"strings"

	"github.com/isangeles/flame/data/res"
)

// Test for unmarshaling area data.
func TestUnmarshalArea(t *testing.T) {
	xmlArea := `<area id="area1">
        <characters>
               <character id="char1" position="10x10" ai="true"/>
        </characters>
        <objects>
                <object id="ob1" position="0x20"></object>
        </objects>
        <subareas>
        </subareas>
</area>`
	data, err := UnmarshalArea(strings.NewReader(xmlArea))
	if err != nil {
		t.Errorf("Unable to unmarshal XML string: %v", err)
		return
	}
	if data.ID != "area1" {
		t.Errorf("Unmarshaled data is invalid: ID: %s != area1", data.ID)
	}
}

// Test for marshaling area data.
func TestMarshalArea(t *testing.T) {
	data := res.AreaData{
		ID: "area1",
	}
	char1 := res.AreaCharData{
		ID:   "ch1",
		PosX: 1.0,
		PosY: 2.0,
		AI:   true,
	}
	data.Characters = append(data.Characters, char1)
	ob1 := res.AreaObjectData{
		ID:   "ob1",
		PosX: 10.0,
		PosY: 20.0,
	}
	data.Objects = append(data.Objects, ob1)
	sa1 := res.AreaData{
		ID: "area1_sub",
	}
	data.Subareas = append(data.Subareas, sa1)
	xmlArea, err := MarshalArea(&data)
	if err != nil {
		t.Errorf("Unable to marshal area data: %v", err)
		return
	}
	if !strings.Contains(xmlArea, "<area id=\"area1\">") ||
		!strings.Contains(xmlArea, "<area id=\"area1_sub\">") {
		t.Errorf("Invalid output XML: %s", xmlArea)
	}
}
