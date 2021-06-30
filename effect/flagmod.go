/*
 * flagmod.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/flag"
)

// Struct for flag modifier.
type FlagMod struct {
	flag   flag.Flag
	off    bool
}

// NewFlagMod create new flag modifier.
func NewFlagMod(data res.FlagModData) *FlagMod {
	fm := new(FlagMod)
	fm.flag = flag.Flag(data.ID)
	fm.off = data.Off
	return fm
}

// Flag returns modifier flag.
func (fm *FlagMod) Flag() flag.Flag {
	return fm.flag
}

// Off checks if flag should be removed.
func (fm *FlagMod) Off() bool {
	return fm.off
}

// Data creates data resource for modifier.
func (fm *FlagMod) Data() res.FlagModData {
	data := res.FlagModData{
		ID: string(fm.Flag()),
		Off: fm.Off(),
	}
	return data
}
