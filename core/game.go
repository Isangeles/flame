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

// core package provides game struct representation.
package core

import (
	"fmt"

	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/scenario"
)

// Struct game representation. Contains game module and PCs.
// Also handles generating of unique serial values for
// all game objects like characters, items, etc.
type Game struct {
	mod *module.Module
	pcs []*character.Character
}

// NewGame returns new instance of game struct.
func NewGame(mod *module.Module, players []*character.Character) *Game {
	g := new(Game)
	g.mod = mod
	g.pcs = players
	// All players to start area.
	startArea := g.Module().Scenario().Area()
	for _, pc := range g.pcs {
		g.ChangePlayerArea(startArea, pc.SerialID())
		g.Module().Chapter().GenerateCharacterSerial(pc)
	}
	return g
}

// Module returns game module.
func (g *Game) Module() *module.Module {
	return g.mod
}

// Player returns player character with specified serial ID
// or nil if no such player character was found.
func (g *Game) Player(serialID string) *character.Character {
	for _, c := range g.pcs {
		if serialID == c.SerialID() {
			return c
		}
	}
	return nil
}

// ChangePlayerArea moves player with specified ID to
// specified area.
func (g *Game) ChangePlayerArea(area *scenario.Area, serialID string) error {
	var pc *character.Character
	for _, c := range g.pcs {
		if serialID == c.SerialID() {
			pc = c
			break
		}
	}
	if pc == nil {
		return fmt.Errorf("player_not_found:%v", serialID)
	}
	area.AddCharacter(pc)
	return nil
}

// PlayerArea returns area for player with specified ID.
func (g *Game) PlayerArea(serialID string) (*scenario.Area, error) {
	var pc *character.Character
	for _, c := range g.pcs {
		if serialID == c.SerialID() {
			pc = c
			break
		}
	}
	if pc == nil {
		return nil, fmt.Errorf("player_not_found:%v", serialID)
	}

	for _, a := range g.mod.Scenario().Areas() {
		if a.ContainsCharacter(pc) {
			return a, nil
		}
	}
	return nil, fmt.Errorf("player_not_found_in_any_scenario_area:%v", serialID)
}

// AreaCharacters returns all characters in current
// area.
func (g *Game) AreaCharacters() []*character.Character {
	return g.Module().Chapter().Scenario().Area().Characters()
}
