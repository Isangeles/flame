/*
 * skill.go
 *
 * Copyright 2019-2024 Dariusz Sikora <ds@isangeles.dev>
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

// Package for skill structs.
package skill

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/req"
	"github.com/isangeles/flame/useaction"
)

// Struct for skills.
type Skill struct {
	id             string
	useAction      *useaction.UseAction
	passiveReqs    []req.Requirement
	passiveEffects []res.EffectData
	owner          User
}

// New creates new skill.
func New(data res.SkillData) *Skill {
	s := new(Skill)
	s.id = data.ID
	if useaction.HasData(data.UseAction) {
		s.useAction = useaction.New(data.UseAction)
	}
	s.passiveReqs = req.NewRequirements(data.Passive.Requirements)
	for _, ed := range data.Passive.Effects {
		data := res.Effect(ed.ID)
		if data == nil {
			log.Err.Printf("use action: effect not found: %s", ed.ID)
			continue
		}
		s.passiveEffects = append(s.passiveEffects, *data)
	}
	return s
}

// Update updates skill.
func (s *Skill) Update(delta int64) {
	if s.UseAction() != nil {
		s.UseAction().Update(delta)
	}
	if s.owner == nil {
		return
	}
	if !s.owner.MeetReqs(s.passiveReqs...) {
		return
	}
passivesAdd:
	for _, ed := range s.passiveEffects {
		for _, e := range s.owner.Effects() {
			if e.ID() == ed.ID {
				continue passivesAdd
			}
		}
		eff := effect.New(ed)
		s.owner.TakeEffect(eff)
	}
}

// ID returns skill ID.
func (s *Skill) ID() string {
	return s.id
}

// UseAction returns skill use action.
func (s *Skill) UseAction() *useaction.UseAction {
	return s.useAction
}

// PassiveRequirements returns list of requirements for
// passive effects.
func (s *Skill) PassiveRequirements() []req.Requirement {
	return s.passiveReqs
}

// PassiveEffects returns all passive effects.
func (s *Skill) PassiveEffects() []res.EffectData {
	return s.passiveEffects
}

// SetOwner sets specified skill user as skill owner.
func (s *Skill) SetOwner(owner User) {
	if s.UseAction() != nil {
		s.UseAction().SetOwner(owner)
	}
	s.owner = owner
}
