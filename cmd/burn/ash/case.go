/*
 * case.go
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

package ash

import (
	"strconv"
	
	"github.com/isangeles/flame/cmd/burn"
)

// Struct for script case.
type ScriptCase struct {
	expr     burn.Expression
	expRes   string
	compType ComparisonType
}

// Type for script case types.
type ComparisonType int

const (
	Greater ComparisonType = iota
	Equal
	Less
	True
)

// NewCase creates new script case.
func NewCase(e burn.Expression, r string, t ComparisonType) *ScriptCase {
	c := new(ScriptCase)
	c.expr = e
	c.expRes = r
	c.compType = t
	return c
}

// Expression returns case expression.
func (c *ScriptCase) Expression() burn.Expression {
	return c.expr
}

// CorrectRes checks if specified result value is
// correct.
func (c *ScriptCase) CorrectRes(r string) bool {
	switch c.compType {
	case Greater:
		n, err := strconv.ParseFloat(r, 64)
		if err != nil {
			return false
		}
		exp, err := strconv.ParseFloat(c.expRes, 64)
		if err != nil {
			return false
		}
		return n > exp
	case Less:
		n, err := strconv.ParseFloat(r, 64)
		if err != nil {
			return false
		}
		exp, err := strconv.ParseFloat(c.expRes, 64)
		if err != nil {
			return false
		}
		return n < exp
	case True:
		return true
	default:
		return false
	}
}





