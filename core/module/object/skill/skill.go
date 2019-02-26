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

package skill

import (
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/req"
	"github.com/isangeles/flame/core/data/res"
)

// Interface for skills.
type Skill struct {
	id, serial string
	name       string
	useReqs    []req.Requirement
	effects    []res.EffectData
	castSec    int // cast time in seconds
	casting    bool
	ready      bool
}

// NewSkill creates new skill with specifie parameters.
func New(data res.SkillData) *Skill {
	s := new(Skill)
	s.id = data.ID
	s.name = data.Name
	s.useReqs = data.UseReqs
	s.effects = data.Effects
	return s
}

// Update updates skill.
func (s *Skill) Update(delta int64) {

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
func (s *Skill) Cast(user effect.Target) error {
	s.casting = true
	return nil
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
