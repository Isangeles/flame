/*
 * quest.go
 *
 * Copyright 2021 Dariusz Sikora <dev@isangeles.pl>
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

package req

import (
	"github.com/isangeles/flame/data/res"
)

// Struct for quest requirement.
type Quest struct {
	questID        string
	questCompleted bool
	meet           bool
}

// NewQuest creates new quest requirement.
func NewQuest(data res.QuestReqData) *Quest {
	q := Quest{
		questID:        data.ID,
		questCompleted: data.Completed,
	}
	return &q
}

// QuestID returns ID of the required quest.
func (q *Quest) QuestID() string {
	return q.questID
}

// QuestCompleted checks if required quest should
// be completed.
func (q *Quest) QuestCompleted() bool {
	return q.questCompleted
}

// Meet checks if requrement is meet.
func (q *Quest) Meet() bool {
	return q.meet
}

// SetMeet sets requirement as meet/not meet.
func (q *Quest) SetMeet(meet bool) {
	q.meet = meet
}

// Data returns data resource for requirement.
func (q *Quest) Data() res.QuestReqData {
	data := res.QuestReqData{
		ID:        q.questID,
		Completed: q.questCompleted,
	}
	return data
}
