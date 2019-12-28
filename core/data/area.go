/*
 * area.go
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
	
	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/log"
)

const (
	AreaFileExt = ".area"
)

// ImportArea imports area from area dir with specified path.
func ImportArea(path string) (*res.AreaData, error) {
	mainFilePath := filepath.FromSlash(fmt.Sprintf("%s/main%s",
		path, AreaFileExt))
	// Open area file.
	file, err := os.Open(mainFilePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open area file: %v", err)
	}
	defer file.Close()
	// Unmarshal area file.
	areaData, err := parsexml.UnmarshalArea(file)
	if err != nil {
		return nil, fmt.Errorf("fail to parse area data: %v", err)
	}
	return areaData, nil
}

// ImportAreaDir imports all areas from directory with specified path.
func ImportAreasDir(path string) ([]*res.AreaData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	areas := make([]*res.AreaData, 0)
	for _, file := range files {
		areaPath := filepath.FromSlash(path + "/" + file.Name())
		area, err := ImportArea(areaPath)
		if err != nil {
			log.Err.Printf("data: areas import: %s: fail to import area: %v",
				areaPath, err)
			continue
		}
		areas = append(areas, area)
	}
	return areas, nil
}
