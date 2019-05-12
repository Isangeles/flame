/*
 * flagmod.go
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
	"github.com/isangeles/flame/core/module/flag"
)

// Struct for flag modifier.
type FlagMod struct {
	flagID string
	flagOn bool
}

// NewFlagMod create new flag modifier.
func NewFlagMod(data res.FlagModData) *FlagMod {
	fm := new(FlagMod)
	fm.flagID = data.ID
	fm.flagOn = data.On
	return fm
}

// Affect modifies targets flags.
func (fm *FlagMod) Affect(source Target, targets ...Target) {
	for _, t := range targets {
		flagger, ok := t.(flag.Flagger)
		if !ok {
			return
		}
		f := flag.Flag(fm.flagID)
		flagger.AddFlag(f)
	}
}

// Undo undos flag modifications on specified targets.
func (fm *FlagMod) Undo(source Target, targets ...Target) {
	for _, t := range targets {
		flagger, ok := t.(flag.Flagger)
		if !ok {
			return
		}
		f := flag.Flag(fm.flagID)
		flagger.RemoveFlag(f)
	}
}
