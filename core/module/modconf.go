/*
 * modconf.go
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

package module

import (
	"path/filepath"
)

// ModConf struct represents module configuration.
type ModConf struct {
	Name, Path, Lang string
	NewcharAttrsMax  int
	NewcharAttrsMin  int
	Chapters         []string
}

// ChaptersPath returns path to module chapters.
func (c ModConf) ChaptersPath() string {
	return filepath.FromSlash(c.Path + "/chapters")
}

// CharactersPath returns path to directory for
// exported characters.
func (c ModConf) CharactersPath() string {
	return filepath.FromSlash(c.Path + "/characters")
}

// ObjectsPath returns path to directory with
// area objects bases.
func (c ModConf) ObjectsPath() string {
	return filepath.FromSlash(c.Path + "/objects")
}

// ItemsPath returns path to directory with
// items bases.
func (c ModConf) ItemsPath() string {
	return filepath.FromSlash(c.Path + "/items")
}

// EffectsPath returns path to directory with
// effects bases.
func (c ModConf) EffectsPath() string {
	return filepath.FromSlash(c.Path + "/effects")
}

// SkillsPath returns path to directory with
// skills base.
func (c ModConf) SkillsPath() string {
	return filepath.FromSlash(c.Path + "/skills")
}

// LangPath returns path to lang directory.
func (c ModConf) LangPath() string {
	return filepath.FromSlash(c.Path + "/lang/" + c.Lang)
}
