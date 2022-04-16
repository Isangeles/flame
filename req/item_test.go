/*
 * item_test.go
 *
 * Copyright 2022 Dariusz Sikora <dev@isangeles.pl>
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

package req

import (
	"testing"

	"github.com/isangeles/flame/data/res"
)

var itemReqData = res.ItemReqData{"item1", 1, true}

// TestNewItem tests creating new item requirement.
func TestNewItem(t *testing.T) {
	req := NewItem(itemReqData)
	if req.ItemID() != "item1" {
		t.Errorf("Invalid item ID: %s != ", req.ItemID())
	}
	if req.ItemAmount() != 1 {
		t.Errorf("Invalid item amount: %d != 1", req.ItemAmount())
	}
	if !req.Charge() {
		t.Errorf("Invalid charge value: %v != true", req.Charge())
	}
}

// TestItemData tests data function of item
// requirement struct.
func TestItemData(t *testing.T) {
	req := NewItem(itemReqData)
	reqData := req.Data()
	if reqData != itemReqData {
		t.Errorf("Invalid requirement data: %v != %v", reqData, itemReqData)
	}
}
