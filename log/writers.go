/*
 * writers.go
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

var (
	InfWriter infoWriter
	ErrWriter errorWriter
	DbgWriter debugWriter
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
	logInfo(s)
	return len(p), nil
}

// Write for error loger.
func (l errorWriter) Write(p []byte) (n int, err error) {
	s := string(p[:])
	logError(s)
	return len(p), nil
}

// Write for debug logger.
func (l debugWriter) Write(p []byte) (n int, err error) {
	s := string(p[:])
	logDebug(s)
	return len(p), nil
}
