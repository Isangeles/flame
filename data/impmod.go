/*
 * impmod.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"path/filepath"

	"github.com/isangeles/flame/data/parsetxt"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module"
)

// ImportModule imports module from directory with specified path.
func ImportModule(path string) (*module.Module, error) {
	// Load module config file.
	confPath := filepath.Join(path, ".module")
	file, err := os.Open(confPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()
	conf := parsetxt.UnmarshalConfig(file)
	conf["path"] = []string{path}
	// Create module.
	data := res.ModuleData{
		Config: conf,
	}
	m := module.New(data)
	err = loadModuleData(m)
	if err != nil {
		return nil, fmt.Errorf("unable to load module data: %v", err)
	}
	err = LoadChapter(m, m.Conf().Chapter)
	if err != nil {
		return nil, fmt.Errorf("unable to load module chapter: %v", err)
	}
	return m, nil
}
