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
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/module"
	"github.com/isangeles/flame/module/area"
	"github.com/isangeles/flame/module/character"
	"github.com/isangeles/flame/module/object"
	"github.com/isangeles/flame/module/serial"
)

var (
	SavegameFileExt = ".savegame"
)

// ExportGame saves specified game to savegame file in specified directory.
func ExportGame(game *flame.Game, path string) error {
	data := gameData(game)
	xml, err := xml.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal game: %v", err)
	}
	// Create savegame file.
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("unable to create savegames dir: %v", err)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to write savegame file: %v", err)
	}
	defer f.Close()
	// Write data to file.
	w := bufio.NewWriter(f)
	w.Write(xml)
	w.Flush()
	log.Dbg.Printf("game saved in: %s", path)
	return nil
}

// ImportGame imports saved game from save file with specified name in
// specified dir.
func ImportGame(mod *module.Module, path, lang string) (*flame.Game, error) {
	doc, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open savegame file: %v", err)
	}
	buf, err := ioutil.ReadAll(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to read save file: %v", err)
	}
	gameData := new(res.GameData)
	err = xml.Unmarshal(buf, gameData)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal savegame data: %v", err)
	}
	// Reset serial objects base.
	serial.Reset()
	// Load chapter with ID from save.
	chapterPath := filepath.Join(mod.Conf().ChaptersPath(), gameData.SavedChapter.ID)
	chapterData, err := ImportChapter(chapterPath, lang)
	if err != nil {
		return nil, fmt.Errorf("unable to import chapter: %v", err)
	}
	chapter := module.NewChapter(mod, chapterData)
	mod.SetChapter(chapter)
	game, err := buildSavedGame(mod, gameData)
	if err != nil {
		return nil, fmt.Errorf("unable to build game from saved data: %v", err)
	}
	return game, nil
}

// ImportGamesDir imports all saved games from save files in
// directory with specified path.
func ImportGamesDir(mod *module.Module, dirPath, lang string) ([]*flame.Game, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Err.Printf("unable to read dir: %v", err)
	}
	games := make([]*flame.Game, 0)
	for _, fInfo := range files {
		if !strings.HasSuffix(fInfo.Name(), SavegameFileExt) {
			continue
		}
		path := filepath.Join(dirPath, fInfo.Name())
		game, err := ImportGame(mod, path, lang)
		if err != nil {
			log.Err.Printf("data savegame load: unable to import saved game: %v", err)
			continue
		}
		games = append(games, game)
	}
	return games, nil
}

// gameData creates data resource for game.
func gameData(g *flame.Game) (data res.GameData) {
	chapter := g.Module().Chapter()
	chapterData := res.SavedChapterData{ID: chapter.ID()}
	for _, a := range chapter.Areas() {
		chapterData.Areas = append(chapterData.Areas, savedAreaData(a))
	}
	data.SavedChapter = chapterData
	return
}

// savedAreaData creates area data resurce for game data.
func savedAreaData(a *area.Area) (data res.SavedAreaData) {
	data.ID = a.ID()
	for _, c := range a.Characters() {
		data.Chars = append(data.Chars, c.Data())
	}
	for _, o := range a.Objects() {
		data.Objects = append(data.Objects, o.Data())
	}
	for _, sa := range a.Subareas() {
		data.Subareas = append(data.Subareas, savedAreaData(sa))
	}
	return
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
		area.AddCharacter(char)
	}
	// Objects.
	for _, obData := range data.Objects {
		ob := object.New(obData)
		area.AddObject(ob)
	}
	// Subareas.
	for _, subareaData := range data.Subareas {
		subarea := buildSavedArea(mod, subareaData)
		area.AddSubarea(subarea)
	}
	return area
}
