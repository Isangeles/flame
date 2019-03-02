/*
 * skill.go
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

// Package for skills.
package skill

import (
	"fmt"
	
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/req"
)

// Interface for skills.
type Skill struct {
	id, serial   string
	name         string
	useReqs      []req.Requirement
	effects      []res.EffectData
	tartype      TargetType
	user         SkillUser
	target       effect.Target
	castTime     int64 // cast time in milliseconds
	castTimeLeft int64 // remaning cast time in milliseconds
	casting      bool
	ready        bool
}

// Type for skills target
// types.
type TargetType int

const (
	// Errors.
	OTHERS_TARGET_ERR = "only_others_target"
	SELF_TARGET_ERR = "only_self_target"
	NO_TARGET_ERR = "no_target"
	REQS_NOT_MET_ERR = "reqs_not_meet"
	// Target types.
	Target_all TargetType = iota
	Target_others
	Target_self
)

// NewSkill creates new skill with specifie parameters.
func New(data res.SkillData) *Skill {
	s := new(Skill)
	s.id = data.ID
	s.name = data.Name
	s.useReqs = data.UseReqs
	s.effects = data.Effects
	s.castTime = int64(data.Cast * 1000) // cast from sec to millisec
	return s
}

// Update updates skill.
func (s *Skill) Update(delta int64) {
	if s.Casting() {
		s.castTimeLeft -= delta
		if s.castTimeLeft <= 0 {
			s.casting = false
			if s.target == nil {
				return
			}
			for _, e := range s.buildEffects() {
				s.target.TakeEffect(e)
			}
		}
	}
}

// ID returns skill ID.
func (s *Skill) ID() string {
	return s.id
}

// Serial returns skill serial
// value.
func (s *Skill) Serial() string {
	return s.serial
}

// SetSerial sets skill serial value.
func (s *Skill) SetSerial(serial string) {
	s.serial = serial
}

// Name returns skill name.
func (s *Skill) Name() string {
	return s.name
}

// SetName sets specified name as
// skill display name.
func (s *Skill) SetName(name string) {
	s.name = name
}

// Cast starts skill casting with specified targetable object
// as skill user.
func (s *Skill) Cast(user SkillUser, target effect.Target) error {
	if s.tartype != Target_all && target == nil {
		return fmt.Errorf(NO_TARGET_ERR)
	}
	if s.tartype == Target_others &&
		user.ID()+user.Serial() == target.ID()+target.Serial() {
		return fmt.Errorf(OTHERS_TARGET_ERR)
	}
	if s.tartype == Target_self &&
		user.ID()+user.Serial() != target.ID()+target.Serial() {
		return fmt.Errorf(SELF_TARGET_ERR)
	}
	s.user = user
	s.target = target
	if !user.MeetReqs(s.useReqs) {
		return fmt.Errorf(REQS_NOT_MET_ERR)
	}
	s.castTimeLeft = s.castTime
	s.casting = true
	return nil
}

// CastTime returns skill casting time int
// milliseconds.
func (s *Skill) CastTime() int64 {
	return s.castTime
}

// Casting checks whether skill is currently casted.
func (s *Skill) Casting() bool {
	return s.casting
}

// Ready checks whether skill is ready, i.e. casting
// was started and funished successfully.
func (s *Skill) Ready() bool {
	return s.ready
}

// effects builds and returns skill effects.
func (s *Skill) buildEffects() []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	for _, ed := range s.effects {
		e := effect.New(ed)
		e.SetSerial(s.ID() + "_" + s.Serial())
		effects = append(effects, e)
	}
	return effects
}
