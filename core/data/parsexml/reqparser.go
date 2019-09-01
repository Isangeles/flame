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
type Reqs struct {
	XMLName      xml.Name      `xml:"reqs"`
	LevelReqs    []LevelReq    `xml:"level-req"`
	GenderReqs   []GenderReq   `xml:"gender-req"`
	FlagReqs     []FlagReq     `xml:"flag-req"`
	ItemReqs     []ItemReq     `xml:"item-req"`
	CurrencyReqs []CurrencyReq `xml:"currency-req"`
}

// Struct for level requirement XML node.
type LevelReq struct {
	XMLName xml.Name `xml:"level-req"`
	Min     int      `xml:"min,value"`
	Max     int      `xml:"max,value"`
}

// Struct for gender requirement XML node.
type GenderReq struct {
	XMLName xml.Name `xml:"gender-req"`
	Type    string   `xml:"type,attr"`
}

// Struct for flag requirement XML node.
type FlagReq struct {
	XMLName xml.Name `xml:"flag-req"`
	ID      string   `xml:"id,attr"`
	Off     bool     `xml:"off,attr"`
}

// Struct for item requirement XML node.
type ItemReq struct {
	XMLName xml.Name `xml:"item-req"`
	ID      string   `xml:"id,attr"`
	Amount  int      `xml:"amount,attr"`
	Charge  bool     `xml:"charge,attr"`
}

// Struct for currency requirement XML node.
type CurrencyReq struct {
	XMLName xml.Name `xml:"currency-req"`
	Amount  int      `xml:"amount,attr"`
	Charge  bool     `xml:"charge,attr"`
}

// xmlReqs parses specified requirements to XML
// reqs node.
func xmlReqs(reqs ...req.Requirement) *Reqs {
	xmlReqs := new(Reqs)
	for _, r := range reqs {
		switch r := r.(type) {
		case *req.LevelReq:
			xmlReq := xmlLevelReq(r)
			xmlReqs.LevelReqs = append(xmlReqs.LevelReqs, *xmlReq)
		case *req.CurrencyReq:
			xmlReq := xmlCurrencyReq(r)
			xmlReqs.CurrencyReqs = append(xmlReqs.CurrencyReqs, *xmlReq)
		}
	}
	return xmlReqs
}

// xmlLevelReq parses specified level requirement to
// XML level req node.
func xmlLevelReq(req *req.LevelReq) *LevelReq {
	xmlReq := new(LevelReq)
	xmlReq.Min = req.MinLevel()
	return xmlReq
}

// xmlCurrencyReq parses specified currency requirement
// to XML currenct req node.
func xmlCurrencyReq(r *req.CurrencyReq) *CurrencyReq {
	xmlReq := new(CurrencyReq)
	xmlReq.Amount = r.Amount()
	xmlReq.Charge = r.Charge()
	return xmlReq
}

// buildReqs creates requirements from specified
// XML data.
func buildReqs(xmlReqs *Reqs) []res.ReqData {
	reqs := make([]res.ReqData, 0)
	// Level reqs.
	for _, xmlReq := range xmlReqs.LevelReqs {
		req := res.LevelReqData{
			Min: xmlReq.Min,
			Max: xmlReq.Max,
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
		req := res.FlagReqData{
			ID:  xmlReq.ID,
			Off: xmlReq.Off,
		}
		reqs = append(reqs, req)
	}
	// Item reqs.
	for _, xmlReq := range xmlReqs.ItemReqs {
		req := res.ItemReqData{
			ID:     xmlReq.ID,
			Amount: xmlReq.Amount,
			Charge: xmlReq.Charge,
		}
		reqs = append(reqs, req)
	}
	// Currency reqs.
	for _, xmlReq := range xmlReqs.CurrencyReqs {
		req := res.CurrencyReqData{
			Amount: xmlReq.Amount,
			Charge: xmlReq.Charge,
		}
		reqs = append(reqs, req)
	}
	return reqs
}
