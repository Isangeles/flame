/*
 * interpreter.go
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

// Flame engine commands interpreter
// @Isangeles
package ci

// Command interface for all commands interpreted by CI
type Command interface {
	Tool() 	     string
	String()     string
	TargetArgs() []string
	OptionArgs() []string
	Args()	     []string
}

const (
	ENGINE_MAN = "engineman"
	MODULE_MAN = "moduleman"
)

// Handles specified command,
// returns response code and message.
func HandleCommand(cmd Command) (int, string) {
	switch cmd.Tool() {
		case ENGINE_MAN:
			return handleEngineCommand(cmd)
		case MODULE_MAN:
			return handleModuleCommand(cmd)
		default:
			return 2, "ERROR_no_such_ci_tool_found:" + cmd.Tool()
	}
}
