/*
 * data.go
 * 
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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

// data package provides connection with external data files like items
// base, savegames, etc.
package data

import (
	"fmt"
	"path/filepath"
	"os"

	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/module"
)

// LoadMod loads module from specified path.
func LoadMod(name, path string) (*module.Module, error) {
	if _, err := os.Stat(path + string(os.PathSeparator) + name);
	os.IsNotExist(err) {
		return nil, fmt.Errorf("module_not_found:'%s' in:'%s'", name, path)
	}
	modConfPath := filepath.FromSlash(path + "/" + name + "/mod.conf")
	confValues, err := text.ReadConfigInt(modConfPath, "new_char_attrs_min",
		"new_char_attrs_max")
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_module_conf:%s", err)
	}
	conf := module.Conf{
		Name:name,
		Path:path,
		NewcharAttrsMin:confValues[0],
		NewcharAttrsMax:confValues[1],
	}
	chaps, err := make([]*module.Chapter, 1), nil//fmt.Errorf("unsupported_yet")
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_module_chapters:%v", err)
	}
	return module.NewModule(conf, chaps), nil
}
