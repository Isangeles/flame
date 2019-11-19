/*
 * areamod.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

package effect

import (
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object"
)

// Interface for area modifier.
type AreaMod struct {
	areaID string
	enterX float64
	enterY float64
}

// NewAreaMod creates new area modifier.
func NewAreaMod(data res.AreaModData) *AreaMod {
	am := new(AreaMod)
	am.areaID = data.ID
	am.enterX = data.EnterX
	am.enterY = data.EnterY
	return am
}

// AreaID returns modifier area ID.
func (am *AreaMod) AreaID() string {
	return am.areaID
}

// EnterPosition returns position for object after
// area change.
func (am *AreaMod) EnterPosition() (float64, float64) {
	return am.enterX, am.enterY
}

// Affect moves all targets to area.
func (am *AreaMod) Affect(source Target, targets ...Target) {
 	for _, t := range targets {
		if c, ok := t.(object.AreaObject); ok {
			c.SetAreaID(am.areaID)
			c.SetPosition(am.enterX, am.enterY)
		}
	}
}

// Undo does nothing.
func (am *AreaMod) Undo(source Target, tagerts ...Target) {
}
