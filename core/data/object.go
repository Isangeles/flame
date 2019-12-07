/*
 * object.go
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
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/effect"
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/log"
)

const (
	ObjectsFileExt = ".objects"
)

// Object creates module area object with specified ID.
func Object(mod *module.Module, obID string) (*object.Object, error) {
	data := res.Object(obID)
	if data == nil {
		return nil, fmt.Errorf("object data not found: %s", obID)
	}
	data.BasicData.HP = data.BasicData.MaxHP
	ob := buildObject(mod, data)
	return ob, nil
}

// ImportObjectsData imports area objects data from base
// file with specified path.
func ImportObjects(path string) ([]*res.ObjectData, error) {
	baseFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fail to open base file: %v", err)
	}
	defer baseFile.Close()
	objects, err := parsexml.UnmarshalObjectsBase(baseFile)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal xml base: %v", err)
	}
	return objects, nil
}

// ImportObjectsDataDir import all area objects data from
// files in directory with specified path.
func ImportObjectsDir(path string) ([]*res.ObjectData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	objects := make([]*res.ObjectData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), ObjectsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(path + "/" + finfo.Name())
		impObjects, err := ImportObjects(basePath)
		if err != nil {
			log.Err.Printf("data: import objects dir: %s: fail to import objects file: %v",
				basePath, err)
			continue
		}
		objects = append(objects, impObjects...)
	}
	return objects, nil
}

// buildObject creates new object from specified data resources.
func buildObject(mod *module.Module, data *res.ObjectData) *object.Object {
	ob := object.NewObject(*data)
	// Effects.
	for _, data := range data.Effects {
		effData := res.Effect(data.ID)
		if effData == nil {
			log.Err.Printf("data: build object: %s: fail to retrieve effect data: %s",
				ob.ID(), data.ID)
			continue
		}
		eff := effect.New(*effData)
		ob.AddEffect(eff)
	}
	return ob
}
