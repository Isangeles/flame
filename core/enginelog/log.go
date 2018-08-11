/*
 * log.go
 * 
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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
 
// Log for engine messagess
// @Isangeles
package enginelog

import (
	"log"
	"os"
	"time"
)

var (
	messages []message
	stderr *log.Logger = log.New(os.Stderr, "flame-core>", 0)
	stdout *log.Logger = log.New(os.Stdout, "flame-core>", 0)
	debug bool
)

// Info registers specified text as info message.
func Info(msg string) {
	m := message{time.Now().String(), msg, INF}
	messages = append(messages, m)
	if debug {
		stdout.Println(msg)
	}
}

// Error registers specified text as error message
func Error(msg string) {
	m := message{time.Now().String(), msg, ERR}
	messages = append(messages, m)
	stderr.Println(msg)
}

// Debug registers specified text as debug message
func Debug(msg string) {
	if debug {
		return
	}
	
	m := message{time.Now().String(), msg, ERR}
	messages = append(messages, m)
	stdout.Println(msg)
}

// EnableDebug enables debug mode
func EnableDebug(d bool) {
	debug = d
}

// Debug checks if debug mode is enabled
func IsDebug() bool {
	return debug
}

