/*
 * talk.go
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/dialog"
	flameconf "github.com/isangeles/flame/config"
)

// talkDialog starts talk CLI dialog with specified
// game dialog.
func talkDialog(d *dialog.Dialog) {
	if game == nil {
		fmt.Printf("%s\n", lang.TextDir(flameconf.LangPath(), "no_game_err"))
		return
	}
	mod := game.Module()
	langPath := mod.Chapter().Conf().LangPath()
	scan := bufio.NewScanner(os.Stdin)
	d.Restart()
	for !d.Finished() {
		fmt.Printf("%s:\n", lang.TextDir(flameconf.LangPath(), "talk_dialog"))
		phase := dialogPhase(d.Texts(), activePC)
		if phase == nil {
			fmt.Printf("%s\n", lang.TextDir(flameconf.LangPath(), "talk_no_phase_err"))
			return
		}
		dlgText := lang.AllText(langPath, "dialogs", phase.ID())[0]
		fmt.Printf("[%s]:%s\n", d.Owner().Name(), dlgText)
		var (
			ans     *dialog.Answer
			ansText string
		)
		for ans == nil {
			fmt.Printf("%s:\n", lang.TextDir(flameconf.LangPath(), "talk_answers"))
			for i, a := range phase.Answers() {
				ansText = lang.AllText(langPath, "dialogs", a.ID())[0]
				fmt.Printf("[%d]%s\n", i, ansText)
			}
			fmt.Printf("%s:", lang.TextDir(flameconf.LangPath(), "talk_answers_select"))
			scan.Scan()
			input := scan.Text()
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("%s:%s\n", lang.TextDir(flameconf.LangPath(), "cli_nan_error"),
					input)
				continue
			}
			if id < 0 || id > len(phase.Answers())-1 {
				fmt.Printf("%s\n", lang.TextDir(flameconf.LangPath(), "talk_no_answer_id_err"))
				continue
			}
			ans = phase.Answers()[id]
		}
		fmt.Printf("[%s]:%s\n", activePC.Name(), ansText)
		d.Next(ans)
	}
}

// dialogPhase selects dialog phase with requirements met by specified character.
func dialogPhase(texts []*dialog.Text, char *character.Character) *dialog.Text {
	for _, t := range texts {
		if char.MeetReqs(t.Requirements()) {
			return t
		}
	}
	return nil
}
