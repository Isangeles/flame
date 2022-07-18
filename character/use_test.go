/*
 * use_test.go
 *
 * Copyright 2022 Dariusz Sikora <ds@isangeles.dev>
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
	"github.com/isangeles/flame/skill"
)

var (
	skillData = res.SkillData{ID: "skill"}
)

// TestUse tests use function.
func TestUse(t *testing.T) {
	char := New(charData)
	skill := skill.New(skillData)
	char.AddSkill(skill)
	err := char.Use(skill)
	if err != nil {
		t.Fatalf("Error while using object: %v", err)
	}
	char.Update(1)
	if char.Cooldown() <= 0 {
		t.Fatalf("Cooldown is not higher than 0")
	}
}

// TestUseReqsNotMeet tests requirements not meet
// error for use function.
func TestUseReqsNotMeet(t *testing.T) {
	char := New(charData)
	reqs := res.ReqsData{
		ItemReqs: []res.ItemReqData{itemReqData},
	}
	var skillData = skillData
	skillData.UseAction.Requirements = reqs
	skill := skill.New(skillData)
	char.AddSkill(skill)
	err := char.Use(skill)
	if err == nil {
		t.Fatalf("No error returned")
	}
	if err != REQS_NOT_MEET {
		t.Fatalf("Invalid error returned: %v", err)
	}
}

// TestUseNotReadyYet test not ready yet error for
// use function.
func TestUseNotReadyYet(t *testing.T) {
	char := New(charData)
	skill := skill.New(skillData)
	char.AddSkill(skill)
	err := char.Use(skill)
	if err != nil {
		t.Fatalf("Error on first use: %v", err)
	}
	char.Update(1)
	err = char.Use(skill)
	if err == nil {
		t.Fatalf("No error returned")
	}
	if err != NOT_READY_YET {
		t.Fatalf("Invalid error returned: %v", err)
	}
}

// TestUseInMove tests in move error for
// use function.
func TestUseInMove(t *testing.T) {
	char := New(charData)
	skill := skill.New(skillData)
	char.AddSkill(skill)
	char.SetDestPoint(10, 10)
	err := char.Use(skill)
	if err == nil {
		t.Fatalf("No error returned")
	}
	if err != IN_MOVE {
		t.Fatalf("Invalid error returned: %v", err)
	}
}
