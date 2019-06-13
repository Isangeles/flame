/*
 * ash.go
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
	"time"

	"github.com/isangeles/flame/cmd/burn"
)

const (
	SCRIPT_FILE_EXT = ".ash"
)

// Run runs specified script.
func Run(scr *Script) error {
	if !executable(scr) {
		return nil
	}
	for i := scr.Position(); i < len(scr.Expressions()); i ++ {
		defer scr.SetPosition(scr.Position()+1)
		e := scr.Expressions()[i]
		if e.Type() == Wait_macro {
			time.Sleep(time.Duration(e.WaitTime()) * time.Millisecond)
			continue
		}
		r, o := burn.HandleExpression(e.BurnExpr())
		if r != 0 {
			return fmt.Errorf("fail_to_run_expr:'%s':[%d]%s",
				e.BurnExpr().String(), r, o)
		}
		if e.Type() == Echo_macro {
			fmt.Printf("%s\n", o)
		}
	}
	return nil
}

// executable checks if specified should be executed
// for specified game.
func executable(scr *Script) bool {
	if len(scr.mainCase) < 1 {
		return true
	}
	return checkCase(scr.mainCase)
}

// checkCase checks specified case for specified
// game.
func checkCase(c string) bool {
	switch {
	case strings.Contains(c, NEAR_KEYWORD):
		ids := strings.Split(c, NEAR_KEYWORD)
		if len(ids) < 2 {
			return false
		}
		// TODO: check range.
		return true
	case c == TRUE_KEYWORD:
		return true
	default:
		return false
	}
}
