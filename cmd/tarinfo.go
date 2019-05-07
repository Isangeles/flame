/*
 * tarinfo.go
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

// targetInfoDialog starts CLI dialog that prints informations
// about current PC target.
func targetInfoDialog() {
	if activePC == nil {
		fmt.Printf("%s\n", lang.TextDir(flameconf.LangPath(), "no_pc_err"))
		return
	}
	tar := activePC.Targets()[0]
	if tar == nil {
		fmt.Printf("%s\n", lang.TextDir(flameconf.LangPath(), "no_tar_err"))
		return
	}
	// Name.
	info := fmt.Sprintf("%s:%s", lang.TextDir(flameconf.LangPath(), "ob_name"),
		tar.Name())
	// Health.
	info += fmt.Sprintf("\n%s:%d", lang.TextDir(flameconf.LangPath(), "ob_health"),
		tar.Health())
	// Mana.
	info += fmt.Sprintf("\n%s:%d", lang.TextDir(flameconf.LangPath(), "ob_mana"),
		tar.Mana())
	// Print.
	fmt.Printf("%s\n", info)
}
