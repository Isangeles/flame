/*
 * main.go
 *
 * Copyright 2020-2024 Dariusz Sikora <ds@isangeles.dev>
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

// Example for moving module character.
package main

import (
	"fmt"
	"time"
	
	"github.com/isangeles/flame"
	"github.com/isangeles/flame/data"
	"github.com/isangeles/flame/area"
	"github.com/isangeles/flame/character"
)

// Main function.
func main() {
	// Import game module from file system.
	modData, err := data.ImportModuleDir("module")
	if err != nil {
		panic(fmt.Errorf("Unable to import module: %v", err))
	}
	mod := flame.NewModule(modData)
	// Create game and start game loop.
	go update(mod)
	// Retrieve chapter area.
	charArea := mod.Chapter().Area("area")
	// Retrieve game character to move from area.
	char := areaCharacter(charArea, "char", "0")
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
func update(mod *flame.Module) {
	var lastUpdate time.Time
	for {
		// Delta.
		dtNano := time.Since(lastUpdate).Nanoseconds()
		delta := dtNano / int64(time.Millisecond) // delta to milliseconds
		// Game.
		mod.Update(delta)
		// Update time.
		lastUpdate = time.Now()
		time.Sleep(time.Duration(16) * time.Millisecond)
	}
}

// areaCharacter returns game character from specified area or nil if character
// with such ID and serial value was not found.
func areaCharacter(a *area.Area, id, serial string) *character.Character {
	for _, object := range a.Objects() {
		if object.ID() == id && object.Serial() == serial {
			char, _ := object.(*character.Character)
			return char
		}
	}
	return nil
}
