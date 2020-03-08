/*
 * areaparser.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

// Struct for XML area node.
type Area struct {
	XMLName    xml.Name        `xml:"area"`
	ID         string          `xml:"id,attr"`
	Characters []AreaCharacter `xml:"characters>character"`
	Objects    []AreaObject    `xml:"objects>object"`
	Subareas   []Area          `xml:"subareas>area"`
}

// Struct for XML object node.
type AreaObject struct {
	XMLName  xml.Name `xml:"object"`
	ID       string   `xml:"id,attr"`
	Position string   `xml:"position,attr"`
}

// Struct for XML area character node.
type AreaCharacter struct {
	XMLName  xml.Name `xml:"character"`
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
		return nil, fmt.Errorf("unable to unmarshal xml: %v", err)
	}
	areaData := buildAreaData(xmlArea)
	return areaData, nil
}

// MarshalArea parses area data to XML string.
func MarshalArea(areaData *res.AreaData) (string, error) {
	xmlArea := xmlArea(areaData)
	out, err := xml.Marshal(xmlArea)
	if err != nil {
		return "", fmt.Errorf("unable to marshal data: %v", err)
	}
	return string(out[:]), nil
}

// xmlArea creates XML struct from specified module
// area data.
func xmlArea(data *res.AreaData) *Area {
	xmlData := new(Area)
	xmlData.ID = data.ID
	// Characters.
	for _, c := range data.Characters {
		xmlChar := new(AreaCharacter)
		xmlChar.ID = c.ID
		xmlChar.Position = MarshalPosition(c.PosX, c.PosY)
		xmlChar.AI = c.AI
		xmlData.Characters = append(xmlData.Characters, *xmlChar)
	}
	// Objects.
	for _, o := range data.Objects {
		xmlOb := new(AreaObject)
		xmlOb.ID = o.ID
		xmlOb.Position = MarshalPosition(o.PosX, o.PosY)
		xmlData.Objects = append(xmlData.Objects, *xmlOb)
	}
	// Subareas.
	for _, sad := range data.Subareas {
		xmlSubarea := xmlArea(&sad)
		xmlData.Subareas = append(xmlData.Subareas, *xmlSubarea)
	}
	return xmlData
}

// buildAreaData creates area data from specified XML data.
func buildAreaData(xmlArea *Area) *res.AreaData {
	area := res.AreaData{ID: xmlArea.ID}
	// Characters.
	for _, xmlChar := range xmlArea.Characters {
		x, y, err := UnmarshalPosition(xmlChar.Position)
		if err != nil {
			log.Err.Printf("xml: build area: %s: build char: %s: unable to parse position: %v",
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
			log.Err.Printf("xml: build area: %s: build object: %s: unable to parse position: %v",
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
