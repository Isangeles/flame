/*
 * target.go
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

package effect

import (
	"github.com/isangeles/flame/core/module/req"
)

// Interfece for effects targets.
type Target interface {
	ID() string
	Serial() string
	Name() string
	Health() int
	SetHealth(val int)
	Mana() int
	SetMana(val int)
	Experience() int
	SetExperience(val int)
	Live() bool
	Effects() []*Effect
	HitEffects() []*Effect
	TakeEffect(e *Effect)
	AddEffect(e *Effect)
	SetPosition(x, y float64)
	Position() (x, y float64)
	SendCombat(msg string)
	CombatLog() chan string
	ChatLog() chan string
	PrivateLog() chan string
	MeetReqs(r ...req.Requirement) bool
}
