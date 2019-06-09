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

	"github.com/isangeles/flame/cmd/burn/syntax"
)

// Struct for Ash script.
type Script struct {
	name      string
	args      []string
	text      string
	mainCase  string
	exprs     []*ScriptExpression
}

const (
	COMMENT_PREFIX = "#"
	BODY_EXPR_SEP  = ";"
	// Keywords.
	NEAR_KEYWORD = "near"
	TRUE_KEYWORD = "true"
	ECHO_KEYWORD = "echo"
)

// NewScript creates new Ash script from specified
// text, returns error in case of syntax error.
func NewScript(text string, args ...string) (*Script, error) {
	s := new(Script)
	s.args = args
	// Remove comment lines.
	for _, l := range strings.Split(text, "\n") {
		if strings.HasPrefix(l, COMMENT_PREFIX) {
			continue
		}
		s.text += l
	}
	// Insert args.
	for i := 1; i < len(s.args); i ++ {
		macro := fmt.Sprintf("@%d", i)
		s.text = strings.ReplaceAll(s.text, macro, s.args[i])
	}
	if !strings.Contains(s.text, "{") {
		return nil, fmt.Errorf("no_script_body")
	}
	// Main case.
	startBrace := strings.Index(s.text, "{")
	mainCase := s.text[:startBrace]
	mainCase = strings.ReplaceAll(mainCase, "{", "")
	s.mainCase = strings.TrimSpace(mainCase)
	// Body.
	body := textBetween(s.text, "{", "}")
	for _, l := range strings.Split(body, BODY_EXPR_SEP) {
		l = strings.TrimSpace(l)
		if len(l) < 1 {
			continue
		}
		echo := false
		if strings.HasPrefix(l, ECHO_KEYWORD) {
			l = textBetween(l, "(", ")")
			echo = true
		}
		expr, err := syntax.NewSTDExpression(l)
		if err != nil {
			return nil, fmt.Errorf("fail_to_parse_script_body:%v", err)
		}
		sExpr := NewExpression(expr, echo)
		s.exprs = append(s.exprs, sExpr)
	}
	return s, nil
}

// Expressions returns all script expressions.
func (s *Script) Expressions() []*ScriptExpression {
	return s.exprs
}

// String returns script text body.
func (s *Script) String() string {
	return s.text
}

// textBetween returns slice from specified text
// between specified start and end sequence or
// the same specified text if start or end sequence
// was not found.
func textBetween(text, start, end string) string {
	startID := strings.Index(text, start)
	if startID < 0 {
		return text
	}
	endID := strings.Index(text, end)
	if endID < 0 {
		return text
	}
	return text[startID+1:endID]
}
