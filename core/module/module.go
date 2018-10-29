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
	"fmt"
	"path/filepath"
	"os"
	
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/data/text" 
	"github.com/isangeles/flame/core/game/object/character"
)

var (
	defaultModulesPath string = filepath.FromSlash("data/modules")
)

// Module struct represents engine module.
type Module struct {
	name, path      string
	newcharAttrsMin int
	newcharAttrsMax int
	
	chapters   map[int]Chapter
	characters map[string]character.Character
}

// DefaultModulesPath returns default path to modules directory.
func DefaultModulesPath() string {
	return defaultModulesPath
}

// NewModule creates new representation of module with specified name and
// with specified path.
func NewModule(name, path string) (*Module, error) {
	//TODO loading module data from directory
	var m Module
	if _, err := os.Stat(path + string(os.PathSeparator) + name);
	os.IsNotExist(err) {
		return nil, fmt.Errorf("module_not_found:'%s' in:'%s'", name, path)
	}
	m.name = name
	m.path = path
	confValues, err := text.ReadConfigInt(m.ConfPath(), "new_char_attrs_min",
		"new_char_attrs_max")
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_module_conf:%s", err)
	}
	m.newcharAttrsMin, m.newcharAttrsMax = confValues[0], confValues[1]
	
	m.characters = make(map[string]character.Character)
	return &m, nil
}

// LoadData loads module data
func (m *Module) LoadData() error {
	chars, err := data.GetCharactersFromXML(m.CharactersBasePath())
	if err != nil {
		return fmt.Errorf("fail_to_load_npcs_base:%v", err)
	}
	m.characters = *chars
	
	return nil
}

// Name returns module name
func (m *Module) Name() string {
	return m.name;
}

// Path returns path to PARENT module directory.
func (m *Module) Path() string {
	return m.path;
}

// FullPath returns path to module directory.
func (m *Module) FullPath() string {
	return filepath.FromSlash(m.Path() + "/" + m.Name())
}

// ConfPath returns path to module configuration file.
func (m *Module) ConfPath() string {
	return filepath.FromSlash(m.FullPath() + "/mod.conf")
}

// CharactersBasePath returns path to XML document with module characters.
func (m *Module) CharactersBasePath() string {
	return filepath.FromSlash("data/modules/" + m.Name() + "/npcs/npcs.base")
}

// NewcharAttrsMin returns minimal amount of
// attributes points for new characer.
func (m *Module) NewcharAttrsMin() int {
	return m.newcharAttrsMin
}

// NewCharAttrsMax return maximal amount of
// attributes points for new character.
func (m *Module) NewcharAttrsMax() int {
	return m.newcharAttrsMax
}

// Checks if module is loaded.
//func (m *Module) Loaded() bool {
//	return m.name != "" && m.path != ""
//}

// GetCharacter returns loaded character with specified ID.
func (m *Module) GetCharacter(id string) character.Character {
	return m.characters[id]
}

// AddCharacter adds character to module characters base.
func (m *Module) AddCharacter(char character.Character) {
	m.characters[char.Id()] = char
}



