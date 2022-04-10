/*
 * req_test.go
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

package res

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

// TestManaReqDataXml tests mana requirement data XML mappings.
func TestManaReqDataXml(t *testing.T) {
	data, err := testData("manareq.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	req := new(ManaReqData)
	err = xml.Unmarshal(data, req)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	if req.Value != 100 {
		t.Errorf("Inavlid mana value in mana requirement data: %d != 100", req.Value)
	}
	if !req.Less {
		t.Errorf("Inavlid mana less in mana requirement data: %v != true", req.Less)
	}
}

// TestManaReqDataJson tests mana requirement data JSON mappings.
func TestManaReqDataJson(t *testing.T) {
	data, err := testData("manareq.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	req := new(ManaReqData)
	err = json.Unmarshal(data, req)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	if req.Value != 100 {
		t.Errorf("Inavlid mana value in mana requirement data: %d != 100", req.Value)
	}
	if !req.Less {
		t.Errorf("Inavlid mana less in mana requirement data: %v != true", req.Less)
	}
}
