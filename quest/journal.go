/*
 * journal.go
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

package quest

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/flag"
	"github.com/isangeles/flame/log"
)

// Struct for character journal.
type Journal struct {
	quests map[string]*Quest
	owner  Quester
}

// NewJournal creates quests journal.
func NewJournal(quester Quester) *Journal {
	j := Journal{
		owner:  quester,
		quests: make(map[string]*Quest),
	}
	return &j
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

// Apply applies specifie data on the journal.
func (j *Journal) Apply(data res.QuestLogData) {
	// Remove quests not present anymore.
	for i, _ := range j.quests {
		found := false
		for _, qd := range data.Quests {
			if qd.ID == i {
				found = true
				break
			}
		}
		if !found {
			delete(j.quests, i)
		}
	}
	// Add/update quests.
	for _, logQuestData := range data.Quests {
		quest := j.quests[logQuestData.ID]
		if quest == nil {
			questData := res.Quest(logQuestData.ID)
			if questData == nil {
				log.Err.Printf("Quest log: Apply: %s#%s: unable to retrieve quest data: %s",
					j.owner.ID(), j.owner.Serial(), logQuestData.ID)
				continue
			}
			quest = New(*questData)
			// Add quest to the quest log.
			j.AddQuest(quest)
		}
		// Restore quest stage.
		for _, s := range quest.Stages() {
			if s.ID() == logQuestData.Stage {
				quest.SetActiveStage(s)
			}
		}
		if quest.ActiveStage() == nil {
			log.Err.Printf("build quest log: %s#%s: quest: %s: unable to set active stage",
				j.owner.ID(), j.owner.Serial(), quest.ID())
		}
	}
}

// Data creates data resource for journal.
func (j *Journal) Data() res.QuestLogData {
	data := res.QuestLogData{}
	for _, q := range j.Quests() {
		questData := res.QuestLogQuestData{
			ID:    q.ID(),
		}
		if q.ActiveStage() != nil {
			questData.Stage = q.ActiveStage().ID()
		}
		data.Quests = append(data.Quests, questData)
	}
	return data
}
