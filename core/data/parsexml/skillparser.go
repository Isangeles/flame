/*
 * skillparser.go
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

package parsexml

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
)

// Struct for skills XML base.
type SkillsBaseXML struct {
	XMLName xml.Name   `xml:"base"`
	Skills  []SkillXML `xml"skill"`
}

// Struct for skill XML node.
type SkillXML struct {
	XMLName xml.Name        `xml:"skill"`
	ID      string          `xml:"id,attr"`
	Cast    int             `xml:"cast,attr"`
	Effects SkillEffectsXML `xml:"effects"`
	Reqs    ReqsXML          `xml:"reqs"`
}

// Struct for skill effects XML node.
type SkillEffectsXML struct {
	XMLName xml.Name         `xml:"effects"`
	Nodes   []SkillEffectXML `xml:"effect"`
}

// Struct for skill effect XML node.
type SkillEffectXML struct {
	XMLName xml.Name `xml:"effect"`
	ID      string   `xml:"id,attr"`
}

// UnmarshalSkillsBase parses specified data to XML skill
// nodes.
func UnmarshalSkillsBase(data io.Reader) ([]SkillXML, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := SkillsBaseXML{}
	err := xml.Unmarshal(doc, &xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v",
			err)
	}
	return xmlBase.Skills, nil
}
