/*
 * command.go
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

package main

import (
	"fmt"
	"strings"
)

// Struct for CLI commands.
type Command struct {
	text, tool string
	targetArgs, optionArgs, args []string
	commandParts []string
}

// Creates new command from specified text input.
// Command structure: '[tool name] -t[target args ...] -o[option args ...] -a[args ...]'.
// Error: If specified input text is not valid text command.
func NewCommand(text string) (Command, error) {
	var c Command
	c.text = text		
	c.commandParts = strings.Split(c.text, " ")
	
	if len(c.commandParts) < 1 {
		return c, fmt.Errorf("command_to_short:'%s'", text)
	}
	
	c.tool = c.commandParts[0]
	for i := 1; i < len(c.commandParts); i++ {
		cPart := c.commandParts[i]
		switch cPart {
		case "-t", "--target":
			for j := i+1; j < len(c.commandParts); j++ {
				cPartArg := c.commandParts[j]
				if strings.HasPrefix(cPartArg, "-") {
					break;
				} 
				c.targetArgs = append(c.targetArgs, cPartArg)
			}
		case "-o", "--option":
			for j := i+1; j < len(c.commandParts); j++ {
				cPartArg := c.commandParts[j]
				if strings.HasPrefix(cPartArg, "-") {
					break;
				} 
				c.optionArgs = append(c.optionArgs, cPartArg)
			}
		case "-a", "--args":
			for j := i+1; j < len(c.commandParts); j++ {
				cPartArg := c.commandParts[j]
				if strings.HasPrefix(cPartArg, "-") {
					break;
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
func (c Command) Tool() string {
	return c.tool
}

// TargetArgs returns slice with target arguments of command.
func (c Command) TargetArgs() []string {
	return c.targetArgs
}

// OptionArgs returns slice with options arguments of command.
func (c Command) OptionArgs() []string {
	return c.optionArgs
}

// Args returns slice with command arguments.
func (c Command) Args() []string {
	return c.args
}

// String return full command text
func (c Command) String() string {
	return c.text
}
