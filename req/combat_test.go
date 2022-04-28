/*
 * combat_test.go
 *
 * Copyright 2022 Dariusz Sikora <dev@isangeles.pl>
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

package req

import (
	"testing"

	"github.com/isangeles/flame/data/res"
)

var combatReqData = res.CombatReqData{true}

// TestNewCombat tests creating new combat requirement.
func TestNewCombat(t *testing.T) {
	req := NewCombat(combatReqData)
	if !req.Combat() {
		t.Errorf("Invalid combat value: %v != true", req.Combat())
	}
}

// TestCombatData tests data function of combat
// requirement struct.
func TestCombatData(t *testing.T) {
	req := NewCombat(combatReqData)
	reqData := req.Data()
	if reqData != combatReqData {
		t.Errorf("Invalid requirement data: %v != %v", reqData, combatReqData)
	}
}
