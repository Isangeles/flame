/*
 * unmarshal.go
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
	"bufio"
	"strings"
)

const (
	commentPrefix = "#"
	idTextSep     = ":"
)

// UnmarshalTextValue retrieves key-values from specified data.
func UnmarshalTextValue(data io.Reader) map[string][]string {
	texts := make(map[string][]string)
	scann := bufio.NewScanner(data)
	for scann.Scan() {
		line := scann.Text()
		if strings.HasPrefix(line, commentPrefix) {
			continue
		}
		sepID := strings.Index(line, idTextSep)
		if sepID > -1 {
			lineID := line[:sepID]
			t := line[sepID+1:]
			texts[lineID] = append(texts[lineID], strings.Split(t, ";")...)
		}
	}
	return texts
}
