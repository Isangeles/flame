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
}

// NewAreaMod creates new area modifier.
func NewAreaMod(data res.AreaModData) *AreaMod {
	am := new(AreaMod)
	am.areaID = data.ID
	return am
}

// AreaID returns modifier area ID.
func (am *AreaMod) AreaID() string {
	return am.areaID
}

// Affect moves all targets to area.
func (am *AreaMod) Affect(source Target, targets ...Target) {
 	for _, t := range targets {
		if c, ok := t.(object.AreaObject); ok {
			c.SetAreaID(am.areaID)
		}
	}
}

// Undo does nothing.
func (am *AreaMod) Undo(source Target, tagerts ...Target) {
}
