/*
 * chapterconf.go
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

package module

import (
	"path/filepath"
)

// Struct for chapter configurtion
// values.
type ChapterConf struct {
	ID, Path    string
	ModulePath  string
	Lang        string
	StartAreaID string
	NextChapter string
}

// FullPath returns path to chapter directory.
func (cc ChapterConf) FullPath() string {
	return filepath.FromSlash(cc.Path)
}

// AreasPath returns path to chapters areas
// directory.
func (cc ChapterConf) AreasPath() string {
	return filepath.FromSlash(cc.FullPath() + "/areas")
}

// LangPath returns path to chapter
// lang directory.
func (cc ChapterConf) LangPath() string {
	return filepath.FromSlash(cc.FullPath() + "/lang" +
		"/" + cc.Lang)
}

// CharactersPath returns path to chapter characters
// directory.
func (cc ChapterConf) CharactersPath() string {
	return filepath.FromSlash(cc.FullPath() + "/characters")
}

// ObjectsPath returns path to chapter objects
// directory.
func (cc ChapterConf) ObjectsPath() string {
	return filepath.FromSlash(cc.FullPath() + "/objects")
}

// DialogsPath returns path to chapter dialogs directory.
func (cc ChapterConf) DialogsPath() string {
	return filepath.FromSlash(cc.FullPath() + "/dialogs")
}

// QuestsPath retruns path to chapter quests directory.
func (cc ChapterConf) QuestsPath() string {
	return filepath.FromSlash(cc.FullPath() + "/quests")
}

// DialogsLangPath returns path to dialogs lang file.
func (cc ChapterConf) DialogsLangPath() string {
	return filepath.FromSlash(cc.LangPath() + "/dialogs")
}

// QuestsLangPath returns path to quests lang file.
func (cc ChapterConf) QuestsLangPath() string {
	return filepath.FromSlash(cc.LangPath() + "/quests")
}
