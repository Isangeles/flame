/*
 * flame.go
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

// flame package provides game struct.
package flame

import (
	"github.com/isangeles/flame/ai"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module"
	"github.com/isangeles/flame/module/area"
	"github.com/isangeles/flame/log"
)

// Struct for game, contains game
// module and PCs.
type Game struct {
	mod    *module.Module
	npcAI  *ai.AI
	paused bool
}

// NewGame creates new game for specified module.
func NewGame(mod *module.Module) *Game {
	g := new(Game)
	g.mod = mod
	g.npcAI = ai.New(g.Module())
	if g.Module().Chapter() != nil {
		g.chapterLoaded(g.Module().Chapter())
	}
	g.Module().SetOnChapterChangedFunc(g.chapterLoaded)
	return g
}

// Update updates game, delta value must be
// time from last update in milliseconds.
func (g *Game) Update(delta int64) {
	if g.paused {
		return
	}
	chapter := g.Module().Chapter()
	// Characters.
	for _, c := range chapter.Characters() {
		c.Update(delta)
	}
	// Area objects.
	for _, o := range chapter.AreaObjects() {
		o.Update(delta)
	}
	g.AI().Update(delta)
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

// AI returns game AI.
func (g *Game) AI() *ai.AI {
	return g.npcAI
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
			areaData, ok := res.Areas[c.AreaID()]
			if !ok {
				log.Err.Printf("area update: %s#%s: area not found: %s\n",
					c.ID(), c.Serial(), c.AreaID())
				c.SetAreaID(currentArea.ID())
				return
			}
			newArea = area.New(areaData)
			chapter.AddAreas(newArea)
		}
		newArea.AddCharacter(c)
		currentArea.RemoveCharacter(c)
	}
}

// chapterLoaded handles new loaded chapter.
func (g *Game) chapterLoaded(c *module.Chapter) {
	for _, a := range c.Areas() {
		g.areaLoaded(a)
	}
	c.SetOnAreaAddedFunc(g.areaLoaded)
}

// areaLoaded handles new loaded area.
func (g *Game) areaLoaded(a *area.Area) {
	chapter := g.Module().Chapter()
	if chapter == nil {
		return
	}
	for _, c := range a.AllCharacters() {
		if c.AI() {
			g.AI().AddCharacters(c)
		}
	}
}
