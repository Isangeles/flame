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
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/req"
)

// Struct for dialog.
type Dialog struct {
	id           string
	finished     bool
	trading      bool
	training     bool
	activeStages []*Stage
	stages       []*Stage
	reqs         []req.Requirement
	owner        Talker
}

// Interface for objects with dialogs.
type Talker interface {
	ID() string
	Name() string
	SendChat(t string)
	Dialogs() []*Dialog
}

// New creates new dialog.
func New(data res.DialogData) *Dialog {
	d := new(Dialog)
	d.id = data.ID
	d.reqs = req.NewRequirements(data.Reqs...)
	for _, sd := range data.Stages {
		p := NewStage(sd)
		d.stages = append(d.stages, p)
		if p.start {
			d.activeStages = append(d.activeStages, p)
		}
	}
	if len(d.activeStages) < 1 {
		d.activeStages = append(d.activeStages, d.stages[0])
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
	if len(d.activeStages) < 1 {
		d.activeStages = append(d.activeStages, d.stages[0])
	}
	d.finished = false
}

// Stages returns all active stages
// of dialog.
func (d *Dialog) Stages() []*Stage {
	return d.activeStages
}

// Next moves dialog forward for specified
// answer.
// Returns error if there is no text
// for specified answer in dialog.
func (d *Dialog) Next(a *Answer) {
	d.trading = a.StartsTrade()
	d.training = a.StartsTraining()
	if a.EndsDialog() {
		d.finished = true
		return
	}
	d.activeStages = make([]*Stage, 0)
	for _, p := range d.stages {
		if p.ordinalID == a.to {
			d.activeStages = append(d.activeStages, p)
		}
	}
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
