/*
 * user.go
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

package skill

import (
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/req"
	"github.com/isangeles/flame/useaction"
)

// Struct for skill users.
type User interface {
	effect.Target
	Skills() []*Skill
	AddSkill(s *Skill)
	RemoveSkill(s *Skill)
	Use(u useaction.Usable) error
	MeetReqs(r ...req.Requirement) bool
}
