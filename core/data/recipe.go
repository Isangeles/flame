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

package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/craft"
	"github.com/isangeles/flame/log"
)

const (
	RecipesFileExt = ".recipes"
)

// Recipe returns new recipe with specified ID or
// error if data was not found.
func Recipe(id string) (*craft.Recipe, error) {
	data := res.Recipe(id)
	if data == nil {
		return nil, fmt.Errorf("recipe data not found: %s", id)
	}
	r := craft.NewRecipe(*data)
	return r, nil
}

// ImportRecipes imports all recipes from base file with
// specified path.
func ImportRecipes(path string) ([]*res.RecipeData, error) {
	doc, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fail to open base file: %v", err)
	}
	defer doc.Close()
	recipes, err := parsexml.UnmarshalRecipes(doc)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal recipes base: %v", err)
	}
	return recipes, nil
}

// ImportRecipesDir imports all recipes from base files in
// directory with specified path.
func ImportRecipesDir(path string) ([]*res.RecipeData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("fail to read dir: %v", err)
	}
	recipes := make([]*res.RecipeData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), RecipesFileExt) {
			continue
		}
		basePath := filepath.FromSlash(path + "/" + finfo.Name())
		rd, err := ImportRecipes(basePath)
		if err != nil {
			log.Err.Printf("data recipes import: %s: fail to import base: %v",
				basePath, err)
		}
		for _, r := range rd {
			recipes = append(recipes, r)
		}
	}
	return recipes, nil
}
