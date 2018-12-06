/*
 * scenarioparser.go
 * 
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/core/module/scenario"
)

// Struct for XML scenario node.
type XMLScenario struct {
	XMLName  xml.Name `xml:"scenario"`
	Id       string   `xml:"id,attr"`
	Mainarea XMLArea  `xml:"mainarea"`
}

// Struct for XML area node.
type XMLArea struct {
	Id string `xml:"id,attr"`
}

// UnmarshalScenarioXML parses scenario from XML data.
func UnmarshalScenarioXML(data io.Reader) (*scenario.Scenario, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlScen := new(XMLScenario)
	err := xml.Unmarshal(doc, xmlScen)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml:%v", err)
	}

	mainarea := scenario.NewArea(xmlScen.Mainarea.Id)
	scen := scenario.NewScenario(xmlScen.Id, mainarea, []*scenario.Area{})

	return scen, nil
}

