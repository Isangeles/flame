/*
 * quest_test.go
 *
 * Copyright 2025 Dariusz Sikora <ds@isangeles.dev>
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

package quest

import (
	"github.com/isangeles/flame/flag"
	"github.com/isangeles/flame/req"
)

// Test quester struct.
type testQuester struct {
	id, serial string
	journal    *Journal
	meetReqs   bool
	flags      []flag.Flag
}

// ID returns ID value.
func (tq *testQuester) ID() string {
	return tq.id
}

// Serial returns serial value.
func (tq *testQuester) Serial() string {
	return tq.serial
}

// SetSerial sets serial value.
func (tq *testQuester) SetSerial(serial string) {
	tq.serial = serial
}

// Journal returns test quester journal.
func (tq *testQuester) Journal() *Journal {
	return tq.journal
}

// MeetReqs returns meetReqs value.
func (tq *testQuester) MeetReqs(reqs ...req.Requirement) bool {
	return tq.meetReqs
}

// AddFlag adds quester flag.
func (tq *testQuester) AddFlag(flag flag.Flag) {
	tq.flags = append(tq.flags, flag)
}

// RemoveFlag removes quester flag.
func (tq *testQuester) RemoveFlag(flg flag.Flag) {
	var flags []flag.Flag
	for _, f := range tq.flags {
		if f.ID() != flg.ID() {
			flags = append(flags, f)
		}
	}
	tq.flags = flags
}

// Flags returns quester flags.
func (tq *testQuester) Flags() []flag.Flag {
	return tq.flags
}
