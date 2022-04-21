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

package req

import (
	"testing"

	"github.com/isangeles/flame/data/res"
)

// TestNewRequirements tests creating new requirements
// from requirements data resource.
func TestNewRequirements(t *testing.T) {
	data := testReqsData()
	reqs := NewRequirements(data)
	var (
		levelReqs         int = 0
		genderReqs            = 0
		flagReqs              = 0
		itemReqs              = 0
		currencyReqs          = 0
		targetRangeReqs       = 0
		killReqs              = 0
		questReqs             = 0
		healthReqs            = 0
		healthPercentReqs     = 0
		manaReqs              = 0
		manaPercentReqs       = 0
		combatReqs            = 0
	)
	for _, r := range reqs {
		switch r.(type) {
		case *Level:
			levelReqs++
		case *Gender:
			genderReqs++
		case *Flag:
			flagReqs++
		case *Item:
			itemReqs++
		case *Currency:
			currencyReqs++
		case *TargetRange:
			targetRangeReqs++
		case *Kill:
			killReqs++
		case *Quest:
			questReqs++
		case *Health:
			healthReqs++
		case *HealthPercent:
			healthPercentReqs++
		case *Mana:
			manaReqs++
		case *ManaPercent:
			manaPercentReqs++
		case *Combat:
			combatReqs++
		}
	}
	if levelReqs != 2 {
		t.Errorf("Invalid number of level requirements: %d != 2", levelReqs)
	}
	if genderReqs != 2 {
		t.Errorf("Invalid number of gender requirements: %d != 2", genderReqs)
	}
	if flagReqs != 2 {
		t.Errorf("Invalid number of flag requirements: %d != 2", flagReqs)
	}
	if itemReqs != 2 {
		t.Errorf("Invalid number of item requirements: %d != 2", itemReqs)
	}
	if currencyReqs != 2 {
		t.Errorf("Invalid number of currency requirements: %d != 2", currencyReqs)
	}
	if targetRangeReqs != 2 {
		t.Errorf("Invalid number of target range requirements: %d != 2", targetRangeReqs)
	}
	if killReqs != 2 {
		t.Errorf("Invalid number of kill requirements: %d != 2", killReqs)
	}
	if questReqs != 2 {
		t.Errorf("Invalid number of quest requirements: %d != 2", questReqs)
	}
	if healthReqs != 2 {
		t.Errorf("Invalid number of health requirements: %d != 2", healthReqs)
	}
	if healthPercentReqs != 2 {
		t.Errorf("Invalid number of health percent requirements: %d != 2", healthPercentReqs)
	}
	if manaReqs != 2 {
		t.Errorf("Invalid number of mana requirements: %d != 2", manaReqs)
	}
	if manaPercentReqs != 2 {
		t.Errorf("Invalid number of mana percent requirements: %d != 2", manaPercentReqs)
	}
	if combatReqs != 2 {
		t.Errorf("Invalid number of combat requirements: %d != 2", combatReqs)
	}
}

