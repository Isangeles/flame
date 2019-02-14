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

package effect

import (
)

// Interface for 'targetable' objects.
type Target interface {
	Health() int
	SetHealth(val int)
	Mana() int
	SetMana(val int)
	Position() (float64, float64)
	SetPosition(x, y float64)
	Live() bool
	Effects() *Effects
	Targets() []Target
}

// Interface for effects.
type Effect interface {
	ID() string
	Serial() string
	SerialID() string
	Affect(owner Target, targets []Target)
	Update(delta int64)
	Expired() bool
}

// Struct for effects container.
type Effects struct {
	owner   Target
	effects map[string]Effect
}

// NewEffects creates new effects container.
func NewEffects(owner Target) *Effects {
	effs := new(Effects)
	effs.owner = owner
	effs.effects = make(map[string]Effect)
	return effs
}

// Add adds specified effect to effects
// container.
func (effs *Effects) Add(effects ...Effect) {
	for _, e := range effects {
		effs.effects[e.SerialID()] = e
	}
}

// Effects returns all effects from container.
func (effs *Effects) Effects() map[string]Effect {
	return effs.effects
}

// Updated updates all effects.
func (effs *Effects) Update(delta int64) {
	for sid, e := range effs.effects {
		e.Update(delta)
		if e.Expired() {
			delete(effs.effects, sid)
		}
		e.Affect(effs.owner, effs.owner.Targets())
	}
}

