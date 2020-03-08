/*
 * objective.go
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

package quest

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/req"
)

// Struct for quest stage
// objective.
type Objective struct {
	id        string
	finisher  bool
	completed bool
	reqs      []req.Requirement
}

// NewObjective creates quest objective.
func NewObjective(data res.QuestObjectiveData) *Objective {
	o := new(Objective)
	o.id = data.ID
	o.finisher = data.Finisher
	o.reqs = req.NewRequirements(data.Reqs...)
	return o
}

// ID returns objective ID.
func (o *Objective) ID() string {
	return o.id
}

// Finisher checks if completing objective
// should complete whole stage.
func (o *Objective) Finisher() bool {
	return o.finisher
}

// Completed checks if objective was marked
// as completed.
func (o *Objective) Completed() bool {
	return o.completed
}

// SetComplete sets objective as complete/not complete.
func (o *Objective) SetComplete(complete bool) {
	o.completed = complete
}

// Reqs returns objective requirements.
func (o *Objective) Reqs() []req.Requirement {
	return o.reqs
}
