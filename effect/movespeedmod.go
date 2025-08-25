/*
 * movespeedmod.go
 *
 * Copyright 2025 Dariusz Sikora <ds@isangeles.dev>
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
	"github.com/isangeles/flame/data/res"
)

// Struct for movement speed modifier.
type MoveSpeedMod struct {
	value int64
}

// NewMoveSpeedMod creates new move speed modifier.
func NewMoveSpeedMod(data res.MoveSpeedModData) *MoveSpeedMod {
	msm := MoveSpeedMod{data.Value}
	return &msm
}

// Value returns the movement speed value of the modifier.
func (msm *MoveSpeedMod) Value() int64 {
	return msm.value
}

// Data returns data resource for modifier.
func (msm *MoveSpeedMod) Data() res.MoveSpeedModData {
	return res.MoveSpeedModData{msm.value}
}
