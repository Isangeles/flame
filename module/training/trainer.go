/*
 * trainer.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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

package training

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/req"
)

// Interface for object with trainings.
type Trainer interface {
	AddTraining(t *TrainerTraining)
	Trainings() []*TrainerTraining
}

// Struct for trainer training.
type TrainerTraining struct {
	*Training
	reqs []req.Requirement
}

// NewTrainerTraining creates new trainer training.
func NewTrainerTraining(training *Training, data res.TrainerTrainingData) *TrainerTraining {
	tt := TrainerTraining{
		Training: training,
		reqs:     req.NewRequirements(data.Reqs),
	}
	return &tt
}

// Requirements returns training requirements specific for
// the trainer.
func (tt *TrainerTraining) Requirements() []req.Requirement {
	reqs := tt.reqs
	for _, r := range tt.UseAction().Requirements() {
		reqs = append(reqs, r)
	}
	return reqs
}

// Data returns data resource for trainer training.
func (tt *TrainerTraining) Data() res.TrainerTrainingData {
	data := res.TrainerTrainingData{
		ID: tt.ID(),
	}
	return data
}
	
