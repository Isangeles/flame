/*
 * chapterconf.go
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
	StartScenID string
	StartAreaID string
	Scenarios   []string
	NextChapter string
}

// FullPath returns path to chapter directory.
func (cc ChapterConf) FullPath() string {
	return filepath.FromSlash(cc.Path)
}

// ScenariosPath returns path to chapter
// scenarios directory.
func (cc ChapterConf) ScenariosPath() string {
	return filepath.FromSlash(cc.FullPath() +
		"/areas/scenarios")
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

// NPCPath returns path to chapter NPCs
// directory.
func (cc ChapterConf) NPCPath() string {
	return filepath.FromSlash(cc.FullPath() + "/npc")
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
