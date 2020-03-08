/*
 * effect.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"os"
	"io/ioutil"
	"strings"
	"path/filepath"
	
	"github.com/isangeles/flame/data/parsexml"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

const (
	EffectsFileExt = ".effects"
)

// ImportEffects imports all XML effects data from effects base
// with specified path.
func ImportEffects(basePath string) ([]*res.EffectData, error) {
	doc, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open effects base file: %v", err)
	}
	defer doc.Close()
	effects, err := parsexml.UnmarshalEffects(doc)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal effects base: %v", err)
	}
	return effects, nil
}

// ImportEffectsDir imports all effects from files in
// specified directory.
func ImportEffectsDir(dirPath string) ([]*res.EffectData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	effects := make([]*res.EffectData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), EffectsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		effs, err := ImportEffects(basePath)
		if err != nil {
			log.Err.Printf("data: effects import: %s: fail to import base: %v",
				basePath, err)
			continue
		}
		for _, e := range effs {
			effects = append(effects, e)
		}
	}
	return effects, nil
}
