/*
 * log.go
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

// Package with engine log.
// To use log, create new log.Loger with one of log
// writers as writer.
// Use PrintStdOut boolean to enable/disable printing
// messages on standard output.
package log

import (
	"fmt"
	"time"
)

var (
	PrintStdOut = false
	messages    []Message
	counter     int64
)

// Messages returns all messages from log.
func Messages() []Message {
	return messages
}

// Info registers specified text as info message.
func logInfo(msg string) {
	m := Message{
		date:  time.Now(),
		text:  msg,
		mType: INF,
	}
	m.id = fmt.Sprintf("%d", counter)
	counter += 1
	messages = append(messages, m)
	if PrintStdOut {
		fmt.Print(msg)
	}
}

// Error registers specified text as error message.
func logError(msg string) {
	m := Message{
		date:  time.Now(),
		text:  msg,
		mType: ERR,
	}
	m.id = fmt.Sprintf("%d", counter)
	counter += 1
	messages = append(messages, m)
	if PrintStdOut {
		fmt.Print(msg)
	}
}

// Debug registers specified text as debug message.
func logDebug(msg string) {
	m := Message{
		date:  time.Now(),
		text:  msg,
		mType: DBG,
	}
	m.id = fmt.Sprintf("%d", counter)
	counter += 1
	messages = append(messages, m)
	if PrintStdOut {
		fmt.Print(msg)
	}
}
