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

package parsexml

import (
	"flag"
	"testing"
	"os"

	"github.com/isangeles/flame/core/module/object/character"
)

var (
	path = flag.String("path", "", "System path of characters XML base to test parse")
)

// Test for unmarshaling of characters base.
// Use '-path' flag to point XML base to test.
func TestUnmarshalCharactersBase(t *testing.T) {
	doc, err := os.Open(*path)
	if err != nil {
		t.Errorf("base_file_not_found:%s\n", *path)
		return
	}
	chars, err := UnmarshalCharactersBase(doc)
	if err != nil {
		t.Errorf("unmarshal_fail:%v\n", err)
		return
	}
	t.Logf("unmarshal_success:chars_base_size:%d\n", len(chars)) 
}

// Test for marshaling of game character.
func TestMarshalCharacter(t *testing.T) {
	char := character.NewCharacter("test_01", "test", 1, character.Male,
		character.Human, character.Friendly, character.NewGuild(""),
		character.Attributes{1, 1, 1, 1, 1}, character.Lawful_good)
	xml, err := MarshalCharacter(char)
	if err != nil {
		t.Errorf("marshal_fail:%v\n", err)
		return
	}
	t.Logf("marshal_success:\n%s\n", xml)
}