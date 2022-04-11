/*
 * mana_test.go
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

var manaReqData = res.ManaReqData{100, true}

// TestNewMana tests creating new mana requirement.
func TestNewMana(t *testing.T) {
	req := NewMana(manaReqData)
	if req.Value() != 100 {
		t.Errorf("Inavlid mana value: %d != 100", req.Value())
	}
	if !req.Less() {
		t.Errorf("Inavlid less value: %v != true", req.Less())
	}
}

// TestManaData tests data function of mana
// requirement struct.
func TestManaData(t *testing.T) {
	req := NewMana(manaReqData)
	reqData := req.Data()
	if reqData != manaReqData {
		t.Errorf("Inavlid requirement data: %v != %v", reqData, manaReqData)
	}
}
