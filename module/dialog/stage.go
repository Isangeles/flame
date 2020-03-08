/*
 * stage.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/effect"	
	"github.com/isangeles/flame/module/req"
)

// Struct for dialog stage.
type Stage struct {
	dialog     *Dialog
	id         string
	text       string
	ordinalID  string
	start      bool
	reqs       []req.Requirement
	talkerMods []effect.Modifier
	ownerMods  []effect.Modifier
	answers    []*Answer
}

// NewStage creates new dialog stage.
func NewStage(dialog *Dialog, data *res.DialogStageData) *Stage {
	s := new(Stage)
	s.dialog = dialog
	s.id = data.ID
	s.ordinalID = data.OrdinalID
	s.start = data.Start
	s.reqs = req.NewRequirements(data.Reqs...)
	s.talkerMods = effect.NewModifiers(data.TalkerMods...)
	s.ownerMods = effect.NewModifiers(data.OwnerMods...)
	for _, ad := range data.Answers {
		a := NewAnswer(s.dialog, ad)
		s.answers = append(s.answers, a)
	}
	s.text = lang.Text(s.ID())
	return s
}

// ID returns dialog stage ID.
func (s *Stage) ID() string {
	return s.id
}

// Answers returns all dialog stage answers.
func (s *Stage) Answers() []*Answer {
	return s.answers
}

// Requirements returns requrements for dialog stage.
func (s *Stage) Requirements() []req.Requirement {
	return s.reqs
}

// TalkerModifiers retruns modifiers for talker.
func (s *Stage) TargetModifiers() []effect.Modifier {
	return s.talkerMods
}

// OnwerModifiers returns modifiers for dialog owner.
func (s *Stage) OwnerModifiers() []effect.Modifier {
	return s.ownerMods
}

// String returns translated text for stage.
func (s *Stage) String() string {
	return s.dialog.dialogText(s.text)
}
