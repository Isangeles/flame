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
// otherwise input is directly send to out(like 'echo')
// Type '$close' to close CLI
// @Isangeles
package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
	
	"github.com/isangeles/flame/core"
	//"github.com/isangeles/flame/core/log"
	"github.com/isangeles/flame/ci"
)

const (
	COMMAND_PREFIX = "$"
	CLOSE_CMD = "close"
	NEW_CHAR_CMD = "newchar"
	NEW_GAME_CMD = "newgame"
	INPUT_INDICATOR = ">"
)

// On init.
func init() {
	err := loadConfig()
	if err != nil {
		log.Printf("flame-cli>fail_to_load_config:%v", err)
	}
}

func main() {
	fmt.Printf("*%s\t%s*\n", core.NAME, core.VERSION)
	fmt.Print(INPUT_INDICATOR)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		input := scan.Text()
		if strings.HasPrefix(input, COMMAND_PREFIX) {
			input := strings.TrimPrefix(input, COMMAND_PREFIX)
			switch input {
			case CLOSE_CMD:
				err := core.SaveConfig()
				if err != nil {
					log.Printf("engine_config_save_fail:%v", err) 
				}
				err = saveConfig()
				if err != nil {
					log.Printf("config_save_fail:%v", err) 
				}
				
				os.Exit(0)
			case NEW_CHAR_CMD:
				createdChar, err := newCharacterDialog()
				if err != nil {
					fmt.Printf("%s\n", err)
					break 
				}
				playableChars = append(playableChars, createdChar)
				core.Mod().AddCharacter(createdChar)
			case NEW_GAME_CMD:
				err := newGameDialog()
				if err != nil {
					fmt.Printf("%s\n", err)
					break
				}
			default:
				cmd, err := NewCommand(input)
				if err != nil {
					fmt.Printf("command_build_error:%v", err)
				}
				code, msg := ci.HandleCommand(cmd)
				log.Printf("flame-ci[%d]>%s\n", code, msg) // uses log to auto print timestamps
			}
		} else {
			fmt.Println(input)
		}
		fmt.Print(INPUT_INDICATOR)
	}
	if err := scan.Err(); err != nil {
		fmt.Printf("input_scanner_init_fail_msg:%v\n", err)  
	}
}
