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
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/req"
)

// Interface for skills.
type Skill struct {
	id, serial string
	name       string
	useReqs    []req.Requirement
	effects    []effect.Effect
	castSec    int // cast time in seconds
	casting    bool
	ready      bool
}

// NewSkill creates new skill with specifie parameters.
func NewSkill(id, name string, cast int, useReqs []req.Requirement,
	effects []effect.Effect) *Skill {
	s := new(Skill)
	s.id = id
	s.name = name
	s.useReqs = useReqs
	s.effects = effects
	return s
}

// Update updates skill.
func (s *Skill) Update(delta int64) {

}

// Cast starts skill casting with specified targetable object
// as skill user.
func (s *Skill) Cast(user object.Object) error {
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
