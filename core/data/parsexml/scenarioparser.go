/*
 * scenarioparser.go
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
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/log"
)

// Struct for XML scenario node.
type Scenario struct {
	XMLName xml.Name `xml:"scenario"`
	ID      string   `xml:"id,attr"`
	Area    Area     `xml:"area"`
}

// Struct for XML area node.
type Area struct {
	ID       string       `xml:"id,attr"`
	NPCs     []AreaChar   `xml:"npcs>char"`
	Objects  []AreaObject `xml:"objects>object"`
	Subareas []Area       `xml:"subareas>area"`
}

// Struct for XML object node.
type AreaObject struct {
	XMLName  xml.Name `xml:"object"`
	ID       string   `xml:"id,attr"`
	Position string   `xml:"position,attr"`
}

// Struct for XML area character node.
type AreaChar struct {
	XMLName  xml.Name `xml:"char"`
	ID       string   `xml:"id,attr"`
	Position string   `xml:"position,attr"`
}

// UnmarshalScenario parses scenario from XML data.
func UnmarshalScenario(data io.Reader) (*res.ModuleScenarioData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlScen := new(Scenario)
	err := xml.Unmarshal(doc, xmlScen)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal xml: %v", err)
	}
	scenData := buildScenarioData(xmlScen)
	return scenData, nil
}

// MarshalScenario parses scenario data to XML string.
func MarshalScenario(scenData *res.ModuleScenarioData) (string, error) {
	xmlScen := xmlScenario(scenData)
	out, err := xml.Marshal(xmlScen)
	if err != nil {
		return "", fmt.Errorf("fail to marshal char: %v", err)
	}
	return string(out[:]), nil
}

// xmlScenario creates XML struct from specified module
// scenario data.
func xmlScenario(scen *res.ModuleScenarioData) *Scenario {
	xmlScen := new(Scenario)
	xmlScen.ID = scen.ID
	xmlScen.Area.ID = scen.Area.ID
	for _, ad := range scen.Area.Subareas {
		xmlArea := new(Area)
		xmlArea.ID = ad.ID
		// Characters.
		xmlChars := xmlArea.NPCs
		for _, npc := range ad.NPCS {
			xmlNPC := new(AreaChar)
			xmlNPC.ID = npc.ID
			xmlNPC.Position = MarshalPosition(npc.PosX, npc.PosY)
			xmlChars = append(xmlChars, *xmlNPC)
		}
		// Objects.
		xmlObjects := xmlArea.Objects
		for _, ob := range ad.Objects {
			xmlOb := new(AreaObject)
			xmlOb.ID = ob.ID
			xmlOb.Position = MarshalPosition(ob.PosX, ob.PosY)
			xmlObjects = append(xmlObjects, *xmlOb)
		}
		xmlScen.Area.Subareas = append(xmlScen.Area.Subareas, *xmlArea)
	}
	return xmlScen
}

// buildScenarioData creates scenario data from specified
// XML data.
func buildScenarioData(xmlScen *Scenario) *res.ModuleScenarioData {
	// Mainarea.
	mainarea := buildAreaData(&xmlScen.Area)
	// Subareas.
	for _, xmlArea := range xmlScen.Area.Subareas {
		area := buildAreaData(&xmlArea)
		mainarea.Subareas = append(mainarea.Subareas, *area)
	}
	scen := res.ModuleScenarioData{
		ID:   xmlScen.ID,
		Area: *mainarea,
	}
	return &scen
}

// buildAreaData creates area data from specified XML data.
func buildAreaData(xmlArea *Area) *res.ModuleAreaData {
	area := res.ModuleAreaData{ID: xmlArea.ID}
	// Characters.
	for _, xmlChar := range xmlArea.NPCs {
		x, y, err := UnmarshalPosition(xmlChar.Position)
		if err != nil {
			log.Err.Printf("xml: build area: %s: build char: %s: fail to parse position: %v",
				xmlArea.ID, xmlChar.ID, err)
			continue
		}
		char := res.AreaCharData{
			ID:   xmlChar.ID,
			PosX: x,
			PosY: y,
		}
		area.NPCS = append(area.NPCS, char)
	}
	// Objects.
	for _, xmlOb := range xmlArea.Objects {
		x, y, err := UnmarshalPosition(xmlOb.Position)
		if err != nil {
			log.Err.Printf("xml: build area: %s: build object: %s: fail to parse position: %v",
				xmlArea.ID, xmlOb.ID, err)
			continue
		}
		ob := res.AreaObjectData{
			ID:   xmlOb.ID,
			PosX: x,
			PosY: y,
		}
		area.Objects = append(area.Objects, ob)
	}
	return &area
}
