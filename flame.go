/*
 * flame.go
 *
 * Copyright 2018-2021 Dariusz Sikora <dev@isangeles.pl>
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

// flame package provides structs for module and chapater.
package flame

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/objects"
)

const (
	Name, Version  = "Flame", "0.1.0-dev"
)

// Module struct represents game module.
type Module struct {
	res              *res.ResourcesData
	conf             *ModuleConfig
	chapter          *Chapter
	onChapterChanged func(c *Chapter)
}

// NewModule creates new game module from specified data.
func NewModule(data res.ModuleData) *Module {
	m := new(Module)
	m.conf = new(ModuleConfig)
	m.Apply(data)
	return m
}

// Update updates module.
func (m *Module) Update(delta int64) {
	if m.Chapter() != nil {
		m.Chapter().Update(delta)
	}
}

// SetChapter sets specified chapter as current chapter.
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
func (m *Module) Conf() *ModuleConfig {
	return m.conf
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

// Resources returns module resources.
func (m *Module) Resources() *res.ResourcesData {
	return m.res
}

// SetOnChapterChangedFunc sets function triggered on chapter change.
func (m *Module) SetOnChapterChangedFunc(f func(c *Chapter)) {
	m.onChapterChanged = f
}

// Apply applies specified data on the module.
// Also, adds module resources to resources
// base in res package.
func (m *Module) Apply(data res.ModuleData) {
	if len(data.Config["id"]) > 0 {
		m.conf.ID = data.Config["id"][0]
	}
	if len(data.Config["path"]) > 0 {
		m.conf.Path = data.Config["path"][0]
	}
	if len(data.Config["chapter"]) > 0 {
		m.conf.Chapter = data.Config["chapter"][0]
	}
	m.res = &data.Resources
	res.Add(*m.res)
	if m.Chapter() == nil || m.Chapter().Conf().ID != data.Chapter.ID {
		chapter := NewChapter(m, data.Chapter)
		m.SetChapter(chapter)
	}
}

// Data creates data resource for module.
func (m *Module) Data() res.ModuleData {
	data := res.ModuleData{ID: m.Conf().ID}
	data.Config = make(map[string][]string)
	data.Config["id"] = []string{m.Conf().ID}
	data.Config["path"] = []string{m.Conf().Path}
	data.Config["chapter"] = []string{m.Chapter().Conf().ID}
	data.Chapter = m.Chapter().Data()
	data.Resources = *m.res
	// Remove old characters and objects from resources, besides basic ones.
	data.Resources.Characters = make([]res.CharacterData, 0)
	for _, c := range m.Resources().Characters {
		if len(c.Serial) < 1 {
			data.Resources.Characters = append(data.Resources.Characters, c)
		}
	}
	data.Resources.Objects = make([]res.ObjectData, 0)
	for _, o := range m.Resources().Objects {
		if len(o.Serial) < 1 {
			data.Resources.Objects = append(data.Resources.Objects, o)
		}
	}
	return data
}
