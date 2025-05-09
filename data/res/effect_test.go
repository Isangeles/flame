/*
 * effect_test.go
 *
 * Copyright 2023-2025 Dariusz Sikora <ds@isangeles.dev>
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

// TestEffectsDataJson tests effects data JSON mappings.
func TestEffectsDataJson(t *testing.T) {
	// Create test effects
	data, err := testData("effects.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	effects := new(EffectsData)
	err = json.Unmarshal(data, effects)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	// Test
	if len(effects.Effects) != 2 {
		t.Errorf("Invalid number of effects: %d != 2", len(effects.Effects))
	}
}

// TestEffectsDataXml tests effects data XML mappings.
func TestEffectsDataXml(t *testing.T) {
	// Create test effects
	data, err := testData("effects.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	effects := new(EffectsData)
	err = xml.Unmarshal(data, effects)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	// Test
	if len(effects.Effects) != 2 {
		t.Errorf("Invalid number of effects: %d != 2", len(effects.Effects))
	}
}

// TestEffectDataJson tests effect data JSON mappings.
func TestEffectDataJson(t *testing.T) {
	// Create test effect
	data, err := testData("effect.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	effect := new(EffectData)
	err = json.Unmarshal(data, effect)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	// Test
	if effect.ID != "effect1" {
		t.Errorf("Invalid ID: %s != 'effect1'", effect.ID)
	}
	if effect.Duration != 1000 {
		t.Errorf("Invalid duration: %d != 1000", effect.Duration)
	}
	if !effect.MeleeHit {
		t.Errorf("Melee hit value is false")
	}
	if !effect.Infinite {
		t.Errorf("Infinite value is false")
	}
	if !effect.Hostile {
		t.Errorf("Hostile value is false")
	}
}

// TestEffectDataXml tests effect data XML mappings.
func TestEffectDataXml(t *testing.T) {
	// Create test effect
	data, err := testData("effect.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	effect := new(EffectData)
	err = xml.Unmarshal(data, effect)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	// Test
	if effect.ID != "effect1" {
		t.Errorf("Invalid ID: %s != 'effect1'", effect.ID)
	}
	if effect.Duration != 1000 {
		t.Errorf("Invalid duration: %d != 1000", effect.Duration)
	}
	if !effect.MeleeHit {
		t.Errorf("Melee hit value is false")
	}
	if !effect.Infinite {
		t.Errorf("Infinite value is false")
	}
	if !effect.Hostile {
		t.Errorf("Hostile value is false")
	}
}

// TestModifiersDataJson tests modifiers data JSON mappings.
func TestModifiersDataJson(t *testing.T) {
	// Create test modifiers
	data, err := testData("modifiers.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	mods := new(ModifiersData)
	err = json.Unmarshal(data, mods)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	// Test
	if len(mods.HealthMods) != 2 {
		t.Errorf("Invalid number of health mods: %d != 2",
			len(mods.HealthMods))
	}
	if len(mods.ManaMods) != 2 {
		t.Errorf("Inavlid number of mana mods: %d != 2",
			len(mods.ManaMods))
	}
	if len(mods.FlagMods) != 2 {
		t.Errorf("Invalid number of flag mods: %d != 2",
			len(mods.FlagMods))
	}
	if len(mods.QuestMods) != 2 {
		t.Errorf("Inavlid number of quest mods: %d != 2",
			len(mods.QuestMods))
	}
	if len(mods.AreaMods) != 2 {
		t.Errorf("Invalid number of area mods: %d != 2",
			len(mods.AreaMods))
	}
	if len(mods.ChapterMods) != 2 {
		t.Errorf("Invalid number of chapter mods: %d != 2",
			len(mods.ChapterMods))
	}
	if len(mods.AddItemMods) != 2 {
		t.Errorf("Invalid number of add item mods: %d != 2",
			len(mods.AddItemMods))
	}
	if len(mods.AddSkillMods) != 2 {
		t.Errorf("Invalid number of add skill mods: %d != 2",
			len(mods.AddSkillMods))
	}
	if len(mods.RemoveItemMods) != 2 {
		t.Errorf("Invalid number of remove item mods: %d != 2",
			len(mods.RemoveItemMods))
	}
	if len(mods.AttributeMods) != 2 {
		t.Errorf("Invalid number of attribute mods: %d != 2",
			len(mods.AttributeMods))
	}
	if len(mods.MemoryMods) != 2 {
		t.Errorf("Invalid number of memory mods: %d != 2",
			len(mods.MemoryMods))
	}
}

// TestModifiersDataXml tests modifiers data XML mappings.
func TestModifiersDataXml(t *testing.T) {
	// Create test modifiers
	data, err := testData("modifiers.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	mods := new(ModifiersData)
	err = xml.Unmarshal(data, mods)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	// Test
	if len(mods.HealthMods) != 2 {
		t.Errorf("Invalid number of health mods: %d != 2",
			len(mods.HealthMods))
	}
	if len(mods.ManaMods) != 2 {
		t.Errorf("Inavlid number of mana mods: %d != 2",
			len(mods.ManaMods))
	}
	if len(mods.FlagMods) != 2 {
		t.Errorf("Invalid number of flag mods: %d != 2",
			len(mods.FlagMods))
	}
	if len(mods.QuestMods) != 2 {
		t.Errorf("Inavlid number of quest mods: %d != 2",
			len(mods.QuestMods))
	}
	if len(mods.AreaMods) != 2 {
		t.Errorf("Invalid number of area mods: %d != 2",
			len(mods.AreaMods))
	}
	if len(mods.ChapterMods) != 2 {
		t.Errorf("Invalid number of chapter mods: %d != 2",
			len(mods.ChapterMods))
	}
	if len(mods.AddItemMods) != 2 {
		t.Errorf("Invalid number of add item mods: %d != 2",
			len(mods.AddItemMods))
	}
	if len(mods.AddSkillMods) != 2 {
		t.Errorf("Invalid number of add skill mods: %d != 2",
			len(mods.AddSkillMods))
	}
	if len(mods.RemoveItemMods) != 2 {
		t.Errorf("Invalid number of remove item mods: %d != 2",
			len(mods.RemoveItemMods))
	}
	if len(mods.AttributeMods) != 2 {
		t.Errorf("Invalid number of attribute mods: %d != 2",
			len(mods.AttributeMods))
	}
	if len(mods.MemoryMods) != 2 {
		t.Errorf("Invalid number of memory mods: %d != 2",
			len(mods.MemoryMods))
	}
}

// TestHealthModDataJson tests health modifier data JSON mappings.
func TestHealthModDataJson(t *testing.T) {
	// Create test modifier
	data, err := testData("healthmod.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	mod := new(HealthModData)
	err = json.Unmarshal(data, mod)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	// Test
	if mod.Min != 1 {
		t.Errorf("Invalid min value: %d != 1", mod.Min)
	}
	if mod.Max != 10 {
		t.Errorf("Invalid max value: %d != 10", mod.Max)
	}
}

// TestHealthModDataXxml tests health modifier data XML mappings.
func TestHealthModDataXml(t *testing.T) {
	// Create test modifier
	data, err := testData("healthmod.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	mod := new(HealthModData)
	err = xml.Unmarshal(data, mod)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	// Test
	if mod.Min != 1 {
		t.Errorf("Invalid min value: %d != 1", mod.Min)
	}
	if mod.Max != 10 {
		t.Errorf("Invalid max value: %d != 10", mod.Max)
	}
}

// TestChapterModDataJson tests chapter modifier data
// JSON mappings.
func TestChapterModDataJson(t *testing.T) {
	// Create test modifier
	data, err := testData("chaptermod.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	mod := new(ChapterModData)
	err = json.Unmarshal(data, mod)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	// Test
	if mod.ID != "ch1" {
		t.Errorf("Invalid ID value: %s != 'ch1'", mod.ID)
	}
}

// TestChapterModDataXml tests chapter modifier data
// XML mapping.
func TestChapterModDataXml(t *testing.T) {
	// Create test modifier
	data, err := testData("chaptermod.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	mod := new(ChapterModData)
	err = xml.Unmarshal(data, mod)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	// Test
	if mod.ID != "ch1" {
		t.Errorf("Invalid ID value: %s != 'ch1'", mod.ID)
	}
}
