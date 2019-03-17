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
	CONFIG_FILE_NAME = ".flame"
)

var (
	langID        = "english" // default eng
	debug         = false
	savegamesPath = "savegames"
	modName       = ""
	modPath       = "data/modules"
	langPath      = "data/lang"
)

// loadConfig loads engine configuration file.
func LoadConfig() error {
	confModVal, err := text.ReadValue(CONFIG_FILE_NAME, "module")
	if err != nil {
		return err
	}
	confValues, err := text.ReadValue(CONFIG_FILE_NAME, "lang", "debug")
	if err != nil {
		SaveConfig() // replace 'corrupted' config with default config
		return fmt.Errorf("fail_to_load_some_conf_values:%v", err)
	}

	langID = confValues["lang"]
	SetDebug(confValues["debug"] == "true")
	// Auto load module.
	modNamePath := strings.Split(confModVal["module"], ";")
	if modNamePath[0] != "" {
		if len(modNamePath) < 2 {
			modName = modNamePath[0]
			modPath = filepath.FromSlash(modPath + "/" + modName)
		} else {
			modPath = modNamePath[1]
			modName = modNamePath[0]
		}
	}
	lang.SetLangPath(langPath + "/" + LangID())

	log.Dbg.Print("config file loaded")
	return nil
}

// saveConfig saves engine configuration in file.
func SaveConfig() error {
	f, err := os.Create(CONFIG_FILE_NAME)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(fmt.Sprintf("%s\n", "# Flame engine configuration file")) // default header
	w.WriteString(fmt.Sprintf("lang:%s;\n", langID))
	w.WriteString(fmt.Sprintf("module:%s;%s;\n", ModuleName(), ModulePath()))
	w.WriteString(fmt.Sprintf("debug:%v;\n", Debug()))
	w.Flush()

	log.Dbg.Print("config file saved")
	//debug.PrintStack()
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
	// TODO check if specified language is supported
	langID = lng
	return nil
}

// SetDebug toggles debug mode.
func SetDebug(dbg bool) {
	debug = dbg
	enginelog.SetDebug(dbg)
}
