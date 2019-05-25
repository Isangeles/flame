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

// AddQuests adds specified quests for character.
func (j *Journal) AddQuest(q *Quest) {
	j.quests[q.ID()] = q
}

// checkQuest checks quest progress.
func (j *Journal) checkQuest(q *Quest) {
	if q.ActiveStage() == nil {
		return
	}
	stage := q.ActiveStage()
	for _, o := range stage.Objectives() {
		if !j.owner.MeetReqs(o.Reqs()...) {
			o.SetComplete(false)
			continue
		}
		o.SetComplete(true)
	}
	if stage.Completed() {
		err := q.SetActiveStage(stage.NextStageID())
		if err != nil {
			return
		}
		if flager, ok := j.owner.(flag.Flagger); ok {
			for _, f := range stage.CompleteFlags() {
				flager.AddFlag(f)
			}
		}
	}
}
