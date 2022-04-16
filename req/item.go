/*
 * item.go
 *
 * Copyright 2019-2022 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/data/res"
)

// Struct for item requirement.
type Item struct {
	itemID string
	amount int
	charge bool
	meet   bool
}

// NewItem creates new item requirement.
func NewItem(data res.ItemReqData) *Item {
	ir := new(Item)
	ir.itemID = data.ID
	ir.amount = data.Amount
	ir.charge = data.Charge
	return ir
}

// ItemID returns ID of required item.
func (ir *Item) ItemID() string {
	return ir.itemID
}

// ItemAmount returns amount of required items.
func (ir *Item) ItemAmount() int {
	return ir.amount
}

// Charge checks if required items should be
// taken from the object after requirement
// check.
func (ir *Item) Charge() bool {
	return ir.charge
}

// Meet checks wheter requirement is set as meet.
func (ir *Item) Meet() bool {
	return ir.meet
}

// SetMeet sets requirement as meet/not meet.
func (ir *Item) SetMeet(meet bool) {
	ir.meet = meet
}

// Data returns data resource for requirement.
func (ir *Item) Data() res.ItemReqData {
	data := res.ItemReqData{
		ID:     ir.ItemID(),
		Amount: ir.ItemAmount(),
		Charge: ir.Charge(),
	}
	return data
}
