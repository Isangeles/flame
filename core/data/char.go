/*
 * char.go
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

package data

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/character"
	"github.com/isangeles/flame/core/module/craft"
	"github.com/isangeles/flame/core/module/quest"
	"github.com/isangeles/flame/log"
)

const (
	CharsFileExt = ".characters"
)

// Character creates module character with specified ID.
func Character(mod *module.Module, charID string) (*character.Character, error) {
	// Get data.
	data := res.Character(charID)
	if data == nil {
		return nil, fmt.Errorf("character data not found: %s", charID)
	}
	// Full build character(with skills, itmes, etc.).
	char := buildCharacter(mod, data)
	return char, nil
}

// ImportCharactersData import characters data from base file
// with specified path.
func ImportCharactersData(path string) ([]*res.CharacterData, error) {
	baseFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fail to open char base file: %v", err)
	}
	defer baseFile.Close()
	chars, err := parsexml.UnmarshalCharacters(baseFile)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal chars base: %v", err)
	}
	return chars, nil
}

// ImportCharactersDataDir imports all characters data from
// files in directory with specified path.
func ImportCharactersDataDir(path string) ([]*res.CharacterData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	chars := make([]*res.CharacterData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), CharsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(path + "/" + finfo.Name())
		impChars, err := ImportCharactersData(basePath)
		if err != nil {
			log.Err.Printf("data: import chars dir: %s: fail to parse char file: %v",
				basePath, err)
			continue
		}
		chars = append(chars, impChars...)
	}
	return chars, nil
}

// ImportCharacters imports characters from base file with
// specified path.
func ImportCharacters(mod *module.Module, path string) ([]*character.Character, error) {
	charFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fail to open char base file: %v", err)
	}
	defer charFile.Close()
	charsData, err := parsexml.UnmarshalCharacters(charFile)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal chars base: %v", err)
	}
	chars := make([]*character.Character, 0)
	for _, charData := range charsData {
		char := buildCharacter(mod, charData)
		chars = append(chars, char)
	}
	return chars, nil
}

// ImportCharactersDir imports all characters files from directory
// with specified path.
func ImportCharactersDir(mod *module.Module, dirPath string) ([]*character.Character, error) {
	chars := make([]*character.Character, 0)
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return chars, fmt.Errorf("fail to read dir: %v", err)
	}
	for _, fInfo := range files {
		if !strings.HasSuffix(fInfo.Name(), CharsFileExt) {
			continue
		}
		charFilePath := filepath.FromSlash(dirPath + "/" + fInfo.Name())
		impChars, err := ImportCharacters(mod, charFilePath)
		if err != nil {
			log.Err.Printf("data char import: %s: fail to parse char file: %v",
				charFilePath, err)
			continue
		}
		for _, c := range impChars {
			chars = append(chars, c)
		}
	}
	return chars, nil
}

// ExportCharacter saves specified character to
// [Module]/characters directory.
func ExportCharacter(char *character.Character, dirPath string) error {
	// Parse character data.
	xml, err := parsexml.MarshalCharacter(char)
	if err != nil {
		return fmt.Errorf("fail to export char: %v", err)
	}
	// Create character file.
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("fail to create chrs dir: %v", err)
	}
	f, err := os.Create(filepath.FromSlash(dirPath+"/"+
		strings.ToLower(char.Name())) + CharsFileExt)
	if err != nil {
		return fmt.Errorf("fail to create char file: %v", err)
	}
	defer f.Close()
	// Write data to file.
	w := bufio.NewWriter(f)
	w.WriteString(xml)
	w.Flush()
	return nil
}

// buildCharacter builds new character from specified data(with items and equipment).
func buildCharacter(mod *module.Module, data *res.CharacterData) *character.Character {
	char := character.New(*data)
	// Quests.
	for _, logQuestData := range data.QuestLog.Quests {
		questData := res.Quest(logQuestData.ID)
		if questData == nil {
			log.Err.Printf("data: build character: %s: fail to retrieve quest: %s",
				char.ID(), logQuestData.ID)
			continue
		}
		// Restore quest stage.
		quest := quest.New(*questData)
		for _, s := range quest.Stages() {
			if s.ID() == logQuestData.Stage {
				quest.SetActiveStage(s)
			}
		}
		if quest.ActiveStage() == nil {
			log.Err.Printf("data: build character: %s: quest: %s: fail to set active stage",
				char.ID(), quest.ID())
		}
		// Add quest to quest log.
		char.Journal().AddQuest(quest)
	}
	// Recipes.
	for _, obRecipeData := range data.Recipes {
		recipeData := res.Recipe(obRecipeData.ID)
		if recipeData == nil {
			log.Err.Printf("data: build character: %s: fail to retrieve recipe: %s",
				char.ID(), obRecipeData.ID)
			continue
		}
		recipe := craft.NewRecipe(*recipeData)
		char.AddRecipe(recipe)
	}
	return char
}
