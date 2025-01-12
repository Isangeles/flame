/*
 * quest.go
 *
 * Copyright 2019-2025 Dariusz Sikora <ds@isangeles.dev>
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

// Struct for quests data.
type QuestsData struct {
	XMLName xml.Name    `xml:"quests" json:"-"`
	Quests  []QuestData `xml:"quest" json:"quests"`
}

// Struct for quest data.
type QuestData struct {
	ID     string           `xml:"id,attr" json:"id"`
	Stages []QuestStageData `xml:"stage" json:"stages"`
}

// Struct for quest stage data.
type QuestStageData struct {
	ID            string               `xml:"id,attr" json:"id"`
	Ordinal       int                  `xml:"ordinal,attr" json:"ordinal"`
	Next          string               `xml:"next,attr" json:"next"`
	Start         bool                 `xml:"start,attr" json:"start"`
	End           bool                 `xml:"end,attr" json:"end"`
	Objectives    []QuestObjectiveData `xml:"objectives>objective" json:"objectives"`
	StartFlags    []FlagData           `xml:"on-start>flags>flag" json:"start-flags"`
	CompleteFlags []FlagData           `xml:"on-complete>flags>flag" json:"complete-flags"`
}

// Struct for quest objective data.
type QuestObjectiveData struct {
	ID       string   `xml:"id,attr" json:"id"`
	Finisher bool     `xml:"finisher,attr" json:"finisher"`
	Reqs     ReqsData `xml:"reqs" json:"reqs"`
}

// Struct for quest log data.
type QuestLogData struct {
	Quests []QuestLogQuestData `xml:"quest" json:"quests"`
}

// Struct for quest data from quest log.
type QuestLogQuestData struct {
	ID    string `xml:"id,attr" json:"id"`
	Stage string `xml:"stage,attr" json:"stage"`
}
