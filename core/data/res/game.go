/*
 * game.go
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

package res

// Struct for game data.
type GameData struct {
	Name    string
	Chapter ChapterData
}

// Struct for game chapter
// data.
type ChapterData struct {
	ID        string
	Scenarios []ScenarioData
}

// Struct for chapter scenario
// data.
type ScenarioData struct {
	ID   string
	Area AreaData
}

// Struct for scenario area
// data.
type AreaData struct {
	ID       string
	Chars    []CharacterData
	Objects  []ObjectData
	Subareas []AreaData
}
