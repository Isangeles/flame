/*
 * miscparser.go
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

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/log"
)

// Struct for misc items base XML node.
type MiscsBaseXML struct {
	XMLName xml.Name  `xml:"base"`
	Items   []MiscXML `xml:"item"`
}

// Struct for misc XML node.
type MiscXML struct {
	XMLName xml.Name `xml:"item"`
	ID      string   `xml:"id,attr"`
	Value   int      `xml:"value,attr"`
	Level   int      `xml:"level,attr"`
	Loot    bool     `xml:"loot,attr"`
}

// UnmarshalMiscItemsBase retrieves misc items data from specified XML data.
func UnmarshalMiscItemsBase(data io.Reader) ([]*res.MiscItemData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(MiscsBaseXML)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v", err)
	}
	miscs := make([]*res.MiscItemData, 0)
	for _, xmlMisc := range xmlBase.Items {
		misc, err := buildMiscData(xmlMisc)
		if err != nil {
			log.Err.Printf("xml:unmarshal_misc_item:build_data_fail:%v", err)
			continue
		}
		miscs = append(miscs, misc)
	}
	return miscs, nil
}

// buildMiscData creates new misc data from specified XML data.
func buildMiscData(xmlMisc MiscXML) (*res.MiscItemData, error) {
	m := res.MiscItemData{
		ID:    xmlMisc.ID,
		Value: xmlMisc.Value,
		Level: xmlMisc.Level,
		Loot:  xmlMisc.Loot,
	}
	return &m, nil
}
