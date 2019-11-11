/*
 * healthmod.go
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
	"fmt"
	
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/rng"
	"github.com/isangeles/flame/core/module/object"
)

// Struct for health modifier.
type HealthMod struct {
	min, max int
}

// NewHealthMod creates new health modifier.
func NewHealthMod(data res.HealthModData) *HealthMod {
	hm := new(HealthMod)
	hm.min = data.Min
	hm.max = data.Max
	return hm
}

// Min returns minimal value of health modifier.
func (hm *HealthMod) Min() int {
	return hm.min
}

// Max returns maximal vallue of healtg modifier.
func (hm *HealthMod) Max() int {
	return hm.max
}

// Affect modifies targets health points.
func (hm *HealthMod) Affect(source Target, targets ...Target) {
	for _, t := range targets {
		t, ok := t.(object.Killable)
		if !ok {
			continue
		}
		val := rng.RollInt(hm.min, hm.max)
		t.SetHealth(t.Health() + val)
		logger, ok := t.(object.Logger)
		if !ok {
			continue
		}
		cmbMsg := fmt.Sprintf("%s:%s:%d", logger.Name(),
			lang.Text("ui", "ob_health"), val)
		logger.SendCombat(cmbMsg)
	}
}

// Undo undos health modification on specified targets.
func (hm *HealthMod) Undo(source Target, targets ...Target) {
}
