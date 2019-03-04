/*
 * lang.go
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

// This package provides easy way to retrieve
// data from transalation files inside data lang
// directory.
package lang

import (
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/core/data/text"
)

const (
	LANG_FILE_EXT = ".lang"
)

var (
	langPath string
	cache    map[string]string
)

// On init.
func init() {
	cache = make(map[string]string)
}

// Text returns text with specified ID from file with specified name in 
// current lang directory.
// In case of error(file/ID not found) returns string with error 
// message.
// Results are cached under lang file + id key, so lang file is open
// only in case when there was no previous requests for specified
// lang file + id pair.
func Text(langFile, id string) string {
	if !strings.HasSuffix(langFile, LANG_FILE_EXT) {
		langFile = langFile + LANG_FILE_EXT
	}
	if cache[langFile + id] != "" {
		return cache[langFile + id]
	}
	fullpath := filepath.FromSlash(LangPath() + "/" + langFile)
	text := text.ReadDisplayText(fullpath, id)[id]
	cache[langFile + id] = text // cache result
	return text
}

// Texts returns map with all values for specified IDs from file
// with specified name in current lang directory(specified IDs as keys).
// In case of error(file/ID not found) returns string with error 
// message.
// Results are cached under lang file + id key, so lang file is open
// only in case when there was no previous requests for specified
// lang file + id pair.
func Texts(langFile string, ids ...string) map[string]string {
	if !strings.HasSuffix(langFile, LANG_FILE_EXT) {
		langFile = langFile + LANG_FILE_EXT
	}
	texts := make(map[string]string)
	uncached := make([]string, 0)
	for _, id := range ids {
		if cache[id] == "" {
			uncached = append(uncached, id)
			continue
		}
		texts[id] = cache[langFile + id]
	}
	fullpath := filepath.FromSlash(LangPath() + "/" + langFile)
	for id, val := range text.ReadDisplayText(fullpath, uncached...) {
		texts[id] = val
	}
	// Cache results.
	for id, t := range texts {
		cache[langFile + id] = t
	}
	return texts
}

// SetLangPath sets specified path as
// current lang directory path.
func SetLangPath(path string) {
	langPath = path
}

// LangPath return path to
// lang directory.
func LangPath() string {
	return langPath
}
