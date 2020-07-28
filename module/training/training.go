/*
 * training.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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

// Package with training structs.
package training

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/useaction"
)

// Struct for training.
type Training struct {
	id        string
	name      string
	info      string
	useAction *useaction.UseAction
}

// New creates new training.
func New(data res.TrainingData) *Training {
	ua := useaction.New(data.Use)
	t := Training{
		id:        data.ID,
		useAction: ua,
	}
	nameInfo := lang.Texts(t.ID())
	t.name = nameInfo[0]
	if len(nameInfo) > 1 {
		t.info = nameInfo[1]
	}
	return &t
}

// ID returns training ID.
func (t *Training) ID() string {
	return t.id
}

// Name returns training name.
func (t *Training) Name() string {
	return t.name
}

// Info return training info.
func (t *Training) Info() string {
	return t.info
}

// UseAction returns training use action.
func (t *Training) UseAction() *useaction.UseAction {
	return t.useAction
}
