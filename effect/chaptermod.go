/*
 * chaptermod.go
 *
 * Copyright 2023 Dariusz Sikora <ds@isangeles.dev>
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

package effect

import (
	"github.com/isangeles/flame/data/res"
)

// Struct for chapter modifier.
type ChapterMod struct {
	chapterID string
}

// NewChapterMod creates new chapter modifier.
func NewChapterMod(data res.ChapterModData) *ChapterMod {
	return &ChapterMod{data.ID}
}

// ChapterID returns ID of chapter to change to.
func (cm *ChapterMod) ChapterID() string {
	return cm.chapterID
}

// Data returns data resource for chapter modifier.
func (cm *ChapterMod) Data() res.ChapterModData {
	return res.ChapterModData{cm.ChapterID()}
}
