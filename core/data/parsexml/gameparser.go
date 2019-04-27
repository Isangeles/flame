/*
 * gameparser.go
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

	"github.com/isangeles/flame/core/data/save"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/log"
)

// Struct for saved game XML node.
type SavedGameXML struct {
	XMLName xml.Name        `xml:"game"`
	Name    string          `xml:"name,attr"`
	Chapter SavedChapterXML `xml:"chapter"`
}

// Struct for saved chapter XML node.
type SavedChapterXML struct {
	XMLName   xml.Name           `xml:"chapter"`
	ID        string             `xml:"id,attr"`
	Scenarios []SavedScenarioXML `xml:"scenario"`
}

// Struct for saved scenario XML node.
type SavedScenarioXML struct {
	XMLName   xml.Name      `xml:"scenario"`
	ID        string        `xml:"id,attr"`
	AreasNode SavedAreasXML `xml:"areas"`
}

// Struct for saved areas XML node.
type SavedAreasXML struct {
	XMLName xml.Name       `xml:"areas"`
	Areas   []SavedAreaXML `xml:"area"`
}

// Struct for saved scenario area XML node.
type SavedAreaXML struct {
	XMLName     xml.Name           `xml:"area"`
	ID          string             `xml:"id,attr"`
	Mainarea    bool               `xml:"mainarea,attr"`
	CharsNode   SavedCharactersXML `xml:"characters"`
	ObjectsNode SavedObjectsXML    `xml:"objects"`
}

// Struct for saved characters XML node.
type SavedCharactersXML struct {
	XMLName    xml.Name       `xml:"characters"`
	Characters []CharacterXML `xml:"char"`
}

// Struct for saved objects XML node.
type SavedObjectsXML struct {
	XMLName xml.Name    `xml:"objects"`
	Nodes   []ObjectXML `xml:"object"`
}

// MarshalGame parses specified game to XML
// savegame data.
func MarshalSaveGame(game *save.SaveGame) (string, error) {
	xmlGame := new(SavedGameXML)
	xmlGame.Name = game.Name
	// Chapter.
	chapter := game.Mod.Chapter()
	if chapter == nil {
		return "", fmt.Errorf("no game chapter set")
	}
	xmlChapter := &xmlGame.Chapter
	xmlChapter.ID = chapter.Conf().ID
	// Scenarios.
	for _, s := range chapter.Scenarios() {
		xmlScenario := SavedScenarioXML{}
		xmlScenario.ID = s.ID()
		// Areas.
		xmlAreas := &xmlScenario.AreasNode
		for _, a := range s.Areas() {
			xmlArea := SavedAreaXML{}
			xmlArea.ID = a.ID()
			if a.ID() == s.Mainarea().ID() {
				xmlArea.Mainarea = true
			}
			// Characters.
			xmlChars := &xmlArea.CharsNode
			for _, c := range a.Characters() {
				xmlChar := xmlCharacter(c)
				serialID := xmlChar.ID + "_" + xmlChar.Serial
				for _, pc := range game.Players {	
					if pc.SerialID() == serialID {
						xmlChar.PC = true
					}
				}
				xmlChars.Characters = append(xmlChars.Characters,
					*xmlChar)
			}
			// Objects.
			xmlObjects := &xmlArea.ObjectsNode
			for _, o := range a.Objects() {
				xmlObject := xmlObject(o)
				xmlObjects.Nodes = append(xmlObjects.Nodes, *xmlObject)
			}			
			xmlAreas.Areas = append(xmlAreas.Areas, xmlArea)
		}
		xmlChapter.Scenarios = append(xmlChapter.Scenarios,
			xmlScenario)
	}
	out, err := xml.Marshal(xmlGame)
	if err != nil {
		return "", fmt.Errorf("fail_to_marshal_game")
	}
	return string(out[:]), nil
}

// UnmarshalGameData parses specified XML data to game
// resource.
func UnmarshalGame(data io.Reader) (*res.GameData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlGame := new(SavedGameXML)
	err := xml.Unmarshal(doc, xmlGame)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v", err)
	}
	game := new(res.GameData)
	game.Name = xmlGame.Name
	// Chapter.
	game.Chapter.ID = xmlGame.Chapter.ID
	// Scenarios.
	for _, xmlScen := range xmlGame.Chapter.Scenarios {
		scen := res.ScenarioData{ID: xmlScen.ID}
		// Areas.
		for _, xmlArea := range xmlScen.AreasNode.Areas {	
			area := res.AreaData{ID: xmlArea.ID}
			// Characters.
			for _, xmlChar := range xmlArea.CharsNode.Characters {
				charData, err := buildCharacterData(&xmlChar)
				if err != nil {
					log.Err.Printf("xml:build_game:area:%s:character:%s:%v",
						xmlArea.ID, xmlChar.ID, err)
					continue
				}
				area.Chars = append(area.Chars, *charData)
			}
			// Objects.
			for _, xmlOb := range xmlArea.ObjectsNode.Nodes {
				obData, err := buildObjectData(&xmlOb)
				if err != nil {
					log.Err.Printf("xml:build_game:area:%s:object:%s:%v",
						xmlArea.ID, xmlOb.ID, err)
					continue
				}
				area.Objects = append(area.Objects, *obData)
			}
			scen.Areas = append(scen.Areas, area)
		}
		game.Chapter.Scenarios = append(game.Chapter.Scenarios, scen)
	}
	return game, nil
}
