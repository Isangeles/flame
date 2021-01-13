/*
 * use.go
 *
 * Copyright 2020-2021 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/skill"
	"github.com/isangeles/flame/module/training"
	"github.com/isangeles/flame/module/useaction"
)

// Use checks requirements and starts cast action for
// specified usable object.
// Returns an error if use requirements are not meet,
// cooldown is active, or character is currently moving.
func (c *Character) Use(ob useaction.Usable) error {
	if ob.UseAction() == nil {
		return nil
	}
	reqs := ob.UseAction().Requirements()
	if t, ok := ob.(*training.TrainerTraining); ok {
		reqs = t.Requirements()
	}
	// Check requirements and cooldown.
	if !c.MeetReqs(reqs...) || c.cooldown > 0 ||
		ob.UseAction().Cooldown() > 0 || c.Moving() {
		if !c.MeetTargetRangeReqs(reqs...) {
			tar := c.Targets()[0]
			c.SetDestPoint(tar.Position())
		}
		return fmt.Errorf("cant cast action")
	}
	c.casted = ob
	return nil
}

// useCasted applies modifiers and effects from specified
// usable object on use action owner, user and user target.
func (c *Character) useCasted(ob useaction.Usable) {
	// Check requirements.
	reqs := ob.UseAction().Requirements()
	if t, ok := ob.(*training.TrainerTraining); ok {
		reqs = t.Requirements()
	}
	if !c.MeetReqs(reqs...) {
		c.PrivateLog().Add(lang.Text("cant_do_right_now"))
		return
	}
	c.ChargeReqs(reqs...)
	// Apply effects and modifiers.
	c.TakeModifiers(c, ob.UseAction().UserMods()...)
	for _, e := range ob.UseAction().UserEffects() {
		c.TakeEffect(e)
	}
	if tar, ok := ob.(effect.Target); ok {
		tar.TakeModifiers(c, ob.UseAction().UserMods()...)
		for _, e := range ob.UseAction().UserEffects() {
			e.SetSource(c.ID(), c.Serial())
			tar.TakeEffect(e)
		}
	}
	if len(c.Targets()) > 0 && c.Targets()[0] != nil {
		tar := c.Targets()[0]
		tar.TakeModifiers(c, ob.UseAction().TargetMods()...)
		for _, e := range ob.UseAction().TargetEffects() {
			e.SetSource(c.ID(), c.Serial())
			tar.TakeEffect(e)
		}
		tar.TakeModifiers(c, ob.UseAction().TargetUserMods()...)
		for _, e := range ob.UseAction().TargetUserEffects() {
			e.SetSource(c.ID(), c.Serial())
			tar.TakeEffect(e)
		}
	} else {
		c.TakeModifiers(c, ob.UseAction().TargetUserMods()...)
		for _, e := range ob.UseAction().TargetUserEffects() {
			e.SetSource(c.ID(), c.Serial())
			c.TakeEffect(e)
		}
	}
	if s, ok := ob.(*skill.Skill); ok && c.onSkillActivated != nil {
		c.onSkillActivated(s)
	}
	ob.UseAction().SetCast(0)
	ob.UseAction().SetCooldown(ob.UseAction().CooldownMax())
	c.cooldown = globalCD
}
