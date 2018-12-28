/*
 * gameparser.go
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

	"github.com/isangeles/flame/core/data/save"
	"github.com/isangeles/flame/core/module"
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
	XMLName   xml.Name           `xml:"area"`
	ID        string             `xml:id,attr"`
	Mainarea  bool               `xml:"mainarea,attr"`
	CharsNode SavedCharactersXML `xml:"characters"`
}

// Struct for saved characters XML node.
type SavedCharactersXML struct {
	XMLName    xml.Name       `xml:"characters"`
	Characters []CharacterXML `xml:"char"`
}

// MarshalGame parses specified game to XML
// savegame data.
func MarshalSaveGame(game *save.SaveGame) (string, error) {
	xmlGame := new(SavedGameXML)
	xmlGame.Name = game.Name
	chapter := game.Mod.Chapter()
	if chapter == nil {
		return "", fmt.Errorf("no game chapter set")
	}
	xmlChapter := &xmlGame.Chapter
	xmlChapter.ID = chapter.Conf().ID
	for _, s := range chapter.Scenarios() {
		xmlScenario := SavedScenarioXML{}
		xmlScenario.ID = s.ID()
		xmlAreas := &xmlScenario.AreasNode
		for _, a := range s.Areas() {
			xmlArea := SavedAreaXML{}
			xmlArea.ID = a.ID()
			if a.ID() == s.Mainarea().ID() {
				xmlArea.Mainarea = true
			}
			xmlChars := &xmlArea.CharsNode
			for _, c := range a.Characters() {
				xmlChar := xmlCharacter(c)
				charSerialID := module.FullSerial(xmlChar.ID,
					xmlChar.Serial)
				for _, pc := range game.Players {	
					if pc.SerialID() == charSerialID {
						xmlChar.PC = true
					}
				}
				xmlChars.Characters = append(xmlChars.Characters,
					*xmlChar)
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

// UnmarshalGame parses specified XML data to saved game
// XML struct.
func UnmarshalGame(data io.Reader) (SavedGameXML, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlGame := new(SavedGameXML)
	err := xml.Unmarshal(doc, xmlGame)
	if err != nil {
		return SavedGameXML{}, fmt.Errorf("fail_to_unmarshal_xml_data:%v",
			err)
	}
	return *xmlGame, nil
}
