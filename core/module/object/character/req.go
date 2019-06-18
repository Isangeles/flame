/*
 * req.go
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

package character

import (
	"github.com/isangeles/flame/core/module/req"
)

// ReqsMeet checks whether all specified requirements
// are meet by character.
func (char *Character) MeetReqs(reqs ...req.Requirement) bool {
	for _, r := range reqs {
		if !char.MeetReq(r) {
			return false
		}
	}
	return true
}

// ChargeReqs takes from character all things that makes
// him to meet specified requirements.
func (c *Character) ChargeReqs(reqs ...req.Requirement) {
	for _, r := range reqs {
		c.ChargeReq(r)
	}
}

// ReqMeet checks whether character meets
// specified requirement.
func (char *Character) MeetReq(r req.Requirement) bool {
	switch r := r.(type) {
	case *req.LevelReq:
		return char.Level() >= r.MinLevel()
	case *req.GenderReq:
		return int(char.Gender()) == r.Type()
	case *req.FlagReq:
		f := char.flags[r.FlagID()]
		if r.FlagOff() {
			return len(f.ID()) == 0
		}
		return len(f.ID()) > 0
	case *req.ItemReq:
		count := 0
		for _, i := range char.Inventory().Items() {
			if i.ID() == r.ItemID() {
				count++
			}
		}
		return count >= r.ItemAmount()
	default:
		return true
	}
}

// ChargeReq takes from character all things that makes
// him to meet specified requirement. Does nothing if
// character don't meet requirement or requirement is
// not 'chargeable'.
func (c *Character) ChargeReq(r req.Requirement) {
	if !c.MeetReq(r) {
		return
	}
	switch r := r.(type) {
	case *req.ItemReq:
		for i := 0; i < r.ItemAmount(); i ++ {
			for _, i := range c.Inventory().Items() {
				if i.ID() != r.ItemID() {
					continue
				}
				c.Inventory().RemoveItem(i)
			}
		}
	}
}
