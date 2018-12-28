/*
 * stdcommand.go
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

package syntax

import (
	"fmt"
	"strings"
)

// Struct for CLI standard commands.
// Standard commands structure:
// '[tool name] -t[target args ...] -o[option args ...] -a[args ...]'.
type STDCommand struct {
	text, tool                   string
	targetArgs, optionArgs, args []string
	commandParts                 []string
}

// Creates new standard command from specified text input.
// Standard command structure:
// '[tool name] -t[target args ...] -o[option args ...] -a[args ...]'.
// Error: If specified input text is not valid text command.
func NewSTDCommand(text string) (*STDCommand, error) {
	c := new(STDCommand)
	c.text = text
	c.commandParts = strings.Split(c.text, " ")

	if len(c.commandParts) < 1 {
		return c, fmt.Errorf("command_to_short:'%s'", text)
	}

	c.tool = c.commandParts[0]
	for i := 1; i < len(c.commandParts); i++ {
		cPart := strings.TrimSpace(c.commandParts[i])
		switch cPart {
		case "-t", "--target":
			for j := i + 1; j < len(c.commandParts); j++ {
				cPartArg := c.commandParts[j]
				if strings.HasPrefix(cPartArg, "-") {
					break
				}
				c.targetArgs = append(c.targetArgs, cPartArg)
			}
		case "-o", "--option":
			for j := i + 1; j < len(c.commandParts); j++ {
				cPartArg := c.commandParts[j]
				if strings.HasPrefix(cPartArg, "-") {
					break
				}
				c.optionArgs = append(c.optionArgs, cPartArg)
			}
		case "-a", "--args":
			for j := i + 1; j < len(c.commandParts); j++ {
				cPartArg := c.commandParts[j]
				if strings.HasPrefix(cPartArg, "-") {
					break
				}
				c.args = append(c.args, cPartArg)
			}
		default:
			continue
		}
	}

	return c, nil
}

// Tool return command tool name.
func (c *STDCommand) Tool() string {
	return c.tool
}

// TargetArgs returns slice with target arguments of command.
func (c *STDCommand) TargetArgs() []string {
	return c.targetArgs
}

// OptionArgs returns slice with options arguments of command.
func (c *STDCommand) OptionArgs() []string {
	return c.optionArgs
}

// Args returns slice with command arguments.
func (c *STDCommand) Args() []string {
	return c.args
}

// AddArgs adds specified text values as
// command args.
func (c *STDCommand) AddArgs(args ...string) {
	c.args = append(c.args, args...)
}

// AddTargetArgs adds specified text values
// as command target args.
func (c *STDCommand) AddTargetArgs(args ...string) {
	c.targetArgs = append(c.targetArgs, args...)
}

// String return full command text.
func (c *STDCommand) String() string {
	return c.text
}