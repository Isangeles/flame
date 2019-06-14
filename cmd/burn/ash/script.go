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
	"strconv"

	"github.com/isangeles/flame/cmd/burn"
	"github.com/isangeles/flame/cmd/burn/syntax"
)

// Struct for Ash script.
type Script struct {
	name      string
	args      []string
	text      string
	mainCase  *ScriptCase
	exprs     []*ScriptExpression
	pos       int
}

const (
	Comment_prefix = "#"
	Body_expr_sep  = ";"
	// Keywords.
	True_keyword   = "true"
	Echo_keyword   = "echo"
	Wait_keyword   = "wait"
	Rawdis_keyword = "rawdis"
)

// NewScript creates new Ash script from specified
// text, returns error in case of syntax error.
func NewScript(text string, args ...string) (*Script, error) {
	s := new(Script)
	s.args = args
	// Remove comment lines.
	for _, l := range strings.Split(text, "\n") {
		if strings.HasPrefix(l, Comment_prefix) {
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
	mainCaseText := s.text[:startBrace]
	mainCaseText = strings.ReplaceAll(mainCaseText, "{", "")
	mainCase, err := parseCase(strings.TrimSpace(mainCaseText))
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_main_case:%v", err)
	}
	s.mainCase = mainCase
	// Body.
	body := textBetween(s.text, "{", "}")
	exprs, err := parseBody(body)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_body:%v", err)
	}
	s.exprs = exprs
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

// Position returns position of currently
// executed script expression.
func (s *Script) Position() int {
	return s.pos
}

// SetPosition sets position of currently
// executed script expression.
func (s *Script) SetPosition(p int) {
	s.pos = p
}

// Finished checks if script is finished.
func (s *Script) Finished() bool {
	return s.Position() >= len(s.Expressions())
}

// MainCase returns script main case.
func (s *Script) MainCase() *ScriptCase {
	return s.mainCase
}

// parseBody creates script expression from script body text.
func parseBody(text string) ([]*ScriptExpression, error) {
	exprs := make([]*ScriptExpression, 0)
	for _, l := range strings.Split(text, Body_expr_sep) {
		l = strings.TrimSpace(l)
		if len(l) < 1 {
			continue
		}
		switch {
		case strings.HasPrefix(l, Echo_keyword):
			l = textBetween(l, "(", ")")
			expr, err := syntax.NewSTDExpression(l)
			if err != nil {
				return nil, fmt.Errorf("fail_to_parse_echo_function:%v", err)
			}
			exprs = append(exprs, NewEchoMacro("", expr))
		case strings.HasPrefix(l, Wait_keyword):
			secText := textBetween(l, "(", ")")
			sec, err := strconv.ParseInt(secText, 32, 64)
			if err != nil {
				return nil, fmt.Errorf("fail_to_parse_wait_function:%v", err)
			}
			exprs = append(exprs, NewWaitMacro(sec * 1000))
		default:
			expr, err := syntax.NewSTDExpression(l)
			if err != nil {
				return nil, fmt.Errorf("fail_to_parse_expr:%v", err)
			}
			sExpr := NewExpression(expr)
			exprs = append(exprs, sExpr)
		}
	}
	return exprs, nil
}

// parseCase creates script case from specified text.
func parseCase(text string) (*ScriptCase, error) {
	if len(text) < 1 || strings.HasPrefix(text, True_keyword) {
		c := NewCase(new(syntax.STDExpression), "", True)
		return c, nil
	}
	var compType ComparisonType
	switch {
	case strings.Contains(text, "<"):
		compType = Less
		exprs := strings.Split(text, "<")
		expr, err := parseCaseExpr(exprs[0])
		if err != nil {
			return nil, fmt.Errorf("fail_to_parse_case_expression:%v", err)
		}
		res := strings.TrimSpace(exprs[1])
		c := NewCase(expr, res, compType)
		return c, nil
	default:
		return nil, fmt.Errorf("unknown case expression:%s", text)
	}
}

// parseCaseExpr creates case expression from specified text.
func parseCaseExpr(text string) (burn.Expression, error) {
	switch {
	case strings.HasPrefix(text, Rawdis_keyword):
		args := strings.Split(textBetween(text, "(", ")"), ",")
		if len(args) < 2 {
			return nil, fmt.Errorf("not enaught args for rawdis")
		}
		exprText := fmt.Sprintf("charman -o show -t %s -a range %s",
			strings.TrimSpace(args[0]), strings.TrimSpace(args[1]))
		expr, err := syntax.NewSTDExpression(exprText)
		if err != nil {
			return nil, fmt.Errorf("fail_to_create_rawdis_exression:%v", err)
		}
		return expr, nil
	default:
		expr, err := syntax.NewSTDExpression(text)
		if err != nil {
			return nil, fmt.Errorf("fail_to_create_std_expression:%v", err)
		}
		return expr, nil
	}
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
