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
	"github.com/isangeles/flame/core/module/object/area"
	"github.com/isangeles/flame/log"
)

const (
	OBJECTS_FILE_EXT = ".objects"
)

// Object creates module area object with specified ID.
func Object(mod *module.Module, obID string) (*area.Object, error) {
	data := res.Object(obID)
	if data == nil {
		return nil, fmt.Errorf("object_data_not_found:%s", obID)
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
		return nil, fmt.Errorf("fail_to_open_base_file:%v", err)
	}
	defer baseFile.Close()
	objects, err := parsexml.UnmarshalObjectsBase(baseFile)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_base:%v", err)
	}
	return objects, nil
}

// ImportObjectsDataDir import all area objects data from
// files in directory with specified path.
func ImportObjectsDir(path string) ([]*res.ObjectData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	objects := make([]*res.ObjectData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), OBJECTS_FILE_EXT) {
			continue
		}
		basePath := filepath.FromSlash(path + "/" + finfo.Name())
		impObjects, err := ImportObjects(basePath)
		if err != nil {
			log.Err.Printf("data:import_objects_dir:%s:fail_to_import_objects_file:%v",
				basePath, err)
			continue
		}
		objects = append(objects, impObjects...)
	}
	return objects, nil
}

// buildObject creates new object from specified data resources.
func buildObject(mod *module.Module, data *res.ObjectData) *area.Object {
	ob := area.NewObject(data.BasicData)
	// Inventory.
	for _, data := range data.Items {
		it, err := Item(mod, data.ID)
		if err != nil {
			log.Err.Printf("data:build_object:%s:fail_to_retrieve_inv_item:%s",
				ob.ID(), data.ID)
			continue
		}
		it.SetSerial(data.Serial)
		ob.Inventory().AddItem(it)
	}
	return ob
}
