/*
 * character_test.go
 *
 * Copyright 2022-2025 Dariusz Sikora <ds@isangeles.dev>
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

var (
	charData        = res.CharacterData{ID: "char", Level: 1, Attributes: res.AttributesData{5, 5, 5, 5, 5}}
	dialogStageData = res.DialogStageData{ID: "dialogStage", Start: true}
	dialogData      = res.DialogData{ID: "dialog", Stages: []res.DialogStageData{dialogStageData}}
)

// TestLive tests live check function.
func TestLive(t *testing.T) {
	// Test live.
	ob := New(charData)
	if !ob.Live() {
		t.Errorf("Character is not live with full health")
	}
	// Test no live.
	ob.SetHealth(0)
	if ob.Live() {
		t.Errorf("Character is live with no health")
	}
}

// TestFighting tests fighting check function.
func TestFighting(t *testing.T) {
	// Create test objects.
	ob := New(charData)
	tar := New(charData)
	// Test no target.
	if ob.Fighting() {
		t.Errorf("Character in the combat with no target")
	}
	// Test attitude.
	ob.SetTarget(tar)
	if ob.Fighting() {
		t.Errorf("Character in the combat with non hostile target")
	}
	tar.SetAttitude(Hostile)
	if !ob.Fighting() {
		t.Errorf("Character not in the combat with hostile target")
	}
	// Test range.
	tar.SetPosition(ob.Attributes().Sight()+1, ob.Attributes().Sight()+1)
	if ob.Fighting() {
		t.Errorf("Character in the combat with out of range target")
	}
	// Test dead target.
	tar.SetPosition(0, 0)
	tar.SetHealth(0)
	if ob.Fighting() {
		t.Errorf("Character in the combat with dead target")
	}
}

// TestAttitudeFor tests function for checking attitude towards specific object.
func TestAttitudeFor(t *testing.T) {
	// Create test objects.
	ob := New(charData)
	tar := New(charData)
	// Test no memory.
	att := ob.AttitudeFor(tar)
	if att != tar.Attitude() {
		t.Errorf("Attitude is not target default attitude: %v != %v", att, tar.Attitude())
	}
	// Test memory.
	tarMem := TargetMemory{tar.ID(), tar.Serial(), Hostile}
	ob.MemorizeTarget(&tarMem)
	att = ob.AttitudeFor(tar)
	if att != Hostile {
		t.Errorf("Invalid attitude for memorized target: %v != %v", att, Hostile)
	}
	// Test dead target.
	tar.SetHealth(0)
	att = ob.AttitudeFor(tar)
	if att != Neutral {
		t.Errorf("Invalid attitude for dead target: %v != %v", att, Neutral)
	}
	// Test hostile for same guild.
	tar.SetHealth(1)
	ob.memory.Delete(tar.ID() + tar.Serial())
	ob.SetAttitude(Hostile)
	tar.SetAttitude(Hostile)
	guild := Guild{"test"}
	ob.SetGuild(guild)
	tar.SetGuild(guild)
	att = ob.AttitudeFor(tar)
	if att != Friendly {
		t.Errorf("Invalid attitude of hostile object for same guild target: %v != %v", att, Friendly)
	}
	// Test neutral object.
	tar.SetGuild(Guild{})
	tar.SetAttitude(Neutral)
	att = ob.AttitudeFor(tar)
	if att != Neutral {
		t.Errorf("Invalid attitude of hostile object for neutral target: %v != %v", att, Neutral)
	}
	// Test hostile object.
	tar.SetAttitude(Friendly)
	att = ob.AttitudeFor(tar)
	if att != Hostile {
		t.Errorf("Invalid attitude of hostile object for friendly target: %v != %v", att, Hostile)
	}
}

// TestDialog tests function for retrieving dialog.
func TestDialog(t *testing.T) {
	// Create test objects
	ob1 := New(charData)
	ob2 := New(charData)
	ob1.AddDialog(dialogData)
	// Test
	dialog := ob1.Dialog(ob2)
	if dialog == nil {
		t.Fatalf("New dialog was not returned")
	}
	startedDialog := ob1.Dialog(ob2)
	if startedDialog == nil || startedDialog != dialog {
		t.Fatalf("Previousely started dialog was not returned")
	}
	dialog.Restart()
	dialog.SetTarget(nil)
	ob1.Update(1) // update to trigger dialog cleanup
	dialog = ob1.Dialog(ob2)
	if dialog == nil || dialog == startedDialog {
		t.Fatalf("New dialog was not returned after ending old one")
	}
}
