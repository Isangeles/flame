/*
 * savegame.go
 *
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/log"
)

var (
	SAVEGAME_FILE_EXT = ".savegame"
)

// SaveGame saves specified game to savegame
// file.
func SaveGame(game *core.Game, dirPath, saveName string) error {
	// Parse game data.
	save := new(save.SaveGame)
	save.Name = saveName
	save.Mod = game.Module()
	save.Players = game.Players()
	xml, err := parsexml.MarshalSaveGame(save)
	if err != nil {
		return fmt.Errorf("fail_to_marshal_game:%v",
			err)
	}
	// Create savegame file.
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("fail_to_create_savegames_dir:%v",
			err)
	}
	filePath := filepath.FromSlash(dirPath + "/" +
		saveName + SAVEGAME_FILE_EXT)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("fail_to_write_savegame_file:%v",
			err)
	}
	defer f.Close()
	// Write data to file.
	w := bufio.NewWriter(f)
	w.WriteString(xml)
	w.Flush()
	log.Dbg.Printf("game_saved_in:%s", filePath)
	return nil
}

// LoadSavedGame loads saved game from save file with specified name in
// specified dir.
func LoadSavedGame(mod *module.Module, dirPath, fileName string) (*save.SaveGame, error) {
	filePath := filepath.FromSlash(dirPath + "/" + fileName +
		SAVEGAME_FILE_EXT)
	doc, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_savegame_file:%v",
			err)
	}
	xmlGame, err := parsexml.UnmarshalGame(doc)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_savegame_data:%v",
			err)
	}
	save, err := buildXMLSavedGame(mod, &xmlGame)
	if err != nil {
		return nil, fmt.Errorf("fail_to_build_game_from_saved_data:%v",
			err)
	}
	return save, nil
}

// LoadSavedGamesDir loads all saved games from save files in
// directory with specified path.
func LoadSavedGamesDir(mod *module.Module, dirPath string) ([]*save.SaveGame, error) {
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
		sav, err := LoadSavedGame(mod, dirPath, fInfo.Name())
		if err != nil {
			log.Err.Printf("data_savegame_load:fail_to_import_save:%v",
				err)
			continue
		}
		saves = append(saves, sav)
	}
	return saves, nil
}

// buildXMLSavedGame build game from data in specified
// saved game XML struct.
func buildXMLSavedGame(mod *module.Module,
	xmlGame *parsexml.SavedGameXML) (*save.SaveGame, error) {
	xmlChapter := &xmlGame.Chapter
	// Load chapter with ID from save.
	err := LoadChapter(mod, xmlChapter.ID)
	if err != nil {
		return nil, fmt.Errorf("fail_to_load_chapter:%v",
			err)
	}
	mod.Chapter().ClearScenarios() // to remove start scenario
	pcs := make([]*character.Character, 0)
	// Build chapter scenarios from save.
	for _, xmlScen := range xmlChapter.Scenarios {
		subareas := make([]*scenario.Area, 0)
		var mainarea *scenario.Area
		for _, xmlArea := range xmlScen.AreasNode.Areas {
			area := scenario.NewArea(xmlArea.ID)
			for _, xmlChar := range xmlArea.CharsNode.Characters {
				// Build chapter NPC.
				char, err := buildXMLCharacter(&xmlChar)
				if err != nil {
					log.Err.Printf("data_build_saved_game:build_char:%s:fail:%v",
						xmlChar.ID, err)
					continue
				}
				posX, posY, err := parsexml.UnmarshalPosition(
					xmlChar.Position)
				if err != nil {
					log.Err.Printf("data_build_saved_game:set_char_pos:%s:fail:%v",
						xmlChar.ID, err)
					continue
				}
				char.SetPosition(posX, posY)
				if xmlChar.PC {
					pcs = append(pcs, char)
				}
				char.SetSerial(xmlChar.Serial)
				area.AddCharacter(char)
			}
			if xmlArea.Mainarea {
				mainarea = area
			} else {
				subareas = append(subareas, area)
			}
		}
		// Create scenario from saved data.
		scen := scenario.NewScenario(xmlScen.ID, mainarea, subareas)
		err := mod.Chapter().AddScenario(scen)
		if err != nil {
			log.Err.Printf("data_build_saved_game:add_chapter_scenario:%s:fail:%v",
				xmlScen.ID, err)
		}
	}
	// Create game from saved data.
	game := new(save.SaveGame)
	game.Name = xmlGame.Name
	game.Mod = mod
	game.Players = pcs
	return game, nil
}
