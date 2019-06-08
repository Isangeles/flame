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
// Uses Burn CI to handle user input and communicate with Flame Engine.
// All commands to be handled by CI must starts with dollar sign($),
// otherwise input is directly send to out(like 'echo').
// Type '$close' to close CLI.
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/isangeles/flame"
	flameconf "github.com/isangeles/flame/config"
	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/module/object/character"

	"github.com/isangeles/flame/cmd/burn"
	"github.com/isangeles/flame/cmd/burn/ash"
	"github.com/isangeles/flame/cmd/burn/syntax"
	"github.com/isangeles/flame/cmd/config"
	"github.com/isangeles/flame/cmd/log"
)

const (
	COMMAND_PREFIX   = "$"
	SCRIPT_PREFIX    = "#"
	CLOSE_CMD        = "close"
	NEW_CHAR_CMD     = "newchar"
	NEW_GAME_CMD     = "newgame"
	NEW_MOD_CMD      = "newmod"
	LOAD_GAME_CMD    = "loadgame"
	IMPORT_CHARS_CMD = "importchars"
	LOOT_TARGET_CMD  = "loot"
	TALK_TARGET_CMD  = "talk"
	FIND_TARGET_CMD  = "target"
	TARGET_INFO_CMD  = "tarinfo"
	QUESTS_CMD       = "quests"
	USE_SKILL_CMD    = "useskill"
	REPEAT_INPUT_CMD = "!"
	INPUT_INDICATOR  = ">"
)

var (
	game        *core.Game
	activePC    *character.Character
	lastCommand string
	lastUpdate  time.Time
)

// On init.
func init() {
	// Load flame config.
	err := flameconf.LoadConfig()
	if err != nil {
		log.Err.Printf("fail_to_load_flame_config:%v", err)
	}
	// Load module.
	err = loadModule(flameconf.ModulePath(), flameconf.LangID())
	if err != nil {
		log.Err.Printf("fail_to_load_module:%v", err)
	}
	// Load CLI config.
	err = config.LoadConfig()
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
		} else if strings.HasPrefix(input, SCRIPT_PREFIX) {
			input := strings.TrimPrefix(input, SCRIPT_PREFIX)
			scrArgs := strings.Split(input, " ")
			executeFile(scrArgs[0], scrArgs...)
		} else if activePC != nil {
			activePC.SendChat(input)
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
		err := flameconf.SaveConfig()
		if err != nil {
			log.Err.Printf("engine_config_save_fail:%v",
				err)
		}
		err = config.SaveConfig()
		if err != nil {
			log.Err.Printf("config_save_fail:%v", err)
		}
		os.Exit(0)
	case NEW_CHAR_CMD:
		createdChar, err := newCharacterDialog(flame.Mod())
		if err != nil {
			log.Err.Printf("%s\n", err)
			break
		}
		playableChars = append(playableChars, createdChar)
	case NEW_GAME_CMD:
		g, err := newGameDialog()
		if err != nil {
			log.Err.Printf("%s:%v", NEW_GAME_CMD, err)
			break
		}
		game = g
		activePC = game.Players()[0]
		lastUpdate = time.Now()
	case NEW_MOD_CMD:
		err := newModDialog()
		if err != nil {
			log.Err.Printf("%s:%v", NEW_MOD_CMD, err)
			break
		}
	case LOAD_GAME_CMD:
		g, err := loadGameDialog()
		if err != nil {
			log.Err.Printf("%s:%v", LOAD_GAME_CMD, err)
			break
		}
		game = g
		activePC = game.Players()[0]
		lastUpdate = time.Now()
	case IMPORT_CHARS_CMD:
		chars, err := data.ImportCharactersDir(flame.Mod(),
			flame.Mod().Conf().CharactersPath())
		if err != nil {
			log.Err.Printf("%s:%v", IMPORT_CHARS_CMD, err)
			break
		}
		log.Inf.Printf("imported_chars:%d\n", len(chars))
		for _, c := range chars {
			playableChars = append(playableChars, c)
		}
	case LOOT_TARGET_CMD:
		err := lootDialog()
		if err != nil {
			log.Err.Printf("%s:%v", LOOT_TARGET_CMD, err)
			break
		}
	case TALK_TARGET_CMD:
		err := talkDialog()
		if err != nil {
			log.Err.Printf("%s:%v", TALK_TARGET_CMD, err)
			break
		}
	case FIND_TARGET_CMD:
		err := targetDialog()
		if err != nil {
			log.Err.Printf("%s:%v", FIND_TARGET_CMD, err)
			break
		}
	case TARGET_INFO_CMD:
		err := targetInfoDialog()
		if err != nil {
			log.Err.Printf("%s:%v", TARGET_INFO_CMD, err)
			break
		}
	case QUESTS_CMD:
		err := questsDialog()
		if err != nil {
			log.Err.Printf("%s:%v", QUESTS_CMD, err)
		}
	case USE_SKILL_CMD:
		err := useSkillDialog()
		if err != nil {
			log.Err.Printf("%s:%v", USE_SKILL_CMD, err)
		}
	case REPEAT_INPUT_CMD:
		execute(lastCommand)
		return
	default: // pass command to CI
		exp, err := syntax.NewSTDExpression(input)
		if err != nil {
			log.Err.Printf("command_build_error:%v", err)
		}
		res, out := burn.HandleExpression(exp)
		log.Inf.Printf("burn[%d]:%s\n", res, out)
	}
}

// executeFile executes script from data/scripts dir.
func executeFile(filename string, args ...string) {
	path := fmt.Sprintf("%s/%s%s", config.ScriptsPath(),
		filename, ash.SCRIPT_FILE_EXT)
	file, err := os.Open(path)
	if err != nil {
		log.Err.Printf("fail_to_open_file:%v", err)
		return
	}
	text, err := ioutil.ReadAll(file)
	if err != nil {
		log.Err.Printf("fail_to_read_file:%v", err)
		return
	}
	scr, err := ash.NewScript(fmt.Sprintf("%s", text), args...)
	if err != nil {
		log.Err.Printf("fail_to_parse_script:%v", err)
		return 
	}
	err = ash.Run(scr)
	if err != nil {
		log.Err.Printf("script_run_fail:%v", err)
		return 
	}
}

// gameLoop handles game updating.
func gameLoop(g *core.Game) {
	dtNano := time.Since(lastUpdate).Nanoseconds()
	delta := dtNano / int64(time.Millisecond) // delta to milliseconds
	g.Update(delta)
	lastUpdate = time.Now()
}

// loadModule loads module with all module data
// from directory with specified path.
func loadModule(path, langID string) error {
	m, err := data.Module(flameconf.ModulePath(), flameconf.LangID())
	if err != nil {
		return fmt.Errorf("fail_to_dir:%v", err)
	}
	// Load module data.
	err = data.LoadModuleData(m)
	if err != nil {
		return fmt.Errorf("fail_to_load_data:%v", err)
	}
	flame.SetModule(m)
	return nil
}
