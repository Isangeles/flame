/*
 * moduleman.go
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

package burn

import (
	"fmt"
	"strconv"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/scenario"
	"github.com/isangeles/flame/core/module/serial"
)

// Handles specified module command,
// returns response code and message.
func handleModuleCommand(cmd Command) (int, string) {
	if flame.Mod() == nil {
		return 3, fmt.Sprintf("%s:no_module_loaded", MODULE_MAN)
	}
	if len(cmd.OptionArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no_option_args", MODULE_MAN)
	}

	switch cmd.OptionArgs()[0] {
	case "show":
		return showModuleOption(cmd)
	case "add":
		return addModuleOption(cmd)
	default:
		return 4, fmt.Sprintf("%s:no_such_option:%s", MODULE_MAN,
			cmd.OptionArgs()[0])
	}
}


// showModuleOption handles show option for moduleman CI tool,
// returns response code and message.
func showModuleOption(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_args_for:%s", MODULE_MAN,
			cmd.OptionArgs()[0])
	}

	switch cmd.Args()[0] {
	case "id", "name":
		return 0, flame.Mod().Conf().ID
	case "area-chars":
		if flame.Game() == nil {
			return 8, fmt.Sprintf("%s:no_game_loaded", MODULE_MAN)
		}
		if len(cmd.TargetArgs()) < 1 {
			return 8, fmt.Sprintf("%s:no_enought_args_for%s",
				MODULE_MAN, cmd.Args()[0])
		}
		areaID := cmd.TargetArgs()[0]
		var area *scenario.Area
		for _, s := range flame.Mod().Chapter().Scenarios() {
			for _, a := range s.Areas() {
				if a.ID() == areaID {
					area = a
				}
			}
		}
		if area == nil {
			return 9, fmt.Sprintf("%s:area_not_found:%s",
				MODULE_MAN, areaID)
		}
		out := ""
		for _, c := range area.Characters() {
			out = fmt.Sprintf("%s %s_%s", out, c.ID(), c.Serial())
		}
		return 0, out
	case "area-objects":
		if flame.Game() == nil {
			return 8, fmt.Sprintf("%s:no_game_loaded", MODULE_MAN)
		}
		if len(cmd.TargetArgs()) < 1 {
			return 8, fmt.Sprintf("%s:no_enought_args_for%s",
				MODULE_MAN, cmd.Args()[0])
		}
		areaID := cmd.TargetArgs()[0]
		var area *scenario.Area
		for _, s := range flame.Mod().Chapter().Scenarios() {
			for _, a := range s.Areas() {
				if a.ID() == areaID {
					area = a
				}
			}
		}
		if area == nil {
			return 9, fmt.Sprintf("%s:area_not_found:%s",
				MODULE_MAN, areaID)
		}
		out := ""
		for _, o := range area.Objects() {
			out = fmt.Sprintf("%s %s_%s", out, o.ID(), o.Serial())
		}
		return 0, out
	case "scenario":
		return 10, "unsupported yet"
	case "res-objects":
		out := ""
		for _, ob := range res.Objects() {
			out = fmt.Sprintf("%s %s", out, ob.BasicData.ID)
		}
		return 0, out
	case "res-dialogs":
		out := ""
		for _, dl := range res.Dialogs() {
			out = fmt.Sprintf("%s %s", out, dl.ID)
		}
		return 0, out
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", ENGINE_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}

// addModuleOption handles add option for moduleman CI tool,
// returns response code and message.
func addModuleOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s", MODULE_MAN,
			cmd.OptionArgs()[0])
	}

	switch cmd.TargetArgs()[0] {
	case "character":
		if len(cmd.Args()) < 3 {
			return 8, fmt.Sprintf("%s:no_enought_args_for%s",
				MODULE_MAN, cmd.TargetArgs()[0])
		}
		id := cmd.Args()[0]
		scenID := cmd.Args()[1]
		areaID := cmd.Args()[2]
		posX, posY := 0.0, 0.0
		if len(cmd.Args()) > 4 {
			var err error
			posX, err = strconv.ParseFloat(cmd.Args()[3], 64)
			if err != nil {
				return 8, fmt.Sprintf("%s:fail_to_parse_x_position:%v",
					MODULE_MAN, err)
			}
			posY, err = strconv.ParseFloat(cmd.Args()[4], 64)
			if err != nil {
				return 8, fmt.Sprintf("%s:fail_to_parse_x_position:%v",
					MODULE_MAN, err)
			}
		}
		char, err := data.Character(flame.Mod(), id)
		if err != nil {
			return 9, fmt.Sprintf("%s:fail_to_retrieve_character:%v", MODULE_MAN,
				err)
		}
		char.SetPosition(posX, posY)
		for _, s := range flame.Mod().Chapter().Scenarios() {
			if s.ID() != scenID {
				continue
			}
			for _, a := range s.Areas() {
				if a.ID() != areaID {
					continue
				}
				serial.AssignSerial(char)
				a.AddCharacter(char)
				return 0, ""
			}
		}
		return 9, fmt.Sprintf("%s:fail_to_found_scenario_area:%s:%s", MODULE_MAN,
			scenID, areaID)
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", ENGINE_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}
