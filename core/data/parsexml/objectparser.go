/*
 * objectparser.go
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

package parsexml

import (
	"encoding/xml"

	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/object/skill"
)

// Struct for XML node with object effects.
type ObjectEffectsXML struct {
	XMLName xml.Name          `xml:"effects"`
	Nodes   []ObjectEffectXML `xml:"effect"`
}

// Strcut for object effects XML node.
type ObjectEffectXML struct {
	XMLName xml.Name              `xml:"effect"`
	ID      string                `xml:"id,attr"`
	Serial  string                `xml:"serial,attr"`
	Time    int64                 `xml:"time,attr"`
	Source  ObjectEffectSourceXML `xml:"sourceSerial,attr"`
}

// Struct for object effect source XML node.
type ObjectEffectSourceXML struct {
	XMLName xml.Name `xml:"source"`
	ID      string   `xml:"id,attr"`
	Serial  string   `xml:"serial,attr"`
}

// Struct for object skills XML node.
type ObjectSkillsXML struct {
	XMLName xml.Name         `xml:"skills"`
	Nodes   []ObjectSkillXML `xml:"skill"`
}

// Struct for object skill XML node.
type ObjectSkillXML struct {
	XMLName  xml.Name `xml:"skill"`
	ID       string   `xml:"id,attr"`
	Serial   string   `xml:"serial,attr"`
	Cooldown int64    `xml:"cooldown,attr"`
}

// xmlObjectEffects parses specified effects to XML
// object effects struct.
func xmlObjectEffects(effs ...*effect.Effect) *ObjectEffectsXML {
	xmlEffs := new(ObjectEffectsXML)
	for _, e := range effs {
		xmlEffSource := ObjectEffectSourceXML{}
		if e.Source() != nil {
			xmlEffSource.ID = e.Source().ID()
			xmlEffSource.Serial = e.Source().Serial()
		}
		xmlEff := ObjectEffectXML{
			ID:     e.ID(),
			Serial: e.Serial(),
			Time:   e.Time(),
			Source: xmlEffSource,
		}
		xmlEffs.Nodes = append(xmlEffs.Nodes, xmlEff)
	}
	return xmlEffs
}

// xmlObjectSkills parses specified skills to XML
// object skills struct.
func xmlObjectSkills(skills ...*skill.Skill) *ObjectSkillsXML {
	xmlSkills := new(ObjectSkillsXML)
	for _, s := range skills {
		xmlSkill := ObjectSkillXML{
			ID:       s.ID(),
			Serial:   s.Serial(),
			Cooldown: s.Cooldown(),
		}
		xmlSkills.Nodes = append(xmlSkills.Nodes, xmlSkill)
	}
	return xmlSkills
}
