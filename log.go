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

package flame

import (
	"github.com/isangeles/flame/core/enginelog"
)

var (
	InfLog infoWriter
	ErrLog errorWriter
	DbgLog debugWriter
)

// infoWriter writes info messages to engine log.
type infoWriter struct {
}

// errorWriter writes error messages to engine log.
type errorWriter struct {
}

// debugWriter writes debug messages to engine log.
type debugWriter struct {
}

// Write for info logger.
func (l infoWriter) Write(p []byte) (n int, err error) {
	s := string(p[:])
	enginelog.Info(s)
	return len(p), nil
}

// Write for error loger.
func (l errorWriter) Write(p []byte) (n int, err error) {
	s := string(p[:])
	enginelog.Error(s)
	return len(p), nil
}

// Write for debug logger.
func (l debugWriter) Write(p []byte) (n int, err error) {
	s := string(p[:])
	enginelog.Debug(s)
	return len(p), nil
}
