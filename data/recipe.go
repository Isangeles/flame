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

package data

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

const (
	RecipesFileExt = ".recipes"
)

// ImportRecipes imports all recipes from base file with
// specified path.
func ImportRecipes(path string) ([]res.RecipeData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open data file: %v", err)
	}
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	defer file.Close()
	data := new(res.RecipesData)
	err = xml.Unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal XML data: %v", err)
	}
	return data.Recipes, nil
}

// ImportRecipesDir imports all recipes from base files in
// directory with specified path.
func ImportRecipesDir(path string) ([]res.RecipeData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	recipes := make([]res.RecipeData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), RecipesFileExt) {
			continue
		}
		basePath := filepath.FromSlash(path + "/" + finfo.Name())
		rd, err := ImportRecipes(basePath)
		if err != nil {
			log.Err.Printf("data recipes import: %s: unable to import base: %v",
				basePath, err)
		}
		for _, r := range rd {
			recipes = append(recipes, r)
		}
	}
	return recipes, nil
}
