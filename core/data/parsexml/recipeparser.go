/*
 * recipeparser.go
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

package parsexml

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/log"
)

// Struct for recipes XML node.
type RecipesBaseXML struct {
	XMLName xml.Name    `xml:"base"`
	Recipes []RecipeXML `xml:"recipe"`
}

// Struct for recipe XML node.
type RecipeXML struct {
	XMLName  xml.Name          `xml:"recipe"`
	ID       string            `xml:"id,attr"`
	Category string            `xml:"category,attr"`
	Results  []RecipeResultXML `xml:"results>result"`
	Reqs     ReqsXML           `xml:"reqs"`
}

// Struct for recipe result XML node.
type RecipeResultXML struct {
	XMLName xml.Name `xml:"result"`
	ID      string   `xml:"id,attr"`
	Amount  int      `xml:"amount,attr"`
}

// UnmarshalRecipesBase retieves all recipes data from specified
// XML data.
func UnmarshalRecipesBase(data io.Reader) ([]*res.RecipeData, error) {
	doc, _ := ioutil.ReadAll(data)
	xmlBase := new(RecipesBaseXML)
	err := xml.Unmarshal(doc, xmlBase)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v", err)
	}
	recipes := make([]*res.RecipeData, 0)
	for _, xmlRecipe := range xmlBase.Recipes {
		recipe, err := buildRecipeData(xmlRecipe)
		if err != nil {
			log.Err.Printf("xml:unmarshal_recipe:build_data_fail:%v", err)
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

// buildRecipeData creates new recipe data from specified XML data.
func buildRecipeData(xmlRecipe RecipeXML) (*res.RecipeData, error) {
	rd := new(res.RecipeData)
	rd.ID = xmlRecipe.ID
	rd.Category = xmlRecipe.Category
	for _, r := range xmlRecipe.Results {
		itd := itemRes(r.ID)
		if itd == nil {
			return nil, fmt.Errorf("result item data not found")
		}
		rrd := res.RecipeResultData{
			ID:     r.ID,
			Amount: r.Amount,
			Item:   itd,
		}
		rd.Results = append(rd.Results, rrd)
	}
	rd.Reqs = buildReqs(&xmlRecipe.Reqs)
	return rd, nil
}

// itemRes returns resources for item with
// specified ID, or nil if no data found.
func itemRes(id string) res.ItemData {
	m := res.MiscItem(id)
	if m != nil {
		return *m
	}
	w := res.Weapon(id)
	if w != nil {
		return *w
	}
	return nil
}
