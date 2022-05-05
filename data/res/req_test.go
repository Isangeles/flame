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

// TestReqsDataJson tests requirements data JSON mappings.
func TestReqsDataJson(t *testing.T) {
	data, err := testData("reqs.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	reqs := new(ReqsData)
	err = json.Unmarshal(data, reqs)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	if len(reqs.LevelReqs) != 2 {
		t.Errorf("Invalid number of level requirements: %d != 2", len(reqs.LevelReqs))
	}
	if len(reqs.GenderReqs) != 2 {
		t.Errorf("Invalid number of gender requirements: %d != 2", len(reqs.GenderReqs))
	}
	if len(reqs.FlagReqs) != 2 {
		t.Errorf("Invalid number of flag requirements: %d != 2", len(reqs.FlagReqs))
	}
	if len(reqs.ItemReqs) != 2 {
		t.Errorf("Invalid number of item requirements: %d != 2", len(reqs.ItemReqs))
	}
	if len(reqs.CurrencyReqs) != 2 {
		t.Errorf("Invalid number of currency requirements: %d != 2", len(reqs.CurrencyReqs))
	}
	if len(reqs.TargetRangeReqs) != 2 {
		t.Errorf("Invalid number of target range requirements: %d != 2", len(reqs.TargetRangeReqs))
	}
	if len(reqs.KillReqs) != 2 {
		t.Errorf("Invalid number of kill requirements: %d != 2", len(reqs.KillReqs))
	}
	if len(reqs.QuestReqs) != 2 {
		t.Errorf("Invalid number of quest requirements: %d != 2", len(reqs.QuestReqs))
	}
	if len(reqs.HealthReqs) != 2 {
		t.Errorf("Invalid number of health requirements: %d != 2", len(reqs.HealthReqs))
	}
	if len(reqs.HealthPercentReqs) != 2 {
		t.Errorf("Invalid number of health percent requirements: %d != 2", len(reqs.HealthPercentReqs))
	}
	if len(reqs.ManaReqs) != 2 {
		t.Errorf("Invalid number of mana requirements: %d != 2", len(reqs.ManaReqs))
	}
	if len(reqs.ManaPercentReqs) != 2 {
		t.Errorf("Invalid number of mana percent requirements: %d != 2", len(reqs.ManaPercentReqs))
	}
	if len(reqs.CombatReqs) != 2 {
		t.Errorf("Invalid number of combat requirements: %d != 2", len(reqs.CombatReqs))
	}
}

// TestReqsDataXml tests requirements data XML mappings.
func TestReqsDataXml(t *testing.T) {
	data, err := testData("reqs.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	reqs := new(ReqsData)
	err = xml.Unmarshal(data, reqs)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	if len(reqs.LevelReqs) != 2 {
		t.Errorf("Invalid number of level requirements: %d != 2", len(reqs.LevelReqs))
	}
	if len(reqs.GenderReqs) != 2 {
		t.Errorf("Invalid number of gender requirements: %d != 2", len(reqs.GenderReqs))
	}
	if len(reqs.FlagReqs) != 2 {
		t.Errorf("Invalid number of flag requirements: %d != 2", len(reqs.FlagReqs))
	}
	if len(reqs.ItemReqs) != 2 {
		t.Errorf("Invalid number of item requirements: %d != 2", len(reqs.ItemReqs))
	}
	if len(reqs.CurrencyReqs) != 2 {
		t.Errorf("Invalid number of currency requirements: %d != 2", len(reqs.CurrencyReqs))
	}
	if len(reqs.TargetRangeReqs) != 2 {
		t.Errorf("Invalid number of target range requirements: %d != 2", len(reqs.TargetRangeReqs))
	}
	if len(reqs.KillReqs) != 2 {
		t.Errorf("Invalid number of kill requirements: %d != 2", len(reqs.KillReqs))
	}
	if len(reqs.QuestReqs) != 2 {
		t.Errorf("Invalid number of quest requirements: %d != 2", len(reqs.QuestReqs))
	}
	if len(reqs.HealthReqs) != 2 {
		t.Errorf("Invalid number of health requirements: %d != 2", len(reqs.HealthReqs))
	}
	if len(reqs.HealthPercentReqs) != 2 {
		t.Errorf("Invalid number of health percent requirements: %d != 2", len(reqs.HealthPercentReqs))
	}
	if len(reqs.ManaReqs) != 2 {
		t.Errorf("Invalid number of mana requirements: %d != 2", len(reqs.ManaReqs))
	}
	if len(reqs.ManaPercentReqs) != 2 {
		t.Errorf("Invalid number of mana percent requirements: %d != 2", len(reqs.ManaPercentReqs))
	}
	if len(reqs.CombatReqs) != 2 {
		t.Errorf("Invalid number of combat requirements: %d != 2", len(reqs.CombatReqs))
	}
}

// TestItemReqDataJson tests item requirement data JSON mappings.
func TestItemReqDataJson(t *testing.T) {
	data, err := testData("itemreq.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	req := new(ItemReqData)
	err = json.Unmarshal(data, req)
	if err != nil {
		t.Fatalf("Unable to unmarshal test data: %v", err)
	}
	if req.ID != "item1" {
		t.Errorf("Invalid item ID: %s != item1", req.ID)
	}
	if req.Amount != 1 {
		t.Errorf("Invalid amount value: %d != 1", req.Amount)
	}
	if !req.Charge {
		t.Errorf("Invalid charge value: %v != true", req.Charge)
	}
}

// TestItemReqDataXml tests item requirement data XML mappings.
func TestItemReqDataXml(t *testing.T) {
	data, err := testData("itemreq.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	req := new(ItemReqData)
	err = xml.Unmarshal(data, req)
	if err != nil {
		t.Fatalf("Unable to unmarshal test data: %v", err)
	}
	if req.ID != "item1" {
		t.Errorf("Invalid item ID: %s != item1", req.ID)
	}
	if req.Amount != 1 {
		t.Errorf("Invalid amount value: %d != 1", req.Amount)
	}
	if !req.Charge {
		t.Errorf("Invalid charge value: %v != true", req.Charge)
	}
}

// TestHealthReqDataJson tests health requirement data JSON mappings.
func TestHealthReqDataJson(t *testing.T) {
	data, err := testData("healthreq.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	req := new(HealthReqData)
	err = json.Unmarshal(data, req)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	if req.Value != 100 {
		t.Errorf("Invalid health value: %d != 100", req.Value)
	}
	if !req.Less {
		t.Errorf("Invalid health less value: %v != true", req.Less)
	}
	if !req.Charge {
		t.Errorf("Invalid charge value: %v != true", req.Charge)
	}
}

// TestHealthReqDataXml tests health requirement data XML mappings.
func TestHealthReqDataXml(t *testing.T) {
	data, err := testData("healthreq.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	req := new(HealthReqData)
	err = xml.Unmarshal(data, req)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	if req.Value != 100 {
		t.Errorf("Invalid health value: %d != 100", req.Value)
	}
	if !req.Less {
		t.Errorf("Invalid health less value: %v != true", req.Less)
	}
	if !req.Charge {
		t.Errorf("Invalid charge value: %v != true", req.Charge)
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
		t.Errorf("Invalid mana value: %d != 100", req.Value)
	}
	if !req.Less {
		t.Errorf("Invalid less value: %v != true", req.Less)
	}
	if !req.Charge {
		t.Errorf("Invalid charge value: %v != true", req.Charge)
	}
}

// TestManaReqDataXml tests mana requirement data XML mappings.
func TestManaReqDataXml(t *testing.T) {
	data, err := testData("manareq.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	req := new(ManaReqData)
	err = xml.Unmarshal(data, req)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	if req.Value != 100 {
		t.Errorf("Invalid mana value: %d != 100", req.Value)
	}
	if !req.Less {
		t.Errorf("Invalid less value: %v != true", req.Less)
	}
	if !req.Charge {
		t.Errorf("Invalid charge value: %v != true", req.Charge)
	}
}

// TestCombatReqDataJson tests combat requirement data JSON mappings.
func TestCombatReqDataJson(t *testing.T) {
	data, err := testData("combatreq.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	req := new(CombatReqData)
	err = json.Unmarshal(data, req)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	if !req.Combat {
		t.Errorf("Invalid combat value: %v != true", req.Combat)
	}
}

// TestCombatReqDataXml tests combat requirement data XML mappings.
func TestCombatReqDataXml(t *testing.T) {
	data, err := testData("combatreq.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	req := new(CombatReqData)
	err = xml.Unmarshal(data, req)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	if !req.Combat {
		t.Errorf("Invalid combat value: %v != true", req.Combat)
	}
}
