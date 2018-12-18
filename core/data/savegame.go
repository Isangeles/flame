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
	"path/filepath"
	"os"

	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/data/parsexml"
)

var (
	SAVEGAME_FILE_EXT = ".savegame"
)

// SaveGame saves specified game to savegame
// file.
func SaveGame(game *core.Game, dirPath, saveName string) error {
	// Parse game data.
	xml, err := parsexml.MarshalGame(game)
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
	f, err := os.Create(filepath.FromSlash(dirPath + "/" +
		saveName + SAVEGAME_FILE_EXT))
	if err != nil {
		return fmt.Errorf("fail_to_write_savegame_file:%v",
			err)
	}
	defer f.Close()
	// Write data to file.
	w := bufio.NewWriter(f)
	w.WriteString(xml)
	w.Flush()
	return nil
}
