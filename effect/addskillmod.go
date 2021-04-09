/*
 * addskillmod.go
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

package effect

import (
	"github.com/isangeles/flame/data/res"
)

// Struct for add skill modifier.
type AddSkillMod struct {
	skillID string
}

// NewAddSkillMod creates new add skill modifier.
func NewAddSkillMod(data res.AddSkillModData) *AddSkillMod {
	asm := AddSkillMod{data.SkillID}
	return &asm
}

// SkillID returns ID of skill to add.
func (asm *AddSkillMod) SkillID() string {
	return asm.skillID
}

// Data creates data resource for modifier.
func (asm *AddSkillMod) Data() res.AddSkillModData {
	data := res.AddSkillModData{asm.SkillID()}
	return data
}
