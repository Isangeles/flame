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
package area

import (
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/core/module/object/item"
)

// Struct for area objects.
type Object struct {
	id, serial string
	name       string
	hp, maxHP  int
	resilience object.Resilience
	posX, posY float64
	inventory  *item.Inventory
	combatlog  chan string
}

// New creates new area object from
// specified data.
func NewObject(data res.ObjectBasicData) *Object {
	ob := Object{
		id:     data.ID,
		serial: data.Serial,
		hp:     data.HP,
		maxHP:  data.MaxHP,
	}
	return &ob
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

// Inventory returns object inventory.
func (ob *Object) Inventory() *item.Inventory {
	return ob.inventory
}
