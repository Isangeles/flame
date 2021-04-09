/*
 * skill.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/useaction"
)

// Struct for skills.
type Skill struct {
	id          string
	useAction   *useaction.UseAction
}

// NewSkill creates new skill with specifie parameters.
func New(data res.SkillData) *Skill {
	s := new(Skill)
	s.id = data.ID
	s.useAction = useaction.New(data.UseAction)
	return s
}

// Update updates skill.
func (s *Skill) Update(delta int64) {
	s.UseAction().Update(delta)
}

// ID returns skill ID.
func (s *Skill) ID() string {
	return s.id
}

// UseAction returns skill use action.
func (s *Skill) UseAction() *useaction.UseAction {
	return s.useAction
}
