/*
 * rng.go
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

// Package for random number generator.
package rng

import (
	"math/rand"
	"time"
)

var (
	rng *rand.Rand
)

// On init.
func init() {
	src := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(src)
}

// RollInt generates random integer from
// specified range.
func RollInt(min, max int) int {
	neg := false
	if min < 1 && max < 1 { // handling negative range
		neg = true
	}
	if min < 1 {
		min *= -1
	}
	if max < 1 {
		max *= -1
	}
	roll := min + rng.Intn(max - min)
	if neg {
		return -roll
	}
	return roll
}
