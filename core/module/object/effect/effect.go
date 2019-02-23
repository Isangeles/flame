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
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/modifier"
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/log"
)

// Struct for effects.
type Effect struct {
	id, serial string
	name       string
	source     object.Target
	target     object.Target
	modifiers  []modifier.Modifier
	subeffects []string
	duration   int64
	time       int64
	sec_timer  int64
}

// NewEffect creates new effect.
func NewEffect(data res.EffectData) *Effect {
	e := new(Effect)
	e.id = data.ID
	e.modifiers = data.Modifiers
	e.subeffects = data.Subeffects
	e.duration = int64(data.Duration * 1000)
	e.SetTimeSeconds(data.Duration)
	return e
}

// Update updates effect.
func (e *Effect) Update(delta int64) {
	if e.target == nil || e.Time() <= 0 {
		return
	}
	e.sec_timer += delta
	if e.sec_timer >= 1000 { // every second
		for _, m := range e.modifiers {
			m.Affect(e.source, e.target)
		}
		e.sec_timer = 0
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
func (e *Effect) SetSource(t object.Target) {
	e.source = t
}

// SetTarget sets specified targertable object
// as effect target.
func (e *Effect) SetTarget(t object.Target) {
	e.target = t
}

// SubEffects creates and returns all subeffects.
func (e *Effect) SubEffects() []*Effect {
	subeffects := make([]*Effect, 0)
	for _, eid := range e.subeffects {
		data := res.Effect(eid)
		if data.ID == "" {
			log.Err.Printf("effect:%s_%s:sub_effect_resource_not_found:%s",
				e.ID(), e.Serial(), eid)
		}
		subeff := NewEffect(data)
		subeffects = append(subeffects, subeff)
	}
	return subeffects
}
