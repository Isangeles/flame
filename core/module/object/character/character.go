/*
 * character.go
 * 
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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
 
// character package provides game character struct and
// other types for game characters.
package character

import (
	"fmt"

	"github.com/isangeles/flame/core/module/object/item"
)

const (
	base_exp = 1000
)

// Character struct represents game character.
type Character struct {
	id            string
	serial        string
	name          string
	level         int
	hp, maxHP     int
	mana, maxMana int
	exp, maxExp   int
	live          bool
	agony         bool
	sex           Gender
	race          Race
	attitude      Attitude
	alignment     Alignment
	guild         Guild
	attributes    Attributes
	posX, posY    float64
	destX, destY  float64
	inventory     *item.Inventory
	equipment     *Equipment
}

// NewCharacter returns new character with specified parameters.
func NewCharacter(id string, name string, level int, sex Gender, race Race,
	attitude Attitude, guild Guild, attributes Attributes,
	alignment Alignment) *Character {
	c := Character{
		id: id,
		name: name,
		sex: sex,
		race: race,
		attitude: attitude,
		guild: guild,
		attributes: attributes,
		alignment: alignment,
	}
	c.live = true
	c.inventory = item.NewInventory(c.Attributes().Lift())
	c.equipment = newEquipment(&c)
	for i := 0; i < level; i++ {
		oldMaxExp := c.MaxExperience()
		c.levelup()
		c.SetExperience(oldMaxExp)
	}
	return &c
}

// Update updates character.
func (c *Character) Update() {
	// Move to dest point.
	if c.InMove() {
		if c.posX < c.destX {
			c.Move(c.posX+1, c.posY) 
		}
		if c.posX > c.destX {
			c.Move(c.posX-1, c.posY)
		}
		if c.posY < c.destY {
			c.Move(c.posX, c.posY+1)
		}
		if c.posY > c.destY {
			c.Move(c.posX, c.posY-1)
		}
	}
	// Check experience value.
	if c.Experience() >= c.MaxExperience() {
		c.levelup()
	}
	// Check health value.
	if c.Health() <= c.agonyHP() {
		c.agony = true
	} else if c.Agony() {
		c.agony = false
	}
	if c.Health() <= 0 {
		c.live = false
	} else if !c.Live() {
		c.live = true
	}
}

// ID returns character ID.
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

// Health returns current value of health
// points.
func (c *Character) Health() int {
	return c.hp
}

// MaxHealth returns maximal value of
// health points.
func (c *Character) MaxHealth() int {
	return c.attributes.Health() + (base_health * c.Level())
}

// Mana returns current value of mana
// points.
func (c *Character) Mana() int {
	return c.mana
}

// MaxMana returns maximal value of mana
// points.
func (c *Character) MaxMana() int {
	return c.attributes.Mana() + (base_mana * c.Level()/2)
}

// Experience returns current value of experience
// points.
func (c *Character) Experience() int {
	return c.exp
}

// MaxExperience returns maximal value of
// experience points.
func (c *Character) MaxExperience() int {
	return c.maxExp
}

// Live checks wheter character is live.
func (c *Character) Live() bool {
	return c.live
}

// Agony check wheter character is in
// agony state.
func (c *Character) Agony() bool {
	return c.agony
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

// DestPoint return current destination point position.
func (c *Character) DestPoint() (float64, float64) {
	return c.destX, c.destY
}

// SightRange returns current sight range.
func (c *Character) SightRange() float64 {
	return c.attributes.Sight()
}

// Inventory returns character inventory.
func (c *Character) Inventory() *item.Inventory {
	return c.inventory
}

// Equipment returns character equipment.
func (c *Character) Equipment() *Equipment {
	return c.equipment
}

// SetName sets specified text as character
// display name.
func (c *Character) SetName(name string) {
	c.name = name
}

// SetHealth sets specified value as current
// amount of health points.
func (c *Character) SetHealth(hp int) {
	c.hp = hp
}

// SetMana sets specified value as current
// amount of mana points.
func (c *Character) SetMana(mana int) {
	c.mana = mana
}

// SetExperience sets specified value as current
// amount of experience points.
func (c *Character) SetExperience(exp int) {
	c.exp = exp
}

// SetPosition sets specified XY position as current
// position of character and current destination point.
func (c *Character) SetPosition(x, y float64) {
	c.posX, c.posY = x, y
	c.destX, c.destY = x, y
}

// Move moves characters to specified XY position
// without changing destination point.
func (c *Character) Move(x, y float64) {
	c.posX, c.posY = x, y
}

// SetDestPoint sets specified XY position as current
// destionation point of character.
func (c *Character) SetDestPoint(x, y float64) {
	c.destX, c.destY = x, y
}

// SetSerial sets specified serial value for this
// character.
func (c *Character) SetSerial(serial string) {
	c.serial = serial
}

// InMove checks whether character is moving.
func (c *Character) InMove() bool {
	if c.posX != c.destX || c.posY != c.destY {
		return true
	} else {
		return false
	}
}

// levelup promotes character to next level.
func (c * Character) levelup() {
	c.level += 1
	c.SetHealth(c.MaxHealth())
	c.SetMana(c.MaxMana())
	c.maxExp = base_exp * c.Level()
}

// agonyHP returns value of health causing
// agony state.
func (c *Character) agonyHP() int {
	return 10 / 100 * c.MaxHealth() 
}
