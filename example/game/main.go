/*
 * main.go
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

// Example of loading module and creating game.
package main

import (
	"fmt"
	
	"github.com/isangeles/flame"
	"github.com/isangeles/flame/config"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/character"
)
var (
	// Example pc data.
	pcData res.CharacterBasicData = res.CharacterBasicData{
		ID:        "pc",
		Name:      "PC",
		Level:     1,
		Sex:       int(character.Male),
		Race:      int(character.Human),
		Attitude:  int(character.Friendly),
		Alignment: int(character.True_neutral),
		Str:       1,
		Con:       1,
		Dex:       1,
		Int:       1,
		Wis:       1,
	}
)

// Main function.
func main() {
	// Load flame config.
	err := config.LoadConfig()
	if err != nil {
		fmt.Printf("fail to load config:%v", err)
	}
	// Get module.
	mod, err := data.Module(config.ModulePath(), config.LangID())
	if err != nil {
		panic(fmt.Sprintf("fail to retrieve module:%v", err))
	}
	// Load module data.
	err = data.LoadModuleData(mod)
	if err != nil {
		panic(fmt.Sprintf("fail to load module data:%v", err))
	}
	// Set module.
	flame.SetModule(mod)
	// Create PC.
	pc := character.New(pcData)
	// Start game.
	g, err := flame.StartGame(pc)
	if err != nil {
		panic(fmt.Sprintf("fail to start game:%v", err))
	}
	// Print game info.
	fmt.Printf("game started\n")
	fmt.Printf("players:\n")
	for _, p := range g.Players() {
		fmt.Printf("%s#%s\n", p.ID(), p.Serial())
	}
}
