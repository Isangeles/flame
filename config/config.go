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
	"github.com/isangeles/flame/core/enginelog"
	"github.com/isangeles/flame/log"
)

const (
	ConfigFileName = ".flame"
)

var (
	langID        = "english" // default eng
	debug         = false
	savegamesPath = "savegames"
	modName       = ""
	modPath       = "data/modules"
	langPath      = "data/lang"
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
		langID = values["lang"][0]
	}
	if len(values["debug"]) > 0 {
		SetDebug(values["debug"][0] == "true")
	}
	if len(values["module"]) > 1 {
		modName = values["module"][0]
		modPath = values["module"][1]
	} else if len(values["module"]) > 0 {
		modName = values["module"][0]
		modPath = filepath.Join(modPath, modName)
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
	conf["lang"] = []string{langID}
	conf["module"] = []string{ModuleName(), ModulePath()}
	conf["debug"] = []string{fmt.Sprintf("%v", Debug())}
	confText := parsetxt.MarshalConfig(conf)
	// Write config text to file.
	w := bufio.NewWriter(file)
	w.WriteString(confText)
	w.Flush()
	log.Dbg.Print("Config file saved")
	return nil
}

// LangID returns current language ID.
func LangID() string {
	return langID
}

// Debug checks whether debug mode is
// enabled.
func Debug() bool {
	return debug
}

// SavegamesPath returns current path
// to savegames directory or errror
// if no module is loaded.
func ModuleSavegamesPath() string {
	return filepath.FromSlash(savegamesPath + "/" + ModuleName())
}

// ModulePath returns path to modules directory.
func ModulePath() string {
	return modPath
}

// ModuleName returns module name from config.
func ModuleName() string {
	return modName
}

// SetLang sets language with specified ID as current language.
func SetLang(lng string) error {
	// TODO: check if specified language is supported.
	langID = lng
	return nil
}

// LangPath returns path to current lang
// directory.
func LangPath() string {
	return filepath.FromSlash(langPath + "/" + langID)
}

// SetDebug toggles debug mode.
func SetDebug(dbg bool) {
	debug = dbg
	enginelog.SetDebug(dbg)
}
