/*
 * quest.go
 *
 * Copyright 2019-2024 Dariusz Sikora <ds@isangeles.dev>
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
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

// ImportQuests imports all auests from base file with
// specified path.
func ImportQuests(path string) ([]res.QuestData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("unable to open data file: %v", err))
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.QuestsData)
	err = unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON data: %v", err)
	}
	return data.Quests, nil
}

// ImportQuestsDir imports all quests from base files in
// directory with specified path.
func ImportQuestsDir(dirPath string) ([]res.QuestData, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("unable to read dir: %v", err))
	}
	quests := make([]res.QuestData, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.FromSlash(dirPath + "/" + file.Name())
		impQuests, err := ImportQuests(filePath)
		if err != nil {
			log.Err.Printf("data quests import: %s: unable to import base: %v",
				filePath, err)
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
	// Marshal quests data.
	json, err := marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal quests: %v", err)
	}
	// Create quests file.
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
	writer.Write(json)
	writer.Flush()
	return nil
}
