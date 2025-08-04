/*
 * effect.go
 *
 * Copyright 2019-2025 Dariusz Sikora <ds@isangeles.dev>
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

// Package for effects.
package effect

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/serial"
)

// Struct for effects.
type Effect struct {
	id, serial string
	source     res.SerialObjectData
	target     res.SerialObjectData
	mods       []Modifier
	dotMods    []Modifier
	duration   int64
	time       int64
	secTimer   int64
	meleeHit   bool
	infinite   bool
	hostile    bool
	started    bool
}

// New creates new effect.
func New(data res.EffectData) *Effect {
	e := new(Effect)
	e.id = data.ID
	e.mods = NewModifiers(data.Modifiers)
	e.dotMods = NewModifiers(data.OverTimeModifiers)
	e.duration = int64(data.Duration)
	e.meleeHit = data.MeleeHit
	e.infinite = data.Infinite
	e.hostile = data.Hostile
	e.SetTime(data.Duration)
	serial.Register(e)
	return e
}

// Update updates effect.
func (e *Effect) Update(delta int64) {
	if e.started && e.Time() <= 0 && !e.Infinite() {
		log.Err.Printf("effect: %s#%s: no time left: %d", e.ID(),
			e.Serial(), e.duration)
		return
	}
	// Fetch target and source objects
	object := serial.Object(e.target.ID, e.target.Serial)
	if object == nil {
		log.Err.Printf("effect: %s#%s: target not found: %s#%s",
			e.ID(), e.Serial(), e.target.ID, e.target.Serial)
		return
	}
	target, ok := object.(Target)
	if !ok {
		log.Err.Printf("effect: %s#%s: target is invalid: %s#%s",
			e.ID(), e.Serial(), e.target.ID, e.target.Serial)
		return
	}
	source := serial.Object(e.source.ID, e.source.Serial)
	// Apply modifiers
	if !e.started {
		target.TakeModifiers(source, e.mods...)
	}
	// Apply over-time modifiers
	e.secTimer += delta
	if !e.started || e.secTimer >= 1000 { // at start and every second after that
		target.TakeModifiers(source, e.dotMods...)
		e.secTimer = 0
	}
	e.started = true
	if !e.Infinite() {
		e.time -= delta
	}
	if !e.Infinite() && e.Time() <= 0 {
		target.RemoveModifiers(source, e.mods...)
	}
}

// ID returns effect ID.
func (e *Effect) ID() string {
	return e.id
}

// Serial returns effect serial value.
func (e *Effect) Serial() string {
	return e.serial
}

// Duration returns effect duration time
// in milliseconds.
func (e *Effect) Duration() int64 {
	return e.duration
}

// Time returns current duration time in
// milliseconds.
func (e *Effect) Time() int64 {
	return e.time
}

// MeleeHit checks if this effect is a melee hit.
func (e *Effect) MeleeHit() bool {
	return e.meleeHit
}

// Infinite checks if effect duration time is infinite.
func (e *Effect) Infinite() bool {
	return e.infinite
}

// Hostile checks if effect is an hostile action towards the target.
func (e *Effect) Hostile() bool {
	return e.hostile
}

// Source returns ID and serial value of effect
// source object.
func (e *Effect) Source() (string, string) {
	return e.source.ID, e.source.Serial
}

// SetSerial sets specified value as
// effect serial value.
func (e *Effect) SetSerial(serial string) {
	e.serial = serial
}

// SetTime sets specified value as effect
// duration time in milliseconds.
func (e *Effect) SetTime(time int64) {
	e.time = time
}

// SetSource sets targetable object with specified ID
// and serial value as effect source.
func (e *Effect) SetSource(id, serial string) {
	e.source.ID, e.source.Serial = id, serial
}

// SetTarget sets specified targertable object
// as effect target.
func (e *Effect) SetTarget(t Target) {
	e.target.ID, e.target.Serial = t.ID(), t.Serial()
}

// Data creates data resource for effect.
func (e *Effect) Data() res.EffectData {
	data := res.EffectData{
		ID:        e.ID(),
		Duration:  e.Duration(),
		Modifiers: ModifiersData(e.mods...),
		OverTimeModifiers: ModifiersData(e.dotMods...),
	}
	return data
}
