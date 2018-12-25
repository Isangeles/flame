/*
 * moduleman.go
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

package burn

import (
	"fmt"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/module/scenario"
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
	default:
		return 4, fmt.Sprintf("%s:no_such_option:%s", MODULE_MAN,
			cmd.OptionArgs()[0])
	}
}


// showModuleOption handles show option for moduleman CI tool,
// returns response code and message.
func showModuleOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s", MODULE_MAN,
			cmd.OptionArgs()[0])
	}

	switch cmd.TargetArgs()[0] {
	case "name":
		return 0, flame.Mod().Name()
	case "chapters":
		return 0, fmt.Sprint(flame.Mod().ChaptersIds())
	case "areachars":
		if flame.Game() == nil { // TODO: better check whether mod has chapter set.
			return 8, fmt.Sprintf("%s:no_game_loaded", MODULE_MAN)
		}
		if len(cmd.TargetArgs()) < 2 {
			return 8, fmt.Sprintf("%s:no_enought_target_args_for%s",
				MODULE_MAN, cmd.TargetArgs()[0])
		}
		areaID := cmd.TargetArgs()[1]
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
		charsList := ""
		for _, c := range area.Characters() {
			charsList = fmt.Sprintf("%s %s", charsList,
				c.SerialID())
		}
		return 0, charsList
	case "scenario":
		return 10, "unsupported yet"
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", ENGINE_MAN,
			cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}
}
