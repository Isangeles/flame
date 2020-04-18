/*
 * data.go
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

// Package with functions for importing/exporting
// data files and directories.
package data

import (
	"fmt"
	"path/filepath"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module"
)

// LoadModuleLang loads translation data for specified
// language for specified module.
func LoadModuleLang(mod *module.Module, lang string) error {
	modLangPath := filepath.Join(mod.Conf().LangPath(), lang)
	err := LoadTranslationData(modLangPath)
	if err != nil {
		return fmt.Errorf("unable to load module translation data: %v", err)
	}
	if mod.Chapter() == nil {
		return nil
	}
	chapterLangPath := filepath.Join(mod.Chapter().Conf().LangPath(), lang)
	err = LoadTranslationData(chapterLangPath)
	if err != nil {
		return fmt.Errorf("unable to load chapter translation data: %v", err)
	}
	return nil
}

// LoadTranslationData loads all lang files from
// from directory with specified path.
func LoadTranslationData(path string) error {
	// Translation.
	langData, err := ImportLangDir(path)
	if err != nil {
		return fmt.Errorf("unable to import lang dir: %v", err)
	}
	resData := res.Translations()
	for _, td := range langData {
		resData = append(resData, td)
	}
	res.SetTranslationData(resData)
	return nil
}
