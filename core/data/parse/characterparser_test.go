/*
 * characterparser_test.go
 * 
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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

package parse

import (
	"flag"
	"testing"
)

var (
	path = flag.String("path", "", "System path of characters XML base to test parse")
)

// Test for characters XML base parsing.
// Use '-path' flag to point XML base to test.
func TestParseCharactersBase(t *testing.T) {
	chars, err := ParseCharactersBaseXML(*path)
	if err != nil {
		t.Errorf("parse_fail:%v\n", err)
		return
	}
	t.Logf("parse_success:chars_base_size:%d\n", len(*chars)) 
}
