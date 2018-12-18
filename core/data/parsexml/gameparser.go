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

	"github.com/isangeles/flame/core"
)

// Struct for saved game XML node.
type SavedGameXML struct {
	XMLName xml.Name        `xml:"game"`
	Chapter SavedChapterXML `xml:"chapter"`
}

// Struct for saved chapter XML node.
type SavedChapterXML struct {
	XMLName   xml.Name           `xml:"chapter"`
	Scenarios []SavedScenarioXML `xml:"scenario"`
}

// Struct for saved scenario XML node.
type SavedScenarioXML struct {
	XMLName   xml.Name      `xml:"scenario"`
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
	CharsNode SavedCharactersXML `xml:"characters"`
}

// Struct for saved characters XML node.
type SavedCharactersXML struct {
	XMLName    xml.Name       `xml:"characters"`
	Characters []CharacterXML `xml:"char"`
}

// MarshalGame parses specified game to XML
// savegame data.
func MarshalGame(game *core.Game) (string, error) {
	xmlGame := new(SavedGameXML)
	xmlChapter := &xmlGame.Chapter
	for _, s := range game.Module().Chapter().Scenarios() {
		xmlScenario := SavedScenarioXML{}
		xmlAreas := &xmlScenario.AreasNode
		for _, a := range s.Areas() {
			xmlArea := SavedAreaXML{}
			xmlChars := &xmlArea.CharsNode
			for _, c := range a.Characters() {
				xmlChar := xmlCharacter(c)
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
