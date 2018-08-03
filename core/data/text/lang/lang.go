/*
 * lang.go
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

// Package for connection with transalation data files.
// @Isangeles
package lang

import (
	"os"
	"path/filepath"

	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/data/text"
)

var (
	mainLangPath string = filepath.FromSlash("data/lang")
)

// GetUIText returns text with specified ID from main UI lang file('ui.lang')
// in 'core.MainLangPath()' directory.
// In case of error(file/ID not found) returns string with error
// message.
func GetUIText(textId string) string {
	return text.ReadDisplayText(MainLangPath() + string(os.PathSeparator) + 
					"ui.lang", textId)[0]
}

// GetUITexts returns all text lines from UI lang file with specified IDs
// In case of error(file/ID not found) returns string with error
// message.
// In case of error(file/ID not found) returns string with error
// message instead of text. 
func GetUITexts(textIDs ...string) []string {
	return text.ReadDisplayText(MainLangPath() + string(os.PathSeparator) + 
					"ui.lang", textIDs...)
}

// MainLangPath return path to main lang direcotry for current language.
func MainLangPath() string {
	return filepath.FromSlash(mainLangPath + "/" + core.LangID())
}
