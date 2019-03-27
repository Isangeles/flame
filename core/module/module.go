/*
 * module.go
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

// Package module provides engine module struct represenation.
// Also handles generating of unique serial values for
// all game objects like characters, items, etc.
package module

import (
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/effect"
)

// Module struct represents engine module.
type Module struct {
	conf    ModConf
	chapter *Chapter
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
	m.chapter = chapter
	return nil
}

// Path returns path to module parent directory.
func (m *Module) Path() string {
	return m.conf.Path
}

// FullPath return full path to module directory.
func (m *Module) FullPath() string {
	return m.conf.Path
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

// Character return character with specified serial
// ID from lodaed module characters or nil if no such
// character was found.
func (m *Module) Character(serialID string) *character.Character {
	return m.Chapter().Character(serialID)
}

// Target returns 'targetable' game object with specified
// serial ID or nil if on object with such ID was found.
func (m *Module) Target(id, serial string) effect.Target {
	char := m.Chapter().Character(id + "_" + serial)
	if char != nil {
		return char
	}
	ob := m.Chapter().AreaObject(id, serial)
	return ob
}
