/*
 * req.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for reqs data.
type ReqsData struct {
	XMLName         xml.Name             `xml:"reqs" json:"-"`
	LevelReqs       []LevelReqData       `xml:"level-req" json:"level-reqs"`
	GenderReqs      []GenderReqData      `xml:"gender-req" json:"gender-reqs"`
	FlagReqs        []FlagReqData        `xml:"flag-req" json:"flag-reqs"`
	ItemReqs        []ItemReqData        `xml:"item-req" json:"item-reqs"`
	CurrencyReqs    []CurrencyReqData    `xml:"currency-req" json:"currency-req"`
	TargetRangeReqs []TargetRangeReqData `xml:"target-range-req" json:"target-range-req"`
	KillReqs        []KillReqData        `xml:"kill-req" json:"kill-req"`
	QuestReqs       []QuestReqData       `xml:"quest-req" json:"quest-req"`
	HealthReqs      []HealthReqData      `xml:"health-req" json:"health-req"`
}

// Struct for level requirement data.
type LevelReqData struct {
	Min int `xml:"min,attr" json:"min"`
	Max int `xml:"max,attr" json:"max"`
}

// Struct for gender requirement data.
type GenderReqData struct {
	Gender string `xml:"type,attr" json:"gender"`
}

// Struct for flag requirement data.
type FlagReqData struct {
	ID  string `xml:"id,attr" json:"id"`
	Off bool   `xml:"off,attr" json:"off"`
}

// Struct for item requirement data.
type ItemReqData struct {
	ID     string `xml:"id,attr" json:"id"`
	Amount int    `xml:"amount,attr" json:"amount"`
	Charge bool   `xml:"charge,attr" json:"charge"`
}

// Struct for currency requirement data.
type CurrencyReqData struct {
	Amount int  `xml:"amount,attr" json:"amount"`
	Charge bool `xml:"charge,attr" json:"charge"`
}

// Struct for target requirement data.
type TargetRangeReqData struct {
	MinRange float64 `xml:"min-range,attr" json:"min-range"`
}

// Struct for kill requirement data.
type KillReqData struct {
	ID     string `xml:"id,attr" json:"id"`
	Amount int    `xml:"amount,attr" json:"amount"`
}

// Struct for quest requirement data.
type QuestReqData struct {
	ID        string `xml:"id,attr" json:"id"`
	Completed bool   `xml:"completed,attr" json:"completed"`
}

// Struct for health requirement data.
type HealthReqData struct {
	Percent int  `xml:"percent,attr" json:"percent"`
	Less    bool `xml:"less,attr" json:"less"`
}
