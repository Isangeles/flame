/*
 * game.go
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
 
// game package provides game struct representation.
// @Isnageles
package game

import (
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/game/object/character"
)

// Struct Game represents game.
type Game struct {
	mod module.Module
	pc  character.Character
}

// NewGame returns new instance of game struct.
func NewGame(mod module.Module, player character.Character) Game {
	return Game{mod, player}
}

