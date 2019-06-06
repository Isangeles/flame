/*
 * script.go
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
	"fmt"
	"strings"
	
	"github.com/isangeles/flame/cmd/burn"
	"github.com/isangeles/flame/cmd/burn/syntax"
)

// Struct for Ash script.
type Script struct {
	name        string
	text        string
	mainCase    string
	expressions []burn.Expression
}

// NewScript creates new Ash script from specified
// text, returns error in case of syntax error.
func NewScript(text string) (*Script, error) {
	s := new(Script)
	s.text = text
	if !strings.Contains(s.text, "{") {
		return nil, fmt.Errorf("no_script_body")
	}
	startBrace := strings.Index(s.text, "{")
	endBrace := strings.Index(s.text, "}")
	mainCase := s.text[:startBrace]
	mainCase = strings.ReplaceAll(mainCase, "{", "")
	s.mainCase = mainCase
	body := s.text[startBrace:endBrace]
	body = strings.ReplaceAll(body, "{", "")
	body = strings.ReplaceAll(body, "}", "")
	for _, l := range strings.Split(body, "\n") {
		if len(l) < 1 {
			continue
		}
		expr, err := syntax.NewSTDExpression(l)
		if err != nil {
			return nil, fmt.Errorf("fail_to_parse_script_body:%v", err)
		}
		s.expressions = append(s.expressions, expr)
	}
	return s, nil
}

// Expressions returns all script expressions.
func (s *Script) Expressions() []burn.Expression {
	return s.expressions
}

// String returns script text body.
func (s *Script) String() string {
	return s.text
}
