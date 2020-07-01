/*
 * useaction.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/req"
)

// Interface for usable game objects.
type Usable interface {
	objects.Object
	UseAction() *UseAction
}

// Struct for use action of usable object.
type UseAction struct {
	object        Usable
	userMods      []effect.Modifier
	objectMods    []effect.Modifier
	userEffects   []res.EffectData
	objectEffects []res.EffectData
	requirements  []req.Requirement
}

// New creates new use action.
func New(ob Usable, data res.UseActionData) *UseAction {
	ua := UseAction{
		object:       ob,
		userMods:     effect.NewModifiers(data.UserMods),
		objectMods:   effect.NewModifiers(data.ObjectMods),
		requirements: req.NewRequirements(data.Requirements),
	}
	for _, ed := range data.UserEffects {
		data := res.Effect(ed.ID)
		if data == nil {
			log.Err.Printf("%s %s: use action: effect not found: %s",
				ua.object.ID(), ua.object.Serial(), ed.ID)
			continue
		}
		ua.userEffects = append(ua.userEffects, *data)
	}
	for _, ed := range data.ObjectEffects {
		data := res.Effect(ed.ID)
		if data == nil {
			log.Err.Printf("%s %s: use action: effect not found: %s",
				ua.object.ID(), ua.object.Serial(), ed.ID)
			continue
		}
		ua.objectEffects = append(ua.objectEffects, *data)
	}
	return &ua
}

// UserMods returns use modifiers for user.
func (ua *UseAction) UserMods() []effect.Modifier {
	return ua.userMods
}

// ObjectMods returns use modifiers for object(use action source).
func (ua *UseAction) ObjectMods() []effect.Modifier {
	return ua.objectMods
}

// UserEffects returns use effects for user.
func (ua *UseAction) UserEffects() (effects []*effect.Effect) {
	for _, ed := range ua.userEffects {
		e := effect.New(ed)
		e.SetSource(ua.object.ID(), ua.object.Serial())
		effects = append(effects, e)
	}
	return
}

// Requirements returns use action requirements.
func (ua *UseAction) Requirements() []req.Requirement {
	return ua.requirements
}

// ObjectEffects returns use effects for object(use action source).
func (ua *UseAction) ObjectEffects() (effects []*effect.Effect) {
	for _, ed := range ua.objectEffects {
		e := effect.New(ed)
		e.SetSource(ua.object.ID(), ua.object.Serial())
		effects = append(effects, e)
	}
	return
}
