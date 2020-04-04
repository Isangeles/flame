/*
 * guild.go
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
	"fmt"
)

// Guild struct represents chracter guild
type Guild struct {
	id string
}

// NewGuild return new guild with specified parameters
func NewGuild(id string) Guild {
	return Guild{id}
}

// ID return guild ID.
func (g Guild) ID() string {
	return g.id
}

// String returns guild ID
func (g Guild) String() string {
	return fmt.Sprintf("%s", g.id)
}
