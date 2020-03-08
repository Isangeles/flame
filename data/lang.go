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

package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/parsetxt"
	"github.com/isangeles/flame/log"
)

const (
	LangFileExt = ".lang"
)

// ImportLang imports all translation data from file with
// specified path.
func ImportLang(path string) ([]*res.TranslationData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fail to open file: %v", err)
	}
	defer file.Close()
	data := parsetxt.UnmarshalLangData(file)
	return data, nil
}

// ImportLangDir imports all translation data from lang
// files in directory with specified path.
func ImportLangDir(path string) ([]*res.TranslationData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	data := make([]*res.TranslationData, 0)
	for _, fileInfo := range files {
		if !strings.HasSuffix(fileInfo.Name(), LangFileExt) {
			continue
		}
		filePath := filepath.FromSlash(path + "/" + fileInfo.Name())
		impData, err := ImportLang(filePath)
		if err != nil {
			log.Err.Printf("data: import lang dir: %s: fail to import lang file: %v",
				filePath, err)
			continue
		}
		data = append(data, impData...)
	}
	return data, nil
}
