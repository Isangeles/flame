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

// Flame engine commands interpreter.
package ci

import (
	"fmt"
	"strings"
)

// Command interface for all commands interpreted by CI.
type Command interface {
	Tool() string
	String() string
	TargetArgs() []string
	OptionArgs() []string
	Args() []string
	AddArgs(args ...string)
	AddTargetArgs(args ...string)
}

const (
	ENGINE_MAN = "engineman"
	MODULE_MAN = "moduleman"
	CHAR_MAN   = "charman"
)

var (
	tools map[string]func(cmd Command) (int, string)
)

// On init.
func init() {
	tools = make(map[string]func(cmd Command) (int, string), 0)
	tools[ENGINE_MAN] = handleEngineCommand
	tools[MODULE_MAN] = handleModuleCommand
	tools[CHAR_MAN] = handleCharCommand
}

// HandleCommand handles specified command,
// returns response code and message.
func HandleCommand(cmd Command) (int, string) {
	for tool, handleFunc := range tools {
		if cmd.Tool() == tool {
			return handleFunc(cmd)
		}
	}
	return 2, fmt.Sprintf("cmd:%s:ERROR_no_such_ci_tool_found:'%s'",
		cmd, cmd.Tool())
}

// AddToolHandler adds specified command handling function as
// CI tool with specified name.
func AddToolHandler(name string, handler func(cmd Command) (int, string)) {
	tools[name] = handler
}

// HandleArgsPipe handles specified commands
// connected with pipe('|').
// Pipe pushes out from command on the left to
// command on right as arguments.
func HandleArgsPipe(cmds ...Command) (res int, out string) {
	for _, cmd := range cmds {
		res, out = pipeArgs(cmd, out)
		if res != 0 {
			return res, out
		}
	}
	return
}

// HandleTargetArgsPipe handles specified commands
// connected with pipe('|').
// Pipe pushes out from command on the left to
// command on right as target arguments.
func HandleTargetArgsPipe(cmds ...Command) (res int, out string) {
	for _, cmd := range cmds {
		res, out = pipeTargetArgs(cmd, out)
		if res != 0 {
			return res, out
		}
	}
	return
}

// pipeArgs pushes specified text(out from previous command)
// to specified command as arguments, and executes
// specified command.
func pipeArgs(cmd Command, out string) (int, string) {
	args := strings.Split(strings.TrimSpace(out), " ")
	cmd.AddArgs(args...)
	return HandleCommand(cmd)
}

// pipeTargetArgs pushes specified text(out from previous command)
// to specified command as target arguments, and executes
// specified command.
func pipeTargetArgs(cmd Command, out string) (int, string) {
	args := strings.Split(strings.TrimSpace(out), " ")
	cmd.AddTargetArgs(args...)
	return HandleCommand(cmd)
}
