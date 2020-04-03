/*
 * training.go
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

// Package with training structs.
package train

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/req"
)

// Interface for trainings.
type Training interface {
	Reqs() []req.Requirement
	Data() res.TrainingData
}

// Interface for object with trainings.
type Trainer interface {
	AddTraining(t Training)
	Trainings() []Training
}

// NewTrainings creates new trainings from specified data.
func NewTrainings(data ...res.TrainingData) (trainings []Training) {
	for _, d := range data {
		switch d := d.(type) {
		case res.AttrsTrainingData:
			atTrain := NewAttrsTraining(d)
			trainings = append(trainings, atTrain)
		}
	}
	return
}
