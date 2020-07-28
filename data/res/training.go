/*
 * train.go
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

package res

import (
	"encoding/xml"
)

// Struct for trainings data.
type TrainingsData struct {
	XMLName   xml.Name       `xml:"trainings" json:"-"`
	Trainings []TrainingData `xml:"training" json:"trainings"`
}

// Struct for training data.
type TrainingData struct {
	ID  string        `xml:"id,attr" json:"id"`
	Use UseActionData `xml:"use" json:"use"`
}

// Struct for trainer training data.
type TrainerTrainingData struct {
	ID   string   `xml:"id,attr" json:"id"`
	Reqs ReqsData `xml:"reqs" json:"reqs"`
}
