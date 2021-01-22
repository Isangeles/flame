/*
 * useaction.go
 *
 * Copyright 2020-2021 Dariusz Sikora <dev@isangeles.pl>
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

// Package with use action struct for usable objects.
package useaction

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/req"
	"github.com/isangeles/flame/module/serial"
)

// Interface for usable game objects.
type Usable interface {
	ID() string
	UseAction() *UseAction
}

// Struct for use action of usable object.
type UseAction struct {
	castMax           int64
	cast              int64
	cooldownMax       int64
	cooldown          int64
	userMods          []effect.Modifier
	objectMods        []effect.Modifier
	targetMods        []effect.Modifier
	targetUserMods    []effect.Modifier
	userEffects       []res.EffectData
	objectEffects     []res.EffectData
	targetEffects     []res.EffectData
	targetUserEffects []res.EffectData
	requirements      []req.Requirement
	owner             res.SerialObjectData
}

// New creates new use action.
func New(data res.UseActionData) *UseAction {
	ua := UseAction{
		castMax:        data.CastMax,
		cast:           data.Cast,
		cooldownMax:    data.CooldownMax,
		cooldown:       data.Cooldown,
		userMods:       effect.NewModifiers(data.UserMods),
		objectMods:     effect.NewModifiers(data.ObjectMods),
		targetMods:     effect.NewModifiers(data.TargetMods),
		targetUserMods: effect.NewModifiers(data.TargetUserMods),
		requirements:   req.NewRequirements(data.Requirements),
	}
	for _, ed := range data.UserEffects {
		data := res.Effect(ed.ID)
		if data == nil {
			log.Err.Printf("use action: effect not found: %s", ed.ID)
			continue
		}
		ua.userEffects = append(ua.userEffects, *data)
	}
	for _, ed := range data.ObjectEffects {
		data := res.Effect(ed.ID)
		if data == nil {
			log.Err.Printf("use action: effect not found: %s", ed.ID)
			continue
		}
		ua.objectEffects = append(ua.objectEffects, *data)
	}
	for _, ed := range data.TargetEffects {
		data := res.Effect(ed.ID)
		if data == nil {
			log.Err.Printf("use action: effect not found: %s", ed.ID)
			continue
		}
		ua.targetEffects = append(ua.targetEffects, *data)
	}
	for _, ed := range data.TargetUserEffects {
		data := res.Effect(ed.ID)
		if data == nil {
			log.Err.Printf("use action: effect not found: %s", ed.ID)
			continue
		}
		ua.targetUserEffects = append(ua.targetUserEffects, *data)
	}
	return &ua
}

// Update updates use action.
func (ua *UseAction) Update(delta int64) {
	if ua.Cooldown() > 0 {
		ua.SetCooldown(ua.Cooldown() - delta)
	}
	if ua.Cooldown() < 0 {
		ua.SetCooldown(0)
	}
}

// CastMax returns maxinal cast time in milliseconds.
func (ua *UseAction) CastMax() int64 {
	return ua.castMax
}

// Cast returns cast time in milliseconds.
func (ua *UseAction) Cast() int64 {
	return ua.cast
}

// SetCast sets current cast time.
func (ua *UseAction) SetCast(cast int64) {
	ua.cast = cast
}

// CooldownMax returns maximal cooldown time in milliseconds.
func (ua *UseAction) CooldownMax() int64 {
	return ua.cooldownMax
}

// Cooldown returns current cooldown time in milliseconds.
func (ua *UseAction) Cooldown() int64 {
	return ua.cooldown
}

// SetCooldown sets cooldown time.
func (ua *UseAction) SetCooldown(cooldown int64) {
	ua.cooldown = cooldown
}

// UserMods returns use modifiers for user.
func (ua *UseAction) UserMods() []effect.Modifier {
	return ua.userMods
}

// ObjectMods returns use modifiers for object(use action source).
func (ua *UseAction) ObjectMods() []effect.Modifier {
	return ua.objectMods
}

// TargetMods returns modifiers for user target.
func (ua *UseAction) TargetMods() []effect.Modifier {
	return ua.targetMods
}

// TargetUserMods returns modifiers for user target or user.
func (ua *UseAction) TargetUserMods() []effect.Modifier {
	return ua.targetUserMods
}

// UserEffects returns use effects for user.
func (ua *UseAction) UserEffects() (effects []*effect.Effect) {
	for _, ed := range ua.userEffects {
		e := effect.New(ed)
		effects = append(effects, e)
	}
	return
}

// ObjectEffects returns use effects for object(use action source).
func (ua *UseAction) ObjectEffects() (effects []*effect.Effect) {
	for _, ed := range ua.objectEffects {
		e := effect.New(ed)
		effects = append(effects, e)
	}
	return
}

// TargetEffects returns use effects for user target.
func (ua *UseAction) TargetEffects() (effects []*effect.Effect) {
	for _, ed := range ua.targetEffects {
		e := effect.New(ed)
		effects = append(effects, e)
	}
	return
}

// TargetUserEffects returns use effects for user target or user.
func (ua *UseAction) TargetUserEffects() (effects []*effect.Effect) {
	for _, ed := range ua.targetUserEffects {
		e := effect.New(ed)
		effects = append(effects, e)
	}
	return
}

// Requirements returns use action requirements.
func (ua *UseAction) Requirements() []req.Requirement {
	return ua.requirements
}

// SetOwner sets specified serial object as use action owner.
func (ua *UseAction) SetOwner(o serial.Serialer) {
	ua.owner.ID, ua.owner.Serial = o.ID(), o.Serial()
}

// Owner returns use action owner.
func (ua *UseAction) Onwer() serial.Serialer {
	return serial.Object(ua.owner.ID, ua.owner.Serial)
}

// Data creates data resource for use action.
func (ua *UseAction) Data() res.UseActionData {
	data := res.UseActionData{
		CastMax:      ua.CastMax(),
		Cast:         ua.Cast(),
		CooldownMax:  ua.CooldownMax(),
		Cooldown:     ua.Cooldown(),
		UserMods:     effect.ModifiersData(ua.UserMods()...),
		ObjectMods:   effect.ModifiersData(ua.ObjectMods()...),
		TargetMods:   effect.ModifiersData(ua.TargetMods()...),
		Requirements: req.RequirementsData(ua.Requirements()...),
	}
	for _, e := range ua.UserEffects() {
		ed := res.UseActionEffectData{e.ID()}
		data.UserEffects = append(data.UserEffects, ed)
	}
	for _, e := range ua.ObjectEffects() {
		ed := res.UseActionEffectData{e.ID()}
		data.ObjectEffects = append(data.ObjectEffects, ed)
	}
	for _, e := range ua.TargetEffects() {
		ed := res.UseActionEffectData{e.ID()}
		data.TargetEffects = append(data.TargetEffects, ed)
	}
	for _, e := range ua.TargetUserEffects() {
		ed := res.UseActionEffectData{e.ID()}
		data.TargetUserEffects = append(data.TargetUserEffects, ed)
	}
	return data
}
