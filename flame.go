/*
 * flame.go
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
 
// Flame engine is modular RPG game engine.
package flame

import (
	"fmt"
	"path/filepath"

	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/character"
)

const (
	NAME, VERSION = "Flame Engine", "0.0.0"
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

// StartGame starts new game for loaded module with specified character
// as PC.
// Error: if no module is loaded.
func StartGame(pcs []*character.Character) (*core.Game, error) {
	if Mod() == nil {
		return nil, fmt.Errorf("no_module_loaded")
	}
	// Load start chapter.
	chapPath := filepath.FromSlash(Mod().ChaptersPath() +
		"/" + Mod().Conf().Chapters[0])
	chapConf, err := data.LoadChapterConf(chapPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_start_chapter_conf:%s:%v",
			chapPath, err)
	}
	chapConf.ID = Mod().Conf().Chapters[0]
	startChap, err := module.NewChapter(Mod(), chapConf)
	if err != nil {
		return nil, fmt.Errorf("fail_to_set_mod_start_chapter:%v",
			err)
	}
	err = mod.SetChapter(startChap) // move to start chapter
	if err != nil {
		return nil, fmt.Errorf("fail_to_move_mod_to_start_chapter:%v",
			err)
	}
	game, err = core.NewGame(mod, pcs)
	if err != nil {
		return nil, fmt.Errorf("fail_to_create_new_game:%v",
			err)
	}
	return game, nil
}
