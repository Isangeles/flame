/*
 * currency.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for currency requirement.
type Currency struct {
	amount int
	charge bool
	meet   bool
}

// NewCurrency creates new currency requirement.
func NewCurrency(data res.CurrencyReqData) *Currency {
	cr := new(Currency)
	cr.amount = data.Amount
	cr.charge = data.Charge
	return cr
}

// Amount returns amount of required currency.
func (cr *Currency) Amount() int {
	return cr.amount
}

// Charge checks is required amount should
// be taken from requirement target.
func (cr *Currency) Charge() bool {
	return cr.charge
}

// Meet checks if requirement is set as met.
func (cr *Currency) Meet() bool {
	return cr.meet
}

// SetMeet sets requirements as meet/not meet.
func (cr *Currency) SetMeet(meet bool) {
	cr.meet = meet
}

// Data returns data resource for requirement.
func (cr *Currency) Data() res.CurrencyReqData {
	data := res.CurrencyReqData{
		Amount: cr.Amount(),
		Charge: cr.Charge(),
	}
	return data
}
