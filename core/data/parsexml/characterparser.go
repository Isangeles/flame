/*
 * characterparser.go
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

package parsexml

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/character"
	"github.com/isangeles/flame/core/module/flag"
	"github.com/isangeles/flame/core/module/quest"
	"github.com/isangeles/flame/log"
)

// Struct for XML characters base.
type Characters struct {
	XMLName    xml.Name    `xml:"characters"`
	Characters []Character `xml:"char"`
}

// Struct for XML character node.
type Character struct {
	XMLName     xml.Name         `xml:"char"`
	ID          string           `xml:"id,attr"`
	Serial      string           `xml:"serial,attr"`
	Name        string           `xml:"name,attr"`
	Gender      string           `xml:"gender,attr"`
	Race        string           `xml:"race,attr"`
	Attitude    string           `xml:"attitude,attr"`
	Alignment   string           `xml:"alignment,attr"`
	Guild       string           `xml:"guild,attr"`
	Level       int              `xml:"level,attr"`
	Attributes  Attributes       `xml:"attributes"`
	PC          bool             `xml:"pc,attr"`
	HP          int              `xml:"hp,attr"`
	Mana        int              `xml:"mana,attr"`
	Exp         int              `xml:"exp,attr"`
	Position    string           `xml:"position,value"`
	DefPosition string           `xml:"default-position,value"`
	Inventory   Inventory        `xml:"inventory"`
	Equipment   Equipment        `xml:"equipment"`
	Effects     []ObjectEffect   `xml:"effects>effect"`
	Skills      ObjectSkills     `xml:"skills"`
	Memory      Memory           `xml:"memory"`
	Dialogs     ObjectDialogs    `xml:"dialogs"`
	Quests      []CharacterQuest `xml:"quests>quest"`
	Flags       []Flag           `xml:"flags>flag"`
	Crafting    []ObjectRecipe   `xml:"crafting>recipe"`
	Trainings   Trainings        `xml:"trainings"`
}

// Struct for equipment XML node.
type Equipment struct {
	XMLName xml.Name        `xml:"equipment"`
	Items   []EquipmentItem `xml:"item"`
}

// Struct for equipment item XML node.
type EquipmentItem struct {
	XMLName xml.Name `xml:"item"`
	ID      string   `xml:"id,attr"`
	Serial  string   `xml:"serial,attr"`
	Slot    string   `xml:"slot"`
}

// Struct for character memory XML node.
type Memory struct {
	XMLName xml.Name       `xml:"memory"`
	Nodes   []TargetMemory `xml:"target"`
}

// Struct for target memory XML node.
type TargetMemory struct {
	XMLName  xml.Name `xml:"target"`
	ID       string   `xml:"id,attr"`
	Serial   string   `xml:"serial,attr"`
	Attitude string   `xml:"attitude,attr"`
}

// Struct for flag XML node.
type Flag struct {
	XMLName xml.Name `xml:"flag"`
	ID      string   `xml:"id,attr"`
}

// Struct for character quest XML node.
type CharacterQuest struct {
	XMLName xml.Name `xml:"quest"`
	ID      string   `xml:"id,attr"`
	Stage   string   `xml:"stage,attr"`
}

// Struct for character attributes node.
type Attributes struct {
	XMLName      xml.Name `xml:"attributes"`
	Strenght     int      `xml:"stringht,attr"`
	Constitution int      `xml:"constitution,attr"`
	Dexterity    int      `xml:"dexterity,attr"`
	Intelligence int      `xml:"inteligence,attr"`
	Wisdom       int      `xml:"wisdom,attr"`
}

// UnmarshalCharacters retrieve all characters data
// from specified XML data.
func UnmarshalCharacters(data io.Reader) ([]*res.CharacterData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(Characters)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal xml data: %v", err)
	}
	chars := make([]*res.CharacterData, 0)
	for _, xmlChar := range xmlBase.Characters {
		char, err := buildCharacterData(&xmlChar)
		if err != nil {
			log.Err.Printf("xml: unmarshal character: build data fail: %v", err)
			continue
		}
		chars = append(chars, char)
	}
	return chars, nil
}

// MarshalCharacter parses game character to XML characters
// base string.
func MarshalCharacter(char *character.Character) (string, error) {
	xmlCharBase := new(Characters)
	xmlChar := xmlCharacter(char)
	xmlCharBase.Characters = append(xmlCharBase.Characters, *xmlChar)
	out, err := xml.Marshal(xmlCharBase)
	if err != nil {
		return "", fmt.Errorf("fail to marshal char: %v", err)
	}
	return string(out[:]), nil
}

// xmlCharacter parses specified game character to
// XML character struct.
func xmlCharacter(char *character.Character) *Character {
	xmlChar := new(Character)
	xmlChar.ID = char.ID()
	xmlChar.Serial = char.Serial()
	xmlChar.Name = char.Name()
	xmlChar.Level = char.Level()
	xmlChar.Gender = marshalGender(char.Gender())
	xmlChar.Race = marshalRace(char.Race())
	xmlChar.Attitude = marshalAttitude(char.Attitude())
	xmlChar.Alignment = marshalAlignment(char.Alignment())
	xmlChar.Attributes = xmlAttributes(char.Attributes())
	xmlChar.HP = char.Health()
	xmlChar.Mana = char.Mana()
	xmlChar.Exp = char.Experience()
	posX, posY := char.Position()
	xmlChar.Position = fmt.Sprintf("%fx%f", posX, posY)
	defX, defY := char.DefaultPosition()
	xmlChar.DefPosition = fmt.Sprintf("%fx%f", defX, defY)
	xmlChar.Inventory = *xmlInventory(char.Inventory())
	xmlChar.Equipment = *xmlEquipment(char.Equipment())
	xmlChar.Effects = xmlObjectEffects(char.Effects()...)
	xmlChar.Skills = *xmlObjectSkills(char.Skills()...)
	xmlChar.Memory = *xmlMemory(char.Memory())
	xmlChar.Dialogs = *xmlObjectDialogs(char.Dialogs()...)
	xmlChar.Quests = xmlQuests(char.Journal().Quests())
	xmlChar.Flags = xmlFlags(char.Flags())
	xmlChar.Crafting = xmlObjectRecipes(char.Recipes()...)
	xmlChar.Trainings = *xmlTrainings(char.Trainings()...)
	return xmlChar
}

// xmlAttributes parses character attributes to XML
// attributes nodes.
func xmlAttributes(attrs character.Attributes) Attributes {
	xmlAttrs := Attributes{
		Strenght:     attrs.Str,
		Constitution: attrs.Con,
		Dexterity:    attrs.Dex,
		Intelligence: attrs.Int,
		Wisdom:       attrs.Wis,
	}
	return xmlAttrs
}

// xmlEquipment parses specified character equipment to
// XML equipment node.
func xmlEquipment(eq *character.Equipment) *Equipment {
	xmlEq := new(Equipment)
	for _, s := range eq.Slots() {
		if s.Item() == nil {
			continue
		}
		xmlEqItem := EquipmentItem{
			ID:     s.Item().ID(),
			Serial: s.Item().Serial(),
			Slot:   MarshalEqSlot(s),
		}
		xmlEq.Items = append(xmlEq.Items, xmlEqItem)
	}
	return xmlEq
}

// xmlMemory parses specified character target memory to
// XML memory node.
func xmlMemory(mem []*character.TargetMemory) *Memory {
	xmlMem := new(Memory)
	for _, am := range mem {
		attAttr := marshalAttitude(am.Attitude)
		xmlAtt := TargetMemory{
			ID:       am.Target.ID(),
			Serial:   am.Target.Serial(),
			Attitude: attAttr,
		}
		xmlMem.Nodes = append(xmlMem.Nodes, xmlAtt)
	}
	return xmlMem
}

// xmlFlags parses specified flags to  XML flags nodes.
func xmlFlags(flags []flag.Flag) (xmlFlags []Flag) {
	for _, f := range flags {
		xmlFlag := Flag{ID: f.ID()}
		xmlFlags = append(xmlFlags, xmlFlag)
	}
	return
}

// xmlQuests parses specified qiests to XML quests nodes.
func xmlQuests(quests []*quest.Quest) (xmlQuests []CharacterQuest) {
	for _, q := range quests {
		xmlQuest := CharacterQuest{
			ID: q.ID(),
		}
		if s := q.ActiveStage(); s != nil {
			xmlQuest.Stage = s.ID()
		}
		xmlQuests = append(xmlQuests, xmlQuest)
	}
	return
}

// buildCharacterData creates character resources from specified
// XML data.
func buildCharacterData(xmlChar *Character) (*res.CharacterData, error) {
	// Basic data.
	baseData := res.CharacterBasicData{
		ID:     xmlChar.ID,
		Serial: xmlChar.Serial,
		Name:   xmlChar.Name,
		Level:  xmlChar.Level,
		Guild:  xmlChar.Guild,
	}
	data := res.CharacterData{BasicData: baseData}
	sex, err := UnmarshalGender(xmlChar.Gender)
	if err != nil {
		return nil, fmt.Errorf("fail to parse gender: %v", err)
	}
	data.BasicData.Sex = int(sex)
	race, err := UnmarshalRace(xmlChar.Race)
	if err != nil {
		return nil, fmt.Errorf("fail to parse race: %v", err)
	}
	data.BasicData.Race = int(race)
	attitude, err := UnmarshalAttitude(xmlChar.Attitude)
	if err != nil {
		return nil, fmt.Errorf("fail to parse attitude: %v", err)
	}
	data.BasicData.Attitude = int(attitude)
	alignment, err := UnmarshalAlignment(xmlChar.Alignment)
	if err != nil {
		return nil, fmt.Errorf("fail to parse alignment: %v", err)
	}
	data.BasicData.Alignment = int(alignment)
	// Attributes.
	attrs := xmlChar.Attributes
	data.BasicData.Str = attrs.Strenght
	data.BasicData.Con = attrs.Constitution
	data.BasicData.Dex = attrs.Dexterity
	data.BasicData.Int = attrs.Intelligence
	data.BasicData.Wis = attrs.Wisdom
	// Flags.
	for _, xmlFlag := range xmlChar.Flags {
		flagData := buildFlagData(xmlFlag)
		data.BasicData.Flags = append(data.BasicData.Flags, flagData)
	}
	// Trainings.
	data.BasicData.Trainings = buildTrainings(&xmlChar.Trainings)
	// Save.
	data.SavedData.PC = xmlChar.PC
	data.SavedData.HP = xmlChar.HP
	data.SavedData.Mana = xmlChar.Mana
	data.SavedData.Exp = xmlChar.Exp
	// Current & default position.
	if xmlChar.Position != "" {
		posX, posY, err := UnmarshalPosition(xmlChar.Position)
		if err != nil {
			return nil, fmt.Errorf("fail to parse position: %v", err)
		}
		data.SavedData.PosX, data.SavedData.PosY = posX, posY
	}
	if xmlChar.DefPosition != "" {
		defX, defY, err := UnmarshalPosition(xmlChar.DefPosition)
		if err != nil {
			return nil, fmt.Errorf("fail to parse default position: %v", err)
		}
		data.SavedData.DefX, data.SavedData.DefY = defX, defY
	}
	// Items.
	data.Inventory = buildInventory(xmlChar.Inventory)
	// Equipment.
	for _, xmlEqIt := range xmlChar.Equipment.Items {
		slot, err := UnmarshalEqSlot(xmlEqIt.Slot)
		if err != nil {
			log.Err.Printf("xml: build character: %s: parse eq item: %s: fail to parse lslot: %v",
				xmlChar.ID, xmlEqIt.ID, err)
			continue
		}
		eqItData := res.EquipmentItemData{
			ID:     xmlEqIt.ID,
			Serial: xmlEqIt.Serial,
			Slot:   int(slot),
		}
		data.EqItems = append(data.EqItems, eqItData)
	}
	// Effects.
	for _, xmlEffect := range xmlChar.Effects {
		effectData := res.ObjectEffectData{
			ID:           xmlEffect.ID,
			Serial:       xmlEffect.Serial,
			Time:         xmlEffect.Time,
			SourceID:     xmlEffect.Source.ID,
			SourceSerial: xmlEffect.Source.Serial,
		}
		data.Effects = append(data.Effects, effectData)
	}
	// Skills.
	for _, xmlSkill := range xmlChar.Skills.Nodes {
		skillData := res.ObjectSkillData{
			ID:     xmlSkill.ID,
			Serial: xmlSkill.Serial,
		}
		data.Skills = append(data.Skills, skillData)
	}
	// Memory.
	for _, xmlAtt := range xmlChar.Memory.Nodes {
		att, err := UnmarshalAttitude(xmlAtt.Attitude)
		if err != nil {
			log.Err.Printf("xml: build character: %s: fail to parse att mem: %s",
				xmlChar.ID, err)
			continue
		}
		attData := res.AttitudeMemoryData{
			ObjectID:     xmlAtt.ID,
			ObjectSerial: xmlAtt.Serial,
			Attitude:     int(att),
		}
		data.Memory = append(data.Memory, attData)
	}
	// Dialogs.
	for _, xmlDialog := range xmlChar.Dialogs.Nodes {
		dialogData := res.ObjectDialogData{
			ID: xmlDialog.ID,
		}
		data.Dialogs = append(data.Dialogs, dialogData)
	}
	// Quests.
	for _, xmlQuest := range xmlChar.Quests {
		questData := res.CharacterQuestData{
			ID:    xmlQuest.ID,
			Stage: xmlQuest.Stage,
		}
		data.Quests = append(data.Quests, questData)
	}
	// Recipes.
	for _, xmlRecipe := range xmlChar.Crafting {
		recipeData := res.ObjectRecipeData{
			ID: xmlRecipe.ID,
		}
		data.Recipes = append(data.Recipes, recipeData)
	}
	return &data, nil
}

// buildFlagData builds flag data from specified XML data.
func buildFlagData(xmlFlag Flag) res.FlagData {
	flagData := res.FlagData{ID: xmlFlag.ID}
	return flagData
}
