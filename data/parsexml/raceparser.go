/*
 * raceparser.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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
)

// Struct for XML races node.
type Races struct {
	XMLName xml.Name `xml:"races"`
	Races   []Race   `xml:"race"`
}

// Struct for XML race node.
type Race struct {
	XMLName  xml.Name `xml:"race"`
	ID       string   `xml:"id,attr"`
	Playable bool     `xml:"playable,attr"`
}

// UnmarshalRaces retrieves races from specified XML data.
func UnmarshalRaces(data io.Reader) ([]res.RaceData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlRaces := new(Races)
	err := xml.Unmarshal(doc, xmlRaces)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal xml data: %v", err)
	}
	races := make([]res.RaceData, 0)
	for _, xmlRace := range xmlRaces.Races {
		rd := buildRaceData(xmlRace)
		races = append(races, rd)
	}
	return races, nil
}

// buildRaceData creates race resource data from specified XML data.
func buildRaceData(xmlRace Race) res.RaceData {
	rd := res.RaceData{
		ID:       xmlRace.ID,
		Playable: xmlRace.Playable,
	}
	return rd
}
