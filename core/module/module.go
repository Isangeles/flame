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
package module

import (
	"fmt"
	"path/filepath"
	"os"
	"strings"

	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/scenario"
)

var (
	defaultModulesPath string = filepath.FromSlash("data/modules")
)

// Module struct represents engine module.
type Module struct {
	conf    Conf
	chapter *Chapter
}

// DefaultModulesPath returns default path to modules directory.
func DefaultModulesPath() string {
	return defaultModulesPath
}

// NewModule creates new instance of module with specified configuration
// and data.
func NewModule(name, path string) (*Module, error) {
	m := new(Module)
	conf, err := loadModConf(name, path)
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_config:%v", err)
	}
	m.conf = conf
	return m, nil
}

// Jumps to next module chapter.
func (m *Module) NextChapter() error {
	// TODO: for now only start chapter.
	c, err := NewChapter(m, m.conf.Chapters[0])
	if err != nil {
		return fmt.Errorf("fail_to_set_next_chapter:%v", err)
	}
	m.chapter = c
	return nil
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

// ChaptersPath returns path to module chapters.
func (m *Module) ChaptersPath() string {
	return filepath.FromSlash(m.FullPath() + "/chapters")
}

// CharactersPath returns path to directory for
// exported characters.
func (m *Module) CharactersPath() string {
	return filepath.FromSlash(m.FullPath() + "/characters")
}

// Chapter returns current module chapter.
func (m *Module) Chapter() *Chapter {
	return m.chapter
}

// Scenario returns current module scenario.
func (m *Module) Scenario() *scenario.Scenario {
	return m.chapter.Scenario()
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

// ChaptersIds returns slice with module chapters
// ID's.
func (m *Module) ChaptersIds() []string {
	return m.conf.Chapters
}

// Character return character with specified serial
// ID from lodaed module characters or nil if no such
// character was found.
func (m *Module) Character(serialID string) *character.Character {
	return m.Chapter().Character(serialID)
}

// loadModConf loads module configuration file
// from specified path.
func loadModConf(name, path string) (Conf, error) {
	if _, err := os.Stat(path + string(os.PathSeparator) + name);
	os.IsNotExist(err) {
		return Conf{}, fmt.Errorf("module_not_found:'%s' in:'%s'",
			name, path)
	}
	modConfPath := filepath.FromSlash(path + "/" + name + "/mod.conf")
	confValues, err := text.ReadConfigInt(modConfPath, "new_char_attrs_min",
		"new_char_attrs_max")
	if err != nil {
		return Conf{}, fmt.Errorf("fail_to_retrieve_int_values:%s",
			err)
	}
	confChapters, err := text.ReadConfigValue(modConfPath, "chapters")
	if err != nil {
		return Conf{}, fmt.Errorf("fail_to_retrieve_chapters_ids:%s",
			err)
	}
	chapters := strings.Split(confChapters[0], ";")
	if len(chapters) < 1 {
		return Conf{}, fmt.Errorf("no_chapters_specified")
	}
	conf := Conf{
		Name:name,
		Path:path,
		NewcharAttrsMin:confValues[0],
		NewcharAttrsMax:confValues[1],
		Chapters:chapters,
	}
	return conf, nil
}


