/*
 * quest.go
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

package quest

import (
	"github.com/isangeles/flame/core/data/res"
)

// Struct for quest.
type Quest struct {
	id          string
	name        string
	info        string
	activeStage *Stage
	stages      []*Stage
}

// NewQuest creates quest.
func NewQuest(data res.QuestData) *Quest {
	q := new(Quest)
	q.id = data.ID
	q.name = data.Name
	q.info = data.Info
	for _, sd := range data.Stages {
		s := NewStage(sd)
		q.stages = append(q.stages, s)
	}
	for _, s := range q.stages {
		if s.Start() {
			q.activeStage = s
		}
	}
	return q
}

// ID returns quest ID.
func (q *Quest) ID() string {
	return q.id
}

// Name returns quest name.
func (q *Quest) Name() string {
	return q.name
}

// Info returns quest info.
func (q *Quest) Info() string {
	return q.info
}

// ActiveStage returns active quest
// stage or nil if there is no active
// stage.
func (q *Quest) ActiveStage() *Stage {
	return q.activeStage
}
