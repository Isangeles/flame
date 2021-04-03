/*
 * quest.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

const (
	QuestsFileExt = ".quests"
)

// ImportQuests imports all auests from base file with
// specified path.
func ImportQuests(path string) ([]res.QuestData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open data file: %v", err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.QuestsData)
	err = xml.Unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal XML data: %v", err)
	}
	return data.Quests, nil
}

// ImportQuestsDir imports all quests from base files in
// directory with specified path.
func ImportQuestsDir(dirPath string) ([]res.QuestData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	quests := make([]res.QuestData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), QuestsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		impQuests, err := ImportQuests(basePath)
		if err != nil {
			log.Err.Printf("data quests import: %s: unable to import base: %v",
				basePath, err)
		}
		for _, q := range impQuests {
			quests = append(quests, q)
		}
	}
	return quests, nil
}

// ExportQuests exports effects to the data file under specified path.
func ExportQuests(path string, quests ...res.QuestData) error {
	data := new(res.QuestsData)
	for _, q := range quests {
		data.Quests = append(data.Quests, q)
	}
	// Marshal races data.
	xml, err := xml.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal quests: %v", err)
	}
	// Create races file.
	if !strings.HasSuffix(path, QuestsFileExt) {
		path += QuestsFileExt
	}
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create quests file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create quests file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	writer := bufio.NewWriter(file)
	writer.Write(xml)
	writer.Flush()
	return nil
}
