/*
 * skill.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/skill"
	"github.com/isangeles/flame/core/module/serial"
	"github.com/isangeles/flame/log"
)

const (
	SkillsFileExt = ".skills"
)

// Skill creates new instance of skill with specified ID
// for specified module, returns error if skill data with such
// ID was not found or module failed to assing serial value for
// skill.
func Skill(id string) (*skill.Skill, error) {
	data := res.Skill(id)
	if data == nil {
		return nil, fmt.Errorf("skill not found: %s", id)
	}
	s := skill.New(*data)
	serial.AssignSerial(s)
	return s, nil
}

// ImportSkills imports all XML skills data from skills base
// with specified path.
func ImportSkills(basePath string) ([]*res.SkillData, error) {
	doc, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open skills base file: %v", err)
	}
	defer doc.Close()
	skills, err := parsexml.UnmarshalSkills(doc)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal skills base: %v", err)
	}
	return skills, nil
}

// ImportSkillsDir imports all skills from files in
// specified directory.
func ImportSkillsDir(dirPath string) ([]*res.SkillData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	skills := make([]*res.SkillData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), SkillsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		impSkills, err := ImportSkills(basePath)
		if err != nil {
			log.Err.Printf("data: skills import: %s: fail to import base: %v",
				basePath, err)
			continue
		}
		for _, s := range impSkills {
			skills = append(skills, s)
		}
	}
	return skills, nil
}
