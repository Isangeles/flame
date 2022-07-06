/*
 * object.go
 *
 * Copyright 2019-2022 Dariusz Sikora <ds@isangeles.dev>
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
	"sync"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/item"
	"github.com/isangeles/flame/objects"
	"github.com/isangeles/flame/serial"
	"github.com/isangeles/flame/useaction"
)

const (
	BaseSightRange = 300.0
)

// Struct for area objects.
type Object struct {
	id, serial      string
	hp, maxHP       int
	resilience      objects.Resilience
	posX, posY      float64
	sightRange      float64
	respawn         int64
	areaID          string
	action          *useaction.UseAction
	inventory       *item.Inventory
	effects         *sync.Map
	chatlog         *objects.Log
	onEffectTaken   func(e *effect.Effect)
	onModifierTaken func(m effect.Modifier)
}

// New creates new area object from
// specified data.
func New(data res.ObjectData) *Object {
	o := Object{
		inventory: item.NewInventory(),
		effects:   new(sync.Map),
		chatlog:   objects.NewLog(),
	}
	o.Apply(data)
	o.inventory.SetCapacity(10)
	// Register serial.
	serial.Register(&o)
	return &o
}

// Update updates object.
func (ob *Object) Update(delta int64) {
	// Effects.
	for _, e := range ob.Effects() {
		e.Update(delta)
		// Remove expired effects.
		if e.Time() <= 0 && !e.Infinite() {
			ob.effects.Delete(e.ID() + e.Serial())
		}
	}
	// Inventory.
	ob.Inventory().Update(delta)
	// Use action.
	if ob.UseAction() != nil {
		ob.UseAction().Update(delta)
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
	// Update ownerships.
	if ob.UseAction() != nil {
		ob.UseAction().SetOwner(ob)
	}
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

// Respawn returns object respawn time in milliseconds.
func (ob *Object) Respawn() int64 {
	return ob.respawn
}

// SetRespawn sets object respawn time in milliseconds.
func (ob *Object) SetRespawn(respawn int64) {
	ob.respawn = respawn
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

// SetAreaID sets area ID for object.
func (ob *Object) SetAreaID(id string) {
	ob.areaID = id
}

// AreaID returns ID of object area.
func (ob *Object) AreaID() string {
	return ob.areaID
}

// UseAction returns object use action.
func (ob *Object) UseAction() *useaction.UseAction {
	return ob.action
}

// Inventory returns object inventory.
func (ob *Object) Inventory() *item.Inventory {
	return ob.inventory
}

// AddEffects adds specified effect to objects.
func (ob *Object) AddEffect(e *effect.Effect) {
	e.SetTarget(ob)
	ob.effects.Store(e.ID()+e.Serial(), e)
}

// RemoveEffect removes specified effect from objects.
func (ob *Object) RemoveEffect(e *effect.Effect) {
	ob.effects.Delete(e.ID() + e.Serial())
}

// Effects returns all obejct effects.
func (ob *Object) Effects() (effects []*effect.Effect) {
	addEffect := func(k, v interface{}) bool {
		e, ok := v.(*effect.Effect)
		if ok {
			effects = append(effects, e)
		}
		return true
	}
	ob.effects.Range(addEffect)
	return
}

// SightRange returns object sight range.
func (ob *Object) SightRange() float64 {
	return ob.sightRange
}

// SetOnEffectTakenFunc sets function triggered on taking new effect.
func (ob *Object) SetOnEffectTakenFunc(f func(e *effect.Effect)) {
	ob.onEffectTaken = f
}

// SetOnModifierTakenFunc sets function triggered on taking new modifier.
func (ob *Object) SetOnModifierTakenFunc(f func(m effect.Modifier)) {
	ob.onModifierTaken = f
}

// ChatLog returns object speech log channel.
func (ob *Object) ChatLog() *objects.Log {
	return ob.chatlog
}
