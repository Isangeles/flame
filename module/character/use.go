/*
 * use.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/useaction"
)

// Use checks requirements and starts cast action for
// specified usable object.
func (c *Character) Use(ob useaction.Usable) {
	if ob.UseAction() == nil {
		return
	}
	// Check requirements.
	if !c.MeetReqs(ob.UseAction().Requirements()...) {
		c.SendPrivate(lang.Text("cant_do_right_now"))
		return
	}
	c.casted = ob
}

// useCasted applies modifiers and effects from specified
// usable object on use action owner, user and user target.
func (c *Character) useCasted(ob useaction.Usable) {
	// Apply effects and modifiers.
	c.TakeModifiers(ob, ob.UseAction().UserMods()...)
	for _, e := range ob.UseAction().UserEffects() {
		c.TakeEffect(e)
	}
	if tar, ok := ob.(effect.Target); ok {
		tar.TakeModifiers(ob, ob.UseAction().UserMods()...)
		for _, e := range ob.UseAction().UserEffects() {
			tar.TakeEffect(e)
		}
	}
	if len(c.Targets()) > 0 && c.Targets()[0] != nil {
		tar := c.Targets()[0]
		tar.TakeModifiers(tar, ob.UseAction().TargetMods()...)
		for _, e := range ob.UseAction().TargetEffects() {
			tar.TakeEffect(e)
		}
		tar.TakeModifiers(tar, ob.UseAction().TargetUserMods()...)
		for _, e := range ob.UseAction().TargetUserEffects() {
			tar.TakeEffect(e)
		}
	} else {
		c.TakeModifiers(ob, ob.UseAction().TargetUserMods()...)
		for _, e := range ob.UseAction().TargetUserEffects() {
			c.TakeEffect(e)
		}
	}
}
