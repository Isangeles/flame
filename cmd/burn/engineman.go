/*
 * engineman.go
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
	"path/filepath"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/config"
	"github.com/isangeles/flame/core/data"
)

// Handles specified engine command,
// returns response code and message.
func handleEngineCommand(cmd Command) (int, string) {
	if len(cmd.OptionArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no_option_args", ENGINE_MAN)
	}

	switch cmd.OptionArgs()[0] {
	case "version":
		return 0, flame.VERSION
	case "lang":
		if len(cmd.Args()) < 1 {
			return 0, config.LangID()
		}

		err := config.SetLang(cmd.Args()[0])
		if err != nil {
			return 8, fmt.Sprintf("%s:lang_set_fail:%v", ENGINE_MAN, err)
		}

		return 0, ""
	case "show":
		return showEngineOption(cmd)
	case "load":
		return loadEngineOption(cmd)
	case "start", "new":
		return startEngineOption(cmd)
	case "set":
		return setEngineOption(cmd)
	case "save":
		return saveEngineOption(cmd)
	default:
		return 4, fmt.Sprintf("%s:no_such_option:%s", ENGINE_MAN,
			cmd.OptionArgs()[0])
	}
}

// showEngineCommand handles show option for engineman CI tool,
// returns response code and message.
func showEngineOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s",
			ENGINE_MAN, cmd.OptionArgs()[0])
	}

	switch cmd.TargetArgs()[0] {
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", ENGINE_MAN,
			cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}
}

// loadEngineOption handles load option for engineman CI tool,
// returns response code and message.
func loadEngineOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s", ENGINE_MAN,
			cmd.OptionArgs()[0])
	}

	switch cmd.TargetArgs()[0] {
	case "module", "mod":
		if len(cmd.Args()) < 1 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				ENGINE_MAN, cmd.TargetArgs()[1])
		}
		var modPath string
		if len(cmd.Args()) > 1 {
			modPath = filepath.FromSlash(cmd.Args()[1] + "/" +
				cmd.Args()[0])
		} else {
			modPath = filepath.FromSlash("data/modules/" +
				cmd.Args()[0])
		}
		m, err := data.Module(modPath, config.LangID())
		if err != nil {
			return 8, fmt.Sprintf("%s:module_load_fail:%s",
				ENGINE_MAN, err)
		}
		flame.SetModule(m)
		return 0, ""
	case "game":
		if len(cmd.Args()) < 1 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				ENGINE_MAN, cmd.TargetArgs()[0])
		}
		if flame.Mod() == nil {
			return 7, fmt.Sprintf("%s:no_module_loaded",
				ENGINE_MAN)
		}
		_, err := flame.LoadGame(cmd.Args()[0])
		if err != nil {
			return 8, fmt.Sprintf("%s:fail_to_load_game:%v",
				ENGINE_MAN, err)
		}
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", ENGINE_MAN,
			cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}
}

// saveEngineOption handles save option for enineman CI tool,
// returns response code and message.
func saveEngineOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s",
			ENGINE_MAN, cmd.OptionArgs()[0])
	}

	switch cmd.TargetArgs()[0] {
	case "config":
		err := config.SaveConfig()
		if err != nil {
			return 8, fmt.Sprintf("%s:config_save_fail:%v",
				ENGINE_MAN, err)
		}

		return 0, ""
	case "game":
		if len(cmd.Args()) < 1 {
			return 7, fmt.Sprintf("%s:not_enought_args_for:%s",
				ENGINE_MAN, cmd.TargetArgs()[0])
		}
		if flame.Mod() == nil {
			return 7, fmt.Sprintf("%s:no_module_loaded",
				ENGINE_MAN)
		}
		savePath := config.ModuleSavegamesPath()
		saveName := cmd.Args()[0]
		err := data.SaveGame(flame.Game(), savePath, saveName)
		if err != nil {
			return 8, fmt.Sprintf("%s:fail_to_save_game:%v",
				ENGINE_MAN, err)
		}
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'",
			ENGINE_MAN, cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}
}

// startEngineOption handles start option for engineman CI tool,
// returns response code and message.
func startEngineOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s",
			ENGINE_MAN, cmd.OptionArgs()[0])
	}

	switch cmd.TargetArgs()[0] {
	case "game":
		if len(cmd.Args()) < 1 {
			return 7, fmt.Sprintf("%s:not_enought_args_for:%s",
				ENGINE_MAN, cmd.OptionArgs()[1])
		}
		if flame.Mod() == nil {
			return 7, fmt.Sprintf("no_module_loaded")
		}
		id, serial := argSerialID(cmd.Args()[0])
		pc := flame.Mod().Chapter().Character(id, serial)
		if pc == nil {
			return 7, fmt.Sprintf("not_found_character_with_id:'%s'",
				cmd.Args()[0])
		}
		_, err := flame.StartGame(pc)
		if err != nil {
			return 8, fmt.Sprintf("%s:new_game_start_fail:%s",
				ENGINE_MAN, err)
		}
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'",
			ENGINE_MAN, cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}
}

// setEngineOption handles 'set' option for engineman CI tool,
// returns response code and message.
func setEngineOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s",
			ENGINE_MAN, cmd.OptionArgs()[0])
	}

	switch cmd.TargetArgs()[0] {
	case "debug":
		if len(cmd.Args()) < 1 {
			return 7, fmt.Sprintf("%s:not_enought_args_for:%s",
				ENGINE_MAN, cmd.TargetArgs()[0])
		}
		arg0 := cmd.Args()[0]
		dbgMode := (arg0 == "true" || arg0 == "on")
		config.SetDebug(dbgMode)
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'",
			ENGINE_MAN, cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}

}
