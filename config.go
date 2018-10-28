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
	"strconv"
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
	newCharAttrMin = 5
	newCharAttrMax = 15
)

// loadConfig loads engine configuration file.
func LoadConfig() error {
	confValues, err := text.ReadConfigValue(CONFIG_FILE_NAME, "lang", 
		"module", "debug", "new_char_attr_min", "new_char_attr_max")
	if err != nil {
		// Terrible hack(possible loop?).
		SaveConfig()
		return LoadConfig()
		//errlog.Printf("fail_to_load_some_conf_values:%v\n", err)
		//return err
	}

	langID = confValues[0]
	
	modNamePath := strings.Split(confValues[1], ";")
	if modNamePath[0] != "" {
		var m *module.Module
		if len(modNamePath) < 2 {
			m, err = module.NewModule(modNamePath[0], module.DefaultModulesPath())
		} else {
			m, err = module.NewModule(modNamePath[0], modNamePath[1])
		}
		if err != nil {
			return err
		}
		err = SetModule(m)
		if err != nil {
			return err
		}
	}
	enginelog.EnableDebug(confValues[2] == "true")
	attrPtsMin, err := strconv.Atoi(confValues[3])
	if err != nil {
		attrPtsMin = 5
	}
	newCharAttrMin = attrPtsMin
	attrPtsMax, err := strconv.Atoi(confValues[4])
	if err != nil {
		attrPtsMax = 15
	}
	newCharAttrMax = attrPtsMax
	
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
	w.WriteString(fmt.Sprintf("new_char_attr_min:%d;\n", newCharAttrMin))
	w.WriteString(fmt.Sprintf("new_char_attr_max:%d;\n", newCharAttrMax))
	w.Flush()
	
	dbglog.Print("config_file_saved")
	//debug.PrintStack()
	return nil
}

// LangID returns current language ID.
func LangID() string {
	return langID
}

// NewCharAttrMin returns minimal amount of attributes points
// for new character.
func NewCharAttrMin() int {
	return newCharAttrMin
}

// NewCharAttrMax returns maximal amount of attributes points
// for new character.
func NewCharAttrMax() int {
	return newCharAttrMax
}

// SetLang sets language with specified ID as current language.
func SetLang(lng string) error {
	// TODO check if specified language is supported
	langID = lng
	return nil 
}
