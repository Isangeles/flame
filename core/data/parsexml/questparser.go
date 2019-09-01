/*
 * questparser.go
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

// Struct for quests base XML node.
type QuestsBase struct {
	XMLName xml.Name `xml:"quests"`
	Quests  []Quest  `xml:"quest"`
}

// Struct for quest XML node.
type Quest struct {
	XMLName xml.Name     `xml:"quest"`
	ID      string       `xml:"id,attr"`
	Stages  []QuestStage `xml:"stage"`
}

// Struct for quest stage XML node.
type QuestStage struct {
	XMLName       xml.Name         `xml:"stage"`
	ID            string           `xml:"id,attr"`
	Ordinal       int              `xml:"ordinal,attr"`
	Next          string           `xml:"next,attr"`
	Objectives    []QuestObjective `xml:"objectives>objective"`
	CompleteFlags []Flag           `xml:"on-complete>flags>flag"`
}

// Struct for quest objective XML node.
type QuestObjective struct {
	XMLName  xml.Name `xml:"objective"`
	ID       string   `xml:"id,attr"`
	Finisher bool     `xml:"finisher,attr"`
	Reqs     Reqs     `xml:"reqs"`
}

// UnmarashalQuests retrieves quests data from specified XML data.
func UnmarshalQuests(data io.Reader) ([]*res.QuestData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(QuestsBase)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v",
			err)
	}
	quests := make([]*res.QuestData, 0)
	for _, xmlQuest := range xmlBase.Quests {
		quest, err := buildQuestData(xmlQuest)
		if err != nil {
			log.Err.Printf("xml:unmarshal_quest:build_data_fail:%v", err)
			continue
		}
		quests = append(quests, quest)
	}
	return quests, nil
}

// buildQuestData creates new quest data from specified XML data.
func buildQuestData(xmlQuest Quest) (*res.QuestData, error) {
	qd := new(res.QuestData)
	qd.ID = xmlQuest.ID
	for _, xmlStage := range xmlQuest.Stages {
		qsd := res.QuestStageData{}
		qsd.ID = xmlStage.ID
		qsd.Ordinal = xmlStage.Ordinal
		qsd.Next = xmlStage.Next
		for _, xmlObjective := range xmlStage.Objectives {
			qod := res.QuestObjectiveData{}
			qod.ID = xmlObjective.ID
			qod.Finisher = xmlObjective.Finisher
			qod.Reqs = buildReqs(&xmlObjective.Reqs)
			qsd.Objectives = append(qsd.Objectives, qod)
		}
		for _, xmlFlag := range xmlStage.CompleteFlags {
			fd := buildFlagData(xmlFlag)
			qsd.CompleteFlags = append(qsd.CompleteFlags, fd)
		}
		qd.Stages = append(qd.Stages, qsd)
	}
	return qd, nil
}
