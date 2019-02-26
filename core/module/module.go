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
	"fmt"

	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/object/item"
)

// Module struct represents engine module.
type Module struct {
	conf    ModConf
	chapter *Chapter
	items   []Serialer
	effects []Serialer
}

// NewModule creates new instance of module with specified configuration
// and data.
func NewModule(conf ModConf) *Module {
	m := new(Module)
	m.conf = conf
	m.items = make([]Serialer, 0)
	return m
}

// Jumps to next module chapter.
func (m *Module) SetChapter(chapter *Chapter) error {
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

// Object returns game object with specified serial ID
// or nil if on object with such ID was found.
func (m *Module) Object(id, serial string) effect.Target {
	char := m.Character(id + "_" + serial)
	if char != nil {
		return char
	}
	return nil
}

// AssignSerial sets unique serial value for
// specified object with serial value.
// Returns error if no active chapter set.
func (m *Module) AssignSerial(ob Serialer) error {
	switch ob := ob.(type) {
	case *character.Character:
		chapter := m.Chapter()
		if chapter == nil {
			return fmt.Errorf("no active chapter set")
		}
		m.assignCharacterSerial(ob)
		return nil
	case item.Item:
		m.assignItemSerial(ob)
		return nil
	case *effect.Effect:
		m.assignEffectSerial(ob)
		return nil
	default:
		return fmt.Errorf("unsupported game object type")
	}
}


// assignCharacterSerial sets unique serial value for specified
// object with serial ID.
func (m *Module) assignCharacterSerial(char *character.Character) {
	chars := m.Chapter().CharactersWithID(char.ID())
	objects := make([]Serialer, 0)
	for _, c := range chars {
		objects = append(objects, c)
	}
	serial := uniqueSerial(objects)
	// Assing serial value to char.
	char.SetSerial(serial)
}

// assignItemSerial assigns unique serial value to
// specified item.
func (m *Module) assignItemSerial(it item.Item) {
	serial := uniqueSerial(m.items)
	it.SetSerial(serial)
	m.items = append(m.items, it)
}

// assignEffectSerial assigns unique serial value to
// specified effect.
func (m *Module) assignEffectSerial(ef *effect.Effect) {
	serial := uniqueSerial(m.effects)
	ef.SetSerial(serial)
	m.effects = append(m.effects, ef)
}


