/*
 * quest.go
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
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/req"
)

// Struct for quest.
type Quest struct {
	id          string
	name, info  string
	completed   bool
	activeStage *Stage
	stages      []*Stage
}

// Interface for objects with quests.
type Quester interface {
	objects.Object
	Journal() *Journal
	MeetReqs(reqs ...req.Requirement) bool
}

const (
	END_QUEST_ID = "end"
)

// New creates quest.
func New(data res.QuestData) *Quest {
	q := new(Quest)
	q.id = data.ID
	for _, sd := range data.Stages {
		s := NewStage(sd)
		q.stages = append(q.stages, s)
	}
	for _, s := range q.stages {
		if s.Start() {
			q.activeStage = s
		}
	}
	nameInfo := lang.Texts(q.ID())
	q.name = nameInfo[0]
	if len(nameInfo) > 1 {
		q.info = nameInfo[1]
	}
	return q
}

// ID returns quest ID.
func (q *Quest) ID() string {
	return q.id
}

// Name returns name of the quest.
func (q *Quest) Name() string {
	return q.name
}

// Info returns info about quest.
func (q *Quest) Info() string {
	return q.info
}

// Completed check if quest was marked
// as completed.
func (q *Quest) Completed() bool {
	return q.completed
}

// SetComplete sets quest as completed/uncompleted.
func (q *Quest) SetComplete(complete bool) {
	q.completed = complete
}

// Stages returns all stages of the quest.
func (q *Quest) Stages() []*Stage {
	return q.stages
}

// ActiveStage returns active quest
// stage or nil if there is no active
// stage.
func (q *Quest) ActiveStage() *Stage {
	return q.activeStage
}

// SetActiveStage sets specified stage as
// active stage.
func (q *Quest) SetActiveStage(s *Stage) {
	q.activeStage = s
}
