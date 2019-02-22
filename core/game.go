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

	"github.com/isangeles/flame/core/data/save"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/log"	
)

// Struct game representation. Contains game module and PCs.
type Game struct {
	mod     *module.Module
	pcs     []*character.Character
	paused bool
}

// NewGame returns new instance of game struct.
func NewGame(mod *module.Module, players []*character.Character) (*Game, error) {
	g := new(Game)
	g.mod = mod
	g.pcs = players
	// Get start scenario.
	chapter := g.Module().Chapter()
	startScen, err := chapter.Scenario(chapter.Conf().StartScenID)
	if err != nil {
		return nil, fmt.Errorf("fail_to_retrieve_start_scenario:%v",
			err)
	}
	// All players to main area of start scenario.
	startArea := startScen.Mainarea()
	for _, pc := range g.pcs {
		g.Module().AssignSerial(pc)
		err := g.ChangePlayerArea(startArea, pc.SerialID())
		if err != nil {
			return nil, fmt.Errorf("fail_to_change_player_area:%v",
				err)
		}
	}
	return g, nil
}

// LoadGame creates game from loaded module
// state and PCs.
func LoadGame(save *save.SaveGame) *Game {
	g := new(Game)
	g.mod = save.Mod
	g.pcs = save.Players
	return g
}

// Update updates game, delta value must be
// time from last update in milliseconds.
func (g *Game) Update(delta int64) {
	if g.paused {
		return
	}
	updateChars := g.Module().Chapter().Characters()
	for _, c := range updateChars {
		c.Update(delta)
	}
}

// Pause toggles game update pause.
func (g *Game) Pause(pause bool) {
	g.paused = pause
	if g.Paused() {
		log.Inf.Printf(lang.Text("ui", "game_pause_info"))
	}
}

// Paused checks whether game is paused.
func (g *Game) Paused() bool {
	return g.paused
}

// Module returns game module.
func (g *Game) Module() *module.Module {
	return g.mod
}

// Players returns all game PCs.
func (g *Game) Players() []*character.Character {
	return g.pcs
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

	area, err := g.Module().Chapter().CharacterArea(pc)
	if err != nil {
		return nil, fmt.Errorf("fail_to_retrive_player_area:%v",
			err)
	}
	return area, nil
}
