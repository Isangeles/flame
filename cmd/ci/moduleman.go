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

package ci

import (
	"fmt"

	"github.com/isangeles/flame"
	//"github.com/isangeles/flame/core/module"
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
	case "scenario":
		scen := flame.Mod().Scenario()
		if scen == nil {
			return 8, fmt.Sprintf("%s:no_current_scenario", MODULE_MAN)
		}
		return 0, fmt.Sprintf("%s(area:%s)", scen.ID(), scen.Area().ID())
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", ENGINE_MAN,
			cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}
}
