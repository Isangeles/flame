/*
 * charman.go
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
	"strings"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/module/flag"
	"github.com/isangeles/flame/core/module/object"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/item"
	"github.com/isangeles/flame/core/module/object/skill"
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
	case "export", "save":
		return exportCharOption(cmd)
	case "add":
		return addCharOption(cmd)
	case "remove":
		return removeCharOption(cmd)
	case "equip":
		return equipCharOption(cmd)
	case "cast":
		return castCharOption(cmd)
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
	case "health", "hp":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.OptionArgs()[1])
		}
		val, err := strconv.Atoi(cmd.Args()[1])
		if err != nil {
			return 8, fmt.Sprintf("%s:invalid_argument:%s", CHAR_MAN,
				cmd.OptionArgs()[1])
		}
		for _, char := range chars {
			char.SetHealth(val)
		}
		return 0, ""
	case "mana":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.OptionArgs()[1])
		}
		val, err := strconv.Atoi(cmd.Args()[1])
		if err != nil {
			return 8, fmt.Sprintf("%s:invalid_argument:%s", CHAR_MAN,
				cmd.OptionArgs()[1])
		}
		for _, char := range chars {
			char.SetMana(val)
		}
		return 0, ""
	case "target":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.OptionArgs()[1])
		}
		serialid := strings.Split(cmd.Args()[1], "_")
		id := serialid[0]
		serial := serialid[len(serialid)-1]
		tar := flame.Game().Module().Target(id, serial)
		if tar == nil {
			return 8, fmt.Sprintf("%s:object_not_found:%s",
				CHAR_MAN, cmd.Args()[1])
		}
		for _, char := range chars {
			char.SetTarget(tar)
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
	case "serial":
		out := ""
		for _, char := range chars {
			out = fmt.Sprintf("%s ", char.Serial())
		}
		return 0, out
	case "items":
		out := ""
		for _, char := range chars {
			for _, it := range char.Inventory().Items() {
				out += fmt.Sprintf("%s_%s ", it.ID(), it.Serial())
			}
		}
		return 0, out
	case "equipment":
		out := ""
		for _, char := range chars {
			for _, it := range char.Equipment().Items() {
				out += fmt.Sprintf("%s_%s:", it.ID(), it.Serial())
				for _, s := range it.Slots() {
					out += fmt.Sprintf("%s ", s.ID())
				}
				out += "\n"
			}
		}
		return 0, out
	case "effects":
		out := ""
		for _, char := range chars {
			for _, e := range char.Effects() {
				out += fmt.Sprintf("%s ", e.ID()+"_"+e.Serial())
			}
		}
		return 0, out
	case "skills":
		out := ""
		for _, char := range chars {
			for _, s := range char.Skills() {
				out += fmt.Sprintf("%s ", s.ID()+"_"+s.Serial())
			}
		}
		return 0, out
	case "dialogs":
		out := ""
		for _, char := range chars {
			for _, d := range char.Dialogs() {
				out += fmt.Sprintf("%s ", d.ID())
			}
		}
		return 0, out
	case "quests":
		out := ""
		for _, char := range chars {
			for _, q := range char.Journal().Quests() {
				out += fmt.Sprintf("%s ", q.ID())
			}
		}
		return 0, out
	case "flags":
		out := ""
		for _, char := range chars {
			for _, f := range char.Flags() {
				out += fmt.Sprintf("%s ", f.ID())
			}
		}
		return 0, out
	case "health", "hp":
		out := ""
		for _, char := range chars {
			out += fmt.Sprintf("%s %d ", out, char.Health())
		}
		return 0, out
	case "max-health", "max-hp":
		out := ""
		for _, char := range chars {
			out += fmt.Sprintf("%d ", char.MaxHealth())
		}
		return 0, out
	case "mana":
		out := ""
		for _, char := range chars {
			out += fmt.Sprintf("%d ", char.Mana())
		}
		return 0, out
	case "max-mana":
		out := ""
		for _, char := range chars {
			out += fmt.Sprintf("%d ", char.MaxMana())
		}
		return 0, out
	case "range":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.Args()[0])
		}
		tar := flame.Game().Module().Character(cmd.Args()[1])
		if tar == nil {
			return 8, fmt.Sprintf("%s:object_not_found:%s",
				CHAR_MAN, cmd.Args()[1])
		}
		out := ""
		for _, char := range chars {
			out += fmt.Sprintf("%f ", object.Range(char, tar))
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
	serialID := cmd.TargetArgs()[0]
	var char *character.Character
	for _, pc := range flame.Game().Players() {
		if pc.ID()+"_"+pc.Serial() == serialID {
			char = pc
		}
	}
	if char == nil {
		return 5, fmt.Sprintf("%s:character_not_found:%s", CHAR_MAN,
			cmd.TargetArgs()[0])
	}
	err := data.ExportCharacter(char, flame.Game().Module().Conf().CharactersPath())
	if err != nil {
		return 8, fmt.Sprintf("%s:%v", CHAR_MAN, err)
	}
	return 0, ""
}

// addCharOption handles 'add' option for charman CI tool.
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
		id := cmd.Args()[1]
		item, err := data.Item(id)
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
	case "effect":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.Args()[0])
		}
		for _, char := range chars {
			effectID := cmd.Args()[1]
			effect, err := data.Effect(flame.Game().Module(), effectID)
			if err != nil {
				return 8, fmt.Sprintf("%s:fail_to_retrieve_effect:%v",
					CHAR_MAN, err)
			}
			char.AddEffect(effect)
		}
		return 0, ""
	case "skill":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.Args()[0])
		}
		for _, char := range chars {
			id := cmd.Args()[1]
			skill, err := data.Skill(id)
			if err != nil {
				return 8, fmt.Sprintf("%s:fail_to_retrieve_skill:%v",
					CHAR_MAN, err)
			}
			char.AddSkill(skill)
		}
		return 0, ""
	case "quest":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.Args()[0])
		}
		for _, char := range chars {
			id := cmd.Args()[1]
			q, err := data.Quest(id)
			if err != nil {
				return 8, fmt.Sprintf("%s:fail_to_retrieve_quest:%v",
					CHAR_MAN, err)
			}
			char.Journal().AddQuest(q)
		}
		return 0, ""
	case "flag":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.Args()[0])
		}
		for _, char := range chars {
			id := cmd.Args()[1]
			flag := flag.Flag(id)
			char.AddFlag(flag)
		}
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}

// removeCharOption handles 'remove' option for charman CI tool.
func removeCharOption(cmd Command) (int, string) {
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
		serialid := strings.Split(cmd.Args()[1], "_")
		id := serialid[0]
		serial := serialid[len(serialid)-1]
		for _, char := range chars {
			for _, i := range char.Inventory().Items() {
				if i.ID() == id && i.Serial() == serial {
					char.Inventory().RemoveItem(i)
				}
			}
		}
		return 0, ""
	case "effect":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.Args()[0])
		}
		for _, char := range chars {
			effectID := cmd.Args()[1]
			effect, err := data.Effect(flame.Game().Module(), effectID)
			if err != nil {
				return 8, fmt.Sprintf("%s:fail_to_retrieve_effect:%v",
					CHAR_MAN, err)
			}
			char.RemoveEffect(effect)
		}
		return 0, ""
	case "skill":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.Args()[0])
		}
		for _, char := range chars {
			id := cmd.Args()[1]
			for _, s := range char.Skills() {
				if s.ID() == id {
					char.RemoveSkill(s)
				}
			}
		}
		return 0, ""
	case "quest":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.Args()[0])
		}
		for _, char := range chars {
			id := cmd.Args()[1]
			for _, q := range char.Journal().Quests() {
				if q.ID() == id {
					char.Journal().RemoveQuest(q)
				}
			}
		}
		return 0, ""
	case "flag":
		if len(cmd.Args()) < 2 {
			return 7, fmt.Sprintf("%s:no_enought_args_for:%s",
				CHAR_MAN, cmd.Args()[0])
		}
		for _, char := range chars {
			id := cmd.Args()[1]
			flag := flag.Flag(id)
			char.RemoveFlag(flag)
		}
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}

// equipCharOption handles 'equip' option for charman CI tool.
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
			serialid := strings.Split(cmd.Args()[1], "_")
			id := serialid[0]
			serial := serialid[len(serialid)-1]
			it := char.Inventory().Item(id, serial)
			if it == nil {
				return 8, fmt.Sprintf("%s:%s:fail_to_retrieve_item_from_inventory:%s",
					CHAR_MAN, char.SerialID(), serialid)
			}
			eit, ok := it.(item.Equiper)
			if !ok {
				return 8, fmt.Sprintf("%s:%s_%s:item_not_equipable:%s",
					CHAR_MAN, char.ID(), char.Serial(), serialid)
			}
			err := char.Equipment().EquipHandRight(eit)
			if err != nil {
				return 8, fmt.Sprintf("%s:%s_%s:fail_to_equip_item:%v",
					CHAR_MAN, char.ID(), char.Serial(), err)
			}
		}
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}

// castCharOption handles 'cast' option for charman CI tool.
func castCharOption(cmd Command) (int, string) {
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
	case "skill":
		for _, char := range chars {
			serialID := cmd.Args()[1]
			var skill *skill.Skill
			for _, s := range char.Skills() {
				if s.ID()+"_"+s.Serial() == serialID {
					skill = s
				}
			}
			if skill == nil {
				return 5, fmt.Sprintf("%s:character:%s:skill_not_known:%s",
					CHAR_MAN, char.ID()+"_"+char.Serial(), serialID)
			}
			char.UseSkill(skill)
		}
		return 0, ""
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", CHAR_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}
