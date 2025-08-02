/*
 * testTarget.go
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
	"github.com/isangeles/flame/flag"
	"github.com/isangeles/flame/serial"
)

// Testing struct to use to test effect targets.
type testTarget struct {
	Health int
	Flags  []flag.Flag
}

// Creates new testing target.
func newTestTarget() *testTarget {
	tt := new(testTarget)
	tt.Health = 100
	return tt
}

func (tt *testTarget) ID() string {
	return ""
}

func (tt *testTarget) SetSerial(s string) {
}

func (tt *testTarget) Serial() string {
	return ""
}

func (tt *testTarget) SetPosition(x, y float64) {

}

func (tt *testTarget) Position() (x, y float64) {
	return 0.0, 0.0
}

func (tt *testTarget) Effects() []*Effect {
	return make([]*Effect, 0)
}

func (tt *testTarget) HitEffects() []*Effect {
	return make([]*Effect, 0)
}

func (tt *testTarget) HitModifiers() []Modifier {
	return make([]Modifier, 0)
}

func (tt *testTarget) TakeEffect(e *Effect) {

}

func (tt *testTarget) RemoveEffect(e *Effect) {

}

// Handles the modifiers.
func (tt *testTarget) TakeModifiers(source serial.Serialer, mods ...Modifier) {
	for _, m := range mods {
		switch mod := m.(type) {
		case *FlagMod:
			tt.Flags = append(tt.Flags, mod.Flag())
		case *HealthMod:
			tt.Health -= mod.Min()
		}
	}
}
