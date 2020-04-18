/*
 * area.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for area data.
type AreaData struct {
	XMLName    xml.Name         `xml:"area" json:"area"`
	ID         string           `xml:"id,attr" json:"id"`
	Characters []AreaCharData   `xml:"characters>character" json:"character"`
	Objects    []AreaObjectData `xml:"objects>object" json:"objects"`
	Subareas   []AreaData       `xml:"subareas>area" json:"subareas"`
}

// Struct for area character data.
type AreaCharData struct {
	ID   string  `xml:"id,attr" json:"id"`
	PosX float64 `xml:"x,attr" json:"pos-x"`
	PosY float64 `xml:"y,attr" json:"pos-y"`
	AI   bool    `xml:"ai,attr" json:"ai"`
}

// Struct for area object data.
type AreaObjectData struct {
	ID   string  `xml:"id,attr" json:"id"`
	PosX float64 `xml:"x,attr" json:"pos-x"`
	PosY float64 `xml:"y,attr" json:"pos-y"`
}
