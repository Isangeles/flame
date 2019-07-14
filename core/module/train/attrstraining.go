/*
 * attrstraining.go
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

package train

import (
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/req"
)

// Structs for attributes training.
type AttrsTraining struct {
	strVal int
	conVal int
	dexVal int
	wisVal int
	intVal int
	reqs   []req.Requirement
}

// NewAttrsTraining creates attributes training.
func NewAttrsTraining(data res.AttrsTrainingData) *AttrsTraining {
	at := AttrsTraining{
		strVal: data.Str,
		conVal: data.Con,
		dexVal: data.Dex,
		wisVal: data.Wis,
		intVal: data.Int,
		reqs:   req.NewRequirements(data.Reqs...),
	}
	return &at
}

// Stringht returns value of strenght training.
func (at *AttrsTraining) Strenght() int {
	return at.strVal
}

// Constitution returns value of constitution training.
func (at *AttrsTraining) Constitution() int {
	return at.conVal
}

// Dexterity returns value of dexterity training.
func (at *AttrsTraining) Dexterity() int {
	return at.dexVal
}

// Wisdom returns value of wisdom training.
func (at *AttrsTraining) Wisdom() int {
	return at.wisVal
}

// Intelligence returns value of intelligence training.
func (at *AttrsTraining) Intelligence() int {
	return at.intVal
}

// Reqs returns training requirements.
func (at *AttrsTraining) Reqs() []req.Requirement {
	return at.reqs
}
		
