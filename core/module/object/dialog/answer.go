/*
 * answer.go
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

package dialog

import (
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/req"
)

// Struct for dialog text answer.
type Answer struct {
	id   string
	to   string
	reqs []req.Requirement
}

// NewAnswer creates new dialog answer.
func NewAnswer(data *res.DialogAnswerData) *Answer {
	a := new(Answer)
	a.id = data.ID
	a.to = data.To
	a.reqs = req.NewRequirements(data.Reqs)
	return a
}

// ID returns answer ID.
func (a *Answer) ID() string {
	return a.id
}

// Requirements returns answer requirements.
func (a *Answer) Requirements() []req.Requirement {
	return a.reqs
}