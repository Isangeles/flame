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
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/req"
	"github.com/isangeles/flame/core/module/serial"
)

// Interface for skills.
type Skill struct {
	id, serial   string
	name         string
	useReqs      []req.Requirement
	effects      []res.EffectData
	tartype      TargetType
	castRange    Range
	user         SkillUser
	target       effect.Target
	castTimeMax  int64 // cast time in milliseconds
	castTime     int64 // remaning cast time in milliseconds
	cooldown     int64 // cooldown time in millisconds
	cooldownMax  int64 // maximal cooldown in milliseconds
	casting      bool
	ready        bool
}

// Type for skills target
// types.
type TargetType int

// Type for skill range.
type Range int

const (
	// Errors.
	OTHERS_TARGET_ERR = "only_others_target"
	SELF_TARGET_ERR = "only_self_target"
	NO_TARGET_ERR = "no_target"
	REQS_NOT_MET_ERR = "reqs_not_meet"
	RANGE_ERR = "user_too_far"
	NOT_READY_ERR = "skill_not_ready"
	// Target types.
	Target_all TargetType = iota
	Target_others
	Target_self
	// Skill ranges.
	Range_touch = iota
	Range_close
	Range_far
	Range_huge
)

// NewSkill creates new skill with specifie parameters.
func New(data res.SkillData) *Skill {
	s := new(Skill)
	s.id = data.ID
	s.name = data.Name
	s.castRange = Range(data.Range)
	s.useReqs = data.UseReqs
	s.effects = data.Effects
	s.castTimeMax = int64(data.Cast * 1000) // cast from sec to millisec
	s.cooldownMax = int64(2 * 1000)
	return s
}

// Update updates skill.
func (s *Skill) Update(delta int64) {
	if s.Casting() {
		s.castTime += delta
		if s.castTime >= s.castTimeMax {
			s.casting = false
			s.cooldown = s.cooldownMax
			s.ready = false
			if s.target == nil {
				return
			}
			for _, e := range s.buildEffects() {
				s.target.TakeEffect(e)
			}
		}
	}
	if !s.Ready() {
		s.cooldown -= delta
		if s.cooldown <= 0 {
			s.ready = true
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
	if object.Range(user, target) > s.castRange.Value() {
		return fmt.Errorf(RANGE_ERR)
	}
	s.user = user
	s.target = target
	if !user.MeetReqs(s.useReqs) {
		return fmt.Errorf(REQS_NOT_MET_ERR)
	}
	s.castTime = 0
	s.casting = true
	return nil
}

// CastTime returns current casting
// time in milliseconds.
func (s *Skill) CastTime() int64 {
	return s.castTime
}

// CastTimeMax returns maximal casting
// time in milliseconds.
func(s *Skill) CastTimeMax() int64 {
	return s.castTimeMax
}

// Casting checks whether skill is
// currently casted.
func (s *Skill) Casting() bool {
	return s.casting
}

// Cooldown retruns current cooldown time.
func (s *Skill) Cooldown() int64 {
	return s.cooldown
}

// CooldownMax retruns cooldown time
// in milliseconds.
func (s *Skill) CooldownMax() int64 {
	return s.cooldownMax
}

// SetCooldown sets specified value as
// current skill cooldown in milliseconds.
func (s *Skill) SetCooldown(c int64) {
	s.cooldown = c
}

// Ready checks whether skill is ready, i.e. casting
// was started and funished successfully.
func (s *Skill) Ready() bool {
	return s.ready
}

// Value returns range value.
func (r Range) Value() float64 {
	switch {
	case r <= Range_touch:
		return 50.0
	case r == Range_close:
		return 100.0
	case r == Range_far:
		return 500.0
	default:
		return 1000.0
	}
}

// effects builds and returns skill effects.
func (s *Skill) buildEffects() []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	for _, ed := range s.effects {
		e := effect.New(ed)
		serial.AssignSerial(e)
		effects = append(effects, e)
	}
	return effects
}
