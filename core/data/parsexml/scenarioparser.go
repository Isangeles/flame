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
type ScenarioXML struct {
	XMLName  xml.Name  `xml:"scenario"`
	ID       string    `xml:"id,attr"`
	Mainarea AreaXML   `xml:"mainarea"`
	Subareas []AreaXML `xml:"area"`
}

// Struct for XML area node.
type AreaXML struct {
	ID      string     `xml:"id,attr"`
	NPCs    NpcsXML    `xml:"npcs"`
	Objects ObjectsXML `xml:"objects"`
}

// Struct for XML npcs node.
type NpcsXML struct {
	XMLName    xml.Name      `xml:"npcs"`
	Characters []AreaCharXML `xml:"char"`
}

// Struct for XML objects node.
type ObjectsXML struct {
	XMLName xml.Name        `xml:"objects"`
	Nodes   []AreaObjectXML `xml:"object"`
}

// Struct for XML object node.
type AreaObjectXML struct {
	XMLName  xml.Name `xml:"object"`
	ID       string   `xml:"id,attr"`
	Position string   `xml:"position,attr"`
}

// Struct for XML area character node.
type AreaCharXML struct {
	XMLName  xml.Name `xml:"char"`
	ID       string   `xml:"id,attr"`
	Position string   `xml:"position,attr"`
}

// UnmarshalScenario parses scenario from XML data.
func UnmarshalScenario(data io.Reader) (*res.ModuleScenarioData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlScen := new(ScenarioXML)
	err := xml.Unmarshal(doc, xmlScen)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml:%v", err)
	}
	scenData := buildScenarioData(xmlScen)
	return scenData, nil
}

// buildScenarioData creates scenario data from specified
// XML data.
func buildScenarioData(xmlScen *ScenarioXML) *res.ModuleScenarioData {
	// Areas.
	areas := make([]res.ModuleAreaData, 0)
	// Mainarea.
	mainarea := buildAreaData(&xmlScen.Mainarea)
	mainarea.Main = true
	areas = append(areas, *mainarea)
	// Subareas.
	for _, xmlArea := range xmlScen.Subareas {
		area := buildAreaData(&xmlArea)
		areas = append(areas, *area)
	}
	scen := res.ModuleScenarioData{
		ID:    xmlScen.ID,
		Areas: areas,
	}
	return &scen
}

// buildAreaData creates area data from specified XML data.
func buildAreaData(xmlArea *AreaXML) *res.ModuleAreaData {
	area := res.ModuleAreaData{ID: xmlArea.ID}
	// Characters.
	for _, xmlChar := range xmlArea.NPCs.Characters {
		x, y, err := UnmarshalPosition(xmlChar.Position)
		if err != nil {
			log.Err.Printf("xml:build_area:%s:build_char:%s:fail_to_parse_position:%v",
				xmlArea.ID, xmlChar.ID, err)
			continue
		}
		char := res.AreaCharData{
			ID: xmlChar.ID,
			PosX: x,
			PosY: y,
		}
		area.NPCS = append(area.NPCS, char)
	}
	// Objects.
	for _, xmlOb := range xmlArea.Objects.Nodes {
		x, y, err := UnmarshalPosition(xmlOb.Position)
		if err != nil {
			log.Err.Printf("xml:build_area:%s:build_object:%s:fail_to_parse_position:%v",
				xmlArea.ID, xmlOb.ID, err)
			continue
		}
		ob := res.AreaObjectData{
			ID: xmlOb.ID,
			PosX: x,
			PosY: y,
		}
		area.Objects = append(area.Objects, ob)
	}
	return &area
}
