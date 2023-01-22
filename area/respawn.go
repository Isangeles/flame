/*
 * respawn.go
 *
 * Copyright 2021-2023 Dariusz Sikora <ds@isangeles.dev>
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
	"sync"
	"time"

	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/serial"
)

// Struct for area respawn.
type Respawn struct {
	area  *Area
	queue *sync.Map
}

// newRespawn creates respawn for area.
func newRespawn(area *Area) *Respawn {
	r := Respawn{
		area:  area,
		queue: new(sync.Map),
	}
	return &r
}

// Update updates respawn.
func (r *Respawn) Update() {
	for _, ob := range r.area.Objects() {
		_, inQueue := r.queue.Load(ob)
		if inQueue || ob.Live() || ob.Respawn() < 1 {
			continue
		}
		r.queue.Store(ob, r.area.Time.Add(time.Duration(ob.Respawn())*time.Millisecond))
	}
	respObject := func(k, v interface{}) bool {
		ob, keyOk := k.(serial.Serialer)
		respTime, valueOk := v.(time.Time)
		if keyOk && valueOk {
			if respTime.Unix() > r.area.Time.Unix() {
				return true
			}
			if char, ok := ob.(*character.Character); ok && !char.Live() {
				r.respawnChar(char)
			}
			r.queue.Delete(ob)
		}
		return true
	}
	r.queue.Range(respObject)
}

// Apply applies respawn data.
func (r *Respawn) Apply(data res.RespawnData) {
	r.queue = new(sync.Map)
	for _, ob := range data.Queue {
		areaOb, _ := r.area.objects.Load(ob.ID + ob.Serial)
		if _, ok := areaOb.(*character.Character); ok {
			r.queue.Store(time.Unix(ob.Time, 0), areaOb)
			continue
		}
	}
}

// Data returns data resource for respawn.
func (r *Respawn) Data() res.RespawnData {
	var data res.RespawnData
	addObject := func(k, v interface{}) bool {
		ob, keyOk := k.(serial.Serialer)
		time, valueOk := v.(time.Time)
		if keyOk && valueOk {
			obData := res.RespawnObject{
				SerialObjectData: res.SerialObjectData{ob.ID(), ob.Serial()},
				Time:             time.Unix(),
			}
			data.Queue = append(data.Queue, obData)
		}
		return true
	}
	r.queue.Range(addObject)
	return data
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
	newChar.SetPosition(char.DefaultPosition())
	newChar.SetDefaultPosition(char.DefaultPosition())
	for _, f := range char.Flags() {
		newChar.AddFlag(f)
	}
	r.area.AddObject(newChar)
	r.area.RemoveObject(char)
}
