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
	"github.com/isangeles/flame/module"
	"github.com/isangeles/flame/module/area"
)

// Struct for game, a wrapper for the game module.
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
	g.Module().Update(delta)
	g.AI().Update(delta)
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
