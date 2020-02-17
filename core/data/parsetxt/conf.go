/*
 * conf.go
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
	"fmt"
	"io"
	"bufio"
	"strings"
)

const (
	commentPrefix = "#"
	keyValuesSep  = ":"
	valuesSep     = ";"
)

// UnmarshalConfig retrieves key-values from specified data.
func UnmarshalConfig(data io.Reader) map[string][]string {
	texts := make(map[string][]string)
	scann := bufio.NewScanner(data)
	for scann.Scan() {
		line := scann.Text()
		if strings.HasPrefix(line, commentPrefix) {
			continue
		}
		sepID := strings.Index(line, keyValuesSep)
		if sepID > -1 {
			lineID := line[:sepID]
			t := line[sepID+1:]
			texts[lineID] = append(texts[lineID], strings.Split(t, valuesSep)...)
		}
	}
	return texts
}

// MarshalConfig parses specified key-values map to
// config string.
func MarshalConfig(keyValues map[string][]string) (config string) {
	for key, values := range keyValues {
		config = fmt.Sprintf("%s%s%s", config, key, keyValuesSep)
		for i, v := range values {
			config = fmt.Sprintf("%s%s", config, v)
			if i < len(values)-1 {
				config += valuesSep
			}
		}
		config = fmt.Sprintf("%s\n", config)
	}
	return strings.TrimSpace(config)
}
