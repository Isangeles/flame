/*
 * dialog.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"encoding/xml"
)

// Struct for dialogs data.
type DialogsData struct {
	XMLName xml.Name     `xml:"dialogs" json:"-"`
	Dialogs []DialogData `xml:"dialog" json:"dialogs'`
}

// Dialog data struct.
type DialogData struct {
	XMLName xml.Name          `xml:"dialog" json:"-"`
	ID      string            `xml:"id,attr" json:"id"`
	Stages  []DialogStageData `xml:"stage" json:"sages"`
	Reqs    ReqsData          `xml:"reqs" json:"reqs"`
}

// Dialog text data struct.
type DialogStageData struct {
	ID         string             `xml:"id,attr" json:"id"`
	OrdinalID  string             `xml:"ordinal,attr" json:"ordinal"`
	Start      bool               `xml:"start,attr" json:"start"`
	Answers    []DialogAnswerData `xml:"answer" json:"answers"`
	Reqs       ReqsData           `xml:"reqs" json:"reqs"`
	TargetMods ModifiersData      `xml:"target>modifiers" json:"target-mods"`
	OwnerMods  ModifiersData      `xml:"owner>modifiers" json:"owner-mods`
}

// Dialog answer data struct.
type DialogAnswerData struct {
	ID         string        `xml:"id,attr" json:"id"`
	To         string        `xml:"to,attr" json:"to"`
	End        bool          `xml:"end,attr" json:"end"`
	Trade      bool          `xml:"trade,attr" json:"trade"`
	Training   bool          `xml:"train,attr" json:"training"`
	Reqs       ReqsData      `xml:"reqs" json:"reqs"`
	TargetMods ModifiersData `xml:"target>modifiers" json:"target-mods"`
	OwnerMods  ModifiersData `xml:"owner>modifiers" json:"owner-mods"`
}
