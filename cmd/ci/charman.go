/*
 * charman.go
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

package ci

import (
	"fmt"
	"strconv"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/data"
)

// handleCharCommand handles specified command for game
// character.
func handleCharCommand(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 3, fmt.Sprintf("%s:no_active_game", CHAR_MAN)
	}
	if len(cmd.OptionArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no_option_args", CHAR_MAN)
	}

	switch cmd.OptionArgs()[0] {
	case "set":
		return setCharOption(cmd)
	case "show":
		return showCharOption(cmd)
	case "export":
		return exportCharOption(cmd)
	default:
		return 4, fmt.Sprintf("%s:no_such_option:%s", CHAR_MAN,
			cmd.OptionArgs()[0])
	}
}

// setCharOption handles set option for charman commands.
func setCharOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 2 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s", CHAR_MAN,
			cmd.OptionArgs()[0])
	}
	char := flame.Game().Player(cmd.TargetArgs()[1])
	if char == nil {
		return 5, fmt.Sprintf("%s:character_not_found:%s", CHAR_MAN,
			cmd.TargetArgs()[1])
	}

	switch cmd.TargetArgs()[0] {
	case "position":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.OptionArgs()[1])
		}
		x, err := strconv.ParseFloat(cmd.Args()[0], 64)
		if err != nil {
			return 8, fmt.Sprintf("%s:invalid_argument:%s", CHAR_MAN,
				cmd.OptionArgs()[0])
		}
		y, err := strconv.ParseFloat(cmd.Args()[1], 64)
		if err != nil {
			return 8, fmt.Sprintf("%s:invalid_argument:%s", CHAR_MAN,
				cmd.OptionArgs()[1])
		}
		char.SetPosition(x, y)
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}
}

// showCharOption handles show option for charman commands.
func showCharOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 2 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s", CHAR_MAN,
			cmd.OptionArgs()[0])
	}
	char := flame.Game().Player(cmd.TargetArgs()[1])
	if char == nil {
		return 5, fmt.Sprintf("%s:character_not_found:%s", CHAR_MAN,
			cmd.TargetArgs()[1])
	}

	switch cmd.TargetArgs()[0] {
	case "position":
		x, y := char.Position()
		return 0, fmt.Sprintf("%fx%f", x, y)
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}
}

// exportEngineOption handles 'export' option for engineman CI tool.
func exportCharOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s",
			CHAR_MAN, cmd.OptionArgs()[0])
	}
	char := flame.Game().Player(cmd.TargetArgs()[0])
	if char == nil {
		return 5, fmt.Sprintf("%s:character_not_found:%s", CHAR_MAN,
			cmd.TargetArgs()[1])
	}

	err := data.ExportCharacter(char, flame.Game().Module().CharactersPath())
	if err != nil {
		return 8, fmt.Sprintf("%s:%v", CHAR_MAN, err)
	}
	return 0, ""
}
