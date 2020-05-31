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
	
	"github.com/isangeles/flame/data/res"
)

// Text returns first text for specified ID.
// Returns error text if translation data for
// specified ID was not found.
func Text(id string) string {
	data, ok := res.Translations[id]
	if !ok {
		return fmt.Sprintf("translation not found: %s", id)
	}
	return data.Texts[0]
}

// Texts returns all texts for specified ID.
// Returns 1-length slice with error text
// if transaltion data for specified ID was
// not found.
func Texts(id string) []string {
	data, ok := res.Translations[id]
	if !ok {
		return []string{fmt.Sprintf("translation not found: %s", id)}
	}
	return data.Texts
}
