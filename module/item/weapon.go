/*
 * weapon.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/req"
	"github.com/isangeles/flame/module/serial"
)

// Struct for weapons.
type Weapon struct {
	id             string
	name, info     string
	serial         string
	value          int
	level          int
	loot           bool
	dmgMin, dmgMax int
	dmgType        objects.Element
	dmgEffects     []res.EffectData
	equipReqs      []req.Requirement
	slots          []Slot
}

// NewWeapon creates new weapon with
// specified parameters.
func NewWeapon(data res.WeaponData) *Weapon {
	w := Weapon{
		id:         data.ID,
		value:      data.Value,
		level:      data.Level,
		loot:       data.Loot,
		dmgMin:     data.DMGMin,
		dmgMax:     data.DMGMax,
		dmgType:    objects.Element(data.DMGType),
		dmgEffects: data.DMGEffects,
	}
	w.equipReqs = req.NewRequirements(data.EQReqs...)
	for _, sid := range data.Slots {
		w.slots = append(w.slots, Slot(sid))
	}
	nameInfo := lang.Texts(w.ID())
	w.name = nameInfo[0]
	if len(nameInfo) > 1 {
		w.info = nameInfo[1]
	}
	serial.AssignSerial(&w)
	return &w
}

// ID returns weapon ID.
func (w *Weapon) ID() string {
	return w.id
}

// Name returns weapon name.
func (w *Weapon) Name() string {
	return w.name
}

// Info return weapon info.
func (w *Weapon) Info() string {
	return w.info
}

// Serial returns weapon serial value.
func (w *Weapon) Serial() string {
	return w.serial
}

// SetSerial sets specified value as serial
// value of weapon.
func (w *Weapon) SetSerial(s string) {
	w.serial = s
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

// DamageType retruns weapon damage type.
func (w *Weapon) DamageType() objects.Element {
	return w.dmgType
}

// DamageEffects returns weapon hit effects.
func (w *Weapon) DamageEffects() []res.EffectData {
	return w.dmgEffects
}

// EquipReqs returns weapon equip requirements.
func (w *Weapon) EquipReqs() []req.Requirement {
	return w.equipReqs
}

// Slots returns type of slots occupated
// by this weapon after equipping.
func (w *Weapon) Slots() []Slot {
	return w.slots
}

// Loot checks whether weapon is 'lootable'.
func (w *Weapon) Loot() bool {
	return w.loot
}
