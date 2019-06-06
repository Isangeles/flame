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
	"strings"

	"github.com/isangeles/flame/cmd/log"
	"github.com/isangeles/flame/core/data/text"
)

const (
	CONFIG_FILE_NAME = ".flame-cli"
)

var (
	restrictMode    = false
	newCharAttrsPts = 10
	newCharSkills   []string
	newCharItems    []string
	scriptsDir      = "data/scripts"
)

// LoadConfig Loads CLI config file.
func LoadConfig() error {
	// Read config file.
	values, err := text.ReadValue(CONFIG_FILE_NAME, "restrict-mode",
		"new-char-attrs", "new-char-skills", "new-char-items")
	if err != nil {
		return fmt.Errorf("fail_to_read_values:%v", err)
	}
	// Set values.
	restrictMode = values["restrict_mode"] == "true"
	for _, sid := range strings.Split(values["new-char-skills"], ";") {
		newCharSkills = append(newCharSkills, sid)
	}
	for _, iid := range strings.Split(values["new-char-items"], ";") {
		newCharItems = append(newCharItems, iid)
	}
	log.Dbg.Println("config file loaded")
	return nil
}

// SaveConfig Saves current config values in config file.
func SaveConfig() error {
	// Create file.
	f, err := os.Create(CONFIG_FILE_NAME)
	if err != nil {
		return err
	}
	defer f.Close()
	// Write values.
	w := bufio.NewWriter(f)
	w.WriteString(fmt.Sprintf("%s\n", "# Flame CLI configuration file")) // default header
	w.WriteString(fmt.Sprintf("restrict-mode:%v\n", restrictMode))
	w.WriteString(fmt.Sprintf("new-char-attrs:%d\n", newCharAttrsPts))
	w.WriteString("new-char-skills:")
	for _, sid := range newCharSkills {
		w.WriteString(sid + ";")
	}
	w.WriteString("\n")
	w.WriteString("new-char-items:")
	for _, iid := range newCharItems {
		w.WriteString(iid + ";")
	}
	w.WriteString("\n")
	// Save.
	w.Flush()
	log.Dbg.Println("config file saved")
	return nil
}

// NewCharAttrs returns amount of attributes
// points for new character.
func NewCharAttrs() int {
	return newCharAttrsPts
}

// NewCharSkills returns IDs of skills
// for new character.
func NewCharSkills() (ids []string) {
	for _, id := range newCharSkills {
		if len(id) < 1 {
			continue
		}
		ids = append(ids, id)
	}
	return
}

// NewCharItems retuns IDs of items
// for new character.
func NewCharItems() (ids []string) {
	for _, id := range newCharItems {
		if len(id) < 1 {
			continue
		}
		ids = append(ids, id)
	}
	return
}

// ScriptsPath returns path to
// scripts directory.
func ScriptsPath() string {
	return scriptsDir
}
