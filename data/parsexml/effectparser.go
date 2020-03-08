/*
 * effectparser.go
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
)

// Struct for effects XML base.
type Effects struct {
	XMLName xml.Name `xml:"effects"`
	Nodes   []Effect `xml:"effect"`
}

// Struct for effect XML node.
type Effect struct {
	XMLName       xml.Name  `xml:"effect"`
	ID            string    `xml:"id,attr"`
	Duration      int64     `xml:"duration,attr"`
	ModifiersNode Modifiers `xml:"modifiers"`
}

// Struct for node with subeffects.
type Subeffects struct {
	XMLName xml.Name `xml:"subeffects"`
	Effects []string `xml:"effect,value"`
}

// UnmarshalEffects retrieves all effects data from
// specified XML data.
func UnmarshalEffects(data io.Reader) ([]*res.EffectData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(Effects)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal xml data: %v", err)
	}
	effects := make([]*res.EffectData, 0)
	for _, xmlEffect := range xmlBase.Nodes {
		effect := buildEffectData(xmlEffect)
		effects = append(effects, effect)
	}
	return effects, nil
}

// buildEffectData builds effect from XML data.
func buildEffectData(xmlEffect Effect) *res.EffectData {
	mods := buildModifiers(&xmlEffect.ModifiersNode)
	data := res.EffectData{
		ID:        xmlEffect.ID,
		Duration:  xmlEffect.Duration * 1000,
		Modifiers: mods,
	}
	return &data
}
