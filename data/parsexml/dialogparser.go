/*
 * dialogparser.go
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

package parsexml

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

// Struct for dialogs XML base node.
type Dialogs struct {
	XMLName xml.Name `xml:"dialogs"`
	Dialogs []Dialog `xml:"dialog"`
}

// Struct for dialog XML node.
type Dialog struct {
	XMLName xml.Name      `xml:"dialog"`
	ID      string        `xml:"id,attr"`
	Stages  []DialogStage `xml:"stage"`
}

// Struct for dialog stage XML node.
type DialogStage struct {
	XMLName    xml.Name  `xml:"stage"`
	ID         string    `xml:"id,attr"`
	Ordinal    string    `xml:"ordinal,attr"`
	Start      bool      `xml:"start,attr"`
	Answers    []Answer  `xml:"answer"`
	Reqs       Reqs      `xml:"reqs"`
	TalkerMods Modifiers `xml:"talker>modifiers"`
	OwnerMods  Modifiers `xml:"owner>modifiers"`
}

// Struct for dialog answer XML node.
type Answer struct {
	XMLName    xml.Name  `xml:"answer"`
	ID         string    `xml:"id,attr"`
	To         string    `xml:"to,attr"`
	Reqs       Reqs      `xml:"reqs"`
	TalkerMods Modifiers `xml:"talker>modifiers"`
	OwnerMods  Modifiers `xml:"owner>modifiers"`
}

// UnmarshalDialogs retrieves dialogs data from specified XML data.
func UnmarshalDialogs(data io.Reader) ([]*res.DialogData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(Dialogs)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal xml data: %v", err)
	}
	dialogs := make([]*res.DialogData, 0)
	for _, xmlDialog := range xmlBase.Dialogs {
		dialog, err := buildDialogData(xmlDialog)
		if err != nil {
			log.Err.Printf("xml: unmarshal dialog: unable to build data: %v", err)
			continue
		}
		dialogs = append(dialogs, dialog)
	}
	return dialogs, nil
}

// buildDialogData creates new dialog data from specified XML data.
func buildDialogData(xmlDialog Dialog) (*res.DialogData, error) {
	dd := new(res.DialogData)
	dd.ID = xmlDialog.ID
	for _, xmlStage := range xmlDialog.Stages {
		dtd := new(res.DialogStageData)
		dtd.ID = xmlStage.ID
		dtd.OrdinalID = xmlStage.Ordinal
		dtd.Start = xmlStage.Start
		for _, xmlAnswer := range xmlStage.Answers {
			dad := new(res.DialogAnswerData)
			dad.ID = xmlAnswer.ID
			dad.To = xmlAnswer.To
			dad.End = xmlAnswer.To == "end"
			dad.Trade = xmlAnswer.To == "trade"
			dad.Training = xmlAnswer.To == "training"
			dad.Reqs = buildReqs(&xmlAnswer.Reqs)
			dad.TalkerMods = buildModifiers(&xmlAnswer.TalkerMods)
			dad.OwnerMods = buildModifiers(&xmlAnswer.OwnerMods)
			dtd.Answers = append(dtd.Answers, dad)
		}
		dtd.Reqs = buildReqs(&xmlStage.Reqs)
		dtd.TalkerMods = buildModifiers(&xmlStage.TalkerMods)
		dtd.OwnerMods = buildModifiers(&xmlStage.OwnerMods)
		dd.Stages = append(dd.Stages, dtd)
	}
	return dd, nil
}
