/*
 * object.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for object data.
type ObjectData struct {
	BasicData ObjectBasicData
	SavedData ObjectSavedData
	Items     []InventoryItemData
	Effects   []ObjectEffectData
}

// Struct for basic object data.
type ObjectBasicData struct {
	ID        string
	Name      string
	Serial    string
	HP, MaxHP int
}

// Struct for saved object data.
type ObjectSavedData struct {
	PosX, PosY float64
}

// Struct for inventory item data
// resource.
type InventoryItemData struct {
	ID         string
	Serial     string
	Trade      bool
	TradeValue int
}

// Struct for object effects data.
type ObjectEffectData struct {
	ID           string
	Serial       string
	Time         int64
	SourceID     string
	SourceSerial string
}

// Struct for object skill data.
type ObjectSkillData struct {
	ID       string
	Serial   string
	Cooldown int64
}

// Struct for object dialog data.
type ObjectDialogData struct {
	ID string
}

// Struct for flag data.
type FlagData struct {
	ID string
}

// Struct for object recipe data.
type ObjectRecipeData struct {
	ID string
}