// TestRequirementsData tests creating requirements data
// resource from requirements.
func TestRequirementsData(t *testing.T) {
	expectedData := testReqsData()
	reqs := NewRequirements(expectedData)
	data := RequirementsData(reqs...)
	if len(data.LevelReqs) != len(expectedData.LevelReqs) {
		t.Errorf("Invalid number of level requirements: %d != %d", len(data.LevelReqs),
			len(expectedData.LevelReqs))
	}
	if len(data.GenderReqs) != len(expectedData.GenderReqs) {
		t.Errorf("Invalid number of gender requirements: %d != %d", len(data.GenderReqs),
			len(expectedData.GenderReqs))
	}
	if len(data.FlagReqs) != len(expectedData.FlagReqs) {
		t.Errorf("Invalid number of flag requirements: %d != %d", len(data.FlagReqs),
			len(expectedData.FlagReqs))
	}
	if len(data.ItemReqs) != len(expectedData.ItemReqs) {
		t.Errorf("Invalid number of item requirements: %d != %d", len(data.ItemReqs),
			len(expectedData.ItemReqs))
	}
	if len(data.CurrencyReqs) != len(expectedData.CurrencyReqs) {
		t.Errorf("Invalid number of currency requirements: %d != %d", len(data.CurrencyReqs),
			len(expectedData.CurrencyReqs))
	}
	if len(data.TargetRangeReqs) != len(expectedData.TargetRangeReqs) {
		t.Errorf("Invalid number of target range requirements: %d != %d", len(data.TargetRangeReqs),
			len(expectedData.TargetRangeReqs))
	}
	if len(data.KillReqs) != len(expectedData.KillReqs) {
		t.Errorf("Invalid number of kill requirements: %d != %d", len(data.KillReqs),
			len(expectedData.KillReqs))
	}
	if len(data.QuestReqs) != len(expectedData.QuestReqs) {
		t.Errorf("Invalid number of quest requirements: %d != %d", len(data.QuestReqs),
			len(expectedData.QuestReqs))
	}
	if len(data.KillReqs) != len(expectedData.KillReqs) {
		t.Errorf("Invalid number of kill requirements: %d != %d", len(data.KillReqs),
			len(expectedData.KillReqs))
	}
	if len(data.HealthReqs) != len(expectedData.HealthReqs) {
		t.Errorf("Invalid number of health requirements: %d != %d", len(data.HealthReqs),
			len(expectedData.HealthReqs))
	}
	if len(data.HealthPercentReqs) != len(expectedData.HealthPercentReqs) {
		t.Errorf("Invalid number of health percent requirements: %d != %d", len(data.HealthPercentReqs),
			len(expectedData.HealthPercentReqs))
	}
	if len(data.ManaReqs) != len(expectedData.ManaReqs) {
		t.Errorf("Invalid number of mana requirements: %d != %d", len(data.ManaReqs),
			len(expectedData.ManaReqs))
	}
	if len(data.ManaPercentReqs) != len(expectedData.ManaPercentReqs) {
		t.Errorf("Invalid number of mana percent requirements: %d != %d", len(data.ManaPercentReqs),
			len(expectedData.ManaPercentReqs))
	}
	if len(data.CombatReqs) != len(expectedData.CombatReqs) {
		t.Errorf("Invalid number of combat requirements: %d != %d", len(data.CombatReqs),
			len(expectedData.CombatReqs))
	}
}

// testReqsData creats test requirements data
// resource.
func testReqsData() res.ReqsData {
	levelReqs := []res.LevelReqData{
		res.LevelReqData{10, 10},
		res.LevelReqData{20, 20},
	}
	genderReqs := []res.GenderReqData{
		res.GenderReqData{"gender1"},
		res.GenderReqData{"gender2"},
	}
	flagReqs := []res.FlagReqData{
		res.FlagReqData{"flag1", false},
		res.FlagReqData{"flag2", true},
	}
	itemReqs := []res.ItemReqData{
		res.ItemReqData{"item1", 1, false},
		res.ItemReqData{"item2", 2, true},
	}
	currencyReqs := []res.CurrencyReqData{
		res.CurrencyReqData{100, false},
		res.CurrencyReqData{50, true},
	}
	targetRangeReqs := []res.TargetRangeReqData{
		res.TargetRangeReqData{5},
		res.TargetRangeReqData{10},
	}
	killReqs := []res.KillReqData{
		res.KillReqData{"object1", 1},
		res.KillReqData{"object2", 2},
	}
	questReqs := []res.QuestReqData{
		res.QuestReqData{"quest1", false},
		res.QuestReqData{"quest2", true},
	}
	healthReqs := []res.HealthReqData{
		res.HealthReqData{100, false},
		res.HealthReqData{50, true},
	}
	healthPercentReqs := []res.HealthPercentReqData{
		res.HealthPercentReqData{100, false},
		res.HealthPercentReqData{50, true},
	}
	manaReqs := []res.ManaReqData{
		res.ManaReqData{10, false, false},
		res.ManaReqData{5, true, true},
	}
	manaPercentReqs := []res.ManaPercentReqData{
		res.ManaPercentReqData{100, false},
		res.ManaPercentReqData{50, true},
	}
	combatReqs := []res.CombatReqData{
		res.CombatReqData{false},
		res.CombatReqData{true},
	}
	data := res.ReqsData{
		LevelReqs:         levelReqs,
		GenderReqs:        genderReqs,
		FlagReqs:          flagReqs,
		ItemReqs:          itemReqs,
		CurrencyReqs:      currencyReqs,
		TargetRangeReqs:   targetRangeReqs,
		KillReqs:          killReqs,
		QuestReqs:         questReqs,
		HealthReqs:        healthReqs,
		HealthPercentReqs: healthPercentReqs,
		ManaReqs:          manaReqs,
		ManaPercentReqs:   manaPercentReqs,
		CombatReqs:        combatReqs,
	}
	return data
}
