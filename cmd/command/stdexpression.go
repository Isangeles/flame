/*
 * stdexpression.go
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

package command

import (
	"fmt"
	"strings"

	"github.com/isangeles/flame/cmd/ci"
)

const (
	STD_ARG_PIPE = " |a "
	STD_TAR_ARG_PIPE = " |t "
)

// Type for standard expressions.
// Syntax:
// Arg pipe delimiter: ' |a '.
// Target arg pipe delimiter: ' |t ',
// e.g. 'moduleman -o show -t areachars [area ID] |t charman -o show -t position' - shows
// positions of all characters in area with specified ID.
type STDExpression struct {
	text     string
	commands []ci.Command
	etype  ci.ExpressionType
}

// NewStdExpression creates new standard expression from
// specified text.
func NewSTDExpression(text string) (STDExpression, error) {
	var exp STDExpression
	exp.text = text
	switch {	
	case strings.Contains(text, STD_TAR_ARG_PIPE):
		exp.etype = ci.PIPE_TAR_ARG_EXP
		cmdsText := strings.Split(text, STD_TAR_ARG_PIPE)
		for _, cmdText := range cmdsText {
			cmd, err := NewSTDCommand(strings.TrimSpace(cmdText))
			if err != nil {
				return exp, fmt.Errorf("command:%s:fail_to_buil_expression_command:%v",
					cmdText, err)
			}
			exp.commands = append(exp.commands, cmd)
		}
		return exp, nil
	default:
		exp.etype = ci.NO_EXP
		cmd, err := NewSTDCommand(strings.TrimSpace(text))
		if err != nil {
			return exp, fmt.Errorf("command:%s:fail_to_buil_expression_command:%v",
				text, err)
		}
		exp.commands = append(exp.commands, cmd)
		return exp, nil
	}
}

// Commands returns all expression commands.
func (exp STDExpression) Commands() []ci.Command {
	return exp.commands
}

// Type returns expression type.
func (exp STDExpression) Type() ci.ExpressionType {
	return exp.etype
}

// Retruns expression text.
func (exp STDExpression) String() string {
	return exp.text
}
