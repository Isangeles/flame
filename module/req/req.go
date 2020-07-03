/*
 * req.go
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

// Package for requirements(e.g. weapon
// equip requirements).
package req

import (
	"github.com/isangeles/flame/data/res"
)

// Interface for requirements.
type Requirement interface {
	Meet() bool
	SetMeet(meet bool)
}

// Interface for requirements targets.
type RequirementsTarget interface {
	MeetReqs(reqs ...Requirement) bool
	ChargeReqs(reqs ...Requirement)
}

// NewRequirements creates new requirements from
// specified data.
func NewRequirements(data res.ReqsData) (reqs []Requirement) {
	for _, d := range data.LevelReqs {
		lreq := NewLevel(d)
		reqs = append(reqs, lreq)
	}
	for _, d := range data.GenderReqs {
		greq := NewGender(d)
		reqs = append(reqs, greq)
	}
	for _, d := range data.FlagReqs {
		freq := NewFlag(d)
		reqs = append(reqs, freq)
	}
	for _, d := range data.ItemReqs {
		ireq := NewItem(d)
		reqs = append(reqs, ireq)
	}
	for _, d := range data.CurrencyReqs {
		creq := NewCurrency(d)
		reqs = append(reqs, creq)
	}
	return
}

// RequirementsData creates data resource for requirements.
func RequirementsData(reqs ...Requirement) (data res.ReqsData) {
	for _, r := range reqs {
		switch r := r.(type) {
		case *Level:
			d := r.Data()
			data.LevelReqs = append(data.LevelReqs, d)
		case *Gender:
			d := r.Data()
			data.GenderReqs = append(data.GenderReqs, d)
		case *Flag:
			d := r.Data()
			data.FlagReqs = append(data.FlagReqs, d)
		case *Item:
			d := r.Data()
			data.ItemReqs = append(data.ItemReqs, d)
		case *Currency:
			d := r.Data()
			data.CurrencyReqs = append(data.CurrencyReqs, d)
		}
	}
	return
}
