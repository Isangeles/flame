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

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	
	"github.com/isangeles/flame/core/data/text"
)

const (
	CONFIG_FILE_NAME = ".flame-cli"
)

var (
	userTools []string // list with names of perrmited user tools
)

// loadConfig Loads CLI config file.
func loadConfig() error {
	file, err := os.Open(CONFIG_FILE_NAME)
	if err != nil {
		return err
	}
	defer file.Close()
	
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, text.COMMENT_PREFIX) {
			continue
		}
		
		// TODO read user tools
	}
	
	return fmt.Errorf("unsupported_yet")
}

// saveConfig Saves current config values in config file.
func saveConfig() error {
	// TODO config saving
	return fmt.Errorf("unsupported_yet")
}
