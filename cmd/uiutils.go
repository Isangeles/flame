/*
 * uiutils.go
 *
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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

package main

import (
	"fmt"

	"github.com/isangeles/flame/core/module/object/character"
)

// charDisplayString returns string with character
// stats and info.
func charDisplayString(char *character.Character) string {
	return fmt.Sprintf("%s:%s,%s,%s:%d,%d,%d,%d,%d",
		char.Name(), char.Race().ID(), char.Gender().ID(),
		"Stats", char.Attributes().Str, char.Attributes().Con,
		char.Attributes().Dex, char.Attributes().Wis,
		char.Attributes().Int)
}
