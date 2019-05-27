/*
 * useskill.go
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
	"github.com/isangeles/flame/core/module/object/skill"
	flameconf "github.com/isangeles/flame/config"
)

// useSkillDialog starts CLI dialog for
// using skills.
func useSkillDialog() error {
	if activePC == nil {
		return fmt.Errorf("no active PC")
	}
	// List skills.
	fmt.Printf("%s:\n", lang.TextDir(flameconf.LangPath(), "use_skill_skills"))
	skills := activePC.Skills()
	for i, s := range skills {
		fmt.Printf("[%d]%s\n", i, s.Name())
	}
	// Select skill.
	scan := bufio.NewScanner(os.Stdin)
	var skill *skill.Skill
	for skill == nil {
		fmt.Printf("%s:", lang.TextDir(flameconf.LangPath(), "use_skill_select"))
		scan.Scan()
		input := scan.Text()
		id, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("%s:%s\n", lang.TextDir(flameconf.LangPath(),
				"cli_nan_error"), input)
			continue
		}
		if id < 0 || id > len(skills)-1 {
			fmt.Printf("%s\n", lang.TextDir(flameconf.LangPath(),
				"use_skill_no_skill_id_err"))
			continue
		}
		skill = skills[id]
	}
	activePC.UseSkill(skill)
	return nil
}
