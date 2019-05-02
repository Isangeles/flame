/*
 * savegame.go
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

	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/save"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/log"
)

var (
	SAVEGAME_FILE_EXT = ".savegame"
)

// SaveGame saves specified game to savegame file.
func SaveGame(game *core.Game, dirPath, saveName string) error {
	// Parse game data.
	save := new(save.SaveGame)
	save.Name = saveName
	save.Mod = game.Module()
	save.Players = game.Players()
	xml, err := parsexml.MarshalSaveGame(save)
	if err != nil {
		return fmt.Errorf("fail_to_marshal_game:%v", err)
	}
	// Create savegame file.
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_savegames_dir:%v", err)
	}
	filePath := filepath.FromSlash(dirPath + "/" + saveName +
		SAVEGAME_FILE_EXT)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("fail_to_write_savegame_file:%v", err)
	}
	defer f.Close()
	// Write data to file.
	w := bufio.NewWriter(f)
	w.WriteString(xml)
	w.Flush()
	log.Dbg.Printf("game_saved_in:%s", filePath)
	return nil
}

// ImportSavedGame imports saved game from save file with specified name in
// specified dir.
func ImportSavedGame(mod *module.Module, dirPath, fileName string) (*save.SaveGame, error) {
	filePath := filepath.FromSlash(dirPath + "/" + fileName)
	if !strings.HasSuffix(filePath, SAVEGAME_FILE_EXT) {
		filePath = filePath + SAVEGAME_FILE_EXT
	}
	doc, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_savegame_file:%v", err)
	}
	gameData, err := parsexml.UnmarshalGame(doc)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_savegame_data:%v",
			err)
	}
	// Load chapter with ID from save.
	err = LoadChapter(mod, gameData.Chapter.ID)
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_chapter:%v", err)
	}
	// Load chapter data(to build quests, characters, erc.).
	err = LoadChapterData(mod.Chapter())
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_chapter_data:%v", err)
	}
	mod.Chapter().ClearScenarios() // to remove start scenario
	save, err := buildSavedGame(mod, gameData)
	if err != nil {
		return nil, fmt.Errorf("fail_to_build_game_from_saved_data:%v",
			err)
	}
	return save, nil
}

// ImportSavedGamesDir imports all saved games from save files in
// directory with specified path.
func ImportSavedGamesDir(mod *module.Module, dirPath string) ([]*save.SaveGame, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Err.Printf("fail_to_read_dir:%v",
			err)
	}
	saves := make([]*save.SaveGame, 0)
	for _, fInfo := range files {
		if !strings.HasSuffix(fInfo.Name(), SAVEGAME_FILE_EXT) {
			continue
		}
		sav, err := ImportSavedGame(mod, dirPath, fInfo.Name())
		if err != nil {
			log.Err.Printf("data_savegame_load:fail_to_import_save:%v",
				err)
			continue
		}
		saves = append(saves, sav)
	}
	return saves, nil
}

// buildSavedGame build saved game from specified data.
func buildSavedGame(mod *module.Module, gameData *res.GameData) (*save.SaveGame, error) {
	chapterData := &gameData.Chapter
	charsData := make([]res.CharacterData, 0)
	objectsData := make([]*res.ObjectData, 0)
	pcs := make([]*character.Character, 0)
	// Scenrios.
	for _, scenData := range chapterData.Scenarios {
		subareas := make([]*scenario.Area, 0)
		mainarea := new(scenario.Area)
		// Areas.
		for _, areaData := range scenData.Areas {
			area := scenario.NewArea(areaData.ID)		 
			// Characters.
			for _, charData := range areaData.Chars {
				charsData = append(charsData, charData) // save data to restore effects later
				char := buildCharacter(mod, &charData)
				// Restore HP, mana & exp.
				char.SetHealth(charData.SavedData.HP)
				char.SetMana(charData.SavedData.Mana)
				char.SetExperience(charData.SavedData.Exp)
				// Restore current and default position.
				char.SetPosition(charData.SavedData.PosX, charData.SavedData.PosY)
				char.SetDefaultPosition(charData.SavedData.DefX, charData.SavedData.DefY)
				if charData.SavedData.PC {
					pcs = append(pcs, char)
				}
				area.AddCharacter(char)
			}
			// Objects.
			for _, obData := range areaData.Objects {
				objectsData = append(objectsData, &obData) // save data to restore effects later
				ob := buildObject(mod, &obData)
				// Restore position.
				ob.SetPosition(obData.SavedData.PosX, obData.SavedData.PosY)
				area.AddObject(ob)
			}
			if areaData.Mainarea {
				mainarea = area
			} else {
				subareas = append(subareas, area)
			}
		}
		// Create scenario from saved data.
		scen := scenario.NewScenario(scenData.ID, mainarea, subareas)
		err := mod.Chapter().AddScenario(scen)
		if err != nil {
			log.Err.Printf("data_build_saved_game:add_chapter_scenario:%s:fail:%v",
				scenData.ID, err)
			continue
		}
	}
	// Restore characters effects & memory.
	for _, cd := range charsData {
		err := restoreCharEffects(mod, &cd)
		if err != nil {
			log.Err.Printf("data:build_saved_game:restore_effects:char%s:%v",
				cd.BasicData.ID, err)
		}
		err = restoreCharMemory(mod, &cd)
		if err != nil {
			log.Err.Printf("data:build_saved_game:restore_memory:char%s:%v",
				cd.BasicData.ID, err)
		}
	}
	// Create game from saved data.
	game := new(save.SaveGame)
	game.Name = gameData.Name
	game.Mod = mod
	game.Players = pcs
	return game, nil
}

// restoreEffects resores effects for module character.
func restoreCharEffects(mod *module.Module, data *res.CharacterData) error {
	char := mod.Character(data.BasicData.ID + "_" + data.BasicData.Serial)
	if char == nil {
		return fmt.Errorf("char_not_found")
	}	
	for _, eData := range data.Effects {
		effect, err := Effect(mod, eData.ID)
		if err != nil {
			log.Err.Printf("data:char:%s:restore_effects:fail_to_create_effect:%v",
				char.ID(), err)
			continue
		}
		effect.SetSerial(eData.Serial)
		effect.SetTime(eData.Time)
		// Restore effect source.
		source := mod.Target(eData.SourceID, eData.SourceSerial)
		if source == nil {
			log.Err.Printf("data:char:%s:restore_effects:fail_to_find_source:%s",
				char.ID(), eData.SourceID + "_" + eData.SourceSerial)
		}
		effect.SetSource(source)
		char.AddEffect(effect)
	}
	return nil
}

// restoreCharMemory restores attitude memory for module character.
func restoreCharMemory(mod *module.Module, data *res.CharacterData) error {
	char := mod.Character(data.BasicData.ID + "_" + data.BasicData.Serial)
	if char == nil {
		return fmt.Errorf("char_not_found")
	}
	for _, memData := range data.Memory {
		tar := mod.Target(memData.ObjectID, memData.ObjectSerial)
		if tar == nil {
			log.Err.Printf("data:char:%s:restore_memory:att_target_not_found:%s_%s",
				char.ID(), memData.ObjectID, memData.ObjectSerial)
			continue
		}
		att := character.Attitude(memData.Attitude)
		char.Memorize(tar, att)
	}
	return nil
}
