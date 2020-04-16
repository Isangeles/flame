/*
 * lang.go
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

package parsetxt

import (
	"io"

	"github.com/isangeles/flame/data/res"
)

// UnmarshalLangData retrives all translation data from
// specified source.
func UnmarshalLangData(data io.Reader) []res.TranslationData {
	values := UnmarshalConfig(data)
	return buildTranslationData(values)
}

// buildTranslationData creates translation data from specified
// key-texts map.
func buildTranslationData(values map[string][]string) (data []res.TranslationData) {
	for k, v := range values {
		td := res.TranslationData{k, v}
		data = append(data, td)
	}
	return
}
