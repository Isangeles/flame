/*
 * area.go
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
	"io"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

const (
	AreaFileExt = ".area"
)

// ImportArea imports area from area dir with specified path.
func ImportArea(path string) (res.AreaData, error) {
	mainFilePath := filepath.FromSlash(fmt.Sprintf("%s/main%s",
		path, AreaFileExt))
	// Open area file.
	file, err := os.Open(mainFilePath)
	if err != nil {
		return res.AreaData{}, fmt.Errorf("unable to open area file: %v", err)
	}
	defer file.Close()
	// Unmarshal area file.
	buf, err := io.ReadAll(file)
	if err != nil {
		return res.AreaData{}, fmt.Errorf("unable to read area file: %v", err)
	}
	data := res.AreaData{}
	err = xml.Unmarshal(buf, &data)
	if err != nil {
		return data, fmt.Errorf("unable to unmarshal XML data: %v", err)
	}
	return data, nil
}

// ImportAreaDir imports all areas from directory with specified path.
func ImportAreasDir(path string) ([]res.AreaData, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	areas := make([]res.AreaData, 0)
	for _, file := range files {
		areaPath := filepath.FromSlash(path + "/" + file.Name())
		area, err := ImportArea(areaPath)
		if err != nil {
			log.Err.Printf("data: areas import: %s: unable to import area: %v",
				areaPath, err)
			continue
		}
		areas = append(areas, area)
	}
	return areas, nil
}

// ExportArea exports area to a new directory under specified
// path.
func ExportArea(path string, data res.AreaData) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("unable to create area dir: %v", err)
	}
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal area data: %v", err)
	}
	areaFilePath := filepath.Join(path, "main"+AreaFileExt)
	areaFile, err := os.Create(areaFilePath)
	if err != nil {
		return fmt.Errorf("unable to create area file: %v", err)
	}
	defer areaFile.Close()
	w := bufio.NewWriter(areaFile)
	w.Write(xmlData)
	w.Flush()
	return nil
}
