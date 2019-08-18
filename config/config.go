/*
 * config.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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
	"strings"

	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/data/text/lang"
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
	// Retrieve values from conf file.
	values, err := text.ReadValue(ConfigFileName, "module", "lang", "debug")
	if err != nil {
		SaveConfig() // replace 'corrupted' config with default config
		return fmt.Errorf("fail to load conf values: %v", err)
	}
	// Set values.
	langID = values["lang"]
	SetDebug(values["debug"] == "true")
	modNamePath := strings.Split(values["module"], ";")
	if modNamePath[0] != "" {
		if len(modNamePath) < 2 {
			modName = modNamePath[0]
			modPath = filepath.FromSlash(modPath + "/" + modName)
		} else {
			modPath = modNamePath[1]
			modName = modNamePath[0]
		}
	}
	lang.SetLangPath(LangPath())
	log.Dbg.Print("config file loaded")
	return nil
}

// SaveConfig saves engine configuration in file.
func SaveConfig() error {
	// Create file.
	f, err := os.Create(ConfigFileName)
	if err != nil {
		return fmt.Errorf("fail to create conf file: %v", err)
	}
	defer f.Close()
	// Write values.
	w := bufio.NewWriter(f)
	w.WriteString(fmt.Sprintf("%s\n", "# Flame engine configuration file")) // default header
	w.WriteString(fmt.Sprintf("lang:%s\n", langID))
	w.WriteString(fmt.Sprintf("module:%s;%s\n", ModuleName(), ModulePath()))
	w.WriteString(fmt.Sprintf("debug:%v\n", Debug()))
	// Save.
	w.Flush()
	log.Dbg.Print("config file saved")
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
