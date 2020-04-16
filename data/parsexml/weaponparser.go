/*
 * weaponparser.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for weapons base XML node.
type Weapons struct {
	XMLName xml.Name `xml:"weapons"`
	Items   []Weapon `xml:"weapon"`
}

// Struct for weapon XML node.
type Weapon struct {
	XMLName xml.Name   `xml:"weapon"`
	ID      string     `xml:"id,attr"`
	Serial  string     `xml:"serial,attr"`
	Value   int        `xml:"value,attr"`
	Level   int        `xml:"level,attr"`
	Loot    bool       `xml:"loot,attr"`
	Damage  Damage     `xml:"damage"`
	Reqs    Reqs       `xml:"reqs"`
	Slots   []ItemSlot `xml:"slots>slot"`
}

// UnmarshalWeaponsBase retrieves weapons data from specified XML data.
func UnmarshalWeapons(data io.Reader) ([]res.WeaponData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(Weapons)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal xml data: %v", err)
	}
	weapons := make([]res.WeaponData, 0)
	for _, xmlWeapon := range xmlBase.Items {
		weapon, err := buildWeaponData(xmlWeapon)
		if err != nil {
			log.Err.Printf("xml: unmarshal weapon: build data unable: %v", err)
			continue
		}
		weapons = append(weapons, weapon)
	}
	return weapons, nil
}

// buildXMLWeapon creates new weapon data from specified XML data.
func buildWeaponData(xmlWeapon Weapon) (res.WeaponData, error) {
	reqs := buildReqs(&xmlWeapon.Reqs)
	slotsID := make([]string, 0)
	for _, s := range xmlWeapon.Slots {
		slotsID = append(slotsID, s.ID)
	}
	hitEffects := make([]res.EffectData, 0)
	for _, xmlEffect := range xmlWeapon.Damage.Effects {
		eff := res.Effect(xmlEffect.ID)
		if eff == nil {
			log.Err.Printf("xml: build weapon: hit effect not found: %s",
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
		DMGType:    xmlWeapon.Damage.Type,
		DMGEffects: hitEffects,
		EQReqs:     reqs,
		Slots:      slotsID,
		Loot:       xmlWeapon.Loot,
	}
	return w, nil
}
