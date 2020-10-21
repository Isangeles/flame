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
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/text"
	"github.com/isangeles/flame/log"
)

const (
	LangFileExt = ".lang"
)

// ImportLang imports all translation data from file with
// specified path.
func ImportLang(path string) ([]res.TranslationData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()
	data, err := text.UnmarshalLangData(file)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal lang file: %v", err)
	}
	return data, nil
}

// ImportLangDir imports all translation data from lang
// files in directory with specified path.
func ImportLangDir(path string) ([]res.TranslationData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	data := make([]res.TranslationData, 0)
	for _, fileInfo := range files {
		if !strings.HasSuffix(fileInfo.Name(), LangFileExt) {
			continue
		}
		filePath := filepath.FromSlash(path + "/" + fileInfo.Name())
		impData, err := ImportLang(filePath)
		if err != nil {
			log.Err.Printf("data: import lang dir: %s: unable to import lang file: %v",
				filePath, err)
			continue
		}
		data = append(data, impData...)
	}
	return data, nil
}

// ImportLangDirs imports all translation data from child directories
// of the directory with a specified path.
func ImportLangDirs(path string) ([]res.TranslationBaseData, error) {
	langDirs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to read lang directory: %v", err)
	}
	data := make([]res.TranslationBaseData, 0)
	for _, langDir := range langDirs {
		langDirPath := filepath.Join(path, langDir.Name())
		langDirData, err := ImportLangDir(langDirPath)
		if err != nil {
			log.Err.Printf("Import lang dirs: unable to import dir: %v", err)
			continue
		}
		base := res.TranslationBaseData{
			ID:           langDir.Name(),
			Translations: langDirData,
		}
		data = append(data, base)
	}
	return data, nil
}

// ExportLang exports translation data to a new file with specified path.
func ExportLang(path string, data ...res.TranslationData) error {
	txt := text.MarshalLangData(data)
	if !strings.HasSuffix(path, LangFileExt) {
		path += LangFileExt
	}
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("unable to create lang file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create lang file: %v", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(txt)
	writer.Flush()
	return nil
}
