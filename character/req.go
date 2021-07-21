/*
 * req.go
 *
 * Copyright 2018-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/item"
	"github.com/isangeles/flame/objects"
	"github.com/isangeles/flame/req"
)

// ReqsMeet checks whether all specified requirements
// are meet by character.
func (c *Character) MeetReqs(reqs ...req.Requirement) bool {
	for _, r := range reqs {
		if !c.MeetReq(r) {
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
func (c *Character) MeetReq(r req.Requirement) bool {
	switch r := r.(type) {
	case *req.Level:
		return c.Level() >= r.MinLevel()
	case *req.Gender:
		return string(c.Gender()) == r.Gender()
	case *req.Flag:
		f := c.flags[r.FlagID()]
		if r.FlagOff() {
			return len(f.ID()) == 0
		}
		return len(f.ID()) > 0
	case *req.Item:
		count := 0
		for _, i := range c.Inventory().Items() {
			if i.ID() == r.ItemID() {
				count++
			}
		}
		return count >= r.ItemAmount()
	case *req.Currency:
		// TODO: currency check.
		val := 0
		for _, it := range c.Inventory().Items() {
			misc, ok := it.(*item.Misc)
			if !ok {
				continue
			}
			if !misc.Currency() {
				continue
			}
			val += misc.Value()
		}
		return val >= r.Amount()
	case *req.TargetRange:
		if len(c.Targets()) < 1 {
			return true
		}
		tar := c.Targets()[0]
		return objects.Range(c, tar) <= r.MinRange()
	case *req.Kill:
		amount := 0
		for _, k := range c.Kills() {
			if k.ID == r.ID() {
				amount ++
			}
		}
		return amount >= r.Amount()
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
	case *req.Item:
		for i := 0; i < r.ItemAmount(); i++ {
			for _, i := range c.Inventory().Items() {
				if i.ID() != r.ItemID() {
					continue
				}
				c.Inventory().RemoveItem(i)
			}
		}
	case *req.Currency:
		// TODO: charge currency items.
	}
}
