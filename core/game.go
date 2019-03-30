/*
 * game.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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

// core package provides game struct.
package core

import (
	"fmt"

	"github.com/isangeles/flame/core/data/save"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/serial"
	"github.com/isangeles/flame/log"
)

// Struct for game, contains game
// module and PCs.
type Game struct {
	mod    *module.Module
	pcs    []*character.Character
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
		serial.AssignSerial(pc)
		startArea.AddCharacter(pc)
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
	go g.listenWorld()
	if g.paused {
		return
	}
	updateChars := g.Module().Chapter().Characters()
	for _, c := range updateChars {
		c.Update(delta)
	}
	updateObjects := g.Module().Chapter().AreaObjects()
	for _, o := range updateObjects {
		o.Update(delta)
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

// listenWorld listens players and near objects
// messages channels ands prints messages to 
// engine log.
func (g *Game) listenWorld() {
	// Players.
	for _, pc := range g.pcs {
		select {
		case msg := <-pc.CombatLog():
			log.Cmb.Printf(msg)
		default:
		}
		// Near objects.
		area, err := g.Module().Chapter().CharacterArea(pc)
		if err != nil {
			continue
		}
		for _, tar := range area.NearTargets(pc, pc.SightRange()) {
			select {
			case msg := <-tar.CombatLog():
				log.Cmb.Printf(msg)
			default:
			}
		}
	}
}
