/*
 * data_test.go
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

package data

import (
	"flag"
	"testing"

	//"github.com/isangeles/flame/core/module/object/character"
)

var (
	path = flag.String("path", "", "System path to directory with chracters files")
)

// Test for importing characters from files in directory.
// Use '-path' flag to point directory with characters
// files to test.
func TestImportCharacters(t *testing.T) {
	chars, err := ImportCharacters(*path)
	if err != nil {
		t.Errorf("fail_to_import_characters:%v\n", err)
		return
	}
	t.Logf("imported_characters:%d\n", len(chars))
}
