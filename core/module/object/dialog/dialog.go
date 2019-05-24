/*
 * dialog.go
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
	"fmt"
	
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/req"
)

// Struct for dialog.
type Dialog struct {
	id          string
	finished    bool
	activeTexts []*Text
	texts       []*Text
	reqs        []req.Requirement
	owner       Talker
}

// Interface for objects with dialogs.
type Talker interface {
	ID() string
	Name() string
	SendChat(t string)
	Dialogs() []*Dialog
}

const (
	END_DIALOG_ID = "end"
)

// New creates new dialog.
func New(data res.DialogData) (*Dialog, error) {
	d := new(Dialog)
	d.id = data.ID
	d.reqs = req.NewRequirements(data.Reqs...)
	for _, sd := range data.Stages {
		t := NewText(sd)
		d.texts = append(d.texts, t)
		if t.start {
			d.activeTexts = append(d.activeTexts, t)
		}
	}
	if len(d.texts) < 1 {
		return nil, fmt.Errorf("no texts")
	}
	if len(d.activeTexts) < 1 {
		d.activeTexts = append(d.activeTexts, d.texts[0])
	}
	return d, nil
}

// ID returns dialog ID.
func (d *Dialog) ID() string {
	return d.id
}

// Restart moves dialog to starting text.
func (d *Dialog) Restart() {
	d.activeTexts = make([]*Text, 0)
	for _, t := range d.texts {
		if t.start {
			d.activeTexts = append(d.activeTexts, t)
		}
	}
	if len(d.activeTexts) < 1 {
		d.activeTexts = append(d.activeTexts, d.texts[0])
	}
	d.finished = false
}

// Texts returns all dialog texts for
// current dialog phase.
func (d *Dialog) Texts() []*Text {
	return d.activeTexts
}

// Next moves dialog forward for specified
// answer. Returns error if there is no text
// for specified answer in dialog.
func (d *Dialog) Next(a *Answer) {
	if a.to == END_DIALOG_ID {
		d.finished = true
		return
	}
	d.activeTexts = make([]*Text, 0)
	for _, t := range d.texts {
		if t.ordinalID == a.to {
			d.activeTexts = append(d.activeTexts, t)
		}
	}
}

// Finished checks if dialog is finished.
func (d *Dialog) Finished() bool {
	return d.finished
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
