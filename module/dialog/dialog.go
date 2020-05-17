/*
 * dialog.go
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
	"strings"
	
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/req"
)

const (
	OwnerNameMacro = "@ownerName"
	TargetNameMacro = "@targetName"
)

// Struct for dialog.
type Dialog struct {
	id           string
	finished     bool
	trading      bool
	training     bool
	activeStage  *Stage
	activeStages []*Stage
	stages       []*Stage
	reqs         []req.Requirement
	owner        Talker
	target       Talker
}

// Interface for objects with dialogs.
type Talker interface {
	ID() string
	Serial() string
	Name() string
	SendChat(t string)
	Dialogs() []*Dialog
	MeetReqs(reqs ...req.Requirement) bool
	TakeModifiers(s objects.Object, mods ...effect.Modifier)
}

// New creates new dialog.
func New(data res.DialogData) *Dialog {
	d := new(Dialog)
	d.id = data.ID
	d.reqs = req.NewRequirements(data.Reqs)
	for _, sd := range data.Stages {
		p := NewStage(d, sd)
		d.stages = append(d.stages, p)
		if p.start {
			d.activeStages = append(d.activeStages, p)
		}
	}
	return d
}

// ID returns dialog ID.
func (d *Dialog) ID() string {
	return d.id
}

// Restart moves dialog to starting phase.
func (d *Dialog) Restart() {
	d.activeStages = make([]*Stage, 0)
	for _, p := range d.stages {
		if p.start {
			d.activeStages = append(d.activeStages, p)
		}
	}
	d.finished = false
}

// Stages returns all stages of the dialog.
func (d *Dialog) Stages() []*Stage {
	return d.stages
}

// SetStage sets specified stage as a active
// stage of the dialog.
func (d *Dialog) SetStage(s *Stage) {
	d.activeStage = s
}

// Stage returns active stage.
func (d *Dialog) Stage() *Stage {
	return d.activeStage
}

// Next moves dialog forward for specified
// answer.
// Returns error if there is no text
// for specified answer in dialog.
func (d *Dialog) Next(a *Answer) {
	if d.Target() == nil {
		return
	}
	d.trading = a.StartsTrade()
	d.training = a.StartsTraining()
	if a.EndsDialog() {
		d.finished = true
		return
	}
	// Search for proper stage for target.
	d.activeStages = make([]*Stage, 0)
	for _, p := range d.stages {
		if p.ordinalID == a.to {
			d.activeStages = append(d.activeStages, p)
		}
	}
	d.activeStage = talkerStage(d.Target(), d.activeStages)
	// Apply modifiers.
	d.Owner().TakeModifiers(d.Target(), a.OwnerModifiers()...)
	d.Target().TakeModifiers(d.Owner(), a.TargetModifiers()...)
	if d.Stage() == nil {
		return
	}
	d.Owner().TakeModifiers(d.Target(), d.Stage().OwnerModifiers()...)
	d.Target().TakeModifiers(d.Owner(), d.Stage().TargetModifiers()...)
}

// Finished checks if dialog is finished.
func (d *Dialog) Finished() bool {
	return d.finished
}

// Trading checks if trade between dialog
// participants should be started.
func (d *Dialog) Trading() bool {
	return d.trading
}

// Training checks if training between dialog
// participants should be started.
func (d *Dialog) Training() bool {
	return d.training
}

// Requirements returns all dialog requirements.
func (d *Dialog) Requirements() []req.Requirement {
	return d.reqs
}

// SetOwner sets dialog owner.
func (d *Dialog) SetOwner(t Talker) {
	d.owner = t
}

// Owner returns dialog owner or nil
// if dialog don't have owner.
func (d *Dialog) Owner() Talker {
	return d.owner
}

// SetTarget sets dialog target.
func (d *Dialog) SetTarget(t Talker) {
	d.target = t
	if d.activeStage == nil {
		d.activeStage = talkerStage(d.Target(), d.activeStages)
	}
}

// Target returns dialog target.
func (d *Dialog) Target() Talker {
	return d.target
}

// dialogText replaces all macros in specified
// text with proper info from owner/target.
func (d *Dialog) dialogText(t string) string {
	text := strings.ReplaceAll(t, OwnerNameMacro, d.Owner().Name())
	text = strings.ReplaceAll(text, TargetNameMacro, d.Target().Name())
	return text
}

// activeStage returns active stage for
// specified object.
func talkerStage(t Talker, stages []*Stage) *Stage {
	for _, s := range stages {
		if t.MeetReqs(s.Requirements()...) {
			return s
		}
	}
	return nil
}
