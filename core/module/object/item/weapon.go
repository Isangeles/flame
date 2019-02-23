/*
 * weapon.go
 *
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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
	"fmt"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/req"
)

// Struct for weapons.
type Weapon struct {
	id             string
	serial         string
	name           string
	value          int
	level          int
	dmgMin, dmgMax int
	equipReqs      []req.Requirement
	slots          []Slot
}

// NewWeapon creates new weapon with
// specified parameters.
func NewWeapon(data res.WeaponData) *Weapon {
	w := Weapon{
		id: data.ID,
		value: data.Value,
		dmgMin: data.DMGMin,
		dmgMax: data.DMGMax,
		level: data.Level,
		equipReqs: data.EQReqs,
	}
	for _, sid := range data.Slots {
		w.slots = append(w.slots, Slot(sid))
	}
	return &w
}

// ID returns weapon ID.
func (w *Weapon) ID() string {
	return w.id
}

// Serial returns weapon serial value.
func (w *Weapon) Serial() string {
	return w.serial
}

// SerialID returns weapon ID and serial
// separated by '_'.
func (w *Weapon) SerialID() string {
	return fmt.Sprintf("%s_%s", w.ID(), w.Serial())
}

// SetSerial sets specified value as serial
// value of weapon.
func (w *Weapon) SetSerial(serial string) {
	w.serial = serial
}

// Name returns item name.
func (w *Weapon) Name() string {
	return w.name
}

// SetName sets specified name as
// item display name.
func (w *Weapon) SetName(name string) {
	w.name = name
}

// Value returns item value.
func (w *Weapon) Value() int {
	return w.value
}

// Level returns item level.
func (w *Weapon) Level() int {
	return w.level
}

// Damge returns minimal and maximal
// damge values.
func (w *Weapon) Damage() (int, int) {
	return w.dmgMin, w.dmgMax
}

// EquipReqs returns weapon equip requirements.
func (w *Weapon) EquipReqs() []req.Requirement {
	return w.equipReqs
}

// Slots returns type of slots occupated by
// this weapon.
func (w *Weapon) Slots() []Slot {
	return w.slots
}
