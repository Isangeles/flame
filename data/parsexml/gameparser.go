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
