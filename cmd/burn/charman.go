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

package burn

import (
	"fmt"
	"strconv"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/item"
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
	case "add":
		return addCharOption(cmd)
	case "equip":
		return equipCharOption(cmd)
	default:
		return 4, fmt.Sprintf("%s:no_such_option:%s", CHAR_MAN,
			cmd.OptionArgs()[0])
	}
}

// setCharOption handles set option for charman commands.
func setCharOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s", CHAR_MAN,
			cmd.OptionArgs()[0])
	}
	if len(cmd.Args()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_args_for:%s", CHAR_MAN,
			cmd.OptionArgs()[0])
	}
	chars := make([]*character.Character, 0)
	for _, charID := range cmd.TargetArgs() {
		char := flame.Game().Module().Character(charID)
		if char == nil {
			return 5, fmt.Sprintf("%s:character_not_found:%s", CHAR_MAN,
				cmd.TargetArgs()[1])
		}
		chars = append(chars, char)
	}

	switch cmd.Args()[0] {
	case "position":
		if len(cmd.Args()) < 3 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.OptionArgs()[1])
		}
		x, err := strconv.ParseFloat(cmd.Args()[1], 64)
		if err != nil {
			return 8, fmt.Sprintf("%s:invalid_argument:%s", CHAR_MAN,
				cmd.OptionArgs()[0])
		}
		y, err := strconv.ParseFloat(cmd.Args()[2], 64)
		if err != nil {
			return 8, fmt.Sprintf("%s:invalid_argument:%s", CHAR_MAN,
				cmd.OptionArgs()[1])
		}
		for _, char := range chars {
			char.SetPosition(x, y)
		}
		return 0, ""
	case "dest":
		if len(cmd.Args()) < 3 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.OptionArgs()[1])
		}
		x, err := strconv.ParseFloat(cmd.Args()[1], 64)
		if err != nil {
			return 8, fmt.Sprintf("%s:invalid_argument:%s", CHAR_MAN,
				cmd.OptionArgs()[0])
		}
		y, err := strconv.ParseFloat(cmd.Args()[2], 64)
		if err != nil {
			return 8, fmt.Sprintf("%s:invalid_argument:%s", CHAR_MAN,
				cmd.OptionArgs()[1])
		}
		for _, char := range chars {
			char.SetDestPoint(x, y)
		}
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.TargetArgs()[0])
	}
}

// showCharOption handles show option for charman commands.
func showCharOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s", CHAR_MAN,
			cmd.OptionArgs()[0])
	}
	if len(cmd.Args()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_args_for:%s", CHAR_MAN,
			cmd.OptionArgs()[0])
	}
	chars := make([]*character.Character, 0)
	for _, id := range cmd.TargetArgs() {
		char := flame.Game().Module().Character(id)
		if char == nil {
			return 5, fmt.Sprintf("%s:character_not_found:%s", CHAR_MAN,
				id)
		}
		chars = append(chars, char)
	}

	switch cmd.Args()[0] {
	case "position":
		out := ""
		for _, char := range chars {		
			x, y := char.Position()
			out = fmt.Sprintf("%s%fx%f ", out,
				x, y)
		}
		return 0, out
	case "id":
		out := ""
		for _, char := range chars {
			out = fmt.Sprintf("%s ", char.ID())
		}
		return 0, out
	case "serialid":
		out := ""
		for _, char := range chars {
			out = fmt.Sprintf("%s ", char.SerialID())
		}
		return 0, out
	case "items":
		out := ""
		for _, char := range chars {
			for _, it := range char.Inventory().Items() {
				out += fmt.Sprintf("%s ", it.SerialID())
			}
		}
		return 0, out
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}

// exportEngineOption handles 'export' option for charman CI tool.
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

// addEngineOption handles 'add' option for charman CI tool.
func addCharOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s",
			CHAR_MAN, cmd.OptionArgs()[0])
	}
	if len(cmd.Args()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_args_for:%s",
			CHAR_MAN, cmd.OptionArgs()[0])
	}
	chars := make([]*character.Character, 0)
	for _, id := range cmd.TargetArgs() {
		char := flame.Game().Module().Character(id)
		if char == nil {
			return 5, fmt.Sprintf("%s:character_not_found:%s", CHAR_MAN,
				id)
		}
		chars = append(chars, char)
	}

	switch cmd.Args()[0] {
	case "item":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
			CHAR_MAN, cmd.Args()[0])
		}
		itemID := cmd.Args()[1]
		item, err := data.Item(flame.Game().Module(), itemID)
		if err != nil {
			return 8, fmt.Sprintf("%s:fail_to_retrieve_item:%v",
				CHAR_MAN, err)
		}
		for _, char := range chars {
			err = char.Inventory().AddItem(item)
			if err != nil {
				return 8, fmt.Sprintf("%s:fail_to_add_item_to_inventory:%v",
					CHAR_MAN, err)
			}
		}
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}

// equipEngineOption handles 'equip' option for charman CI tool.
func equipCharOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s",
			CHAR_MAN, cmd.OptionArgs()[0])
	}
	if len(cmd.Args()) < 2 {
		return 5, fmt.Sprintf("%s:no_enought_args_for:%s",
			CHAR_MAN, cmd.OptionArgs()[0])
	}
	chars := make([]*character.Character, 0)
	for _, id := range cmd.TargetArgs() {
		char := flame.Game().Module().Character(id)
		if char == nil {
			return 5, fmt.Sprintf("%s:character_not_found:%s", CHAR_MAN,
				id)
		}
		chars = append(chars, char)
	}

	switch cmd.Args()[0] {
	case "hand-right":
		for _, char := range chars {
			itID := cmd.Args()[1]
			it := char.Inventory().Item(itID)
			if it == nil {
				return 8, fmt.Sprintf("%s:%s:fail_to_retrieve_item_from_inventory:%s",
					CHAR_MAN, char.SerialID(), itID)
			}
			eit, ok := it.(item.Equiper)
			if !ok {
				return 8, fmt.Sprintf("%s:%s:item_not_equipable:%s",
					CHAR_MAN, char.SerialID(), it.SerialID())
			}
			err := char.Equipment().EquipHandRight(eit)
			if err != nil {
				return 8, fmt.Sprintf("%s:%s:fail_to_equip_item:%v",
					CHAR_MAN, char.SerialID(), err)
			}
		}
		return 0, "" 
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}
