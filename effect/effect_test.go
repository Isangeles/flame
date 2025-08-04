/*
 * effect.go
 *
 * Copyright 2025 Dariusz Sikora <ds@isangeles.dev>
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

package effect

import (
	"testing"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/serial"
)

// Test for applying effect modifiers through update function.
func TestEffectUpdateModifiers(t *testing.T) {
	// Create test target and effect
	ob := newTestTarget()
	serial.Register(ob)
	effData := res.EffectData{Duration: 1000}
	modData := res.FlagModData{ID: "flag", Off: false}
	effData.Modifiers.FlagMods = append(effData.Modifiers.FlagMods, modData)
	dotData := res.HealthModData{Min: 1, Max: 1}
	effData.OverTimeModifiers.HealthMods = append(effData.OverTimeModifiers.HealthMods, dotData)
	eff := New(effData)
	// Apply effect
	eff.SetTarget(ob)
	eff.Update(1)
	// Test modifiers
	if len(ob.Flags) != 1 {
		t.Errorf("No modifier flag after applying the effect")
	}
	// Test DOT modifiers
	if ob.Health != 99 {
		t.Fatalf("DOT modifier not applied on start")
	}
	eff.Update(1000)
	if ob.Health != 98 {
		t.Fatalf("DOT modifier not ticking")
	}
	// Test finished effect
	if eff.Time() > 0 {
		t.Fatalf("Effect not finished")
	}
	if len(ob.Flags) != 0 {
		t.Fatalf("Flag modifier not removed after finishing the effect")
	}
}
