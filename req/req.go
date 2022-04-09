/*
 * req.go
 *
 * Copyright 2018-2022 Dariusz Sikora <dev@isangeles.pl>
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
	for _, d := range data.TargetRangeReqs {
		trreq := NewTargetRange(d)
		reqs = append(reqs, trreq)
	}
	for _, d := range data.KillReqs {
		kreq := NewKill(d)
		reqs = append(reqs, kreq)
	}
	for _, d := range data.QuestReqs {
		qreq := NewQuest(d)
		reqs = append(reqs, qreq)
	}
	for _, d := range data.HealthPercentReqs {
		hreq := NewHealthPercent(d)
		reqs = append(reqs, hreq)
	}
	for _, d := range data.ManaReqs {
		mreq := NewMana(d)
		reqs = append(reqs, mreq)
	}
	for _, d := range data.CombatReqs {
		creq := NewCombat(d)
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
		case *TargetRange:
			d := r.Data()
			data.TargetRangeReqs = append(data.TargetRangeReqs, d)
		case *Kill:
			d := r.Data()
			data.KillReqs = append(data.KillReqs, d)
		case *Quest:
			d := r.Data()
			data.QuestReqs = append(data.QuestReqs, d)
		case *HealthPercent:
			d := r.Data()
			data.HealthPercentReqs = append(data.HealthPercentReqs, d)
		case *Mana:
			d := r.Data()
			data.ManaReqs = append(data.ManaReqs, d)
		case *Combat:
			d := r.Data()
			data.CombatReqs = append(data.CombatReqs, d)
		}
	}
	return
}
