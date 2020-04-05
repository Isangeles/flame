/*
 * craft.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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

// Struct for recipe data.
type RecipeData struct {
	ID       string
	Category string
	Results  []RecipeResultData
	Reqs     ReqsData
	Cast     int64
}

// Struct for recipe result data.
type RecipeResultData struct {
	ID     string
	Amount int
	Item   ItemData
}

// Struct for crafting data.
type CraftingData struct {
	Recipes []CraftingRecipeData `xml:"recipe"`
}

// Struct for crafting recipe data.
type CraftingRecipeData struct {
	ID string `xml:"id,attr"`
}
