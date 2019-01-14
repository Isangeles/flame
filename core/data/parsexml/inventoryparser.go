/*
 * inventoryparser.go
 * 
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/core/module/object/item"
)

// Struct for inventory XML node.
type InventoryNodeXML struct {
	XMLName  xml.Name               `xml:"inventory"`
	Capacity int                    `xml:"cap,attr"`
	Items    []InventoryItemNodeXML `xml:"item"`
}

// Struct for inventory item XML node.
type InventoryItemNodeXML struct {
	XMLName xml.Name `xml:"item"`
	ID      string   `xml:"id,attr"`
	Serial  string   `xml:"serial,attr"`
}

// xmlInventory parses specified inventory to XML
// inventory node.
func xmlInventory(inv *item.Inventory) *InventoryNodeXML {
	xmlInv := new(InventoryNodeXML)
	xmlInv.Capacity = inv.Capacity()
	for _, i := range inv.Items() {
		xmlInvItem := InventoryItemNodeXML{
			ID:     i.ID(),
			Serial: i.Serial(),
		}
		xmlInv.Items = append(xmlInv.Items, xmlInvItem)
	}
	return xmlInv
}
