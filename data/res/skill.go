/*
 * skill.go
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

import (
	"encoding/xml"
)

// Struct for skills data.
type SkillsData struct {
	XMLName xml.Name    `xml:"skills" json:"-"`
	Skills  []SkillData `xml:"skill", json:"skills"`
}

// Struct for skill data.
type SkillData struct {
	ID       string            `xml:"id,attr" json:"id"`
	Name     string            `xml:"name,attr" json:"name"`
	UseAction UseActionData    `xml:"use" json:"use"`
}
