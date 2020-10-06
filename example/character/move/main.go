/*
 * main.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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

// Example for moving game character.
package main

import (
	"fmt"
	"time"
	
	"github.com/isangeles/flame"
	"github.com/isangeles/flame/data"
	"github.com/isangeles/flame/module"
	"github.com/isangeles/flame/module/area"
	"github.com/isangeles/flame/module/character"
)

// Main function.
func main() {
	// Import game module from file system.
	modData, err := data.ImportModule("data/modules/test")
	if err != nil {
		panic(fmt.Errorf("Unable to import module: %v", err))
	}
	mod := module.New()
	mod.Apply(modData)
	// Create game and start game loop.
	game := flame.NewGame(mod)
	go update(game)
	// Retrieve chapter area.
	gameArea := game.Module().Chapter().Area("area1_main")
	// Retrieve game character to move from area.
	char := areaCharacter(gameArea, "char", "0")
	if char == nil {
		panic(fmt.Errorf("Character not found: char 0"))
	}
	// Set destination point for the character.
	char.SetDestPoint(10, 10)
	for {
		x, y := char.Position()
		fmt.Printf("Character position: %fx%f\n", x, y)
		if x == 10 && y == 10 {
			break
		}
	}
}

// update updates specified game.
func update(game *flame.Game) {
	var lastUpdate time.Time
	for {
		// Delta.
		dtNano := time.Since(lastUpdate).Nanoseconds()
		delta := dtNano / int64(time.Millisecond) // delta to milliseconds
		// Game.
		game.Update(delta)
		// Update time.
		lastUpdate = time.Now()
	}
}

// areaCharacter returns game character from specified area or nil if character
// with such ID and serial value was not found.
func areaCharacter(a *area.Area, id, serial string) *character.Character {
	for _, c := range a.Characters() {
		if c.ID() == id && c.Serial() == serial {
			return c
		}
	}
	return nil
}
