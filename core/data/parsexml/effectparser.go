/*
 * effectparser.go
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
	"github.com/isangeles/flame/core/module/object/effect"
)

// Struct for effects XML base.
type EffectsBaseXML struct {
	XMLName xml.Name    `xml:"base"`
	Nodes   []EffectXML `xml:"effect"`
}

// Struct for effect XML node.
type EffectXML struct {
	XMLName       xml.Name         `xml:"effect"`
	ID            string           `xml:"id,attr"`
	Duration      int64            `xml:"duration,attr"`
	ModifiersNode ModifiersNodeXML `xml:"modifiers"`
}

// Struct for node with subeffects.
type SubeffectsXML struct {
	XMLName xml.Name `xml:"subeffects"`
	Effects []string `xml:"effect,value"`
}

// UnmarshalEffectsBase retrieves all effects data from
// specified XML data.
func UnmarshalEffectsBase(data io.Reader) ([]*res.EffectData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(EffectsBaseXML)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v",
			err)
	}
	effects := make([]*res.EffectData, 0) 
	for _, xmlEffect := range xmlBase.Nodes {
		effect := buildEffectData(xmlEffect)
		effects = append(effects, effect)
	}
	return effects, nil
}

// buildEffectData builds effect from XML data.
func buildEffectData(xmlEffect EffectXML) *res.EffectData {
	mods := buildModifiers(&xmlEffect.ModifiersNode)
	data := res.EffectData{
		ID: xmlEffect.ID,
		Duration: xmlEffect.Duration * 1000,
	}
	for _, m := range mods {
		switch mod := m.(type) {
		case effect.HealthMod:
			mData := res.HealthModData{mod.Min, mod.Max}
			data.HealthMods = append(data.HealthMods, mData)
		case effect.HitMod:
			mData := res.HitModData{}
			data.HitMods = append(data.HitMods, mData)
		}
	}
	return &data
}
