/*
 * effect.go
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

package res

// Struct for effect data resource.
type EffectData struct {
	ID         string
	Name       string
	Duration   int64
	Modifiers  ModifiersData
}

// Struct for modifiers data resource.
type ModifiersData struct {
	HealthMods []HealthModData
	FlagMods   []FlagModData
	QuestMods  []QuestModData
	AreaMods   []AreaModData
}

// Struct for health modifier
// data.
type HealthModData struct {
	Min, Max int
}

// Struct for flag modifier
// data.
type FlagModData struct {
	ID string
	On bool
}

// Struct for quest modifier
// data.
type QuestModData struct {
	ID string
}

// Struct for area modifier
// data.
type AreaModData struct {
	ID     string
	EnterX float64
	EnterY float64
}
