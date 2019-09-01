/*
 * armorparser.go
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

// Struct for armors node.
type Armors struct {
	XMLName xml.Name `xml:"armors"`
	Armors  []Armor  `xml:"armor"`
}

// Struct for armor node.
type Armor struct {
	XMLName   xml.Name      `xml:"armor"`
	ID        string        `xml:"id,attr"`
	Serial    string        `xml:"seral,attr"`
	Value     int           `xml:"value,attr"`
	Level     int           `xml:"level,attr"`
	Armor     int           `xml:"armor,attr"`
	Slots     string        `xml:"slots,attr"`
	Loot      bool          `xml:"loot,attr"`
	Reqs      Reqs          `xml:"eq>reqs"`
	EQEffects ObjectEffects `xml:"eq>effects"`
}

// UnmarshalArmors retrieves armor data from specified XML data.
func UnmarshalArmors(data io.Reader) ([]*res.ArmorData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(Armors)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal xml data: %v", err)
	}
	armors := make([]*res.ArmorData, 0)
	for _, xmlArmor := range xmlBase.Armors {
		armorData, err := buildArmorData(xmlArmor)
		if err != nil {
			log.Err.Printf("xml: unmarshal armor: build data fail: %v", err)
			continue
		}
		armors = append(armors, armorData)
	}
	return armors, nil
}

// buildArmorData creates armor resource from specified armor node.
func buildArmorData(xmlArmor Armor) (*res.ArmorData, error) {
	reqs := buildReqs(&xmlArmor.Reqs)
	slots, err := UnmarshalItemSlots(xmlArmor.Slots)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal slot types: %v", err)
	}
	slotsID := make([]int, 0)
	for _, s := range slots {
		slotsID = append(slotsID, int(s))
	}
	eqEffects := make([]res.EffectData, 0)
	for _, xmlEffect := range xmlArmor.EQEffects.Nodes {
		eff := res.Effect(xmlEffect.ID)
		if eff == nil {
			log.Err.Printf("xml: build armor: hit effect not found: %s",
				xmlEffect.ID)
			continue
		}
		eqEffects = append(eqEffects, *eff)
	}
	ad := res.ArmorData{
		ID:        xmlArmor.ID,
		Value:     xmlArmor.Value,
		Level:     xmlArmor.Level,
		Armor:     xmlArmor.Armor,
		EQEffects: eqEffects,
		EQReqs:    reqs,
		Slots:     slotsID,
		Loot:      xmlArmor.Loot,
	}
	return &ad, nil
}
