/*
 * cli.go
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

// Command line interface for flame engine.
// Uses flame command interpreter(CI) to handle user input and communicate
// with Flame Engine.
// All commands to be handled by CI must starts with dollar sign($),
// otherwise input is directly send to out(like 'echo').
// Type '$close' to close CLI.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	
	"github.com/isangeles/flame"
	"github.com/isangeles/flame/config"
	"github.com/isangeles/flame/cmd/burn"
	"github.com/isangeles/flame/cmd/burn/syntax"
	"github.com/isangeles/flame/cmd/log"
	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/data"
)

const (
	COMMAND_PREFIX   = "$"
	CLOSE_CMD        = "close"
	NEW_CHAR_CMD     = "newchar"
	NEW_GAME_CMD     = "newgame"
	LOAD_GAME_CMD    = "loadgame"
	IMPORT_CHARS_CMD = "importchars"
	REPEAT_INPUT_CMD = "!"
	INPUT_INDICATOR  = ">"
)

var (
	game        *core.Game
	lastCommand string
	lastUpdate  time.Time
)

// On init.
func init() {
	// Load flame config.
	err := config.LoadConfig()
	if err != nil {
		log.Err.Printf("fail_to_load_flame_config:%v", err)
	}
	// Load module.
	m, err := data.Module(config.ModulePath(), config.LangID())
	if err != nil {
		log.Err.Printf("fail_to_load_module:%v", err)
	}
	flame.SetModule(m)
	// Load CLI config.
	err = loadConfig()
	if err != nil {
		log.Err.Printf("fail_to_load_config:%v", err)
	}
}

func main() {
	fmt.Printf("*%s\t%s*\n", flame.NAME, flame.VERSION)
	fmt.Print(INPUT_INDICATOR)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		input := scan.Text()
		if strings.HasPrefix(input, COMMAND_PREFIX) {
			cmd := strings.TrimPrefix(input, COMMAND_PREFIX)
			execute(cmd)
			lastCommand = cmd
		} else {
			log.Inf.Println(input)
		}
		fmt.Print(INPUT_INDICATOR)
		// Game update on input.
		if game != nil {
			go gameLoop(game)
		}
	}
	if err := scan.Err(); err != nil {
		log.Err.Printf("input_scanner_init_fail_msg:%v\n", err)
	}
}

// execute passes specified command to CI.
func execute(input string) {
	switch input {
	case CLOSE_CMD:
		err := config.SaveConfig()
		if err != nil {
			log.Err.Printf("engine_config_save_fail:%v",
				err)
		}
		err = saveConfig()
		if err != nil {
			log.Err.Printf("config_save_fail:%v", err)
		}

		os.Exit(0)
	case NEW_CHAR_CMD:
		createdChar, err := newCharacterDialog()
		if err != nil {
			log.Err.Printf("%s\n", err)
			break
		}
		playableChars = append(playableChars, createdChar)
	case NEW_GAME_CMD:
		g, err := newGameDialog()
		if err != nil {
			log.Err.Printf("%s\n", err)
			break
		}
		game = g
		lastUpdate = time.Now()
	case LOAD_GAME_CMD:
		g, err := loadGameDialog()
		if err != nil {
			log.Err.Printf("%s", err)
			break
		}
		game = g
		lastUpdate = time.Now()
	case IMPORT_CHARS_CMD:
		if flame.Mod() == nil {
			log.Err.Printf("no_module_loaded")
			break
		}
		chars, err := data.ImportCharactersDir(flame.Mod(),
			flame.Mod().Conf().CharactersPath())
		if err != nil {
			log.Err.Printf("fail_to_import_module_characters:%v\n", err)
			break
		}
		log.Inf.Printf("imported_chars:%d\n", len(chars))
		for _, c := range chars {
			playableChars = append(playableChars, c)
		}
	case REPEAT_INPUT_CMD:
		execute(lastCommand)
		return
	default:
		exp, err := syntax.NewSTDExpression(input)
		if err != nil {
			log.Err.Printf("command_build_error:%v", err)
		}
		res, out := burn.HandleExpression(exp)
		log.Inf.Printf("burn[%d]:%s\n", res, out)
	}
}

// gameLoop handles game updating.
func gameLoop(g *core.Game) {
	dtNano := time.Since(lastUpdate).Nanoseconds()
	delta := dtNano / int64(time.Millisecond) // delta to milliseconds
	g.Update(delta)
	lastUpdate = time.Now()
}
