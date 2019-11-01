/*
 * journal.go
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
	"github.com/isangeles/flame/core/module/flag"
)

// Struct for character journal.
type Journal struct {
	quests map[string]*Quest
	owner  Quester
}

// NewJournal creates quests journal.
func NewJournal(quester Quester) *Journal {
	j := new(Journal)
	j.owner = quester
	j.quests = make(map[string]*Quest)
	return j
}

// Update updates journal quests.
func (j *Journal) Update(delta int64) {
	for _, q := range j.quests {
		if q.Completed() {
			continue
		}
		j.checkQuest(q)
	}
}

// Quests return all quests.
func (j *Journal) Quests() (qs []*Quest) {
	for _, q := range j.quests {
		qs = append(qs, q)
	}
	return
}

// AddQuests adds specified quest.
func (j *Journal) AddQuest(q *Quest) {
	j.quests[q.ID()] = q
}

// RemoveQuest removes specified quest.
func (j *Journal) RemoveQuest(q *Quest) {
	delete(j.quests, q.ID())
}

// checkQuest checks quest progress.
func (j *Journal) checkQuest(q *Quest) {
	if q.ActiveStage() == nil {
		return
	}
	// Check previous stages.
	for _, s := range q.Stages() {
		if s.Ordinal() > q.ActiveStage().Ordinal() {
			continue
		}
		j.checkStage(s)
		// Move back to previous incompleted stage.
		if !s.Completed() {
			q.SetActiveStage(s)
		}
	}
	// Check active stage objectives.
	j.checkStage(q.ActiveStage())
	// Move to next stage.
	if q.ActiveStage().Completed() {
		if flager, ok := j.owner.(flag.Flagger); ok {
			for _, f := range q.ActiveStage().CompleteFlags() {
				flager.AddFlag(f)
			}
		}
		if q.ActiveStage().Last() {
			q.SetComplete(true)
			return
		}
		var nextStage *Stage
		for _, s := range q.Stages() {
			if s.ID() == q.ActiveStage().NextStageID() {
				nextStage = s
			}
		}
		q.SetActiveStage(nextStage)
	}
}

// checkStage checks progress of specified
// quest stage.
func (j *Journal) checkStage(s *Stage) {
	for _, o := range s.Objectives() {
		if !j.owner.MeetReqs(o.Reqs()...) {
			o.SetComplete(false)
			continue
		}
		o.SetComplete(true)
	}
}
