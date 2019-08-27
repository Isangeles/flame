/*
 * quest.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/quest"
	"github.com/isangeles/flame/log"
)

const (
	QuestsFileExt = ".quests"
)

// Quest returns new quest with specified ID or
// error if data was not found or quest creation
// failed.
func Quest(id string) (*quest.Quest, error) {
	data := res.Quest(id)
	if data == nil {
		return nil, fmt.Errorf("res not found: %s", id)
	}
	q := quest.New(*data)
	return q, nil
}

// ImportQuests imports all auests from base file with
// specified path.
func ImportQuests(basePath string) ([]*res.QuestData, error) {
	doc, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open base file: %v", err)
	}
	defer doc.Close()
	quests, err := parsexml.UnmarshalQuestsBase(doc)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal quests base: %v", err)
	}
	return quests, nil
}

// ImportQuestsDir imports all quests from base files in
// directory with specified path.
func ImportQuestsDir(dirPath string) ([]*res.QuestData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	quests := make([]*res.QuestData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), QuestsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		impQuests, err := ImportQuests(basePath)
		if err != nil {
			log.Err.Printf("data quests import: %s: fail to import base: %v",
				basePath, err)
		}
		for _, q := range impQuests {
			quests = append(quests, q)
		}
	}
	return quests, nil
}
