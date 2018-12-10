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
)

// Struct for XML scenario node.
type ScenarioXML struct {
	XMLName  xml.Name  `xml:"scenario"`
	ID       string    `xml:"id,attr"`
	Mainarea AreaXML   `xml:"mainarea"`
	Subareas []AreaXML `xml:"area"`
}

// Struct for XML area node.
type AreaXML struct {
	ID      string  `xml:"id,attr"`
	NPCs    NpcsXML `xml:"npcs"`
}

// Struct for XML npcs node.
type NpcsXML struct {
	XMLName    xml.Name      `xml:"npcs"`
	Characters []AreaCharXML `xml:"char"`
}

// Struct for XML area character node.
type AreaCharXML struct {
	XMLName  xml.Name `xml:"char"`
	ID       string   `xml:"id,attr"`
	Position string   `xml:"position,attr"`
}

// UnmarshalScenario parses scenario from XML data.
func UnmarshalScenario(data io.Reader) (*ScenarioXML, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlScen := new(ScenarioXML)
	err := xml.Unmarshal(doc, xmlScen)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml:%v", err)
	}
	return xmlScen, nil
}

