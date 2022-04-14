/*
 * req_test.go
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

package character

import (
	"testing"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/req"
)

var (
	charData      = res.CharacterData{ID: "char"}
	healthReqData = res.HealthReqData{10, false}
	manaReqData   = res.ManaReqData{10, false}
)

// TestMeetReqHealth tests meet requirement check functions
// for health requirement.
func TestMeetReqHealth(t *testing.T) {
	// Meet
	char := New(charData)
	char.SetHealth(15)
	healthReq := req.NewHealth(healthReqData)
	if !char.MeetReq(healthReq) {
		t.Errorf("Requirement should be meet: required health: %d, character health: %d",
			healthReq.Value(), char.Health())
	}
	// Not meet.
	char.SetHealth(5)
	healthReq = req.NewHealth(healthReqData)
	if char.MeetReq(healthReq) {
		t.Errorf("Requirement should not be meet: required health: %d, character health: %d",
			healthReq.Value(), char.Health())
	}
}

// TestMeetReqManaMeet tests meet requirement check function
// for mana requirement.
func TestMeetReqMana(t *testing.T) {
	// Meet.
	char := New(charData)
	char.SetMana(15)
	manaReq := req.NewMana(manaReqData)
	if !char.MeetReq(manaReq) {
		t.Errorf("Requirement should be meet: required mana: %d, character mana: %d",
			manaReq.Value(), char.Mana())
	}
	// Not meet.
	char.SetMana(5)
	manaReq = req.NewMana(manaReqData)
	if char.MeetReq(manaReq) {
		t.Errorf("Requirement should not be meet: required mana: %d, character mana: %d",
			manaReq.Value(), char.Mana())
	}
}
