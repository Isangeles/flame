/*
 * race.go
 *
 * Copyright 2021 Dariusz Sikora <dev@isangeles.pl>
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

import (
	"encoding/xml"
)

// Struct for races data.
type RacesData struct {
	XMLName xml.Name   `xml:"races" json:"-"`
	Races   []RaceData `xml:"race" json:"races"`
}

// Struct for race data.
type RaceData struct {
	ID       string            `xml:"id,attr" json:"id"`
	Playable bool              `xml:"playable,attr" json:"playable"`
	Skills   []ObjectSkillData `xml:"skills>skill" json:"skills"`
}
