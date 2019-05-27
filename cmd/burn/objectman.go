/*
 * objectman.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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
	"strings"
	
	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/module/object/area"
)

// handleObjectCommand handles specified command for game
// object.
func handleObjectCommand(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 3, fmt.Sprintf("%s:no_active_game", OBJECT_MAN)
	}
	if len(cmd.OptionArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no_option_args", OBJECT_MAN)
	}

	switch cmd.OptionArgs()[0] {
	case "show":
		return showObjectOption(cmd)
	case "add":
		return addObjectOption(cmd)
	default:
		return 4, fmt.Sprintf("%s:no_such_option:%s", OBJECT_MAN,
			cmd.OptionArgs()[0])
	}
}

// showCharOption handles show option for charman commands.
func showObjectOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s", OBJECT_MAN,
			cmd.OptionArgs()[0])
	}
	if len(cmd.Args()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_args_for:%s", OBJECT_MAN,
			cmd.OptionArgs()[0])
	}
	objects := make([]*area.Object, 0)
	for _, arg := range cmd.TargetArgs() {
		serialid := strings.Split(arg, "_")
		id := serialid[0]
		serial := serialid[len(serialid)-1]
		ob := flame.Game().Module().Chapter().AreaObject(id, serial)
		if ob == nil {
			return 5, fmt.Sprintf("%s:object_not_found:%s", OBJECT_MAN, id)
		}
		objects = append(objects, ob)
	}
	switch cmd.Args()[0] {
	case "health", "hp":
		out := ""
		for _, ob := range objects {
			out = fmt.Sprintf("%s %d", out, ob.Health())
		}
		return 0, out
	case "items":
		out := ""
		for _, ob := range objects {
			for _, it := range ob.Inventory().Items() {
				out = fmt.Sprintf("%s %s_%s", out, it.ID(), it.Serial())
			}
		}
		return 0, out
	default:
		return 6, fmt.Sprintf("%s:no_vaild_target_for_%s:'%s'", OBJECT_MAN,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}
}

// addObjectOption handles 'add' option for objectman tool.
func addObjectOption(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_target_args_for:%s", OBJECT_MAN,
			cmd.OptionArgs()[0])
	}
	if len(cmd.Args()) < 1 {
		return 5, fmt.Sprintf("%s:no_enought_args_for:%s", OBJECT_MAN,
			cmd.OptionArgs()[0])
	}
	objects := make([]*area.Object, 0)
	for _, arg := range cmd.TargetArgs() {
		serialid := strings.Split(arg, "_")
		id := serialid[0]
		serial := serialid[len(serialid)-1]
		ob := flame.Game().Module().Chapter().AreaObject(id, serial)
		if ob == nil {
			return 5, fmt.Sprintf("%s:object_not_found:%s", OBJECT_MAN, id)
		}
		objects = append(objects, ob)
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
				OBJECT_MAN, err)
		}
		for _, ob := range objects {
			err := ob.Inventory().AddItem(item)
			if err != nil {
				return 8, fmt.Sprintf("%s:fail_to_add_item_to_inventory:%v",
					OBJECT_MAN, err)
			}
		}
		return 0, ""
	default:
		return 4, fmt.Sprintf("%s:no_such_option:%s", OBJECT_MAN,
			cmd.OptionArgs()[0])
	}
}
