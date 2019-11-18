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
 *w
 */

package parsexml

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/area"
	"github.com/isangeles/flame/core/module/craft"
	"github.com/isangeles/flame/core/module/dialog"
	"github.com/isangeles/flame/core/module/effect"
	"github.com/isangeles/flame/core/module/skill"
	"github.com/isangeles/flame/log"
)

// Struct for XML objects base node.
type Objects struct {
	XMLName xml.Name `xml:"objects"`
	Nodes   []Object `xml:"object"`
}

// Struct for XML object node.
type Object struct {
	XMLName   xml.Name       `xml:"object"`
	ID        string         `xml:"id,attr"`
	Serial    string         `xml:"serial,attr"`
	HP        int            `xml:"hp,attr"`
	MaxHP     int            `xml:"max-hp,attr"`
	Position  string         `xml:"position,value"`
	Action    ObjectAction   `xml:"action"`
	Inventory Inventory      `xml:"inventory"`
	Effects   []ObjectEffect `xml:"effects>effect"`
}

// Strcut for object effects XML node.
type ObjectEffect struct {
	XMLName xml.Name           `xml:"effect"`
	ID      string             `xml:"id,attr"`
	Serial  string             `xml:"serial,attr"`
	Time    int64              `xml:"time,attr"`
	Source  ObjectEffectSource `xml:"source"`
}

// Struct for object effect source XML node.
type ObjectEffectSource struct {
	XMLName xml.Name `xml:"source"`
	ID      string   `xml:"id,attr"`
	Serial  string   `xml:"serial,attr"`
}

// Struct for object skills XML node.
type ObjectSkills struct {
	XMLName xml.Name      `xml:"skills"`
	Nodes   []ObjectSkill `xml:"skill"`
}

// Struct for object skill XML node.
type ObjectSkill struct {
	XMLName  xml.Name `xml:"skill"`
	ID       string   `xml:"id,attr"`
	Serial   string   `xml:"serial,attr"`
	Cooldown int64    `xml:"cooldown,attr"`
}

// Struct for object dialogs XML node.
type ObjectDialogs struct {
	XMLName xml.Name       `xml:"dialogs"`
	Nodes   []ObjectDialog `xml:"dialog"`
}

// Struct for object dialog XML node.
type ObjectDialog struct {
	XMLName xml.Name `xml:"dialog"`
	ID      string   `xml:"id,attr"`
}

// Struct for object recipe XML node.
type ObjectRecipe struct {
	XMLName xml.Name `xml:"recipe"`
	ID      string   `xml:"id,attr"`
}

// Struct for object action XML node.
type ObjectAction struct {
	XMLName  xml.Name  `xml:"action"`
	SelfMods Modifiers `xml:"self>modifiers"`
	UserMods Modifiers `xml:"user>modifiers"`
}

// UnmarshalObjectsBaseXML parses specified data to XML
// object nodes.
func UnmarshalObjectsBase(data io.Reader) ([]*res.ObjectData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(Objects)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal xml data: %v", err)
	}
	objects := make([]*res.ObjectData, 0)
	for _, xmlObject := range xmlBase.Nodes {
		object, err := buildObjectData(&xmlObject)
		if err != nil {
			log.Err.Printf("xml: unmarshal object: %s: build data fail: %v",
				xmlObject.ID, err)
			continue
		}
		objects = append(objects, object)
	}
	return objects, nil
}

// xmlObject parses specified area object to XML
// object struct.
func xmlObject(ob *area.Object) *Object {
	xmlOb := new(Object)
	xmlOb.ID = ob.ID()
	xmlOb.Serial = ob.Serial()
	xmlOb.HP = ob.Health()
	xmlOb.MaxHP = ob.MaxHealth()
	posX, posY := ob.Position()
	xmlOb.Position = fmt.Sprintf("%fx%f", posX, posY)
	xmlOb.Action = xmlObjectAction(ob.Action())
	xmlOb.Inventory = *xmlInventory(ob.Inventory())
	xmlOb.Effects = xmlObjectEffects(ob.Effects()...)
	return xmlOb
}

// xmlObjectEffects parses specified effects to XML
// object effects struct.
func xmlObjectEffects(effs ...*effect.Effect) []ObjectEffect {
	var xmlEffs []ObjectEffect
	for _, e := range effs {
		xmlEffSource := ObjectEffectSource{}
		if e.Source() != nil {
			xmlEffSource.ID = e.Source().ID()
			xmlEffSource.Serial = e.Source().Serial()
		}
		xmlEff := ObjectEffect{
			ID:     e.ID(),
			Serial: e.Serial(),
			Time:   e.Time(),
			Source: xmlEffSource,
		}
		xmlEffs = append(xmlEffs, xmlEff)
	}
	return xmlEffs
}

// xmlObjectSkills parses specified skills to XML
// object skills struct.
func xmlObjectSkills(skills ...*skill.Skill) *ObjectSkills {
	xmlSkills := new(ObjectSkills)
	for _, s := range skills {
		xmlSkill := ObjectSkill{
			ID:       s.ID(),
			Serial:   s.Serial(),
			Cooldown: s.Cooldown(),
		}
		xmlSkills.Nodes = append(xmlSkills.Nodes, xmlSkill)
	}
	return xmlSkills
}

// xmlObjectDialogs parses specified dialogs to XML
// object dialogs struct.
func xmlObjectDialogs(dialogs ...*dialog.Dialog) *ObjectDialogs {
	xmlDialogs := new(ObjectDialogs)
	for _, d := range dialogs {
		xmlDialog := ObjectDialog{
			ID: d.ID(),
		}
		xmlDialogs.Nodes = append(xmlDialogs.Nodes, xmlDialog)
	}
	return xmlDialogs
}

// xmlObjectRecipes parses specified recipes to XML nodes.
func xmlObjectRecipes(recipes ...*craft.Recipe) []ObjectRecipe {
	xmlRecipes := make([]ObjectRecipe, 0)
	for _, r := range recipes {
		xmlRecipe := ObjectRecipe{
			ID: r.ID(),
		}
		xmlRecipes = append(xmlRecipes, xmlRecipe)
	}
	return xmlRecipes
}

// xmlObjectAction parses specified object action to XML node.
func xmlObjectAction(action area.ObjectAction) ObjectAction {
	var xmlAction ObjectAction
	xmlAction.SelfMods = xmlModifiers(action.SelfMods...)
	xmlAction.UserMods = xmlModifiers(action.UserMods...)
	return xmlAction
}

// buildObjectData creates object data from specified XML
// data.
func buildObjectData(xmlOb *Object) (*res.ObjectData, error) {
	// Basic data.
	baseData := res.ObjectBasicData{
		ID:     xmlOb.ID,
		Serial: xmlOb.Serial,
		HP:     xmlOb.HP,
		MaxHP:  xmlOb.MaxHP,
	}
	// Action.
	baseData.Action.SelfMods = buildModifiers(&xmlOb.Action.SelfMods)
	baseData.Action.UserMods = buildModifiers(&xmlOb.Action.UserMods)
	data := res.ObjectData{BasicData: baseData}
	// Position.
	if xmlOb.Position != "" {
		posX, posY, err := UnmarshalPosition(xmlOb.Position)
		if err != nil {
			return nil, fmt.Errorf("fail to parse position: %v", err)
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
	for _, xmlEffect := range xmlOb.Effects {
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
