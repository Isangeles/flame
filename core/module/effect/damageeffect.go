 /*
 * damageeffect.go
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

// Struct for damgae effect.
type DamageEffect struct {
	dmgMin, dmgMax int
	durationSec    int
	expired        bool
}

// NewDamageEffect creates new damage effect withs specified
// min/max damage and duration time in seconds.
func NewDamageEffect(dmgMin, dmgMax, duration int) *DamageEffect {
	de := new(DamageEffect)
	de.dmgMin = dmgMin
	de.dmgMax = dmgMax
	de.durationSec = duration
	return de
}

// Affect does damage to specified targets.x
func (de *DamageEffect) Affect(targets []Target) {
	for _, t := range targets {
		dmg := de.rollDMG()
		t.SetHealth(-dmg)
	}
}

// Update updates effects duration time.
func (de *DamageEffect) Update(delta int64) {
	
}

// Expired checks whether effecs is expired.
func (de *DamageEffect) Expired() bool {
	return de.expired
}

// rollDMG returns random damge value between
// dmgMin and dmgMax values.
func (de *DamageEffect) rollDMG() int {
	// TODO: roll damgae value.
	return de.dmgMin
}
