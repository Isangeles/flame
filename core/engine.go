/*
 * engine.go
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
 
// Flame engine core
// @Isangeles
package core

import (
	"fmt"
	"os"

	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/game"
	"github.com/isangeles/flame/core/enginelog"
	//"github.com/isangeles/flame/core/game/object/character"
)

const (
	NAME, VERSION = "Flame Engine", "0.0.0"
)

var (
	mod module.Module
	gm  game.Game
)

// On init.
func init() {
	err := loadConfig()
	if err != nil {
		enginelog.Error(fmt.Sprintf("config_load_fail:%s\n", err))
	}
}

// LoadModule loads module with specified name from default directory(data/modules).
// Error: if specified module name was invalid
func LoadModule(name, path string) error {
	//TODO loading module data from directory
	if _, err := os.Stat(path + string(os.PathSeparator) + name); os.IsNotExist(err) {
		return fmt.Errorf("module_not_found:'%s' in:'%s'", name, path)
	}
	
	mod = module.NewModule(name, path)
	return nil
}

// Mod returns loaded module
func Mod() *module.Module {
	return &mod
}

// StartGame starts new game for loaded module.
// pcId:  ID of game character in loaded module base for player.
// Error: if no module is loaded.
func StartGame(pcId string) error {
	if !mod.IsLoaded() {
		return fmt.Errorf("no_module_loaded")
	}
	
	pc := mod.GetCharacter(pcId)
	if pc.Id() == "" {
		return fmt.Errorf("not_found_character_with_id:'%s'", pcId)
	}
	
	gm = game.NewGame(mod, pc)
	return nil
}
