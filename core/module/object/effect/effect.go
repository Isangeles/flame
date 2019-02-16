/*
 * effect.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/core/module/modifier"
)

// Struct for effects.
type Effect struct {
	id, serial string
	name       string
	source     modifier.Target
	target     modifier.Target
	modifiers  []modifier.Modifier
	duration   int64
	time       int64
}

// NewEffect creates new effect.
func NewEffect(id string, modifiers []modifier.Modifier, duration int64) *Effect {
	e := new(Effect)
	e.id = id
	e.modifiers = modifiers
	e.duration = int64(duration * 1000)
	e.SetTimeSeconds(duration)
	return e
}

// Update updates effect.
func (e *Effect) Update(delta int64) {
	if e.target == nil || e.Time() <= 0 {
		return
	}
	for _, m := range e.modifiers {
		m.Affect(e.target)
	}
	e.time -= delta
}

// ID returns effect ID.
func (e *Effect) ID() string {
	return e.id
}

// Serial returns effect serial value.
func (e *Effect) Serial() string {
	return e.serial
}

// Name returns effect display name.
func (e *Effect) Name() string {
	return e.name
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

// SetSerial sets specified value as
// effect serial value.
func (e *Effect) SetSerial(serial string) {
	e.serial = serial
}

// SetName sets specified text as effect
// display name.
func (e *Effect) SetName(name string) {
	e.name = name
}

// SetTimeSeconds sets specified value as effect
// duration time in seconds.
func (e *Effect) SetTimeSeconds(time int64) {
	e.time = int64(time * 1000)
}

// SetSource sets specified targetable object
// as effect source.
func (e *Effect) SetSource(t modifier.Target) {
	e.source = t
}

// SetTarget sets specified targertable object
// as effect target.
func (e *Effect) SetTarget(t modifier.Target) {
	e.target = t
}
