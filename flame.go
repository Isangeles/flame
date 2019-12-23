/*
 * flame.go
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

// Main package that allows starting new game for current module.
package flame

import (
	"fmt"

	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/character"
	"github.com/isangeles/flame/core/module/serial"
)

const (
	Name, Version = "Flame Engine", "0.0.0"
)

var (
	mod  *module.Module
	game *core.Game
)

// SetModule sets specified module as current module.
func SetModule(m *module.Module) {
	mod = m
}

// Mod returns currently loaded module or nil
// if no module is loaded.
func Mod() *module.Module {
	return mod
}

// Game returns currently active game or nil
// if no game is active.
func Game() *core.Game {
	return game
}

// SetGame sets specified game as
// current game.
func SetGame(g *core.Game) {
	game = g
}

// StartGame starts new game for loaded module with specified
// characters as PCs.
func StartGame(pcs ...*character.Character) (*core.Game, error) {
	if Mod() == nil {
		return nil, fmt.Errorf("no module loaded")
	}
	// Load start chapter for module.
	err := data.LoadChapter(Mod(), Mod().Conf().StartChapter)
	if err != nil {
		return nil, fmt.Errorf("fail to load start chapter: %v", err)
	}
	// Load chapter data(to build quests, characters, erc.).
	err = data.LoadChapterData(Mod().Chapter())
	if err != nil {
		return nil, fmt.Errorf("fail to load module data: %v", err)
	}
	// Load chapter start area.
	chapter := Mod().Chapter()
	err = data.LoadArea(Mod(), chapter.Conf().StartAreaID)
	if err != nil {
		return nil, fmt.Errorf("fail to load start area: %v", err)
	}
	// Create new game.
	game, err = core.NewGame(mod)
	if err != nil {
		return nil, fmt.Errorf("fail to create game: %v", err)
	}
	SetGame(game)
	// All players to start area.
	startArea := chapter.Area(chapter.Conf().StartAreaID)
	if startArea == nil {
		return nil, fmt.Errorf("start area not found: %s: %v",
			chapter.Conf().StartAreaID)
	}
	for _, pc := range pcs {
		serial.AssignSerial(pc)
		startArea.AddCharacter(pc)
	}
	// Add players to game.
	for _, pc := range pcs {
		game.AddPlayer(pc)
	}
	return game, nil
}
