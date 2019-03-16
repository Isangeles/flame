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

package main

import (
	"fmt"
	"os"
	"bufio"
	
	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/cmd/log"
)

const (
	CONFIG_FILE_NAME = ".flame-cli"
)

var (
	restrictMode bool = false
)

// loadConfig Loads CLI config file.
func loadConfig() error {
	confValues, err := text.ReadValue(CONFIG_FILE_NAME, "restrict_mode")
	if err != nil {
		return err
	}
	restrictMode = confValues["restrict_mode"] == "true"
	log.Dbg.Println("config file loaded")
	return nil
}

// saveConfig Saves current config values in config file.
func saveConfig() error {
	f, err := os.Create(CONFIG_FILE_NAME)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(fmt.Sprintf("%s\n", "# Flame CLI configuration file")) // default header
	w.WriteString(fmt.Sprintf("restrict_mode:%v;\n", restrictMode))
	w.Flush()
	log.Dbg.Println("config file saved")
	return nil
}
