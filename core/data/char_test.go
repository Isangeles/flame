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
)

var (
	file = flag.String("file", "", "System path to characters base file")
	dir  = flag.String("dir", "", "System path to directory with characters bases")
)

// Test for importing characters from characters base file.
// Use '-file' flag to point file with characters
// to test.
func TestImportCharacters(t *testing.T) {
	chars, err := ImportCharacters(*file)
	if err != nil {
		t.Errorf("fail_to_import_characters:%v\n", err)
		return
	}
	t.Logf("import_charasters_file_success_imported_characters:%d\n",
		len(chars))
}

// Test for importing characters from file in directory.
// Use '-dir' flag to point directory with characters
// bases to test.
func TestImportCharactersDir(t *testing.T) {
	chars, err := ImportCharactersDir(*dir)
	if err != nil {
		t.Errorf("fail_to_import_characters_dir:%v\n", err)
		return
	}
	t.Logf("import_characters_dir_success:imported_chars:%d\n",
		len(chars))
}
