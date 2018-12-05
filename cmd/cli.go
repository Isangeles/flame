/*
 * cli.go
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

// Command line interface for flame engine.
// Uses flame command interpreter(CI) to handle user input and communicate
// with Flame Engine.
// All commands to be handle by CI must starts with generic sum sign($),
// otherwise input is directly send to out(like 'echo').
// Type '$close' to close CLI.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/enginelog"
	"github.com/isangeles/flame/cmd/ci"
	"github.com/isangeles/flame/cmd/command"
)

const (
	COMMAND_PREFIX   = "$"
	CLOSE_CMD        = "close"
	NEW_CHAR_CMD     = "newchar"
	NEW_GAME_CMD     = "newgame"
	IMPORT_CHARS_CMD = "importchars"
	REPEAT_INPUT_CMD = "!"
	INPUT_INDICATOR  = ">"
)

var (
	stdout *log.Logger = log.New(enginelog.InfLog, "flame-cli>", 0)
	stderr *log.Logger = log.New(enginelog.ErrLog, "flame-cli>", 0)
	dbglog *log.Logger = log.New(enginelog.DbgLog, "flame-cli-debug>", 0)

	game *core.Game

	lastCommand string
)

// On init.
func init() {
	err := flame.LoadConfig()
	if err != nil {
		stderr.Printf("fail_to_load_flame_config:%v", err)
	}
	err = loadConfig()
	if err != nil {
		stderr.Printf("fail_to_load_config:%v", err)
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
			stdout.Println(input)
		}
		fmt.Print(INPUT_INDICATOR)
	}
	if err := scan.Err(); err != nil {
		fmt.Printf("input_scanner_init_fail_msg:%v\n", err)
	}
}

// execute passes specified command to CI.
func execute(input string) {
	switch input {
	case CLOSE_CMD:
		err := flame.SaveConfig()
		if err != nil {
			stderr.Printf("engine_config_save_fail:%v",
				err)
		}
		err = saveConfig()
		if err != nil {
			stderr.Printf("config_save_fail:%v", err)
		}

		os.Exit(0)
	case NEW_CHAR_CMD:
		createdChar, err := newCharacterDialog()
		if err != nil {
			stderr.Printf("%s\n", err)
			break
		}
		playableChars = append(playableChars, createdChar)
	case NEW_GAME_CMD:
		g, err := newGameDialog()
		if err != nil {
			stderr.Printf("%s\n", err)
			break
		}
		game = g
	case IMPORT_CHARS_CMD:
		if flame.Mod() == nil {
			stderr.Printf("no_module_loaded")
			break
		}
		chars, err := data.ImportCharactersDir(flame.Mod().CharactersPath())
		if err != nil {
			stderr.Printf("fail_to_import_module_characters:%v\n", err)
			break
		}
		stdout.Printf("imported_chars:%d\n", len(chars))
		for _, c := range chars {
			playableChars = append(playableChars, c)
		}
	case REPEAT_INPUT_CMD:
		execute(lastCommand)
		return
	default:
		cmd, err := command.NewStdCommand(input)
		if err != nil {
			stderr.Printf("command_build_error:%v", err)
		}
		code, msg := ci.HandleCommand(cmd)
		stdout.Printf("CI[%d]:%s\n", code, msg)
	}
}
