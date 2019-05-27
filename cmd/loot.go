/*
 * cli.go
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
)

// lootDialog start CLI dialog current
// PC target loot.
func lootDialog() error {
	if game == nil {
		return fmt.Errorf("no_game_started")
	}
	if len(game.Players()) < 1 {
		return fmt.Errorf("no_players")
	}
	pc := game.Players()[0]
	tar := pc.Targets()[0]
	if tar == nil {
		return fmt.Errorf("no_target")
	}
	if !tar.Live() {
		return fmt.Errorf("tar_not_lootable")
	}
	for _, it := range tar.Inventory().Items() {
		if !it.Loot() {
			continue
		}
		pc.Inventory().AddItem(it)
		tar.Inventory().RemoveItem(it)
	}
	return nil
}
