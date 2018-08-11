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
	
	"github.com/isangeles/flame/core"
	//"github.com/isangeles/flame/core/module"
)

// Handles specified module command,
// returns response code and message.
func handleModuleCommand(cmd Command) (int, string) {
	if !core.Mod().Loaded() {
		return 3, fmt.Sprintf("%s:no_module_loaded", MODULE_MAN)
	} 
	if len(cmd.OptionArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no_option_args", MODULE_MAN)
	}
	
	switch cmd.OptionArgs()[0] {
		case "name":
			return 0, core.Mod().Name()
		default:
			return 4, fmt.Sprintf("%s:no_such_option:%s", MODULE_MAN, cmd.OptionArgs()[0])
	}
} 
