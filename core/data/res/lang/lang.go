/*
 * lang.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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

// Package for easy retrieval of translation data.
package lang

import (
	"fmt"
	
	"github.com/isangeles/flame/core/data/res"
)

// Text returns first text for specified ID.
func Text(id string) string {
	data := res.Translation(id)
	if data == nil {
		return fmt.Sprintf("translation not found: %s", id)
	}
	return data.Texts[0]
}

// Texts returns all texts for specified ID.
func Texts(id string) []string {
	data := res.Translation(id)
	if data == nil {
		return []string{fmt.Sprintf("translation not found: %s", id)}
	}
	return data.Texts
}
