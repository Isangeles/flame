/*
 * log.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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
// To use log, create new log.Loger with one of
// log writers as writer.
package enginelog

import (
	"fmt"
	"time"
)

var (
	messages []message
	debuging = false
)

// Messages returns all messages from log.
func Messages() []message {
	return messages
}

// SetDebug toggles debug mode.
func SetDebug(dbg bool) {
	debuging = dbg
}

// Info registers specified text as info message.
func logInfo(msg string) {
	m := message{time.Now(), msg, INF}
	messages = append(messages, m)
	if debuging {	
		fmt.Print(msg)
	}
}

// Error registers specified text as error message.
func logError(msg string) {
	m := message{time.Now(), msg, ERR}
	messages = append(messages, m)
	fmt.Print(msg)
}

// Debug registers specified text as debug message.
func logDebug(msg string) {
	if !debuging {
		return
	}
	m := message{time.Now(), msg, ERR}
	messages = append(messages, m)
	fmt.Print(msg)
}
