/*
 * module.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/core/module/effect"
	"github.com/isangeles/flame/core/module/objects"
)

// Module struct represents engine module.
type Module struct {
	conf    *Config
	chapter *Chapter
	onChapterChanged func(c *Chapter)
}

// New creates new instance of module with specified
// configuration and data.
func New(conf Config) *Module {
	m := new(Module)
	m.conf = &conf
	return m
}

// Jumps to next module chapter.
func (m *Module) SetChapter(chapter *Chapter) {
	m.chapter = chapter
	if m.onChapterChanged != nil {
		m.onChapterChanged(m.Chapter())
	}
}

// Chapter returns current module chapter.
func (m *Module) Chapter() *Chapter {
	return m.chapter
}

// Conf returns module configuration.
func (m *Module) Conf() *Config {
	return m.conf
}

// Target returns 'targetable' game object with specified
// serial ID or nil if on object with such ID was found.
func (m *Module) Target(id, serial string) effect.Target {
	char := m.Chapter().Character(id, serial)
	if char != nil {
		return char
	}
	ob := m.Chapter().AreaObject(id, serial)
	if ob != nil {
		return ob
	}
	return nil
}

// Object returns game object with specified ID and serial
// or nil if no such object was found.
func (m *Module) Object(id, serial string) objects.Object {
	char := m.Chapter().Character(id, serial)
	if char != nil {
		return char
	}
	ob := m.Chapter().AreaObject(id, serial)
	if ob != nil {
		return ob
	}
	return nil
}

// SetOnChapterChangedFunc sets function triggered on chapter change.
func (m *Module) SetOnChapterChangedFunc(f func(c *Chapter)) {
	m.onChapterChanged = f
}
