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

// Variable with language ID, "english" by default.
var ID = "english"

// Text returns first text for specified ID and current Lang variable.
// Returns error text if translation data for
// specified ID was not found.
func Text(id string) string {
	data, found := Translation(id)
	if !found {
		return fmt.Sprintf("translation not found: %s", id)
	}
	return data.Texts[0]
}

// Texts returns all texts for specified ID and current Lang variable.
// Returns 1-length slice with error text
// if transaltion data for specified ID was
// not found.
func Texts(id string) []string {
	data, found := Translation(id)
	if !found {
		return []string{fmt.Sprintf("translation not found: %s", id)}
	}
	return data.Texts
}

// AddTranslation add specified translation data to the translation
// base for current Lang variable.
func AddTranslation(data res.TranslationData) {
	base := res.TranslationBase(ID)
	if base == nil {
		base := new(res.TranslationBaseData)
		base.ID = ID
		res.TranslationBases = append(res.TranslationBases, base)
	}
	base.Translations = append(base.Translations, data)
}

// Translation returns translation data for specified ID and current
// Lang variable. Second return argument indicates whether data was
// found or not.
func Translation(id string) (data res.TranslationData, found bool) {
	base := res.TranslationBase(ID)
	if base == nil {
		return data, false
	}
	for _, t := range base.Translations {
		if t.ID != id {
			continue
		}
		return t, true
	}
	return data, false
}
