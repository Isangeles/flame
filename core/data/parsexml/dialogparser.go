/*
 * dialogparser.go
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
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/log"
)

// Struct for dialogs XML base node.x
type DialogsBaseXML struct {
	XMLName xml.Name    `xml:"base"`
	Dialogs []DialogXML `xml:"dialog"`
}

// Struct for dialog XML node.
type DialogXML struct {
	XMLName xml.Name  `xml:"dialog"`
	ID      string    `xml:"id,attr"`
	Texts   []TextXML `xml:"text"`
}

// Struct for dialog text XML node.
type TextXML struct {
	XMLName   xml.Name     `xml:"text"`
	ID        string       `xml:"id,attr"`
	Ordinal   string       `xml:"ordinal,attr"`
	Start     bool         `xml:"start,attr"`
	Answers   []AnswerXML  `xml:"answer"`
	Reqs      ReqsXML      `xml:"reqs"`
	TalkerMods ModifiersXML `xml:"talker>modifiers"`
}

// Struct for dialog answer XML node.
type AnswerXML struct {
	XMLName xml.Name `xml:"answer"`
	ID      string   `xml:"id,attr"`
	To      string   `xml:"to,attr"`
	Reqs    ReqsXML  `xml:"reqs"`
}

// UnmarshalDialogsBase retrieves dialogs data from specified XML data.
func UnmarshalDialogsBase(data io.Reader) ([]*res.DialogData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(DialogsBaseXML)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v", err)
	}
	dialogs := make([]*res.DialogData, 0)
	for _, xmlDialog := range xmlBase.Dialogs {
		dialog, err := buildDialogData(xmlDialog)
		if err != nil {
			log.Err.Printf("xml:unmarshal_dialog:build_data_fail:%v", err)
			continue
		}
		dialogs = append(dialogs, dialog)
		fmt.Printf("xml_dlg:%s:texts:%v\n", xmlDialog.ID, xmlDialog.Texts)
	}
	return dialogs, nil
}

// buildDialogData creates new dialog data from specified xml data.
func buildDialogData(xmlDialog DialogXML) (*res.DialogData, error) {
	dd := new(res.DialogData)
	dd.ID = xmlDialog.ID
	for _, xmlText := range xmlDialog.Texts {
		dtd := new(res.DialogTextData)
		dtd.ID = xmlText.ID
		dtd.OrdinalID = xmlText.Ordinal
		dtd.Start = xmlText.Start
		for _, xmlAnswer := range xmlText.Answers {
			dad := new(res.DialogAnswerData)
			dad.ID = xmlAnswer.ID
			dad.To = xmlAnswer.To
			dad.Reqs = buildReqs(&xmlAnswer.Reqs)
			dtd.Answers = append(dtd.Answers, dad)
		}
		dtd.Reqs = buildReqs(&xmlText.Reqs)
		dd.Texts = append(dd.Texts, dtd)
	}
	return dd, nil
}
