/*
 * loggers.go
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

// Package for engine loggers.
package log

import (
	"log"
	
	"github.com/isangeles/flame/core/enginelog"
)

var (
	Inf *log.Logger = log.New(enginelog.InfLog, "flame-core>", 0)
	Err *log.Logger = log.New(enginelog.ErrLog, "flame-core>", 0)
	Dbg *log.Logger = log.New(enginelog.DbgLog, "flame-debug>", 0)
)

