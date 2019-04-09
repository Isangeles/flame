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
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/area"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/object/skill"
	"github.com/isangeles/flame/log"
)

// Struct for XML objects base node.
type ObjectsBaseXML struct {
	XMLName xml.Name    `xml:"base"`
	Nodes   []ObjectXML `xml:"object"`
}

// Struct for XML object node.
type ObjectXML struct {
	XMLName   xml.Name         `xml:"object"`
	ID        string           `xml:"id,attr"`
	Serial    string           `xml:"serial,attr"`
	HP        int              `xml:"hp,attr"`
	MaxHP     int              `xml:"max-hp,attr"`
	Position  string           `xml:"position,value"`
	Inventory InventoryXML     `xml:"inventory"`
	Effects   ObjectEffectsXML `xml:"effects"`
}

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
	Source  ObjectEffectSourceXML `xml:"source"`
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

// UnmarshalObjectsBaseXML parses specified data to XML
// object nodes.
func UnmarshalObjectsBase(data io.Reader) ([]*res.ObjectData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(ObjectsBaseXML)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v",
			err)
	}
	objects := make([]*res.ObjectData, 0)
	for _, xmlObject := range xmlBase.Nodes {
		object, err := buildObjectData(&xmlObject)
		if err != nil {
			log.Err.Printf("xml:unmarshal_object:%s:build_data_fail:%v",
				xmlObject.ID, err)
			continue
		}
		objects = append(objects, object)
	}
	return objects, nil
}

// xmlObject parses specified area object to XML
// object struct.
func xmlObject(ob *area.Object) *ObjectXML {
	xmlOb := new(ObjectXML)
	xmlOb.ID = ob.ID()
	xmlOb.Serial = ob.Serial()
	xmlOb.HP = ob.Health()
	xmlOb.MaxHP = ob.MaxHealth()
	posX, posY := ob.Position()
	xmlOb.Position = fmt.Sprintf("%fx%f", posX, posY)
	xmlOb.Inventory = *xmlInventory(ob.Inventory())
	xmlOb.Effects = *xmlObjectEffects(ob.Effects()...)
	return xmlOb
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

// buildObjectData creates object data from specified XML
// data.
func buildObjectData(xmlOb *ObjectXML) (*res.ObjectData, error) {
	// Basic data.
	baseData := res.ObjectBasicData{
		ID:     xmlOb.ID,
		Serial: xmlOb.Serial,
		HP:     xmlOb.HP,
		MaxHP:  xmlOb.MaxHP,
	}
	data := res.ObjectData{BasicData: baseData}
	// Position.
	if xmlOb.Position != "" {
		posX, posY, err := UnmarshalPosition(xmlOb.Position)
		if err != nil {
			return nil, fmt.Errorf("fail_to_parse_position:%v", err)
		}
		data.SavedData.PosX, data.SavedData.PosY = posX, posY
	}
	// Items.
	for _, xmlIt := range xmlOb.Inventory.Items {
		itData := res.InventoryItemData{
			ID:     xmlIt.ID,
			Serial: xmlIt.Serial,
		}
		data.Items = append(data.Items, itData)
	}
	// Effects.
	for _, xmlEffect := range xmlOb.Effects.Nodes {
		effectData := res.ObjectEffectData{
			ID:           xmlEffect.ID,
			Serial:       xmlEffect.Serial,
			Time:         xmlEffect.Time,
			SourceID:     xmlEffect.Source.ID,
			SourceSerial: xmlEffect.Source.Serial,
		}
		data.Effects = append(data.Effects, effectData)
	}
	return &data, nil
}
