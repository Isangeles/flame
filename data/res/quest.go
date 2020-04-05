/*
 * quest.go
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

package res

// Struct for quest data.
type QuestData struct {
	ID     string
	Stages []QuestStageData
}

// Struct for quest stage data.
type QuestStageData struct {
	ID            string
	Ordinal       int
	Next          string
	Objectives    []QuestObjectiveData
	CompleteFlags []FlagData
}

// Struct for quest objective data.
type QuestObjectiveData struct {
	ID       string
	Finisher bool
	Reqs     ReqsData
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
