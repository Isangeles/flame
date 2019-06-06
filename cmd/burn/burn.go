/*
 * burn.go
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

// Burn is Flame engine commands interpreter.
package burn

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

// Interaface for all expressions interpreted by CI.
// Expressions are multiple commands connected by
// special character(e.g. pipe). 
type Expression interface {
	Commands() []Command
	Type() ExpressionType
	String() string
}

// Type for expression type.
type ExpressionType int

const (
	// CI tools names.
	ENGINE_MAN = "engineman"
	MODULE_MAN = "moduleman"
	CHAR_MAN   = "charman"
	OBJECT_MAN = "objectman"

	ID_SERIAL_SEP = "#"
	// Expr types.
	PIPE_ARG_EXP ExpressionType = iota
	PIPE_TAR_ARG_EXP
	SEQ_EXP
	NO_EXP
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
	tools[OBJECT_MAN] = handleObjectCommand
}

// AddToolHandler adds specified command handling function as
// CI tool with specified name.
func AddToolHandler(name string, handler func(cmd Command) (int, string)) {
	tools[name] = handler
}

// HandleCommand handles specified command,
// returns response code and message.
func HandleCommand(cmd Command) (int, string) {
	for tool, handleFunc := range tools {
		if cmd.Tool() == tool {
			return handleFunc(cmd)
		}
	}
	return 2, fmt.Sprintf("no_such_ci_tool_found:'%s'",
		cmd.Tool())
}

// HandleExpression handles specified expression,
// returns response code and massage.
func HandleExpression(exp Expression) (int, string) {
	switch(exp.Type()) {
	case PIPE_ARG_EXP:
		return HandleArgsPipe(exp.Commands()...)
	case PIPE_TAR_ARG_EXP:
		return HandleTargetArgsPipe(exp.Commands()...)
	case NO_EXP:
		return HandleCommand(exp.Commands()[0])
	default:
		return 2, fmt.Sprintf("unknow expression type")
	}
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

// argSerialID parses specified command arg to
// game object ID and serial value.
// Argument format: [id]ID_SERIAL_SEP[serial].
// In case of error returns empty serial.
func argSerialID(arg string) (string, string) {
	serialid := strings.Split(arg, ID_SERIAL_SEP)
	if len(serialid) < 2 {
		return arg, ""
	}
	return serialid[0], serialid[1]
}
