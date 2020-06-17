/*
 * misc.go
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

package item

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/serial"
)

// Struct for miscellaneous items.
type Misc struct {
	id         string
	name, info string
	serial     string
	value      int
	level      int
	loot       bool
	currency   bool
	consumable bool
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
		useMods:    effect.NewModifiers(data.UseMods),
	}
	// Name & info.
	nameInfo := lang.Texts(m.ID())
	m.name = nameInfo[0]
	if len(nameInfo) > 1 {
		m.info = nameInfo[1]
	}
	// Use effects.
	for _, ed := range data.UseEffects {
		data := res.Effect(ed.ID)
		if data != nil {
			m.useEffects = append(m.useEffects, *data)
		}
	}
	// Serial.
	serial.Register(&m)
	return &m
}

// ID returns item ID.
func (m *Misc) ID() string {
	return m.id
}

// Name returns item name.
func (m *Misc) Name() string {
	return m.name
}

// Info returns item info.
func (m *Misc) Info() string {
	return m.info
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

// Loot checks if item is 'lootable'.
func (m *Misc) Loot() bool {
	return m.loot
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

// Use applies item use modifiers on specified
// user.
func (m *Misc) Use(user effect.Target) {
	for _, ed := range m.useEffects {
		e := effect.New(ed)
		e.SetSource(m.ID(), m.Serial())
		user.TakeEffect(e)
	}
	user.TakeModifiers(nil, m.useMods...)
}
