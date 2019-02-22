/*
 * defence.go
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

package object

// Interfece for 'atackable' objects.
type Target interface {
	Health() int
	Mana() int
	SetMana(val int)
	Experience() int
	SetExperience(val int)
	Live() bool
	TakeHit(h Hit)
}

// Struct for hit.
type Hit struct {
	Source Target
	Type   HitType
	Damage int
}

// Struct for hit type.
type HitType int

const (
	Hit_normal HitType = iota
	Hit_fire
	Hit_frost
	Hit_nature
	Hit_magic
)
