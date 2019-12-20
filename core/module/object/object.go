/*
 * object.go
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

// Package for area objects.
package object

import (
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/effect"
	"github.com/isangeles/flame/core/module/flag"
	"github.com/isangeles/flame/core/module/item"
	"github.com/isangeles/flame/core/module/objects"
	"github.com/isangeles/flame/core/module/serial"
	"github.com/isangeles/flame/log"
)

// Struct for area objects.
type Object struct {
	id, serial string
	name       string
	hp, maxHP  int
	resilience objects.Resilience
	posX, posY float64
	action     *Action
	inventory  *item.Inventory
	effects    map[string]*effect.Effect
	flags      map[string]flag.Flag
	chatlog    chan string
	combatlog  chan string
	privatelog chan string
}

// New creates new area object from
// specified data.
func New(data res.ObjectData) *Object {
	ob := Object{
		id:     data.BasicData.ID,
		name:   data.BasicData.Name,
		maxHP:  data.BasicData.MaxHP,
	}
	ob.SetHealth(ob.MaxHealth())
	if len(data.BasicData.Action.SelfMods) > 0 || len(data.BasicData.Action.UserMods) > 0 {
		ob.action = NewAction(data.BasicData.Action)
	}
	ob.inventory = item.NewInventory(data.Inventory)
	ob.inventory.SetCapacity(10)
	ob.effects = make(map[string]*effect.Effect)
	ob.flags = make(map[string]flag.Flag)
	ob.chatlog = make(chan string, 1)
	ob.combatlog = make(chan string, 3)
	// Add effects.
	for _, data := range data.Effects {
		effData := res.Effect(data.ID)
		if effData == nil {
			log.Err.Printf("object: %s: fail to retrieve effect data: %s",
				ob.ID(), data.ID)
			continue
		}
		eff := effect.New(*effData)
		ob.AddEffect(eff)
	}
	// Assign serial.
	serial.AssignSerial(&ob)
	return &ob
}

// Update updates object.
func (ob *Object) Update(delta int64) {
	// Effects.
	for serial, e := range ob.effects {
		e.Update(delta)
		// Remove expired effects.
		if e.Time() <= 0 {
			delete(ob.effects, serial)
		}
	}
}

// ID returns object ID.
func (ob *Object) ID() string {
	return ob.id
}

// Serial returns object serial
// value.
func (ob *Object) Serial() string {
	return ob.serial
}

// SetSerial sets specified text as
// object serial value.
func (ob *Object) SetSerial(s string) {
	ob.serial = s
}

// Name returns object name.
func (ob *Object) Name() string {
	return ob.name
}

// SetName sets object name.
func (ob *Object) SetName(s string) {
	ob.name = s
}

// Position returns object XY position.
func (ob *Object) Position() (float64, float64) {
	return ob.posX, ob.posY
}

// SetPosition sets specified values as
// XY position.
func (ob *Object) SetPosition(x, y float64) {
	ob.posX, ob.posY = x, y
}

// Health retruns current object
// health value.
func (ob *Object) Health() int {
	return ob.hp
}

// SetHealth sets specified value
// as current object health.
func (ob *Object) SetHealth(h int) {
	ob.hp = h
}

// MaxHealth returns maximal health
// value.
func (ob *Object) MaxHealth() int {
	return ob.maxHP
}

// SetMaxHealth sets specified value
// as maximal health value.
func (ob *Object) SetMaxHealth(h int) {
	ob.maxHP = h
}

// Live checks whether object is
// not destroed(HP higher then 0).
func (ob *Object) Live() bool {
	return ob.Health() > 0
}

// Mana returns 0, object do not
// have mana. Function to satisfy
// effect targer interface.
func (ob *Object) Mana() int {
	return 0
}

// MaxMana returns 0, object do not
// have mana. Function to satisfy
// effect targer interface.
func (ob *Object) MaxMana() int {
	return 0
}

// SetMana does nothing, object do not
// have mana. Function to satisfy
// effect targer interface.
func (ob *Object) SetMana(v int) {
}

// Experience returns 0, object do not
// have experience. Function to satisfy
// effect targer interface.
func (ob *Object) Experience() int {
	return 0
}

// SetExperience does nothing, object do
// not have experience. Function to satisfy
// effect targer interface.
func (ob *Object) SetExperience(v int) {
}

// Action returns object action struct.
func (ob *Object) Action() *Action {
	return ob.action
}

// Inventory returns object inventory.
func (ob *Object) Inventory() *item.Inventory {
	return ob.inventory
}

// AddEffects adds specified effect to objects.
func (ob *Object) AddEffect(e *effect.Effect) {
	e.SetTarget(ob)
	ob.effects[e.ID()+e.Serial()] = e
}

// RemoveEffect removes specified effect from objects.
func (ob *Object) RemoveEffect(e *effect.Effect) {
	delete(ob.effects, e.ID()+e.Serial())
}

// Effects returns all obejct effects.
func (ob *Object) Effects() []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	for _, e := range ob.effects {
		effects = append(effects, e)
	}
	return effects
}

// AddFlag adds specified flag.
func (ob *Object) AddFlag(f flag.Flag) {
	ob.flags[f.ID()] = f
}

// Flags returns all object flags.
func (ob *Object) Flags() (flags []flag.Flag) {
	for _, f := range ob.flags {
		flags = append(flags, f)
	}
	return
}

// SendCmb sends specified text message to
// comabt log channel.
func (ob *Object) SendCombat(msg string) {
	select {
	case ob.combatlog <- msg:
	default:
	}
}

// CombatLog returns object combat log
// channel.
func (ob *Object) CombatLog() chan string {
	return ob.combatlog
}

// ChatLog returns object speech log channel.
func (ob *Object) ChatLog() chan string {
	return ob.chatlog
}

// PrivateLog returns object private log channel.
func (ob *Object) PrivateLog() chan string {
	return ob.privatelog
}
