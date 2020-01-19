/*
 * armor.go
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
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/data/res/lang"
	"github.com/isangeles/flame/core/module/effect"
	"github.com/isangeles/flame/core/module/req"
	"github.com/isangeles/flame/core/module/serial"
)

// Struct for armor items.
type Armor struct {
	id        string
	name      string
	info      string
	serial    string
	level     int
	value     int
	loot      bool
	armor     int
	eqEffects []res.EffectData
	eqReqs    []req.Requirement
	slots     []Slot
}

// NewArmor creates new armor from specified
// armor data.
func NewArmor(data res.ArmorData) *Armor {
	a := Armor{
		id:        data.ID,
		level:     data.Level,
		value:     data.Value,
		loot:      data.Loot,
		armor:     data.Armor,
		eqEffects: data.EQEffects,
	}
	a.eqReqs = req.NewRequirements(data.EQReqs...)
	for _, sid := range data.Slots {
		a.slots = append(a.slots, Slot(sid))

	}
	nameInfo := lang.Texts(a.ID())
	a.name = nameInfo[0]
	if len(nameInfo) > 1 {
		a.info = nameInfo[1]
	}
	serial.AssignSerial(&a)
	return &a
}

// ID returns armor ID.
func (a *Armor) ID() string {
	return a.id
}

// Name returns armor name.
func (a *Armor) Name() string {
	return a.name
}

// Info returns armor info.
func (a *Armor) Info() string {
	return a.info
}

// Serial returns armor serial value.
func (a *Armor) Serial() string {
	return a.serial
}

// SetSerial sets specified value as serial
// value of armor.
func (a *Armor) SetSerial(s string) {
	a.serial = s
}

// Level returns armor level.
func (a *Armor) Level() int {
	return a.level
}

// Value returns armor value.
func (a *Armor) Value() int {
	return a.value
}

// Armor returns armor rating
// value.
func (a *Armor) Armor() int {
	return a.armor
}

// EquipEffects returns armor equip effects
func (a *Armor) EquipEffects() (effs []*effect.Effect) {
	for _, ed := range a.eqEffects {
		e := effect.New(ed)
		effs = append(effs, e)
	}
	return
}

// EquipReqs returns armor equip requirements.
func (a *Armor) EquipReqs() []req.Requirement {
	return a.eqReqs
}

// Slots returns types of slots occupated
// by this armor after equipping.
func (a *Armor) Slots() []Slot {
	return a.slots
}

// Loot checks if armor is 'lootable'.
func (a *Armor) Loot() bool {
	return a.loot
}
