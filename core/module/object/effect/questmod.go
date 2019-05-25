/*
 * questmod.go
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

package effect

import (
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/quest"
	"github.com/isangeles/flame/log"
)

// Struct for quest modifier.
type QuestMod struct {
	questID string
}

// NewQuestMod creates new quest modifier.
func NewQuestMod(data res.QuestModData) *QuestMod {
	qm := new(QuestMod)
	qm.questID = data.ID
	return qm
}

// Affect modifiers targets quests.
func (qm *QuestMod) Affect(source Target, targets ...Target) {
	for _, t := range targets {
		quester, ok := t.(quest.Quester)
		if !ok {
			return
		}
		qData := res.Quest(qm.questID)
		if qData == nil {
			log.Err.Printf("q_mod:quest_data_not_found:%s", qm.questID)
			return
		}
		q := quest.New(*qData)
		quester.Journal().AddQuest(q)
	}
}

// Undo undos modifications on specified target.
func (qm *QuestMod) Undo(source Target, targets ...Target) {
}
