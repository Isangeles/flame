/*
 * combat.go
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

package character

import (
	"fmt"

	"github.com/isangeles/flame/config"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/item"
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/serial"
)

// Damage retruns min and max damage value,
// including weapons, active effects, etc.
func (c *Character) Damage() (int, int) {
	min, max := c.Attributes().Damage()
	var rightHandItem item.Equiper
	for _, s := range c.Equipment().Slots() {
		if s.Type() != HandRight {
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
		if s.Type() != HandRight {
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

// DamageEffects returns character damage effects.
func (c *Character) DamageEffects() []*effect.Effect {
	effects := make([]*effect.Effect, 0)
	var rightHandItem item.Equiper
	for _, s := range c.Equipment().Slots() {
		if s.Type() != HandRight {
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

// HitEffects returns all character hit effects.
func (c *Character) HitEffects() []*effect.Effect {
	dmgMin, dmgMax := c.Damage()
	healthMod := res.HealthModData{-dmgMin, -dmgMax}
	mods := res.ModifiersData{
		HealthMods: []res.HealthModData{healthMod},
	}
	hitData := res.EffectData{
		ID:        c.ID() + c.Serial() + "_hit",
		Modifiers: mods,
	}
	hitEffect := c.buildEffects(hitData)
	effects := c.DamageEffects()
	effects = append(effects, hitEffect...)
	return effects
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
	}
	// Send combat message.
	msg := fmt.Sprintf("%s: %s: %s", c.Name(), lang.Text("ob_effect"), e.Name())
	if config.Debug { // add effect serial ID to combat message
		msg = fmt.Sprintf("%s(%s_%s)", msg, e.ID(), e.Serial())
	}
	c.SendCombat(msg)
}
