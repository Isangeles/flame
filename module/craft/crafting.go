/*
 * crafting.go
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

// Package for crafting structs.
package craft

import (
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
	"github.com/isangeles/flame/module/objects"
)

// Interface for objects with crafting.
type Crafter interface {
	objects.Object
	Crafting() *Crafting
}

// Struct fo crafting object.
type Crafting struct {
	owner   Crafter
	recipes map[string]*Recipe
}

// NewCrafting creates new crafting object.
func NewCrafting(data res.CraftingData, object Crafter) *Crafting {
	c := new(Crafting)
	c.owner = object
	c.recipes = make(map[string]*Recipe)
	for _, craftRecipeData := range data.Recipes {
		recipeData := res.Recipe(craftRecipeData.ID)
		if recipeData == nil {
			log.Err.Printf("crafting: %s#%s: unable to retrieve recipe: %s",
				c.owner.ID(), c.owner.Serial(), craftRecipeData.ID)
			continue
		}
		recipe := NewRecipe(*recipeData)
		c.AddRecipes(recipe)
	}
	return c
}

// Recipes returns all recipes in crafting object.
func (c *Crafting) Recipes() (recipes []*Recipe) {
	for _, r := range c.recipes {
		recipes = append(recipes, r)
	}
	return
}

// AddRecipes adds specified recipes to crafting object.
func (c *Crafting) AddRecipes(recipes ...*Recipe) {
	for _, r := range recipes {
		c.recipes[r.ID()] = r
	}
}

// Data creates data resource for crafting.
func (c *Crafting) Data() res.CraftingData {
	data := res.CraftingData{}
	for _, r := range c.Recipes() {
		recipeData := res.CraftingRecipeData{r.ID()}
		data.Recipes = append(data.Recipes, recipeData)
	}
	return data
}
