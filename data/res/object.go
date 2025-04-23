/*
 * object.go
 *
 * Copyright 2019-2025 Dariusz Sikora <ds@isangeles.dev>
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

package res

import (
	"time"
)

// Struct for object effects data.
type ObjectEffectData struct {
	ID           string `xml:"id,attr" json:"id"`
	Serial       string `xml:"serial,attr" json:"serial"`
	Time         int64  `xml:"time,attr" json:"time"`
	SourceID     string `xml:"source-id,attr" json:"source-id"`
	SourceSerial string `xml:"source-serial,attr" json:"source-serial"`
}

// Struct for object skill data.
type ObjectSkillData struct {
	ID       string `xml:"id,attr" json:"id"`
	Cooldown int64  `xml:"cooldown,attr" json:"cooldown"`
}

// Struct for object dialog data.
type ObjectDialogData struct {
	ID        string `xml:"id,attr" json:"id"`
	StartedID string `xml:started-id,attr" json:"started-id"`
	Stage     string `xml:"stage,attr" json:"stage"`
}

// Struct for object log data.
type ObjectLogData struct {
	Messages []ObjectLogMessageData `xml:"message" json:"messages"`
}

// Struct for object log message data.
type ObjectLogMessageData struct {
	Translated bool      `xml:"translated,attr" json:"translated"`
	Text       string    `xml:"text,attr" json:"text"`
	Time       time.Time `xml:"time,attr" json:"time"`
}

// Struct for flag data.
type FlagData struct {
	ID string `xml:"id,attr" json:"id"`
}

// Struct for kill data.
type KillData struct {
	ID         string `xml:"id,attr" json:"id"`
	Serial     string `xml:"serial,attr" json:"serial"`
	Experience int    `xml:"experience,attr" json:"experience"`
}

// Struct for data of object with serial ID.
type SerialObjectData struct {
	ID     string `xml:"id,attr" json:"id"`
	Serial string `xml:"serial,attr" json:"serial"`
}
