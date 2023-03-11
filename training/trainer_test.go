/*
 * trainer_test.go
 *
 * Copyright 2023 Dariusz Sikora <ds@isangeles.dev>
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
	"testing"

	"github.com/isangeles/flame/data/res"
)

var (
	itemReqData         = res.ItemReqData{ID: "testItem", Amount: 1}
	reqsData            = res.ReqsData{ItemReqs: []res.ItemReqData{itemReqData}}
	trainerTrainingData = res.TrainerTrainingData{
		ID:   "test",
		Reqs: reqsData,
	}
)

// TestTrainerTrainingRequirements tests requirements function
// of TrainerTraining struct.
func TestTrainerTrainingRequirements(t *testing.T) {
	training := New(trainingData)
	trainerTraining := NewTrainerTraining(training, trainerTrainingData)
	if len(trainerTraining.Requirements()) < 1 {
		t.Errorf("No requirements")
	}
}

// TestTrainerTrainingData test creating data resource
// for TrainerTraining struct.
func TestTrainerTrainingData(t *testing.T) {
	training := New(trainingData)
	trainerTraining := NewTrainerTraining(training, trainerTrainingData)
	data := trainerTraining.Data()
	if data.ID != trainerTraining.ID() {
		t.Errorf("Invalid data ID: %s != %s", data.ID, trainerTraining.ID())
	}
	if len(data.Reqs.ItemReqs) != 1 {
		t.Errorf("Invalid number of item requirements in data resource: %d != 1",
			len(data.Reqs.ItemReqs))
	}
}
