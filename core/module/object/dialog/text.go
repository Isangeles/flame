/*
 * text.go
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

package dialog

import (
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/effect"	
	"github.com/isangeles/flame/core/module/req"
)

// Structoo for dialog text.
type Text struct {
	id         string
	ordinalID  string
	start      bool
	reqs       []req.Requirement
	talkerMods []effect.Modifier
	ownerMods  []effect.Modifier
	answers    []*Answer
}

// NewText creates new dialog text.
func NewText(data *res.DialogStageData) *Text {
	t := new(Text)
	t.id = data.ID
	t.ordinalID = data.OrdinalID
	t.start = data.Start
	t.reqs = req.NewRequirements(data.Reqs...)
	t.talkerMods = effect.NewModifiers(data.TalkerMods...)
	t.ownerMods = effect.NewModifiers(data.OwnerMods...)
	for _, ad := range data.Answers {
		a := NewAnswer(ad)
		t.answers = append(t.answers, a)
	}
	return t
}

// ID returns dialog text ID.
func (t *Text) ID() string {
	return t.id
}

// Answers returns all dialog text answers.
func (t *Text) Answers() []*Answer {
	return t.answers
}

// Requirements requrements for dialog text.
func (t *Text) Requirements() []req.Requirement {
	return t.reqs
}

// TalkerModifiers retruns modifiers for talker.
func (t *Text) TalkerModifiers() []effect.Modifier {
	return t.talkerMods
}

// OnwerModifiers returns modifiers for dialog owner.
func (t *Text) OwnerModifiers() []effect.Modifier {
	return t.ownerMods
}
