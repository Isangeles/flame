/*
 * stage.go
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

// Struct for quest stage.
type Stage struct {
	id         string
	name, info string
	start      bool
	next       string
	objectives []*Objective
}

// NewStage creates quest stage.
func NewStage(data res.QuestStageData) *Stage {
	s := new(Stage)
	s.id = data.ID
	s.info = data.Info
	s.start = data.Start
	s.next = data.Next
	for _, od := range data.Objectives {
		o := NewObjective(od)
		s.objectives = append(s.objectives, o)
	}
	return s
}

// ID returns stage ID.
func (s *Stage) ID() string {
	return s.id
}

// Name returns stage name.
func (s *Stage) Name() string {
	return s.name
}

// Info returns stage info.
func (s *Stage) Info() string {
	return s.info
}

// Start checks if stage is quest
// start stage.
func (s *Stage) Start() bool {
	return s.start
}

// NextStageID returns ID of next quest
// stage after completing this one.
func (s *Stage) NextStageID() string {
	return s.next
}
