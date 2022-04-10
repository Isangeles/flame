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
		t.Errorf("Inavlid number of level requrements: %d != 2", len(reqs.LevelReqs))
	}
	if len(reqs.GenderReqs) != 2 {
		t.Errorf("Inavlid number of gender requrements: %d != 2", len(reqs.GenderReqs))
	}
	if len(reqs.FlagReqs) != 2 {
		t.Errorf("Inavlid number of flag requrements: %d != 2", len(reqs.FlagReqs))
	}
	if len(reqs.ItemReqs) != 2 {
		t.Errorf("Inavlid number of item requrements: %d != 2", len(reqs.ItemReqs))
	}
	if len(reqs.CurrencyReqs) != 2 {
		t.Errorf("Inavlid number of currency requrements: %d != 2", len(reqs.CurrencyReqs))
	}
	if len(reqs.TargetRangeReqs) != 2 {
		t.Errorf("Inavlid number of target range requrements: %d != 2", len(reqs.TargetRangeReqs))
	}
	if len(reqs.KillReqs) != 2 {
		t.Errorf("Inavlid number of kill requrements: %d != 2", len(reqs.KillReqs))
	}
	if len(reqs.QuestReqs) != 2 {
		t.Errorf("Inavlid number of quest requrements: %d != 2", len(reqs.QuestReqs))
	}
	if len(reqs.HealthPercentReqs) != 2 {
		t.Errorf("Inavlid number of health percent requrements: %d != 2", len(reqs.HealthPercentReqs))
	}
	if len(reqs.ManaReqs) != 2 {
		t.Errorf("Inavlid number of mana requrements: %d != 2", len(reqs.ManaReqs))
	}
	if len(reqs.ManaPercentReqs) != 2 {
		t.Errorf("Inavlid number of mana percent requrements: %d != 2", len(reqs.ManaPercentReqs))
	}
	if len(reqs.CombatReqs) != 2 {
		t.Errorf("Inavlid number of combat requrements: %d != 2", len(reqs.CombatReqs))
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
		t.Errorf("Inavlid number of level requrements: %d != 2", len(reqs.LevelReqs))
	}
	if len(reqs.GenderReqs) != 2 {
		t.Errorf("Inavlid number of gender requrements: %d != 2", len(reqs.GenderReqs))
	}
	if len(reqs.FlagReqs) != 2 {
		t.Errorf("Inavlid number of flag requrements: %d != 2", len(reqs.FlagReqs))
	}
	if len(reqs.ItemReqs) != 2 {
		t.Errorf("Inavlid number of item requrements: %d != 2", len(reqs.ItemReqs))
	}
	if len(reqs.CurrencyReqs) != 2 {
		t.Errorf("Inavlid number of currency requrements: %d != 2", len(reqs.CurrencyReqs))
	}
	if len(reqs.TargetRangeReqs) != 2 {
		t.Errorf("Inavlid number of target range requrements: %d != 2", len(reqs.TargetRangeReqs))
	}
	if len(reqs.KillReqs) != 2 {
		t.Errorf("Inavlid number of kill requrements: %d != 2", len(reqs.KillReqs))
	}
	if len(reqs.QuestReqs) != 2 {
		t.Errorf("Inavlid number of quest requrements: %d != 2", len(reqs.QuestReqs))
	}
	if len(reqs.HealthPercentReqs) != 2 {
		t.Errorf("Inavlid number of health percent requrements: %d != 2", len(reqs.HealthPercentReqs))
	}
	if len(reqs.ManaReqs) != 2 {
		t.Errorf("Inavlid number of mana requrements: %d != 2", len(reqs.ManaReqs))
	}
	if len(reqs.ManaPercentReqs) != 2 {
		t.Errorf("Inavlid number of mana percent requrements: %d != 2", len(reqs.ManaPercentReqs))
	}
	if len(reqs.CombatReqs) != 2 {
		t.Errorf("Inavlid number of combat requrements: %d != 2", len(reqs.CombatReqs))
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
		t.Errorf("Inavlid mana value in mana requirement data: %d != 100", req.Value)
	}
	if !req.Less {
		t.Errorf("Inavlid mana less in mana requirement data: %v != true", req.Less)
	}
}
