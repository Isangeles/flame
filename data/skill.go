/*
 * skill.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

const (
	SkillsFileExt = ".skills"
)

// ImportSkills imports all XML skills data from skills base
// with specified path.
func ImportSkills(path string) ([]res.SkillData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open data file: %v", err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.SkillsData)
	err = xml.Unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal XML data: %v", err)
	}
	return data.Skills, nil
}

// ImportSkillsDir imports all skills from files in
// specified directory.
func ImportSkillsDir(dirPath string) ([]res.SkillData, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	skills := make([]res.SkillData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), SkillsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		impSkills, err := ImportSkills(basePath)
		if err != nil {
			log.Err.Printf("data: skills import: %s: unable to import base: %v",
				basePath, err)
			continue
		}
		for _, s := range impSkills {
			skills = append(skills, s)
		}
	}
	return skills, nil
}

// ExportSkills exports skills to data file under specified path.
func ExportSkills(path string, skills ...res.SkillData) error {
	data := new(res.SkillsData)
	for _, s := range skills {
		data.Skills = append(data.Skills, s)
	}
	// Marshal races data.
	xml, err := xml.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal skills: %v", err)
	}
	// Create races file.
	if !strings.HasSuffix(path, SkillsFileExt) {
		path += SkillsFileExt
	}
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create skills file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create skills file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	writer := bufio.NewWriter(file)
	writer.Write(xml)
	writer.Flush()
	return nil
}
