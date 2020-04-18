/*
 * main.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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

// Example of loading module and creating game.
package main

import (
	"fmt"
	
	"github.com/isangeles/flame"
	"github.com/isangeles/flame/config"
	"github.com/isangeles/flame/data"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module"
	"github.com/isangeles/flame/module/character"
)

// Main function.
func main() {
	// Load flame config.
	err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("Unable to load config: %v", err))
	}
	// Import module from config.
	modData, err := data.ImportModule(config.ModulePath())
	if err != nil {
		panic(fmt.Errorf("Unable to import module: %v", err))
	}
	mod := module.New(modData)
	// Create game.
	game := flame.NewGame(mod)
	// Create PC.
	// Example pc data.
	pcData := res.CharacterData{
		ID:        "pc",
		Name:      "PC",
		Level:     1,
		Sex:       string(character.Male),
		Race:      "rHuman",
		Attitude:  string(character.Friendly),
		Alignment: string(character.TrueNeutral),
	}
	pcData.Attributes = res.AttributesData{
		Str:       2,
		Con:       3,
		Dex:       4,
		Int:       5,
		Wis:       6,
	}
	pc := character.New(pcData)
	// Add PC to start area and set position.
	chapterConf := game.Module().Chapter().Conf()
	startArea := game.Module().Chapter().Area(chapterConf.StartArea)
	if startArea == nil {
		panic(fmt.Errorf("Start area not found: %s",
			chapterConf.StartArea))
	}
	startArea.AddCharacter(pc)
	pc.SetPosition(chapterConf.StartPosX, chapterConf.StartPosY)
	// Print game info.
	fmt.Printf("Game started\n")
	fmt.Printf("Characters:\n")
	for _, c := range game.Module().Chapter().Characters() {
		fmt.Printf("%s#%s\n", c.ID(), c.Serial())
	}
}
