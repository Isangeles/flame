/*
 * phase.go
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
	"github.com/isangeles/flame/core/module/effect"	
	"github.com/isangeles/flame/core/module/req"
)

// Struct for dialog phase.
type Phase struct {
	id         string
	ordinalID  string
	start      bool
	reqs       []req.Requirement
	talkerMods []effect.Modifier
	ownerMods  []effect.Modifier
	answers    []*Answer
}

// NewPhase creates new dialog phase.
func NewPhase(data *res.DialogStageData) *Phase {
	p := new(Phase)
	p.id = data.ID
	p.ordinalID = data.OrdinalID
	p.start = data.Start
	p.reqs = req.NewRequirements(data.Reqs...)
	p.talkerMods = effect.NewModifiers(data.TalkerMods...)
	p.ownerMods = effect.NewModifiers(data.OwnerMods...)
	for _, ad := range data.Answers {
		a := NewAnswer(ad)
		p.answers = append(p.answers, a)
	}
	return p
}

// ID returns dialog phase ID.
func (p *Phase) ID() string {
	return p.id
}

// Answers returns all dialog phase answers.
func (p *Phase) Answers() []*Answer {
	return p.answers
}

// Requirements returns requrements for dialog phase.
func (p *Phase) Requirements() []req.Requirement {
	return p.reqs
}

// TalkerModifiers retruns modifiers for talker.
func (p *Phase) TalkerModifiers() []effect.Modifier {
	return p.talkerMods
}

// OnwerModifiers returns modifiers for dialog owner.
func (p *Phase) OwnerModifiers() []effect.Modifier {
	return p.ownerMods
}
