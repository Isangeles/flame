/*
 * skill.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"errors"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/req"
)

// Struct for skills.
// All times(cast time, cooldown, etc.)
// are in milliseconds.
type Skill struct {
	id          string
	name        string
	useReqs     []req.Requirement
	effects     []res.EffectData
	tartype     TargetType
	castRange   Range
	user        effect.Target
	target      effect.Target
	castTimeMax int64
	castTime    int64
	cooldown    int64
	cooldownMax int64
	casting     bool
	casted      bool
	ready       bool
}

// Type for skills target
// types.
type TargetType string

// Type for skill range.
type Range string

var (
	// Errors.
	OtherTargetOnly = errors.New("other taget only")
	SelfTargetOnly  = errors.New("self target only")
	NoTarget        = errors.New("no target")
	ReqsNotMet      = errors.New("requirements not met")
	TooFar          = errors.New("target too far")
	NotReady        = errors.New("not ready")
)

const (
	// Target types.
	TargetAll TargetType = TargetType("skillTarAll")
	TargetOthers = TargetType("skillTarOthers")
	TargetSelf = TargetType("skillTarSelf")
	// Skill ranges.
	RangeTouch = Range("skillRangeTouch")
	RangeClose = Range("skillRangeClose")
	RangeFar = Range("skillRangeFar")
	RangeHuge = Range("skillRangeHuge")
)

// NewSkill creates new skill with specifie parameters.
func New(data res.SkillData) *Skill {
	s := new(Skill)
	s.id = data.ID
	s.name = data.Name
	s.castRange = Range(data.Range)
	s.castTimeMax = data.Cast
	s.cooldownMax = data.Cooldown
	s.useReqs = req.NewRequirements(data.UseReqs)
	if len(s.name) < 1 {
		s.name = lang.Text(s.id)
	}
	for _, sed := range data.Effects {
		ed := res.Effect(sed.ID)
		if ed == nil {
			continue
		}
		s.effects = append(s.effects, *ed)
	}
	return s
}

// Update updates skill.
func (s *Skill) Update(delta int64) {
	if !s.Ready() {
		s.cooldown -= delta
		if s.cooldown <= 0 {
			s.ready = true
		}
		return
	}
	if s.Casting() {
		s.castTime += delta
		if s.castTime >= s.castTimeMax {
			s.casting = false
			s.casted = true
		}
	}
}

// Activate activates skill.
func (s *Skill) Activate() {
	s.cooldown = s.cooldownMax
	s.ready = false
	s.casted = false
	if s.target == nil {
		return
	}
	for _, e := range s.buildEffects(s.effects) {
		s.target.TakeEffect(e)
	}
}

// ID returns skill ID.
func (s *Skill) ID() string {
	return s.id
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
func (s *Skill) Cast(user User, target effect.Target) error {
	if !s.Ready() {
		return NotReady
	}
	if s.tartype != TargetAll && target == nil {
		return NoTarget
	}
	if s.tartype == TargetOthers &&
		user.ID()+user.Serial() == target.ID()+target.Serial() {
		return OtherTargetOnly
	}
	if s.tartype == TargetSelf &&
		user.ID()+user.Serial() != target.ID()+target.Serial() {
		return SelfTargetOnly
	}
	if objects.Range(user, target) > s.castRange.Value() {
		return TooFar
	}
	s.user = user
	s.target = target
	if user, ok := s.user.(req.RequirementsTarget); ok &&
		!user.MeetReqs(s.useReqs...) {
		return ReqsNotMet
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
func (s *Skill) CastTimeMax() int64 {
	return s.castTimeMax
}

// Casting checks whether skill is
// currently casted.
func (s *Skill) Casting() bool {
	return s.casting
}

// Casted check whether skill was
// successfully casted but not
// activated yet.
func (s *Skill) Casted() bool {
	return s.casted
}

// StopCast stops skill casting.
func (s *Skill) StopCast() {
	s.casting = false
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

// Ready checks whether skill is ready
// to cast.
func (s *Skill) Ready() bool {
	return s.ready
}

// SetReady toggles skill ready state.
func (s *Skill) SetReady(ready bool) {
	s.ready = ready
}

// Value returns range value.
func (r Range) Value() float64 {
	switch {
	case r <= RangeTouch:
		return 50.0
	case r == RangeClose:
		return 100.0
	case r == RangeFar:
		return 500.0
	default:
		return 1000.0
	}
}

// effects builds and returns skill effects.
func (s *Skill) buildEffects(effectsData []res.EffectData) []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	for _, ed := range effectsData {
		e := effect.New(ed)
		e.SetSource(s.user.ID(), s.user.Serial())
		effects = append(effects, e)
	}
	return effects
}
