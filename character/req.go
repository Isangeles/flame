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

package character

import (
	"github.com/isangeles/flame/flag"
	"github.com/isangeles/flame/item"
	"github.com/isangeles/flame/objects"
	"github.com/isangeles/flame/req"
)

// ReqsMeet checks whether all specified requirements
// are meet by character.
func (c *Character) MeetReqs(reqs ...req.Requirement) bool {
	for _, r := range reqs {
		if !c.meetReq(r) {
			return false
		}
	}
	return true
}

// ChargeReqs takes from character all things that make
// the character meet specified requirements.
// All requirements that are not 'chargeable' will be
// ignored.
func (c *Character) ChargeReqs(reqs ...req.Requirement) {
	for _, r := range reqs {
		if r := r.(req.Chargeable); r.Charge() {
			c.chargeReq(r)
		}
	}
}

// meetReq checks whether character meets
// specified requirement.
func (c *Character) meetReq(r req.Requirement) bool {
	switch r := r.(type) {
	case *req.Level:
		return c.Level() >= r.MinLevel()
	case *req.Gender:
		return string(c.Gender()) == r.Gender()
	case *req.Flag:
		f := flag.Flag(r.FlagID())
		if r.FlagOff() {
			return !c.HasFlag(f)
		}
		return c.HasFlag(f)
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
				amount++
			}
		}
		return amount >= r.Amount()
	case *req.Quest:
		for _, q := range c.Journal().Quests() {
			if q.ID() == r.QuestID() && q.Completed() == r.QuestCompleted() {
				return true
			}
		}
		return false
	case *req.Health:
		if r.Less() {
			return c.Health() < r.Value()
		}
		return c.Health() >= r.Value()
	case *req.HealthPercent:
		hp := (c.Health() * 100) / c.MaxHealth()
		if r.Less() {
			return hp < r.Value()
		}
		return hp >= r.Value()
	case *req.Mana:
		if r.Less() {
			return c.Mana() < r.Value()
		}
		return c.Mana() >= r.Value()
	case *req.ManaPercent:
		mp := (c.Mana() * 100) / c.MaxMana()
		if r.Less() {
			return mp < r.Value()
		}
		return mp >= r.Value()
	case *req.Combat:
		return c.Fighting() == r.Combat()
	default:
		return true
	}
}

// chargeReq takes from character all things that make
// this character meet specified 'chargeable' requirement.
// Does nothing if character don't meet requirement.
func (c *Character) chargeReq(r req.Chargeable) {
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
