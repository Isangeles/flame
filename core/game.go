/*
 * game.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/area"
	"github.com/isangeles/flame/core/module/character"
	"github.com/isangeles/flame/log"
)

// Struct for game, contains game
// module and PCs.
type Game struct {
	mod    *module.Module
	pcs    map[string]*character.Character
	ai     *AI
	paused bool
}

// NewGame creates new game for specified module.
func NewGame(mod *module.Module) (*Game, error) {
	g := new(Game)
	g.mod = mod
	g.pcs = make(map[string]*character.Character)
	g.ai = NewAI(g)
	if g.mod.Chapter() == nil {
		return nil, fmt.Errorf("no active module chapter")
	}
	// Chapter NPCs under AI control.
	for _, c := range g.mod.Chapter().Characters() {
		g.ai.AddCharacter(c)
	}
	// Events.
	g.mod.Chapter().SetOnAreaAddedFunc(g.onModAreaAdded)
	return g, nil
}

// Update updates game, delta value must be
// time from last update in milliseconds.
func (g *Game) Update(delta int64) {
	if g.paused {
		return
	}
	// Characters.
	updateChars := g.Module().Chapter().Characters()
	for _, c := range updateChars {
		c.Update(delta)
	}
	// Area objects.
	updateObjects := g.Module().Chapter().AreaObjects()
	for _, o := range updateObjects {
		o.Update(delta)
	}
	// AI.
	g.ai.Update(delta)
	// Objects area.
	g.updateObjectsArea()
}

// Pause toggles game update pause.
func (g *Game) Pause(pause bool) {
	g.paused = pause
}

// Paused checks whether game is paused.
func (g *Game) Paused() bool {
	return g.paused
}

// Module returns game module.
func (g *Game) Module() *module.Module {
	return g.mod
}

// AddPlayer adds specified character to game as player.
func (g *Game) AddPlayer(char *character.Character) {
	g.pcs[char.ID()+char.Serial()] = char
	g.ai.RemoveCharacter(char)
}

// Players returns all game PCs.
func (g *Game) Players() (pcs []*character.Character) {
	for _, pc := range g.pcs {
		pcs = append(pcs, pc)
	}
	return
}

// AI returns game AI.
func (g *Game) AI() *AI {
	return g.ai
}

// updateObjectsArea checks and moves game objects to
// proper areas, if needed.
func (g *Game) updateObjectsArea() {
	chapter := g.Module().Chapter()
	if chapter == nil {
		return
	}
	for _, c := range chapter.Characters() {
		currentArea := chapter.CharacterArea(c)
		if currentArea != nil && currentArea.ID() == c.AreaID() {
			continue
		}
		var newArea *area.Area
		// Search for area in current chapter.
		for _, a := range chapter.Areas() {
			if a.ID() == c.AreaID() {
				newArea = a
				break
			}
			for _, sa := range a.AllSubareas() {
				if sa.ID() == c.AreaID() {
					newArea = sa
					break
				}
			}
		}
		if newArea == nil {
			// Search for area data in res package.
			areaData := res.Area(c.AreaID())
			if areaData != nil {
				newArea = area.New(*areaData)
			}
			chapter.AddAreas(newArea)
		}
		if newArea == nil {
			log.Err.Printf("area update: %s#%s: area not found: %s\n",
				c.ID(), c.Serial(), c.AreaID())
			c.SetAreaID(currentArea.ID())
			return
		}
		newArea.AddCharacter(c)
		currentArea.RemoveCharacter(c)
	}
}

// Triggered after adding new area to module chapter.
func (g *Game) onModAreaAdded(a *area.Area) {
	for _, c := range a.AllCharacters() {
		if g.pcs[c.ID()+c.Serial()] != nil {
			continue
		}
		g.ai.AddCharacter(c)
	}
}
