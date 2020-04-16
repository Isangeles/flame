/*
 * object.go
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
	ObjectsFileExt = ".objects"
)

// ImportObjectsData imports area objects data from base
// file with specified path.
func ImportObjects(path string) ([]res.ObjectData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open data file: %v", err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.ObjectsData)
	err = xml.Unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal XML: %v", err)
	}
	return data.Objects, nil
}

// ImportObjectsDataDir import all area objects data from
// files in directory with specified path.
func ImportObjectsDir(path string) ([]res.ObjectData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	objects := make([]res.ObjectData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), ObjectsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(path + "/" + finfo.Name())
		impObjects, err := ImportObjects(basePath)
		if err != nil {
			log.Err.Printf("data: import objects dir: %s: unable to import objects file: %v",
				basePath, err)
			continue
		}
		objects = append(objects, impObjects...)
	}
	return objects, nil
}
