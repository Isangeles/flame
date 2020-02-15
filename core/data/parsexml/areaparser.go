/*
 * areaparser.go
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

// Struct for XML area node.
type Area struct {
	ID         string       `xml:"id,attr"`
	Characters []AreaChar   `xml:"npcs>char"`
	Objects    []AreaObject `xml:"objects>object"`
	Subareas   []Area       `xml:"subareas>area"`
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
	AI       bool     `xml:"ai,attr"`
}

// UnmarshalArea parses area from XML data to resource.
func UnmarshalArea(data io.Reader) (*res.AreaData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlArea := new(Area)
	err := xml.Unmarshal(doc, xmlArea)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal xml: %v", err)
	}
	areaData := buildAreaData(xmlArea)
	return areaData, nil
}

// MarshalArea parses scenario data to XML string.
func MarshalArea(areaData *res.AreaData) (string, error) {
	xmlArea := xmlArea(areaData)
	out, err := xml.Marshal(xmlArea)
	if err != nil {
		return "", fmt.Errorf("fail to marshal data: %v", err)
	}
	return string(out[:]), nil
}

// xmlArea creates XML struct from specified module
// area data.
func xmlArea(areaData *res.AreaData) *Area {
	xmlArea := new(Area)
	xmlArea.ID = areaData.ID
	for _, sad := range areaData.Subareas {
		xmlSubarea := new(Area)
		xmlArea.ID = sad.ID
		// Characters.
		xmlChars := xmlArea.Characters
		for _, npc := range sad.Characters {
			xmlNPC := new(AreaChar)
			xmlNPC.ID = npc.ID
			xmlNPC.Position = MarshalPosition(npc.PosX, npc.PosY)
			xmlNPC.AI = npc.AI
			xmlChars = append(xmlChars, *xmlNPC)
		}
		// Objects.
		xmlObjects := xmlArea.Objects
		for _, ob := range sad.Objects {
			xmlOb := new(AreaObject)
			xmlOb.ID = ob.ID
			xmlOb.Position = MarshalPosition(ob.PosX, ob.PosY)
			xmlObjects = append(xmlObjects, *xmlOb)
		}
		xmlArea.Subareas = append(xmlArea.Subareas, *xmlSubarea)
	}
	return xmlArea
}

// buildAreaData creates area data from specified XML data.
func buildAreaData(xmlArea *Area) *res.AreaData {
	area := res.AreaData{ID: xmlArea.ID}
	// Characters.
	for _, xmlChar := range xmlArea.Characters {
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
			AI:   xmlChar.AI,
		}
		area.Characters = append(area.Characters, char)
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
	// Subareas.
	for _, xmlSubarea := range xmlArea.Subareas {
		subarea := buildAreaData(&xmlSubarea)
		area.Subareas = append(area.Subareas, *subarea)
	}
	return &area
}
