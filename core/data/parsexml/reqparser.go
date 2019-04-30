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
)

// Struct for requirements XML node.
type ReqsXML struct {
	XMLName   xml.Name      `xml:"reqs"`
	LevelReqs []LevelReqXML `xml:"levelReq"`
}

// Struct for level requirement XML node.
type LevelReqXML struct {
	XMLName  xml.Name `xml:"levelReq"`
	MinLevel int      `xml:"min,value"`
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
	return reqs
}
