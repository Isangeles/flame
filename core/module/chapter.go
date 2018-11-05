/*
 * chapter.go
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

package module

import (
	"github.com/isangeles/flame/core/module/scenario"
)

// Chapter struct represents module chapter
type Chapter struct {
	id        string
	scenarios []*scenario.Scenario
} 

// NewChapters creates new instance of module chapter.
func NewChapter(id string, scenarios []*scenario.Scenario) *Chapter {
	c := new(Chapter)
	c.id = id
	c.scenarios = scenarios
	return c
}

// Id returns chapter ID.
func (c *Chapter) Id() string {
	return c.id
}
