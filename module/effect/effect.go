/*
 * effect.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/serial"
	"github.com/isangeles/flame/log"
)

// Struct for effects.
type Effect struct {
	id, serial       string
	name             string
	srcID, srcSerial string
	tarID, tarSerial string
	modifiers        []Modifier
	duration         int64
	time             int64
	secTimer         int64
}

// New creates new effect.
func New(data res.EffectData) *Effect {
	e := new(Effect)
	e.id = data.ID
	e.name = lang.Text(e.ID())
	e.modifiers = NewModifiers(data.Modifiers)
	e.duration = int64(data.Duration)
	e.SetTime(data.Duration)
	if len(e.name) < 1 {
		e.name = lang.Text(e.id)
	}
	serial.Register(e)
	return e
}

// Update updates effect.
func (e *Effect) Update(delta int64) {
	object := serial.Object(e.tarID, e.tarSerial)
	if object == nil || e.Time() <= 0 {
		log.Err.Printf("effect: %s#%s: target not found: %s#%s",
			e.ID(), e.Serial(), e.tarID, e.tarSerial)
		return
	}
	target, ok := object.(Target)
	if !ok {
		log.Err.Printf("effect: %s#%s: target is invalid: %s#%s",
			e.ID(), e.Serial(), e.tarID, e.tarSerial)
		return
	}
	source := serial.Object(e.srcID, e.srcSerial)
	e.secTimer += delta
	if e.time == e.duration || e.secTimer >= 1000 { // at start and every second after that
		target.TakeModifiers(source, e.modifiers...)
		e.secTimer = 0
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

// Source returns ID and serial value of effect
// source object.
func (e *Effect) Source() (string, string) {
	return e.srcID, e.srcSerial
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

// SetTime sets specified value as effect
// duration time in milliseconds.
func (e *Effect) SetTime(time int64) {
	e.time = time
}

// SetSource sets specified targetable object
// as effect source.
func (e *Effect) SetSource(t Target) {
	e.srcID, e.srcSerial = t.ID(), t.Serial()
}

// SetTarget sets specified targertable object
// as effect target.
func (e *Effect) SetTarget(t Target) {
	e.tarID, e.tarSerial = t.ID(), t.Serial()
}

// Data creates data resource for effect.
func (e *Effect) Data() res.EffectData {
	data := res.EffectData{
		ID:        e.ID(),
		Duration:  e.Duration(),
		Modifiers: ModifiersData(e.modifiers...),
	}
	return data
}
