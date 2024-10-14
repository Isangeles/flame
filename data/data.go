/*
 * data.go
 *
 * Copyright 2018-2024 Dariusz Sikora <ds@isangeles.dev>
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

// Package with functions for importing/exporting
// data files and directories.
package data

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"path/filepath"

	"github.com/isangeles/flame/data/res"
)

var jsonData = false

// LoadTranslationData loads all lang files from
// from directory with specified path.
func LoadTranslationData(path string) error {
	// Translation.
	langData, err := ImportLangDir(path)
	if err != nil {
		return fmt.Errorf("Unable to import lang dir: %v", err)
	}
	lang := filepath.Base(path)
	base := res.TranslationBase(lang)
	if base == nil {
		return fmt.Errorf("Translation base not found: %s", lang)
	}
	base.Translations = append(base.Translations, langData...)
	return nil
}

// unmarshal decodes specified data buffer with
// XML or JSON decoder, depnding on value of jsonData
// variable.
func unmarshal(buffer []byte, dataStruct any) error {
	if jsonData {
		return json.Unmarshal(buffer, dataStruct)
	}
	return xml.Unmarshal(buffer, dataStruct)
}

// marshal encodes specified data into JSON or
// XML format, depending on the value of jsonData
// variable.
func marshal(data any) ([]byte, error) {
	if jsonData {
		return json.Marshal(data)
	}
	return xml.Marshal(data)
}
