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
	"path/filepath"

	"github.com/isangeles/flame/core/module/object/character"
)

var (
	defaultModulesPath string = filepath.FromSlash("data/modules")
)

// Module struct represents engine module.
type Module struct {
	conf    ModConf
	chapter *Chapter
}

// DefaultModulesPath returns default path to modules directory.
func DefaultModulesPath() string {
	return defaultModulesPath
}

// NewModule creates new instance of module with specified configuration
// and data.
func NewModule(conf ModConf) *Module {
	m := new(Module)
	m.conf = conf
	return m
}

// Jumps to next module chapter.
func (m *Module) SetChapter(chapter *Chapter) error {
	// TODO: for now only start chapter.
	/*
	chapPath := filepath.FromSlash(m.ChaptersPath() +
		"/" + m.conf.Chapters[0])
	chapConf, err := LoadChapterConf(chapPath)
	if err != nil {
		return fmt.Errorf("fail_to_read_chapter_conf:%s:%v",
			chapPath, err)
	}
	chapConf.ID = m.conf.Chapters[0]
	c, err := NewChapter(m, chapConf)
	if err != nil {
		return fmt.Errorf("fail_to_set_next_chapter:%v", err)
	}
        */
	m.chapter = chapter
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
	return m.conf.Path
}

// ChaptersPath returns path to module chapters.
func (m *Module) ChaptersPath() string {
	return m.conf.ChaptersPath()
}

// CharactersPath returns path to directory for
// exported characters.
func (m *Module) CharactersPath() string {
	return m.conf.CharactersPath()
}

// Chapter returns current module chapter.
func (m *Module) Chapter() *Chapter {
	return m.chapter
}

// Conf returns module configuration.
func (m *Module) Conf() ModConf {
	return m.conf
}

// LangID return ID of current module
// language.
func (m *Module) LangID() string {
	return m.conf.Lang
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
