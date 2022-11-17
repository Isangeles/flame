/*
 * moduleconfig.go
 *
 * Copyright 2018-2022 Dariusz Sikora <dev@isangeles.pl>
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

package flame

import (
	"path/filepath"
)

// ModuleConfig struct represents module configuration.
type ModuleConfig struct {
	ID      string
	Path    string
	Chapter string
}

// ChaptersPath returns path to module chapters.
func (c ModuleConfig) ChaptersPath() string {
	return filepath.Join(c.Path, "chapters")
}

// CharactersPath returns path to directory for
// exported characters.
func (c ModuleConfig) CharactersPath() string {
	return filepath.Join(c.Path, "characters")
}

// ObjectsPath returns path to directory with
// area objects bases.
func (c ModuleConfig) ObjectsPath() string {
	return filepath.Join(c.Path, "objects")
}

// ItemsPath returns path to directory with
// items bases.
func (c ModuleConfig) ItemsPath() string {
	return filepath.Join(c.Path, "items")
}

// EffectsPath returns path to directory with
// effects bases.
func (c ModuleConfig) EffectsPath() string {
	return filepath.Join(c.Path, "effects")
}

// SkillsPath returns path to directory with
// skills base.
func (c ModuleConfig) SkillsPath() string {
	return filepath.Join(c.Path, "skills")
}

// RecipesPath returns path to directory with
// recipes base.
func (c ModuleConfig) RecipesPath() string {
	return filepath.Join(c.Path, "recipes")
}

// RacesPath returns path to directory with
// races data files.
func (c ModuleConfig) RacesPath() string {
	return filepath.Join(c.CharactersPath(), "races")
}

// LangPath returns path to lang directory.
func (c ModuleConfig) LangPath() string {
	return filepath.Join(c.Path, "lang")
}
