/*
 * journal_test.go
 *
 * Copyright 2025 Dariusz Sikora <ds@isangeles.dev>
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
	"testing"

	"github.com/isangeles/flame/data/res"
)

var (
	flagData  = res.FlagData{ID: "flag1"}
	stageData = res.QuestStageData{ID: "stage1", Start: true, StartFlags: []res.FlagData{flagData}}
	questData = res.QuestData{ID: "quest1", Stages: []res.QuestStageData{stageData}}
)

// Tests adding new quest to the journal.
func TestJournalAddQuest(t *testing.T) {
	// Create object and journal.
	object := new(testQuester)
	object.journal = NewJournal(object)
	quest := New(questData)
	// Test adding quest.
	object.journal.AddQuest(quest)
	questPresent := false
	for _, q := range object.journal.Quests() {
		questPresent = q.ID() == quest.ID()
		if questPresent {
			break
		}
	}
	if !questPresent {
		t.Errorf("Quest not added")
	}
	// Test adding start flag.
	flagPresent := false
	for _, flag := range object.Flags() {
		flagPresent = flag.ID() == flagData.ID
		if flagPresent {
			break
		}
	}
	if !flagPresent {
		t.Errorf("Start flag not added")
	}
}
