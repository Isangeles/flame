/*
 * conf.go
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

package module

import (
	"path/filepath"
)

// Conf struct represents module configuration.
type Conf struct {
	Name, Path, Lang string
	NewcharAttrsMax  int
	NewcharAttrsMin  int
	Chapters         []string
}

// FullPath return full path to module directory.
func (c Conf) FullPath() string {
	return filepath.FromSlash(c.Path + "/" + c.Name)
}

// ChaptersPath returns path to module chapters.
func (c Conf) ChaptersPath() string {
	return filepath.FromSlash(c.FullPath() + "/chapters")
}

// CharactersPath returns path to directory for
// exported characters.
func (c Conf) CharactersPath() string {
	return filepath.FromSlash(c.FullPath() + "/characters")
}
