/*
 * useaction.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for use action data.
type UseActionData struct {
	CastMax           int64                 `xml:"cast-max,attr" json:"cast-max"`
	Cast              int64                 `xml:"cast,attr" json:"cast"`
	CooldownMax       int64                 `xml:"cooldown-max,attr" json:"cooldown-max"`
	Cooldown          int64                 `xml:"cooldown,attr" json:"cooldown"`
	UserMods          ModifiersData         `xml:"user>modifiers" json:"user-mods"`
	ObjectMods        ModifiersData         `xml:"object>modifiers" json:"object-mods"`
	TargetMods        ModifiersData         `xml:"target>modifiers" json:"target-mods"`
	TargetUserMods    ModifiersData         `xml:"target-user>modifiers" json:"target-user-mods"`
	UserEffects       []UseActionEffectData `xml:"user>effects>effect" json:"user-effects"`
	ObjectEffects     []UseActionEffectData `xml:"object>effects>effect" json:"object-effects"`
	TargetEffects     []UseActionEffectData `xml:"target>effects>effect" json:"target-effects"`
	TargetUserEffects []UseActionEffectData `xml:"target-user>effects>effect" json:"target-user-effects"`
	Requirements      ReqsData              `xml:"reqs" json:"reqs"`
}

// Struct for use action effect data.
type UseActionEffectData struct {
	ID string `xml:"id,attr" json:"id"`
}
