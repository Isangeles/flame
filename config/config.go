/*
 * config.go
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

package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/core/data/parsetxt"
	"github.com/isangeles/flame/log"
)

const (
	Name, Version  = "Flame Engine", "0.0.0"
	ConfigFileName = ".flame"
)

var (
	Lang   = "english" // default eng
	Debug  = false
	Module = ""
)

// LoadConfig loads engine configuration file.
func LoadConfig() error {
	// Load config file.
	file, err := os.Open(ConfigFileName)
	if err != nil {
		SaveConfig() // save default config
		return fmt.Errorf("unable to open config file: %v", err)
	}
	values := parsetxt.UnmarshalConfig(file)
	// Set values.
	if len(values["lang"]) > 0 {
		Lang = values["lang"][0]
	}
	if len(values["debug"]) > 0 {
		Debug = values["debug"][0] == "true"
	}
	if len(values["module"]) > 0 {
		Module = values["module"][0]
	}
	log.Dbg.Print("Config file loaded")
	return nil
}

// SaveConfig saves engine configuration in file.
func SaveConfig() error {
	// Create file.
	file, err := os.Create(ConfigFileName)
	if err != nil {
		return fmt.Errorf("unable to create conf file: %v", err)
	}
	defer file.Close()
	// Marshal config.
	conf := make(map[string][]string)
	conf["lang"] = []string{Lang}
	conf["module"] = []string{Module}
	conf["debug"] = []string{fmt.Sprintf("%v", Debug)}
	confText := parsetxt.MarshalConfig(conf)
	// Write config text to file.
	w := bufio.NewWriter(file)
	w.WriteString(confText)
	w.Flush()
	log.Dbg.Print("Config file saved")
	return nil
}

// SavegamesPath returns current path
// to savegames directory or errror
// if no module is loaded.
func ModuleSavegamesPath() string {
	return filepath.Join("savegames", Module)
}

// LangPath returns path to current lang
// directory.
func LangPath() string {
	return filepath.Join("data/lang", Lang)
}

// ModulePath returns path to current module.
func ModulePath() string {
	return filepath.Join("data/modules", Module)
}
