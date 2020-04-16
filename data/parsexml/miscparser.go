/*
 * miscparser.go
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

package parsexml

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

// Struct for misc items base XML node.
type Miscs struct {
	XMLName xml.Name `xml:"miscs"`
	Items   []Misc   `xml:"misc"`
}

// Struct for misc XML node.
type Misc struct {
	XMLName  xml.Name `xml:"misc"`
	ID       string   `xml:"id,attr"`
	Value    int      `xml:"value,attr"`
	Level    int      `xml:"level,attr"`
	Loot     bool     `xml:"loot,attr"`
	Currency bool     `xml:"currency,attr"`
}

// UnmarshalMiscItems retrieves misc items data from specified XML data.
func UnmarshalMiscItems(data io.Reader) ([]res.MiscItemData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(Miscs)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal xml data: %v", err)
	}
	miscs := make([]res.MiscItemData, 0)
	for _, xmlMisc := range xmlBase.Items {
		misc, err := buildMiscData(xmlMisc)
		if err != nil {
			log.Err.Printf("xml: unmarshal misc item: unable to build data: %v", err)
			continue
		}
		miscs = append(miscs, misc)
	}
	return miscs, nil
}

// buildMiscData creates new misc data from specified XML data.
func buildMiscData(xmlMisc Misc) (res.MiscItemData, error) {
	m := res.MiscItemData{
		ID:       xmlMisc.ID,
		Value:    xmlMisc.Value,
		Level:    xmlMisc.Level,
		Loot:     xmlMisc.Loot,
		Currency: xmlMisc.Currency,
	}
	return m, nil
}
