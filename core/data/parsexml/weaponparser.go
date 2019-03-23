/*
 * weaponparser.go
 *
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for weapons base XML doc.
type WeaponsBaseXML struct {
	XMLName xml.Name    `xml:"base"`
	Items   []WeaponXML `xml:"item"`
}

// Struct for weapon XML node.
type WeaponXML struct {
	XMLName    xml.Name      `xml:"item"`
	ID         string        `xml:"id,attr"`
	Serial     string        `xml:"serial,attr"`
	Value      int           `xml:"value,attr"`
	Level      int           `xml:"level,attr"`
	Damage     DamageXML     `xml:"damage"`
	Reqs       ReqsXML       `xml:"reqs"`
	Slots      string        `xml:"slots,attr"`
}

// UnmarshalWeaponsBase retrieves weapons data from specified
// XML data.
func UnmarshalWeaponsBase(data io.Reader) ([]*res.WeaponData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(WeaponsBaseXML)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v", err)
	}
	weapons := make([]*res.WeaponData, 0)
	for _, xmlWeapon := range xmlBase.Items {
		weapon, err := buildWeaponData(xmlWeapon)
		if err != nil {
			log.Err.Printf("xml:unmarshal_weapon:build_data_fail:%v", err)
			continue
		}
		weapons = append(weapons, weapon)
	}
	return weapons, nil
}

// buildXMLWeapon creates new weapon data from specified XML data.
func buildWeaponData(xmlWeapon WeaponXML) (*res.WeaponData, error) {
	reqs := buildReqs(&xmlWeapon.Reqs)
	slots, err := UnmarshalItemSlots(xmlWeapon.Slots)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_slot_types:%v", err)
	}
	slotsID := make([]int, 0)
	for _, s := range slots {
		slotsID = append(slotsID, int(s))
	}
	dmgType, err := UnmarshalHitType(xmlWeapon.Damage.Type)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_damage_type:%v", err)
	}
	hitEffects := make([]res.EffectData, 0)
	for _, xmlEffect := range xmlWeapon.Damage.Effects.Nodes {
		eff := res.Effect(xmlEffect.ID)
		if eff == nil {
			log.Err.Printf("xml:build_weapon:hit_effect_not_found:%s",
				xmlEffect.ID)
			continue
		}
		hitEffects = append(hitEffects, *eff)
	}
	w := res.WeaponData{
		ID:         xmlWeapon.ID,
		Value:      xmlWeapon.Value,
		Level:      xmlWeapon.Level,
		DMGMin:     xmlWeapon.Damage.Min,
		DMGMax:     xmlWeapon.Damage.Max,
		DMGType:    int(dmgType),
		DMGEffects: hitEffects,
		EQReqs:     reqs,
		Slots:      slotsID,
	}
	return &w, nil
}
