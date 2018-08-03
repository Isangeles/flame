/*
 * module.go
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
 
// Package module provides engine module struct represenation.
// @Isangeles
package module

import (
	"path/filepath"
	
	"github.com/isangeles/flame/core/game/object/character"
)

var (
	defaultModulesPath string = filepath.FromSlash("data/modules")
)

// Module struct represents engine module.
type Module struct {
	name, path string
	
	characters map[string]character.Character
}

// NewModule creates new representation of module with specified name and
// with specified path.
func NewModule(name, path string) Module {
	m := Module{name: name, path: path}
	m.characters = make(map[string]character.Character)
	return m
}

// Name returns module name
func (m *Module) Name() string {
	return m.name;
}

// Path returns module path
func (m *Module) Path() string {
	return m.path;
}

// Checks if module is loaded.
func (m *Module) IsLoaded() bool {
	return m.name != "" && m.path != ""
}

// GetCharacter returns loaded character with specified ID.
func (m *Module) GetCharacter(id string) character.Character {
	return m.characters[id]
}

// AddCharacter adds character to module characters base.
func (m *Module) AddCharacter(char character.Character) {
	m.characters[char.Id()] = char
}

// DefaultModulesPath returns default path to modules directory.
func DefaultModulesPath() string {
	return defaultModulesPath
}


