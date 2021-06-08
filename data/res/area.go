/*
 * area.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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
	XMLName    xml.Name         `xml:"area" json:"-"`
	ID         string           `xml:"id,attr" json:"id"`
	Time       string           `xml:"time,attr" json:"time"`
	Weather    string           `xml:"weather,attr" json:"weather"`
	Characters []AreaCharData   `xml:"characters>character" json:"character"`
	Objects    []AreaObjectData `xml:"objects>object" json:"objects"`
	Subareas   []AreaData       `xml:"subareas>area" json:"subareas"`
}

// Struct for area character data.
type AreaCharData struct {
	ID     string     `xml:"id,attr" json:"id"`
	Serial string     `xml:"serial,attr" json:"serial"`
	PosX   float64    `xml:"x,attr" json:"pos-x"`
	PosY   float64    `xml:"y,attr" json:"pos-y"`
	DestX  float64    `xml:"dest-x,attr" json:"dest-pos-x"`
	DestY  float64    `xml:"dest-y,attr" json:"dest-pos-y"`
	DefX   float64    `xml:"def-x,attr" json:"def-pos-x"`
	DefY   float64    `xml:"def-y,attr" json:"def-pos-y"`
	AI     bool       `xml:"ai,attr" json:"ai"`
	Flags  []FlagData `xml:"flags>flag" json:"flags"`
}

// Struct for area object data.
type AreaObjectData struct {
	ID     string  `xml:"id,attr" json:"id"`
	Serial string  `xml:"serial,attr" json:"serial"`
	PosX   float64 `xml:"x,attr" json:"pos-x"`
	PosY   float64 `xml:"y,attr" json:"pos-y"`
}
