/*
 * action.go
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

package object

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/effect"	
)

// Struct for object action.
type Action struct {
	selfMods []effect.Modifier
	userMods []effect.Modifier
}

// NewAction creates new object action.
// Returns nil if specified data does
// not contain any data for modifiers.
func NewAction(data res.ObjectActionData) *Action {
	a := new(Action)
	a.selfMods = effect.NewModifiers(data.SelfMods)
	a.userMods = effect.NewModifiers(data.UserMods)
	if len(a.SelfMods()) < 1 && len(a.UserMods()) < 1 {
		return nil
	}
	return a
}

// SelfMods returns action modifiers for action owner.
func (a *Action) SelfMods() []effect.Modifier {
	return a.selfMods
}

// UserMods returns action modifiers for action user.
func (a *Action) UserMods() []effect.Modifier {
	return a.userMods
}
