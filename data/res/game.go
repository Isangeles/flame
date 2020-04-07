/*
 * game.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

package res

import (
	"encoding/xml"
)

// Struct for game data.
type GameData struct {
	XMLName      xml.Name         `xml:"game" json:"-"`
	SavedChapter SavedChapterData `xml:"chapter" json:"chapter"`
}

// Struct for game chapter
// data.
type SavedChapterData struct {
	ID    string          `xml:"id,attr" json:"id"`
	Areas []SavedAreaData `xml:"area" json:"areas"`
}

// Struct for scenario area
// data.
type SavedAreaData struct {
	ID       string          `xml:"id,attr" json:"id"`
	Chars    []CharacterData `xml:"characters>char" json:"chars"`
	Objects  []ObjectData    `xml:"objects>object" json:"objects"`
	Subareas []SavedAreaData `xml:"subareas>area" json:"areas"`
}
