/*
 * lang_test.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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
	"testing"

	"github.com/isangeles/flame/data/res"
)

// Test for exporting translation data for lang file.
func TestExportLang(t *testing.T) {
	data1 := res.TranslationData{
		ID:    "item1",
		Texts: []string{"Item 1", "Item description 1"},
	}
	data2 := res.TranslationData{
		ID:    "item2",
		Texts: []string{"Item 2", "Item description 2"},
	}
	err := ExportLang("testlang", data1, data2)
	if err != nil {
		t.Errorf("Unable to export translation data: %v", err)
	}
}
