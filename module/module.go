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
package module

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/objects"
)

// Module struct represents engine module.
type Module struct {
	Res              res.ResourcesData
	conf             *Config
	chapter          *Chapter
	onChapterChanged func(c *Chapter)
}

// New creates new instance of module from specified data, adds
// module resources to resources base in res package.
func New(data res.ModuleData) *Module {
	m := new(Module)
	m.conf = new(Config)
	if len(data.Config["id"]) > 0 {
		m.conf.ID = data.Config["id"][0]
	}
	if len(data.Config["path"]) > 0 {
		m.conf.Path = data.Config["path"][0]
	}
	if len(data.Config["chapter"]) > 0 {
		m.conf.Chapter = data.Config["chapter"][0]
	}
	m.Res = data.Resources
	res.Add(m.Res)
	chapter := NewChapter(m, data.Chapter)
	m.SetChapter(chapter)
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

// Data creates data resource for module.
func (m *Module) Data() res.ModuleData {
	data := res.ModuleData{ID: m.Conf().ID}
	data.Config = make(map[string][]string)
	data.Config["id"] = []string{m.Conf().ID}
	data.Config["path"] = []string{m.Conf().Path}
	data.Config["chapter"] = []string{m.Chapter().Conf().ID}
	data.Chapter = m.Chapter().Data()
	data.Resources = m.Res
	return data
}
