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
	
	//"github.com/isangeles/flame/core/data/text" 
	"github.com/isangeles/flame/core/module/object/character"
)

var (
	defaultModulesPath string = filepath.FromSlash("data/modules")
)

// Module struct represents engine module.
type Module struct {
	conf           Conf
	chapters       []*Chapter
	currentChapter *Chapter
}

// DefaultModulesPath returns default path to modules directory.
func DefaultModulesPath() string {
	return defaultModulesPath
}

// NewModule creates new instance of module with specified configuration
// and data.
func NewModule(conf Conf, chapters []*Chapter) (*Module) {
	m := new(Module)
	m.conf = conf
	m.chapters = chapters
	return m
}

// Name returns module name
func (m *Module) Name() string {
	return m.conf.Name;
}

// Path returns path to module parent directory.
func (m *Module) Path() string {
	return m.conf.Path
}

// FullPath return full path to module directory.
func (m *Module) FullPath() string {
	return filepath.FromSlash(m.Path() + "/" + m.Name())
}

// CharactersBasePath returns path to XML document with module characters.
/*
func (m *Module) CharactersBasePath() string {
	return filepath.FromSlash("data/modules/" + m.Name() + "/npcs/npcs.base")
}
*/

// NewcharAttrsMin returns minimal amount of
// attributes points for new characer.
func (m *Module) NewcharAttrsMin() int {
	return m.conf.NewcharAttrsMin
}

// NewCharAttrsMax return maximal amount of
// attributes points for new character.
func (m *Module) NewcharAttrsMax() int {
	return m.conf.NewcharAttrsMax
}

// Character return character with specified ID from module
// character or nil if no such character found.
func (m *Module) Character(id string) *character.Character {
	// TODO: search module characters.
	return nil
}



