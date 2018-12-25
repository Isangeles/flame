/*
 * loadmodule_test.go
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
	"flag"
	"testing"
)

// TestCommand represents CI command.
type TestCommand struct {
	tool                 string
	target, option, args []string
}

func (t TestCommand) Tool() string {
	return t.tool
}

func (t TestCommand) String() string {
	return fmt.Sprintf("tool:%s, target:%s, option%s, args:%s",
		t.tool, t.target, t.option, t.args)
}

func (t TestCommand) TargetArgs() []string {
	return t.target
}

func (t TestCommand) OptionArgs() []string {
	return t.option
}


func (t TestCommand) Args() []string {
	return t.args
}


var (
	path = flag.String("path", "", "System path of module to test")
	modulename = flag.String("module", "", "Module name")
)

// Test for module loading.
func TestLoadModule(t *testing.T) {
	command := TestCommand{"engineman", []string{"module"}, []string{"load"},
		[]string{*modulename, *path}}
	res, out := HandleCommand(command)
	if res > 0 {
		t.Errorf("command_error[%d]:%s", res, out)
	} else {
		t.Logf("command_succes[%d]:%s", res, out)
	}
}
