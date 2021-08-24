/*
 * respawn.go
 *
 * Copyright 2021 Dariusz Sikora <dev@isangeles.pl>
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

package area

import (
	"time"

	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/object"
	"github.com/isangeles/flame/objects"
)

// Struct for area respawn.
type Respawn struct {
	area  *Area
	queue map[objects.Object]time.Time
}

// newRespawn creates respawn for area.
func newRespawn(area *Area) *Respawn {
	r := Respawn{
		area:  area,
		queue: make(map[objects.Object]time.Time),
	}
	return &r
}

// Update updates respawn.
func (r *Respawn) Update() {
	for _, char := range r.area.Characters() {
		_, inQueue := r.queue[char]
		if inQueue || char.Live() || char.Respawn() < 1 {
			continue
		}
		r.queue[char] = r.area.time.Add(time.Duration(char.Respawn()) * time.Millisecond)
	}
	for _, ob := range r.area.Objects() {
		_, inQueue := r.queue[ob]
		if inQueue || ob.Live() || ob.Respawn() < 1 {
			continue
		}
		r.queue[ob] = r.area.time.Add(time.Duration(ob.Respawn()) * time.Millisecond)
	}
	for ob, respTime := range r.queue {
		if respTime.Unix() > r.area.time.Unix() {
			continue
		}
		if char, ok := ob.(*character.Character); ok && !char.Live() {
			r.respawnChar(char)
		}
		if ob, ok := ob.(*object.Object); ok && !ob.Live() {
			r.respawnObject(ob)
		}
		delete(r.queue, ob)
	}
}

// respawnChar respawns specified character.
func (r *Respawn) respawnChar(char *character.Character) {
	charData := res.Character(char.ID(), "")
	if charData == nil {
		log.Err.Printf("Area: %s: respawn: %s: character data not found",
			r.area.ID(), char.ID())
		return
	}
	newChar := character.New(*charData)
	newChar.SetRespawn(char.Respawn())
	newChar.SetPosition(char.Position())
	newChar.SetDefaultPosition(char.DefaultPosition())
	r.area.AddCharacter(newChar)
}

// respawnObject respawns specified object.
func (r *Respawn) respawnObject(ob *object.Object) {
	obData := res.Object(ob.ID(), "")
	if obData == nil {
		log.Err.Printf("Area: %s: respawn: %s: object data not found",
			r.area.ID(), ob.ID())
		return
	}
	newOb := object.New(*obData)
	newOb.SetRespawn(ob.Respawn())
	newOb.SetPosition(ob.Position())
	r.area.AddObject(newOb)
}
