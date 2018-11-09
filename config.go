/*
 * config.go
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

package flame

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	//"runtime/debug"

	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/enginelog"
	"github.com/isangeles/flame/core/module"
)

const (
	CONFIG_FILE_NAME = ".flame"
)

var (
	langID = "english" // default eng
)

// loadConfig loads engine configuration file.
func LoadConfig() error {
	confModVal, err := text.ReadConfigValue(CONFIG_FILE_NAME, "module")
	if err != nil {
		return err
	}
	modNamePath := strings.Split(confModVal[0], ";")
	if modNamePath[0] != "" {
		var m *module.Module
		if len(modNamePath) < 2 {
			m, err = module.NewModule(modNamePath[0],
				module.DefaultModulesPath())
		} else {
			m, err = module.NewModule(modNamePath[0], modNamePath[1])
		}
		if err != nil {
			return err
		}
		SetModule(m)
		if err != nil {
			return err
		}
	}
	confValues, err := text.ReadConfigValue(CONFIG_FILE_NAME, "lang", "debug")
	if err != nil {
		SaveConfig() // replace 'corrupted' config with default config
		return fmt.Errorf("fail_to_load_some_conf_values:%v", err)
	}

	langID = confValues[0]
	enginelog.EnableDebug(confValues[1] == "true")
	
	dbglog.Print("config_file_loaded")
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
	if mod != nil {
		w.WriteString(fmt.Sprintf("module:%s;%s;\n", mod.Name(), mod.Path()))
	} else {
		w.WriteString(fmt.Sprintf("module:;;\n"))
	}
	w.WriteString(fmt.Sprintf("debug:%v;\n", enginelog.IsDebug()))
	w.Flush()
	
	dbglog.Print("config_file_saved")
	//debug.PrintStack()
	return nil
}

// LangID returns current language ID.
func LangID() string {
	return langID
}

// SetLang sets language with specified ID as current language.
func SetLang(lng string) error {
	// TODO check if specified language is supported
	langID = lng
	return nil 
}

// SetDebug toggles debug mode.
func SetDebug(dbg bool) {
	enginelog.EnableDebug(dbg)
}
