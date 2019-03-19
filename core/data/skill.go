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
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/skill"
	"github.com/isangeles/flame/core/module/serial"
	"github.com/isangeles/flame/log"
)

const (
	SKILLS_FILE_EXT = ".skills"
)

// Skill creates new instance of skill with specified ID
// for specified module, returns error if skill data with such
// ID was not found or module failed to assing serial value for
// skill.
func Skill(mod *module.Module, id string) (*skill.Skill, error) {
	data := res.Skill(id)
	if data.ID == "" {
		return nil, fmt.Errorf("skill_not_found:%s", id)
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
		return nil, fmt.Errorf("fail_to_open_skills_base_file:%v", err)
	}
	defer doc.Close()
	xmlSkills, err := parsexml.UnmarshalSkillsBase(doc)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_skills_base:%v", err)
	}
	skills := make([]*res.SkillData, 0)
	for _, xmlSkill := range xmlSkills {
		s := buildXMLSkillData(xmlSkill)
		skills = append(skills, s)
	}
	return skills, nil
}

// ImportSkillsDir imports all skills from files in
// specified directory.
func ImportSkillsDir(dirPath string) ([]*res.SkillData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	skills := make([]*res.SkillData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), SKILLS_FILE_EXT) {
			continue
		}
		basePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		impSkills, err := ImportSkills(basePath)
		if err != nil {
			log.Err.Printf("data:skills_import:%s:fail_to_import_base:%v",
				basePath, err)
			continue
		}
		for _, s := range impSkills {
			skills = append(skills, s)
		}
	}
	return skills, nil
}

// buildXMLSkillData builds skill from XML data.
func buildXMLSkillData(xmlSkill parsexml.SkillXML) *res.SkillData {
	reqs := buildXMLReqs(&xmlSkill.Reqs)
	effects := make([]res.EffectData, 0)
	for _, xmlEffect := range xmlSkill.Effects.Nodes {
		eff := res.Effect(xmlEffect.ID)
		if eff == nil {
			log.Err.Printf("data:build_xml_skill_data:effect_data_not_found:%s",
				xmlEffect.ID)
			continue
		}
		effects = append(effects, *eff)
	}
	skillRange, err := parsexml.UnmarshalSkillRange(xmlSkill.Range)
	if err != nil {
		log.Err.Printf("data:build_xml_skill_data:fail_to_parse_range:%v",
			err)
	}
	data := res.SkillData{
		ID:       xmlSkill.ID,
		Cast:     int64(xmlSkill.CastSec * 1000),
		Cooldown: int64(xmlSkill.CooldownSec * 1000),
		Range:    int(skillRange),
		Effects:  effects,
		UseReqs:  reqs,
	}
	return &data
}
