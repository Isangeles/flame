/*
 * item_test.go
 *
 * Copyright 2023 Dariusz Sikora <ds@isangeles.dev>
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

package res

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

// TestInventoryDataJson tests inventory data JSON mappings.
func TestInventoryDataJson(t *testing.T) {
	data, err := testData("inventory.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	inv := new(InventoryData)
	err = json.Unmarshal(data, inv)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	if len(inv.Items) != 2 {
		t.Errorf("Invalid amount of inventory items: %d != 2",
			len(inv.Items))
	}
}

// TestInventoryDataXml tests inventory data XML mappings.
func TestInventoryDataXml(t *testing.T) {
	data, err := testData("inventory.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	inv := new(InventoryData)
	err = xml.Unmarshal(data, inv)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	if len(inv.Items) != 2 {
		t.Errorf("Invalid amount of inventory items: %d != 2",
			len(inv.Items))
	}
}

// TestInventoryItemDataJson tests inventory item data JSON mappings.
func TestInventoryItemDataJson(t *testing.T) {
	data, err := testData("inventoryitem.json")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	item := new(InventoryItemData)
	err = json.Unmarshal(data, item)
	if err != nil {
		t.Fatalf("Unable to unmarshal JSON data: %v", err)
	}
	if item.ID != "item1" {
		t.Errorf("Invalid ID: %s != item1", item.ID)
	}
	if item.Serial != "0" {
		t.Errorf("Invalid serial: %s != 0", item.Serial)
	}
	if !item.Loot {
		t.Errorf("Loot flag false")
	}
	if !item.Trade {
		t.Errorf("Trade flag false")
	}
	if item.TradeValue != 10 {
		t.Errorf("Invalid trade value: %d != 10", item.TradeValue)
	}
	if item.Random != 0.5 {
		t.Errorf("Invalid random value: %f != 0.5", item.Random)
	}
}

// TestInventoryItemDataXml tests inventory item data XML mappings.
func TestInventoryItemDataXml(t *testing.T) {
	data, err := testData("inventoryitem.xml")
	if err != nil {
		t.Fatalf("Unable to retrieve test data: %v", err)
	}
	item := new(InventoryItemData)
	err = xml.Unmarshal(data, item)
	if err != nil {
		t.Fatalf("Unable to unmarshal XML data: %v", err)
	}
	if item.ID != "item1" {
		t.Errorf("Invalid ID: %s != item1", item.ID)
	}
	if item.Serial != "0" {
		t.Errorf("Invalid serial: %s != 0", item.Serial)
	}
	if !item.Loot {
		t.Errorf("Loot flag false")
	}
	if !item.Trade {
		t.Errorf("Trade flag false")
	}
	if item.TradeValue != 10 {
		t.Errorf("Invalid trade value: %d != 10", item.TradeValue)
	}
	if item.Random != 0.5 {
		t.Errorf("Invalid random value: %f != 0.5", item.Random)
	}
}
