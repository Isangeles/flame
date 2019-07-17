/*
 * trainingparser.go
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

package parsexml

import (
	"encoding/xml"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/train"
)

// Struct for training XML node.
type TrainingsXML struct {
	XMLName    xml.Name           `xml:"trainings"`
	AttrsTrain []AttrsTrainingXML `xml:"attrs-train"`
}

// Struct for attributes training XML node.
type AttrsTrainingXML struct {
	XMLName xml.Name `xml:"attrs-train"`
	Str     int      `xml:"str,attr"`
	Con     int      `xml:"con,attr"`
	Dex     int      `xml:"dex,attr"`
	Wis     int      `xml:"wis,attr"`
	Int     int      `xml:"int,attr"`
	Reqs    ReqsXML  `xml:"reqs"`
}

// xmlTrainings parses specified trainings to XML
// trainings node.
func xmlTrainings(trainings ...train.Training) *TrainingsXML {
	xmlTrainings := new(TrainingsXML)
	for _, t := range trainings {
		switch t := t.(type) {
		case *train.AttrsTraining:
			xmlAt := xmlAttrsTraining(t)
			xmlTrainings.AttrsTrain = append(xmlTrainings.AttrsTrain, *xmlAt)
		}
	}
	return xmlTrainings
}

// xmlAttrsTraining parser specified attributes training to
// XML attributes training node.
func xmlAttrsTraining(at *train.AttrsTraining) *AttrsTrainingXML {
	xmlTrain := new(AttrsTrainingXML)
	xmlTrain.Str = at.Strenght()
	xmlTrain.Con = at.Constitution()
	xmlTrain.Dex = at.Dexterity()
	xmlTrain.Wis = at.Wisdom()
	xmlTrain.Int = at.Intelligence()
	xmlTrain.Reqs = *xmlReqs(at.Reqs()...)
	return xmlTrain
}

// buildTraining creates training from specified XML data.
func buildTrainings(xmlTrainings *TrainingsXML) []res.TrainingData {
	train := make([]res.TrainingData, 0)
	for _, xmlTrain := range xmlTrainings.AttrsTrain {
		atd := res.AttrsTrainingData{
			Str:  xmlTrain.Str,
			Con:  xmlTrain.Con,
			Dex:  xmlTrain.Dex,
			Wis:  xmlTrain.Wis,
			Int:  xmlTrain.Int,
			Reqs: buildReqs(&xmlTrain.Reqs),
		}
		train = append(train, atd)
	}
	return train
}
