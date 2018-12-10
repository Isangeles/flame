/*
 * character.go
 * 
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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
 
// character package provides game character struct representation and
// other types for game characters.
package character

import (
	"fmt"
)

// Character struct represents game character.
type Character struct {
	id         string
	serial     string
	name       string
	level      int
	sex        Gender
	race       Race
	attitude   Attitude
	alignment  Alignment
	guild      Guild
	attributes Attributes
	posX, posY float64
	sight      float64
}

// NewCharacter returns new character with specified parameters.
func NewCharacter(id string, name string, level int, sex Gender, race Race,
	attitude Attitude, guild Guild, attributes Attributes,
	alignment Alignment) *Character {
	c := Character{
		id: id,
		name: name,
		level: level,
		sex: sex,
		race: race,
		attitude: attitude,
		guild: guild,
		attributes: attributes,
		alignment: alignment,
		sight: 300,
	}
	return &c
}

// Id returns character ID.
func (c *Character) ID() string {
	return c.id
}

// Serial returns serial value.
func (c *Character) Serial() string {
	return c.serial
}

// SerialId returns character ID and serial value
// in form: [ID]_[serial].
func (c *Character) SerialID() string {
	return fmt.Sprintf("%s_%s", c.ID(), c.serial)
}

// Name returns character name.
func (c *Character) Name() string {
	return c.name
}

// Level returns character level.
func (c *Character) Level() int {
	return c.level
}

// Gender returns character gender.
func (c *Character) Gender() Gender {
	return c.sex
}

// Race returns character race.
func (c *Character) Race() Race {
	return c.race
}

// Attitude returns character attitude.
func (c *Character) Attitude() Attitude {
	return c.attitude
}

// Guild returns character guild.
func (c *Character) Guild() Guild {
	return c.guild
}

// Attributes returns character attributes.
func (c *Character) Attributes() Attributes {
	return c.attributes
}

// Alignment returns character alignment
func (c *Character) Alignment() Alignment {
	return c.alignment
}

// Position returns current character position.
func (c *Character) Position() (float64, float64) {
	return c.posX, c.posY
}

// SightRange returns current sight range.
func (c *Character) SightRange() float64 {
	return c.sight
}

// SetPosition sets specified XY position as current
// position of character.
func (c *Character) SetPosition(x, y float64) {
	c.posX, c.posY = x, y
}

// SetSerial sets specified serial value for this
// character.
func (c *Character) SetSerial(serial string) {
	c.serial = serial
}

// HasSerial checks whether character has
// serial value.
func (c *Character) HasSerial() bool {
	return c.Serial() != ""
}

// String returns string with character parameters spearated by ', '.
func (c *Character) String() string {
	return fmt.Sprintf("%s, %d, %v, %v, %v, %v, %s",
		c.id, c.level, c.sex, c.race, c.attitude, c.guild, c.attributes,
		c.alignment) 
}
