/*
 * reqparser.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/core/module/req"
	"github.com/isangeles/flame/log"
)

// Struct for requirements XML node.
type ReqsXML struct {
	XMLName    xml.Name       `xml:"reqs"`
	LevelReqs  []LevelReqXML  `xml:"level-req"`
	GenderReqs []GenderReqXML `xml:"gender-req"`
	FlagReqs   []FlagReqXML   `xml:"flag-req"`
	ItemReqs   []ItemReqXML   `xml:"item-req"`
}

// Struct for level requirement XML node.
type LevelReqXML struct {
	XMLName  xml.Name `xml:"level-req"`
	MinLevel int      `xml:"min,value"`
}

// Struct for gender requirement XML node.
type GenderReqXML struct {
	XMLName xml.Name `xml:"gender-req"`
	Type    string   `xml:"type,attr"`
}

// Struct for flag requirement XML node.
type FlagReqXML struct {
	XMLName xml.Name `xml:"flag-req"`
	ID      string   `xml:"id,attr"`
}

// Struct for item requirement XML node.
type ItemReqXML struct {
	XMLName xml.Name `xml:"item-req"`
	ID      string   `xml:"id,attr"`
	Amount  int      `xml:"amount,attr"`
}

// xmlLevelReq parses specified level requirement to
// XML level req node.
func xmlLevelReq(req *req.LevelReq) *LevelReqXML {
	xmlReq := new(LevelReqXML)
	xmlReq.MinLevel = req.MinLevel()
	return xmlReq
}

// buildReqs creates requirements from specified
// XML data.
func buildReqs(xmlReqs *ReqsXML) []res.ReqData {
	reqs := make([]res.ReqData, 0)
	// Level reqs.
	for _, xmlReq := range xmlReqs.LevelReqs {
		req := res.LevelReqData{
			Min: xmlReq.MinLevel,
			Max: -1, // TODO: support max value
		}
		reqs = append(reqs, req)
	}
	// Gender reqs.
	for _, xmlReq := range xmlReqs.GenderReqs {
		gen, err := UnmarshalGender(xmlReq.Type)
		if err != nil {
			log.Err.Printf("xml:parse_req:fail_to_parse_gender:%v", err)
		}
		req := res.GenderReqData{
			Type: int(gen),
		}
		reqs = append(reqs, req)
	}
	// Flag reqs.
	for _, xmlReq := range xmlReqs.FlagReqs {
		req := res.FlagReqData{xmlReq.ID}
		reqs = append(reqs, req)
	}
	// Item reqs.
	for _, xmlReq := range xmlReqs.ItemReqs {
		req := res.ItemReqData{xmlReq.ID, xmlReq.Amount}
		reqs = append(reqs, req)
	}
	return reqs
}
