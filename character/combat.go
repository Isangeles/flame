/*
 * combat.go
 *
 * Copyright 2019-2022 Dariusz Sikora <ds@isangeles.dev>
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
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/item"
	"github.com/isangeles/flame/objects"
	"github.com/isangeles/flame/serial"
)

// AddKill adds specified kill on character kill list.
func (c *Character) AddKill(kill res.KillData) {
	c.kills = append(c.kills, kill)
	c.SetExperience(c.Experience() + kill.Experience)
}

// Kills returns all character kill records.
func (c *Character) Kills() []res.KillData {
	return c.kills
}

// Damage retruns min and max damage value,
// including weapons, active effects, etc.
func (c *Character) Damage() (int, int) {
	min, max := c.Attributes().Damage()
	var rightHandItem item.Equiper
	for _, s := range c.Equipment().Slots() {
		if s.Type() != item.Hand {
			continue
		}
		rightHandItem = s.Item()
		break
	}
	if rightHandItem != nil {
		if w, ok := rightHandItem.(*item.Weapon); ok {
			dmgMin, dmgMax := w.Damage()
			min += dmgMin
			max += dmgMax
		}
	}
	return min, max
}

// DamgaeType returns type of damage caused by
// character.
func (c *Character) DamageType() objects.Element {
	var rightHandItem item.Equiper
	for _, s := range c.Equipment().Slots() {
		if s.Type() != item.Hand {
			continue
		}
		rightHandItem = s.Item()
		break
	}
	if rightHandItem != nil {
		if w, ok := rightHandItem.(*item.Weapon); ok {
			w.DamageType()
		}
	}
	return objects.ElementNone
}

// HitEffects returns character hit effects.
func (c *Character) HitEffects() []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	var rightHandItem item.Equiper
	for _, s := range c.Equipment().Slots() {
		if s.Type() != item.Hand {
			continue
		}
		rightHandItem = s.Item()
		break
	}
	if rightHandItem != nil {
		if w, ok := rightHandItem.(*item.Weapon); ok {
			dmgEffects := c.buildEffects(w.DamageEffects()...)
			effects = append(effects, dmgEffects...)
		}
	}
	return effects
}

// HitModifiers returns character hit modifiers.
func (c *Character) HitModifiers() []effect.Modifier {
	dmgMin, dmgMax := c.Damage()
	healthMod := res.HealthModData{-dmgMin, -dmgMax}
	mods := res.ModifiersData{
		HealthMods: []res.HealthModData{healthMod},
	}
	return effect.NewModifiers(mods)
}

// takeEffects adds specified effects
func (c *Character) TakeEffect(e *effect.Effect) {
	// TODO: handle resists
	// Add effect.
	c.AddEffect(e)
	// Memorize source as hostile.
	mem := TargetMemory{Attitude: Hostile}
	mem.TargetID, mem.TargetSerial = e.Source()
	c.MemorizeTarget(&mem)
	// Add hit effects.
	source := serial.Object(e.Source())
	if s, ok := source.(effect.Target); ok && e.MeleeHit() {
		for _, e := range s.HitEffects() {
			c.TakeEffect(e)
		}
		c.TakeModifiers(s, s.HitModifiers()...)
	}
	if c.onEffectTaken != nil {
		c.onEffectTaken(e)
	}
}
