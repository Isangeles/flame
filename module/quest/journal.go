/*
 * journal.go
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
	"github.com/isangeles/flame/module/flag"
	"github.com/isangeles/flame/log"
)

// Struct for character journal.
type Journal struct {
	quests map[string]*Quest
	owner  Quester
}

// NewJournal creates quests journal.
func NewJournal(data res.QuestLogData, quester Quester) *Journal {
	j := new(Journal)
	j.owner = quester
	j.quests = make(map[string]*Quest)
	for _, logQuestData := range data.Quests {
		questData, ok := res.Quests[logQuestData.ID]
		if !ok {
			log.Err.Printf("build quest log: %s#%s: fail to retrieve quest data: %s",
				j.owner.ID(), j.owner.Serial(), logQuestData.ID)
			continue
		}
		// Restore quest stage.
		quest := New(questData)
		for _, s := range quest.Stages() {
			if s.ID() == logQuestData.Stage {
				quest.SetActiveStage(s)
			}
		}
		if quest.ActiveStage() == nil {
			log.Err.Printf("build quest log: %s#%s: quest: %s: fail to set active stage",
				j.owner.ID(), j.owner.Serial(), quest.ID())
		}
		// Add quest to quest log.
		j.AddQuest(quest)
	}
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
