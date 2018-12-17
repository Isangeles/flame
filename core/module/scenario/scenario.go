/*
 * scenario.go
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

// scenario package provides structs and functions for
// game world areas.
package scenario

import (
	"fmt"
)

// Scenario struct represents area scenario
type Scenario struct {
	id       string
	mainarea *Area
	areas    []*Area
}

// NewScenario returns new instance of scenario.
func NewScenario(id string, mainarea *Area, subareas []*Area) (*Scenario) {
	s := new(Scenario)
	s.id = id
	s.mainarea = mainarea
	s.areas = append(s.areas, s.Mainarea())
	s.areas = append(s.areas, subareas...)
	return s
}

// ID returns scenario ID.
func (s *Scenario) ID() string {
	return s.id
}

// Mainarea returns main area of
// scenario.
func (s *Scenario) Mainarea() *Area {
	return s.mainarea
}

// Area returns current scenario area.
func (s *Scenario) Area(id string) (*Area, error) {
	for _, a := range s.areas {
		if a.ID() == id {
			return a, nil
		}
	}
	return nil, fmt.Errorf("area_not_found:%s", id)
}

// Areas returns all scenario areas.
func (s *Scenario) Areas() []*Area {
	return s.areas
}
