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
	"github.com/isangeles/flame/cmd/burn"
)

// Struct for script expressions.
type ScriptExpression struct {
	burnExpr burn.Expression
	exprType ExpressionType
	echoText string
	waitTime int64
}

type ExpressionType int

const (
	Expr ExpressionType = iota
	Echo_macro
	Wait_macro
)

// NewExpression returns new script expression for specified
// burn expression.
func NewExpression(expr burn.Expression) *ScriptExpression {
	se := new(ScriptExpression)
	se.exprType = Expr
	se.burnExpr = expr
	return se
}

// NewEchoMacro returns new script expression for
// echo macro.
func NewEchoMacro(text string, expr burn.Expression) *ScriptExpression{
	se := new(ScriptExpression)
	se.exprType = Echo_macro
	se.burnExpr = expr
	se.echoText = text
	return se
}

// NewWaitMacro returns new script expression for wait
// macro.
func NewWaitMacro(millis int64) *ScriptExpression {
	se := new(ScriptExpression)
	se.exprType = Wait_macro
	se.waitTime = millis
	return se
}

// Type returns expression type.
func (se *ScriptExpression) Type() ExpressionType {
	return se.exprType
}

// WaitMillis returns number of milliseconds to wait
// before expression execution.
func (se *ScriptExpression) WaitTime() int64 {
	return se.waitTime
}

// BurnExpr returns burn expression or empty struct
// if there is no burn expression in this script
// expression.
func (se *ScriptExpression) BurnExpr() burn.Expression {
	return se.burnExpr
}
