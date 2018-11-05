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
// @Isangeles
package flame

import (
	"fmt"
	"log"
	
	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/enginelog"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/character"
)

const (
	NAME, VERSION = "Flame Engine", "0.0.0"
)

var (
	inflog *log.Logger = log.New(enginelog.InfLog, "flame-core>", 0)
	errlog *log.Logger = log.New(enginelog.ErrLog, "flame-core>", 0)
	dbglog *log.Logger = log.New(enginelog.DbgLog, "flame-debug>", 0)
	
	mod *module.Module
	g   *core.Game
)

// SetModule sets specified module as current module.
func SetModule(m *module.Module) {
	mod = m
}

// Mod returns loaded module
func Mod() *module.Module {
	return mod
}

// StartGame starts new game for loaded module with specified character
// as PC.
// Error: if no module is loaded.
func StartGame(pcs []*character.Character) error {
	if mod == nil {
		return fmt.Errorf("no_module_loaded")
	}
	
	g = core.NewGame(mod, pcs)
	return nil
}
