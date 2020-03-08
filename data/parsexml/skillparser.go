/*
 * skillparser.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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

package parsexml

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

// Struct for skills XML base.
type Skills struct {
	XMLName xml.Name `xml:"skills"`
	Skills  []Skill  `xml:"skill"`
}

// Struct for skill XML node.
type Skill struct {
	XMLName     xml.Name       `xml:"skill"`
	ID          string         `xml:"id,attr"`
	CastSec     int            `xml:"cast,attr"`
	CooldownSec int            `xml:"cooldown,attr"`
	Range       string         `xml:"range,attr"`
	Melee       bool           `xml:"melee,attr"`
	Spell       bool           `xml:"spell,attr"`
	Effects     []ObjectEffect `xml:"effects>effect"`
	Reqs        Reqs           `xml:"reqs"`
}

// UnmarshalSkills retrieves skills data from specified XML data.
func UnmarshalSkills(data io.Reader) ([]*res.SkillData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(Skills)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal xml data: %v", err)
	}
	skills := make([]*res.SkillData, 0)
	for _, xmlSkill := range xmlBase.Skills {
		skill, err := buildSkillData(xmlSkill)
		if err != nil {
			log.Err.Printf("xml: unmarshal character: build data unable: %v", err)
			continue
		}
		skills = append(skills, skill)
	}
	return skills, nil
}

// buildSkillData builds skill from XML data.
func buildSkillData(xmlSkill Skill) (*res.SkillData, error) {
	reqs := buildReqs(&xmlSkill.Reqs)
	effects := make([]res.EffectData, 0)
	for _, xmlEffect := range xmlSkill.Effects {
		eff := res.Effect(xmlEffect.ID)
		if eff == nil {
			log.Err.Printf("xml: build skill data: effect data not found: %s",
				xmlEffect.ID)
			continue
		}
		effects = append(effects, *eff)
	}
	skillRange, err := UnmarshalSkillRange(xmlSkill.Range)
	if err != nil {
		return nil, fmt.Errorf("unable to parse range: %v", err)
	}
	data := res.SkillData{
		ID:       xmlSkill.ID,
		Cast:     int64(xmlSkill.CastSec * 1000),
		Cooldown: int64(xmlSkill.CooldownSec * 1000),
		Range:    int(skillRange),
		Melee:    xmlSkill.Melee,
		Spell:    xmlSkill.Spell,
		Effects:  effects,
		UseReqs:  reqs,
	}
	return &data, nil
}
