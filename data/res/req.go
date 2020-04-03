/*
 * req.go
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

package res

// Struct for reqs data.
type ReqsData struct {
	LevelReqs    []LevelReqData
	GenderReqs   []GenderReqData
	FlagReqs     []FlagReqData
	ItemReqs     []ItemReqData
	CurrencyReqs []CurrencyReqData
}

// Struct for level requirement data.
type LevelReqData struct {
	Min, Max int
}

// Struct for gender requirement data.
type GenderReqData struct {
	Gender string
}

// Struct for flag requirement data.
type FlagReqData struct {
	ID  string
	Off bool
}

// Struct for item requirement data.
type ItemReqData struct {
	ID     string
	Amount int
	Charge bool
}

// Struct for currency requirement data.
type CurrencyReqData struct {
	Amount int
	Charge bool
}
