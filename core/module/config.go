/*
 * config.go
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

package module

import (
	"path/filepath"
)

// Config struct represents module configuration.
type Config struct {
	ID      string
	Path    string
	Lang    string
	Chapter string
}

// ChaptersPath returns path to module chapters.
func (c Config) ChaptersPath() string {
	return filepath.Join(c.Path, "chapters")
}

// CharactersPath returns path to directory for
// exported characters.
func (c Config) CharactersPath() string {
	return filepath.Join(c.Path, "characters")
}

// ObjectsPath returns path to directory with
// area objects bases.
func (c Config) ObjectsPath() string {
	return filepath.Join(c.Path, "objects")
}

// ItemsPath returns path to directory with
// items bases.
func (c Config) ItemsPath() string {
	return filepath.Join(c.Path, "items")
}

// EffectsPath returns path to directory with
// effects bases.
func (c Config) EffectsPath() string {
	return filepath.Join(c.Path, "effects")
}

// SkillsPath returns path to directory with
// skills base.
func (c Config) SkillsPath() string {
	return filepath.Join(c.Path, "skills")
}

// RecipesPath returns path to directory with
// recipes base.
func (c Config) RecipesPath() string {
	return filepath.Join(c.Path, "recipes")
}

// LangPath returns path to lang directory.
func (c Config) LangPath() string {
	return filepath.Join(c.Path, "lang", c.Lang)
}

// ItemsLangPath returns path to items lang file.
func (c Config) ItemsLangPath() string {
	return filepath.Join(c.LangPath(), "items")
}

// ChatLangPath returns path to random chat lang file.
func (c Config) ChatLangPath() string {
	return filepath.Join(c.LangPath(), "random_chat")
}

// RecipesLangPath returns path to recipes lang file.
func (c Config) RecipesLangPath() string {
	return filepath.Join(c.LangPath(), "recipes")
}
