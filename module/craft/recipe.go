/*
 * recipe.go
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

package craft

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/useaction"
)

// Struct for recipes.
type Recipe struct {
	id        string
	category  string
	useAction *useaction.UseAction
}

// NewRecipe creates new crafting recipe.
func NewRecipe(data res.RecipeData) *Recipe {
	r := Recipe{
		id:        data.ID,
		category:  data.Category,
		useAction: useaction.New(data.UseAction),
	}
	return &r
}

// Update updates recipe.
func (r *Recipe) Update(delta int64) {
	r.UseAction().Update(delta)
}

// ID returns recipe ID.
func (r *Recipe) ID() string {
	return r.id
}

// Category returns ID of recipe category.
func (r *Recipe) Category() string {
	return r.category
}

// UseAction returns recipe use action.
func (r *Recipe) UseAction() *useaction.UseAction {
	return r.useAction
}
