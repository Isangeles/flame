/*
 * stage.go
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
	"github.com/isangeles/flame/module/flag"
)

// Struct for quest stage.
type Stage struct {
	id            string
	info          string
	ordinal       int
	start         bool
	last          bool
	next          string
	objectives    []*Objective
	completeFlags []flag.Flag
}

// NewStage creates quest stage.
func NewStage(data res.QuestStageData) *Stage {
	s := new(Stage)
	s.id = data.ID
	s.ordinal = data.Ordinal
	s.start = s.ordinal == 0
	s.last = data.Next == "end" 
	s.next = data.Next
	// Objectives.
	for _, od := range data.Objectives {
		o := NewObjective(od)
		s.objectives = append(s.objectives, o)
	}
	// Flags.
	for _, fd := range data.CompleteFlags {
		f := flag.Flag(fd.ID)
		s.completeFlags = append(s.completeFlags, f)
	}
	s.info = lang.Text(s.ID())
	return s
}

// ID returns stage ID.
func (s *Stage) ID() string {
	return s.id
}

// Info returns stage info.
func (s *Stage) Info() string {
	return s.info
}

// Ordinal returns stage ordinal number.
func (s *Stage) Ordinal() int {
	return s.ordinal
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

// Last checks whether stage is
// last stage.
func (s *Stage) Last() bool {
	return s.last
}

// Objectives returns all objectives of quest
// stage.
func (s *Stage) Objectives() []*Objective {
	return s.objectives
}

// Completed check if stage is completed.
// Stage is completed if all objectives are
// marked as completed or at least one, which
// is marked as finisher.
func (s *Stage) Completed() bool {
	for _, o := range s.objectives {
		if o.Completed() && o.Finisher() {
			return true
		}
	}
	for _, o := range s.objectives {
		if !o.Completed() {
			return false
		}
	}
	return true
}

// CompleteFlags returns flags for finishing stage.
func (s *Stage) CompleteFlags() []flag.Flag {
	return s.completeFlags
}
