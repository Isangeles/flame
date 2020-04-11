/*
 * objectparser.go
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
 *w
 */

package parsexml

import (
	"encoding/xml"
)

// Struct for XML objects base node.
type Objects struct {
	XMLName xml.Name `xml:"objects"`
	Nodes   []Object `xml:"object"`
}

// Struct for XML object node.
type Object struct {
	XMLName   xml.Name       `xml:"object"`
	ID        string         `xml:"id,attr"`
	Serial    string         `xml:"serial,attr"`
	HP        int            `xml:"hp,attr"`
	MaxHP     int            `xml:"max-hp,attr"`
	PosX      float64        `xml:"position-x,attr"`
	PosY      float64        `xml:"position-y,attr"`
	Action    ObjectAction   `xml:"action"`
	Inventory Inventory      `xml:"inventory"`
	Effects   []ObjectEffect `xml:"effects>effect"`
}

// Strcut for object effects XML node.
type ObjectEffect struct {
	XMLName xml.Name           `xml:"effect"`
	ID      string             `xml:"id,attr"`
	Serial  string             `xml:"serial,attr"`
	Time    int64              `xml:"time,attr"`
	Source  ObjectEffectSource `xml:"source"`
}

// Struct for object effect source XML node.
type ObjectEffectSource struct {
	XMLName xml.Name `xml:"source"`
	ID      string   `xml:"id,attr"`
	Serial  string   `xml:"serial,attr"`
}

// Struct for object skills XML node.
type ObjectSkills struct {
	XMLName xml.Name      `xml:"skills"`
	Nodes   []ObjectSkill `xml:"skill"`
}

// Struct for object skill XML node.
type ObjectSkill struct {
	XMLName  xml.Name `xml:"skill"`
	ID       string   `xml:"id,attr"`
	Serial   string   `xml:"serial,attr"`
	Cooldown int64    `xml:"cooldown,attr"`
}

// Struct for object dialogs XML node.
type ObjectDialogs struct {
	XMLName xml.Name       `xml:"dialogs"`
	Nodes   []ObjectDialog `xml:"dialog"`
}

// Struct for object dialog XML node.
type ObjectDialog struct {
	XMLName xml.Name `xml:"dialog"`
	ID      string   `xml:"id,attr"`
}

// Struct for object recipe XML node.
type ObjectRecipe struct {
	XMLName xml.Name `xml:"recipe"`
	ID      string   `xml:"id,attr"`
}

// Struct for object action XML node.
type ObjectAction struct {
	XMLName  xml.Name  `xml:"action"`
	SelfMods Modifiers `xml:"self>modifiers"`
	UserMods Modifiers `xml:"user>modifiers"`
}
