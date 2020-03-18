/*
 * race.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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

package character

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/res/lang"
)

// Struct for character race.
type Race struct {
	id       string
	name     string
	playable bool
}

// NewRace creates new race.
func NewRace(data res.RaceData) *Race {
	r := new(Race)
	r.id = data.ID
	r.playable = data.Playable
	r.name = lang.Text(r.ID())
	return r
}

// ID retruns race ID.
func (r *Race) ID() string {
	return r.id
}

// Playable checks if race is playable.
func (r *Race) Playable() bool {
	return r.playable
}
