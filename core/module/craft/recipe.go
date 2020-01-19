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
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/data/res/lang"
	"github.com/isangeles/flame/core/module/item"
	"github.com/isangeles/flame/core/module/req"
	"github.com/isangeles/flame/core/module/serial"
)

// Struct for recipes.
type Recipe struct {
	id          string
	name, info  string
	catID       string
	res         []res.RecipeResultData
	reqs        []req.Requirement
	castTime    int64
	castTimeMax int64
	casting     bool
	casted      bool
}

// NewRecipe creates new crafting recipe.
func NewRecipe(data res.RecipeData) *Recipe {
	r := new(Recipe)
	r.id = data.ID
	r.catID = data.Category
	r.res = data.Results
	r.reqs = req.NewRequirements(data.Reqs...)
	r.castTimeMax = data.Cast
	nameInfo := lang.Texts(r.ID())
	r.name = nameInfo[0]
	if len(nameInfo) > 1 {
		r.info = nameInfo[1]
	}
	return r
}

// Update updates recipe.
func (r *Recipe) Update(delta int64) {
	if r.casting {
		r.castTime += delta
		if r.castTime >= r.castTimeMax {
			r.casting = false
			r.casted = true
		}
	}
}

// Make creates, assigns serials and
// returns items from recipe.
func (r *Recipe) Make() []item.Item {
	r.casted = false
	items := make([]item.Item, 0)
	for _, resData := range r.res {
		switch d := resData.Item.(type) {
		case res.WeaponData:
			it := item.NewWeapon(d)
			serial.AssignSerial(it)
			items = append(items, it)
		case res.MiscItemData:
			it := item.NewMisc(d)
			serial.AssignSerial(it)
			items = append(items, it)
		}
	}
	return items
}

// Cast starts casting make
// recipe action.
func (r *Recipe) Cast() {
	r.castTime = 0
	r.casting = true
}

// ID returns recipe ID.
func (r *Recipe) ID() string {
	return r.id
}

// Name returns recipe name.
func (r *Recipe) Name() string {
	return r.name
}

// Info returns recipe info.
func (r *Recipe) Info() string {
	return r.info
}

// CategoryID returns ID of recipe category.
func (r *Recipe) CategoryID() string {
	return r.catID
}

// Reqs returns recipe requirements
func (r *Recipe) Reqs() []req.Requirement {
	return r.reqs
}

// Result returns recipe result data.
func (r *Recipe) Result() []res.RecipeResultData {
	return r.res
}

// CastTime returns current casting
// time in millisecond.
func (r *Recipe) CastTime() int64 {
	return r.castTime
}

// CastTimeMax returns maximal casting
// time in milliseconds.
func (r *Recipe) CastTimeMax() int64 {
	return r.castTimeMax
}

// Casting checks if recipe make action
// is active.
func (r *Recipe) Casting() bool {
	return r.casting
}

// Casted checks if recipe make action
// was finished.
func (r *Recipe) Casted() bool {
	return r.casted
}
