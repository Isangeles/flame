/*
 * message.go
 *
 * Copyright 2018-2020 Dariusz Sikora <dev@isangeles.pl>
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

package log

import (
	"time"
)

// Enum type for messages types.
// 0 - info, 1 - error, 2 - debug.
type MessageType int

const (
	INF MessageType = iota
	ERR
	DBG
)

// Struct for log messages.
type Message struct {
	id      string
	date    time.Time
	text    string
	mType   MessageType
}

// ID returns message ID.
func (m Message) ID() string {
	return m.id
}

// Text returns text content of message.
func (m Message) String() string {
	return m.text
}

// Date returns message date.
func (m Message) Date() time.Time {
	return m.date
}
