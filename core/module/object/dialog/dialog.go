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
	currentText *Text
	texts       []*Text
	reqs        []req.Requirement
	owner       Talker
}

// Interface for objects with dialogs.
type Talker interface {
	Name() string
	SendChat(t string)
	Dialogs() []*Dialog
}

const (
	END_DIALOG_ID = "end"
)

// NewDialog creates new dialog.
func NewDialog(data res.DialogData) (*Dialog, error) {
	d := new(Dialog)
	d.id = data.ID
	d.reqs = req.NewRequirements(data.Reqs...)
	for _, td := range data.Texts {
		t := NewText(td)
		d.texts = append(d.texts, t)
		if t.start {
			d.currentText = t
		}
	}
	if len(d.texts) < 1 {
		return nil, fmt.Errorf("no_texts")
	}
	if d.currentText == nil {
		d.currentText = d.texts[0]
	}
	return d, nil
}

// ID returns dialog ID.
func (d *Dialog) ID() string {
	return d.id
}

// Restart moves dialog to starting text.
func (d *Dialog) Restart() {
	for _, t := range d.texts {
		if t.start {
			d.currentText = t
		}
	}
}

// Text returns current dialog text.
func (d *Dialog) Text() *Text {
	return d.currentText
}

// Next moves dialog forward for specified
// answer. Returns error if there is no text
// for specified answer in dialog.
func (d *Dialog) Next(a *Answer) {
	if a.to == END_DIALOG_ID {
		d.finished = true
	}
	for _, t := range d.texts {
		if t.id == a.to {
			d.currentText = t
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
