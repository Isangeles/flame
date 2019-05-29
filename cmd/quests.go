/*
 * quests.go
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
	"fmt"

	"github.com/isangeles/flame/core/data/text/lang"
	flameconf "github.com/isangeles/flame/config"
)

// questsDialog starts quests journal CLI dialog.
func questsDialog() error {
	if activePC == nil {
		return fmt.Errorf("no active PC")
	}
	fmt.Printf("%s:\n", lang.TextDir(flameconf.LangPath(), "quests_list"))
	mod := game.Module()
	questsLangPath := mod.Chapter().Conf().QuestsLangPath()
	for i, q := range activePC.Journal().Quests() {
		questInfo := lang.AllText(questsLangPath, q.ID())
		fmt.Printf("[%d]%s\n", i, questInfo[0])
		if len(questInfo) > 1 {
			fmt.Printf("\t%s\n", questInfo[1])
		}
		if q.ActiveStage() == nil {
			continue
		}
		stageInfo := lang.AllText(questsLangPath, q.ActiveStage().ID())
		fmt.Printf("\t%s\n", stageInfo[0])
	}
	return nil
}
