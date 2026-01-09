/*
 * spawn.go
 *
 * Copyright 2021-2026 Dariusz Sikora <ds@isangeles.dev>
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

// Struct for area spawn.
type Spawn struct {
	area         *Area
	respawnQueue *sync.Map
	despawnQueue *sync.Map
}

// newSpawn creates spawn for the area.
func newSpawn(area *Area) *Spawn {
	r := Spawn{
		area:  area,
		respawnQueue: new(sync.Map),
		despawnQueue: new(sync.Map),
	}
	return &r
}

// Update updates spawn.
func (r *Spawn) Update() {
	// Fill the queues
	for _, ob := range r.area.Objects() {
		_, inQueue := r.respawnQueue.Load(ob)
		if !inQueue && !ob.Live() && ob.Respawn() > 0 {
			r.respawnQueue.Store(ob, r.area.Time.Add(time.Duration(ob.Respawn())*time.Millisecond))
		}
		_, inQueue = r.despawnQueue.Load(ob)
		if !inQueue && ob.OpenLoot() && len(ob.Inventory().Items()) < 1 && ob.Despawn() > 0 {
			r.despawnQueue.Store(ob, r.area.Time.Add(time.Duration(ob.Despawn())*time.Millisecond))
		}
	}
	// Respawn
	respObject := func(k, v interface{}) bool {
		ob, keyOk := k.(serial.Serialer)
		respTime, valueOk := v.(time.Time)
		if !keyOk || !valueOk || respTime.Unix() > r.area.Time.Unix(){
			return true
		}
		if char, ok := ob.(*character.Character); ok && !char.Live() {
			r.respawnChar(char)
		}
		r.respawnQueue.Delete(ob)
		return true
	}
	r.respawnQueue.Range(respObject)
	// Despawn
	despObject := func(k, v interface{}) bool {
		ob, keyOk := k.(serial.Serialer)
		despTime, valueOk := v.(time.Time)
		if !keyOk || !valueOk || despTime.Unix() > r.area.Time.Unix() {
			return true
		}
		if char, ok := ob.(*character.Character); ok {
			r.area.RemoveObject(char)
		}
		r.despawnQueue.Delete(ob)
		return true
	}
	r.despawnQueue.Range(despObject)
}

// Apply applies respawn data.
func (r *Spawn) Apply(data res.SpawnData) {
	r.respawnQueue = new(sync.Map)
	for _, ob := range data.RespawnQueue {
		areaOb, _ := r.area.objects.Load(ob.ID + ob.Serial)
		if _, ok := areaOb.(*character.Character); ok {
			r.respawnQueue.Store(time.Unix(ob.Time, 0), areaOb)
			continue
		}
	}
	for _, ob := range data.DespawnQueue {
		areaOb, _ := r.area.objects.Load(ob.ID + ob.Serial)
		if _, ok := areaOb.(*character.Character); ok {
			r.despawnQueue.Store(time.Unix(ob.Time, 0), areaOb)
			continue
		}
	}
}

// Data returns data resource for respawn.
func (r *Spawn) Data() res.SpawnData {
	var data res.SpawnData
	addSpawnObject := func(k, v interface{}) bool {
		ob, keyOk := k.(serial.Serialer)
		time, valueOk := v.(time.Time)
		if !keyOk || !valueOk {
			return true
		}
		obData := res.SpawnObject{
			SerialObjectData: res.SerialObjectData{ob.ID(), ob.Serial()},
			Time:             time.Unix(),
		}
		data.RespawnQueue = append(data.RespawnQueue, obData)
		return true
	}
	r.respawnQueue.Range(addSpawnObject)
	addDespawnObject := func(k, v interface{}) bool {
		ob, keyOk := k.(serial.Serialer)
		time, valueOk := v.(time.Time)
		if !keyOk || !valueOk {
			return true
		}
		obData := res.SpawnObject{
			SerialObjectData: res.SerialObjectData{ob.ID(), ob.Serial()},
			Time:             time.Unix(),
		}
		data.DespawnQueue = append(data.DespawnQueue, obData)
		return true
	}
	r.despawnQueue.Range(addDespawnObject)
	return data
}

// respawnChar respawns specified character.
func (r *Spawn) respawnChar(char *character.Character) {
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
