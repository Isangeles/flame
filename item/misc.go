/*
 * misc.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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

package item

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/serial"
	"github.com/isangeles/flame/useaction"
)

// Struct for miscellaneous items.
type Misc struct {
	id         string
	serial     string
	value      int
	level      int
	loot       bool
	currency   bool
	consumable bool
	useAction  *useaction.UseAction
	useEffects []res.EffectData
	useMods    []effect.Modifier
}

// NewMisc creates new misc item.
func NewMisc(data res.MiscItemData) *Misc {
	m := Misc{
		id:         data.ID,
		value:      data.Value,
		loot:       data.Loot,
		currency:   data.Currency,
		consumable: data.Consumable,
	}
	// Serial.
	serial.Register(&m)
	// Use action.
	m.useAction = useaction.New(data.UseAction)
	m.useAction.SetOwner(&m)
	return &m
}

// Update updates item.
func (m *Misc) Update(delta int64) {
	if m.UseAction() != nil {
		m.UseAction().Update(delta)
	}
}

// ID returns item ID.
func (m *Misc) ID() string {
	return m.id
}

// Serial returns item serial value.
func (m *Misc) Serial() string {
	return m.serial
}

// SetSerial sets item serial value.
func (m *Misc) SetSerial(s string) {
	m.serial = s
}

// Value return item value.
func (m *Misc) Value() int {
	return m.value
}

// Level return item level.
func (m *Misc) Level() int {
	return m.level
}

// Currency check if item can be
// used as currency.
func (m *Misc) Currency() bool {
	return m.currency
}

// Consumable checks if item should be
// deleted after use.
func (m *Misc) Consumable() bool {
	return m.consumable
}

// UseAction returns items use action.
func (m *Misc) UseAction() *useaction.UseAction {
	return m.useAction
}
