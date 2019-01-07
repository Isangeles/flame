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
	
	"github.com/isangeles/flame/core/module/req"
	"github.com/isangeles/flame/core/module/object/item"
)

// Struct for weapons base XML doc.
type WeaponsBaseXML struct {
	XMLName xml.Name        `xml:"base"`
	Items   []WeaponNodeXML `xml:"item"`
}

// Struct for weapon XML node.
type WeaponNodeXML struct {
	XMLName xml.Name      `xml:"item"`
	ID      string        `xml:"id,attr"`
	Serial  string        `xml:"serial,attr"`
	Value   int           `xml:"value,attr"`
	Level   int           `xml:"level,attr"`
	Damage  DamageNodeXML `xml:"damage"`
	Reqs    ReqsNodeXML   `xml:"reqs"`
	Slots   string        `xml:"slots,attr"`
}

// UnmarshalWeaponsBase parses specified data to
// XML weapon nodes.
func UnmarshalWeaponsBase(data io.Reader) ([]WeaponNodeXML, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlWeaponsBase := new(WeaponsBaseXML)
	err := xml.Unmarshal(doc, xmlWeaponsBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v",
			err)
	}
	return xmlWeaponsBase.Items, nil
}

// xmlWeapon parses specified weapon item to XML
// weapon node.
func xmlWeapon(w *item.Weapon) *WeaponNodeXML {
	xmlWeapon := new(WeaponNodeXML)
	xmlWeapon.ID = w.ID()
	xmlWeapon.Serial = w.Serial()
	xmlWeapon.Value = w.Value()
	xmlWeapon.Level = w.Level()
	xmlWeapon.Damage.Min, xmlWeapon.Damage.Max = w.Damage()
	reqs := &xmlWeapon.Reqs
	for _, r := range w.EquipReqs() {
		switch r := r.(type) {
		case *req.LevelReq:
			req := xmlLevelReq(r)
			reqs.LevelReqs = append(reqs.LevelReqs, *req)
		}
	}
	for _, s := range w.Slots() {
		xmlWeapon.Slots = fmt.Sprintf("%s %s", xmlWeapon.Slots,
			s.ID())
	}
	return xmlWeapon
}
