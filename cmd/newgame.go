/*
 * newgame.go
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/object/character"
)

var (
	playableChars []*character.Character
)

// newGameDialog starts CLI dialog for new game.
func newGameDialog() (*core.Game, error) {
	if flame.Mod() == nil {
		return nil, fmt.Errorf("no_module_loaded")
	}
	if len(playableChars) < 1 {
		return nil, fmt.Errorf(lang.Text("ui", "cli_newgame_no_chars_err"))
	}
	var pc *character.Character
	scan := bufio.NewScanner(os.Stdin)
	for accept := false; !accept; {
		fmt.Printf("%s:\n", lang.Text("ui", "cli_newgame_chars"))
		for i, c := range playableChars {
			fmt.Printf("[%d]%v\n", i, charDisplayString(c))
		}
		fmt.Printf("%s:", lang.Text("ui", "cli_newgame_select_char"))
		for scan.Scan() {
			input := scan.Text()
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("%s:%s\n",
					lang.Text("ui", "cli_nan_error"), input)
			}
			if id >= 0 && id < len(playableChars) {
				pc = playableChars[id]
				break
			}
		}

		fmt.Printf("%s:%v\n", lang.Text("ui", "cli_newgame_summary"),
			charDisplayString(pc))
		fmt.Printf("%s:", lang.Text("ui", "cli_accept_dialog"))
		scan.Scan()
		input := scan.Text()
		if input != "r" {
			accept = true
		}
	}
	var pcs []*character.Character
	pcs = append(pcs, pc)
	g, err := flame.StartGame(pcs...)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", lang.Text("ui", "cli_newgame_start_err"), err)
	}
	return g, nil 
}
