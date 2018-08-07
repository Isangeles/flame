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

package core

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	
	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/enginelog"
)

const (
	CONFIG_FILE_NAME = ".flame"
)

var (
	langID string = "english" // default eng
)

// loadConfig loads engine configuration file.
func LoadConfig() error {
	confValues, err := text.ReadConfigValue(CONFIG_FILE_NAME, "lang", 
							"module")
	if err != nil {
		return err
	}
	
	langID = confValues[0]
	
	modNamePath := strings.Split(confValues[1], ";")
	if modNamePath[0] != "" {
		var m module.Module
		if len(modNamePath) < 2 {
			m, err = module.NewModule(modNamePath[0], module.DefaultModulesPath())
		} else {
			m, err = module.NewModule(modNamePath[0], modNamePath[1])
		}
		if err != nil {
			enginelog.Error(fmt.Sprintf("config_module_load_fail:", err))
		}
		err = SetModule(m)
		if err != nil {
			enginelog.Error(fmt.Sprintf("config_module_load_fail:", err))
		}
	}
	
	enginelog.Info("config_file_loaded")
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
	w.WriteString(fmt.Sprintf("%s\n", "#Flame engine config file")) // default header
	w.WriteString(fmt.Sprintf("lang:%s;\n", langID))
	w.WriteString(fmt.Sprintf("module:%s;%s;\n", mod.Name(), mod.Path()))
	
	w.Flush()
	
	enginelog.Info("config_file_saved")
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
