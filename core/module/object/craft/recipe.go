/*
 * recipe.go
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

// Package for crafting structs.
package craft

import (
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/req"
)

// Struct for recipes.
type Recipe struct {
	id    string
	catID string
	res   []res.RecipeResultData
	reqs  []req.Requirement
}

// NewRecipe creates new crafting recipe.
func NewRecipe(data res.RecipeData) *Recipe {
	r := new(Recipe)
	r.id = data.ID
	r.catID = data.Category
	r.res = data.Results
	r.reqs = req.NewRequirements(data.Reqs...)
	return r
}

// ID returns recipe ID.
func (r *Recipe) ID() string {
	return r.id
}

// CategoryID returns ID of recipe category.
func (r *Recipe) CategoryID() string {
	return r.catID
}

// Reqs returns recipe requirements
func (r *Recipe) Reqs() []req.Requirement {
	return r.reqs
}
