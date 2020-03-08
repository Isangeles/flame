/*
 * save.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/data/parsexml"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module"
	"github.com/isangeles/flame/module/area"
	"github.com/isangeles/flame/module/character"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/object"
	"github.com/isangeles/flame/log"
)

var (
	SavegameFileExt = ".savegame"
)

// ExportGame saves specified game to savegame file in specified directory.
func ExportGame(game *flame.Game, dirPath, saveName string) error {
	xml, err := parsexml.MarshalGame(game)
	if err != nil {
		return fmt.Errorf("unable to marshal game: %v", err)
	} 
	// Create savegame file.
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create savegames dir: %v", err)
	}
	filePath := filepath.FromSlash(dirPath + "/" + saveName +
		SavegameFileExt)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to write savegame file: %v", err)
	}
	defer f.Close()
	// Write data to file.
	w := bufio.NewWriter(f)
	w.WriteString(xml)
	w.Flush()
	log.Dbg.Printf("game saved in: %s", filePath)
	return nil
}

// ImportGame imports saved game from save file with specified name in
// specified dir.
func ImportGame(mod *module.Module, dirPath, fileName string) (*flame.Game, error) {
	filePath := filepath.FromSlash(dirPath + "/" + fileName)
	if !strings.HasSuffix(filePath, SavegameFileExt) {
		filePath = filePath + SavegameFileExt
	}
	doc, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open savegame file: %v", err)
	}
	gameData, err := parsexml.UnmarshalGame(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal savegame data: %v", err)
	}
	// Load chapter with ID from save.
	err = LoadChapter(mod, gameData.SavedChapter.ID)
	if err != nil {
		return nil, fmt.Errorf("unable to load chapter: %v", err)
	}
	game, err := buildSavedGame(mod, gameData)
	if err != nil {
		return nil, fmt.Errorf("unable to build game from saved data: %v", err)
	}
	return game, nil
}

// ImportGamesDir imports all saved games from save files in
// directory with specified path.
func ImportGamesDir(mod *module.Module, dirPath string) ([]*flame.Game, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Err.Printf("unable to read dir: %v", err)
	}
	games := make([]*flame.Game, 0)
	for _, fInfo := range files {
		if !strings.HasSuffix(fInfo.Name(), SavegameFileExt) {
			continue
		}
		game, err := ImportGame(mod, dirPath, fInfo.Name())
		if err != nil {
			log.Err.Printf("data savegame load: unable to import saved game: %v", err)
			continue
		}
		games = append(games, game)
	}
	return games, nil
}

// buildSavedGame build saved game from specified data.
func buildSavedGame(mod *module.Module, gameData *res.GameData) (*flame.Game, error) {
	chapterData := &gameData.SavedChapter
	// Areas.
	for _, areaData := range chapterData.Areas {
		// Create area from saved data.
		area := buildSavedArea(mod, areaData)
		mod.Chapter().AddAreas(area)
	}
	// Restore objects effects and memory.
	for _, areaData := range chapterData.Areas {
		restoreAreaEffects(mod, areaData)
		restoreAreaMemory(mod, areaData)
	}
	// Create game from saved data.
	game := flame.NewGame(mod)
	return game, nil
}

// buildSavedArea creates area from saved data.
func buildSavedArea(mod *module.Module, data res.SavedAreaData) *area.Area {
	areaData := res.AreaData{
		ID: data.ID,
	}
	area := area.New(areaData)
	// Characters.
	for _, charData := range data.Chars {
		char := character.New(charData)
		// Restore serial.
		char.SetSerial(charData.BasicData.Serial)
		// Restore HP, mana & exp.
		char.SetHealth(charData.SavedData.HP)
		char.SetMana(charData.SavedData.Mana)
		char.SetExperience(charData.SavedData.Exp)
		// Restore current and default position.
		char.SetPosition(charData.SavedData.PosX, charData.SavedData.PosY)
		char.SetDefaultPosition(charData.SavedData.DefX, charData.SavedData.DefY)
		// Char to area.
		area.AddCharacter(char)
	}
	// Objects.
	for _, obData := range data.Objects {
		// Build object.
		ob := object.New(obData)
		// Restore serial.
		ob.SetSerial(obData.BasicData.Serial)
		// Restore health & position.
		ob.SetHealth(obData.SavedData.HP)
		ob.SetPosition(obData.SavedData.PosX, obData.SavedData.PosY)
		// Object to area.
		area.AddObject(ob)
	}
	// Subareas.
	for _, subareaData := range data.Subareas {
		subarea := buildSavedArea(mod, subareaData)
		area.AddSubarea(subarea)
	}
	return area
}

// restoreAreaEffects restores saved effects for characters and objects.
func restoreAreaEffects(mod *module.Module, data res.SavedAreaData) {
	for _, charData := range data.Chars {
		char := mod.Chapter().Character(charData.BasicData.ID, charData.BasicData.Serial)
		if char == nil {
			log.Err.Printf("data: save: restore effects: module char not found: %s#%s",
				charData.BasicData.ID, charData.BasicData.Serial)
			continue
		}
		for _, obEffectData := range charData.Effects {
			effectData := res.Effect(obEffectData.ID)
			if effectData == nil {
				log.Err.Printf("data: char: %s: restore effects: unable to create effect: %s",
					char.ID(), obEffectData.ID)
				continue
			}
			effect := effect.New(*effectData)
			effect.SetSerial(obEffectData.Serial)
			effect.SetTime(obEffectData.Time)
			// Restore effect source.
			source := mod.Target(obEffectData.SourceID, obEffectData.SourceSerial)
			if source == nil {
				log.Err.Printf("data: char: %s: restore effects: unable to find source: %s#%s",
					char.ID(), obEffectData.SourceID, obEffectData.SourceSerial)
			}
			effect.SetSource(source)
			char.AddEffect(effect)
		}
	}
	for _, subareaData := range data.Subareas {
		restoreAreaEffects(mod, subareaData)
	}
}

// restoreAreaMemory restores saved memory for characters.
func restoreAreaMemory(mod *module.Module, data res.SavedAreaData) {
	for _, charData := range data.Chars {
		char := mod.Chapter().Character(charData.BasicData.ID, charData.BasicData.Serial)
		if char == nil {
			log.Err.Printf("data: save: restore effects: module char not found: %s#%s",
				charData.BasicData.ID, charData.BasicData.Serial)
			continue
		}
		for _, memData := range charData.Memory {
			tar := mod.Target(memData.ObjectID, memData.ObjectSerial)
			if tar == nil {
				log.Err.Printf("data: char: %s: restore memory: att target not found: %s#%s",
					char.ID(), memData.ObjectID, memData.ObjectSerial)
				continue
			}
			att := character.Attitude(memData.Attitude)
			mem := character.TargetMemory{
				Target:   tar,
				Attitude: att,
			}
			char.MemorizeTarget(&mem)
		}
	}
	for _, subareaData := range data.Subareas {
		restoreAreaMemory(mod, subareaData)
	}
}
