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

// UnmarshalWeaponsBase parses specified data to
// XML weapon nodes.
func UnmarshalWeaponsBase(data io.Reader) ([]WeaponXML, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlWeaponsBase := new(WeaponsBaseXML)
	err := xml.Unmarshal(doc, xmlWeaponsBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v",
			err)
	}
	return xmlWeaponsBase.Items, nil
}
