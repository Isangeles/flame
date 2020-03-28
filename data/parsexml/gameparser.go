/*
 * gameparser.go
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

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/module/area"
)

// Struct for saved game XML node.
type SavedGame struct {
	XMLName xml.Name     `xml:"game"`
	Chapter SavedChapter `xml:"chapter"`
}

// Struct for saved chapter XML node.
type SavedChapter struct {
	XMLName xml.Name    `xml:"chapter"`
	ID      string      `xml:"id,attr"`
	Areas   []SavedArea `xml:"area"`
}

// Struct for saved area area XML node.
type SavedArea struct {
	XMLName    xml.Name    `xml:"area"`
	ID         string      `xml:"id,attr"`
	Mainarea   bool        `xml:"mainarea,attr"`
	Characters []Character `xml:"characters>char"`
	Objects    []Object    `xml:"objects>object"`
	Subareas   []SavedArea `xml:"subareas>area"`
}

// MarshalGame parses specified game to XML savegame data.
func MarshalGame(game *flame.Game) (string, error) {
	xmlGame := new(SavedGame)
	// Chapter.
	chapter := game.Module().Chapter()
	if chapter == nil {
		return "", fmt.Errorf("no game chapter set")
	}
	xmlChapter := &xmlGame.Chapter
	xmlChapter.ID = chapter.Conf().ID
	// Areas.
	for _, a := range chapter.Areas() {
		xmlArea := SavedArea{}
		xmlArea.ID = a.ID()
		xmlArea = *marshalGameArea(game, a)
		xmlChapter.Areas = append(xmlChapter.Areas, xmlArea)
	}
	out, err := xml.Marshal(xmlGame)
	if err != nil {
		return "", fmt.Errorf("unable to marshal data")
	}
	return string(out[:]), nil
}

// UnmarshalGameData parses specified XML data to game
// resource.
func UnmarshalGame(data io.Reader) (*res.GameData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlGame := new(SavedGame)
	err := xml.Unmarshal(doc, xmlGame)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal xml data: %v", err)
	}
	game := new(res.GameData)
	// Chapter.
	game.SavedChapter.ID = xmlGame.Chapter.ID
	// Areas.
	for _, xmlArea := range xmlGame.Chapter.Areas {
		area := *unmarshalGameArea(xmlArea)
		game.SavedChapter.Areas = append(game.SavedChapter.Areas, area)
	}
	return game, nil
}

// marshalGameArea parses specified area to saved area node.
func marshalGameArea(game *flame.Game, a *area.Area) *SavedArea {
	xmlArea := new(SavedArea)
	xmlArea.ID = a.ID()
	// Characters.
	for _, c := range a.Characters() {
		xmlChar := xmlCharacter(c)
		xmlArea.Characters = append(xmlArea.Characters, *xmlChar)
	}
	// Objects.
	for _, o := range a.Objects() {
		xmlObject := xmlObject(o)
		xmlArea.Objects = append(xmlArea.Objects, *xmlObject)
	}
	// Subareas.
	for _, sa := range a.Subareas() {
		xmlSubarea := marshalGameArea(game, sa)
		xmlArea.Subareas = append(xmlArea.Subareas, *xmlSubarea)
	}
	return xmlArea
}

// unmarshalGameArea parses specified saved area node to area data.
func unmarshalGameArea(xmlArea SavedArea) *res.SavedAreaData {
	area := new(res.SavedAreaData)
	area.ID = xmlArea.ID
	// Characters.
	for _, xmlChar := range xmlArea.Characters {
		charData, err := buildCharacterData(&xmlChar)
		if err != nil {
			log.Err.Printf("xml: build game: area: %s: character: %s: %v",
				xmlArea.ID, xmlChar.ID, err)
			continue
		}
		area.Chars = append(area.Chars, *charData)
	}
	// Objects.
	for _, xmlOb := range xmlArea.Objects {
		obData, err := buildObjectData(&xmlOb)
		if err != nil {
			log.Err.Printf("xml: build game: area: %s: object: %s: %v",
				xmlArea.ID, xmlOb.ID, err)
			continue
		}
		area.Objects = append(area.Objects, *obData)
	}
	// Subareas.
	for _, xmlSa := range xmlArea.Subareas {
		saData := unmarshalGameArea(xmlSa)
		area.Subareas = append(area.Subareas, *saData)
	}
	return area
}
