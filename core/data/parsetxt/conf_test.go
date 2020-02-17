/*
 * conf_test.go
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
	"testing"
	"strings"
)

// Test for unmarshaling of config data.
func TestUnmarshalConfig(t *testing.T) {
	config := `#comment1
key1:value1
key2:value2;value3
key3:value4;value5;value6
#comment2`
	keyValues := UnmarshalConfig(strings.NewReader(config))
	for key, values := range keyValues {
		t.Logf("%s: ", key)
		for _, v := range values {
			t.Logf("%s ", v)
		}
	}
	if len(keyValues) != 3 {
		t.Errorf("Invalid number of keys from unmarshaled config string: %d != 3\n",
			len(keyValues))
	}
}

// Test for marshaling config data.
func TestMarshalConfig(t *testing.T) {
	keyValues := make(map[string][]string)
	keyValues["key1"] = []string{"value1"}
	keyValues["key2"] = []string{"value2", "value3"}
	config := MarshalConfig(keyValues)
	t.Logf("config:\n%s", config)
	if !strings.Contains(config, "key1:value1") || !strings.Contains(config, "key2:value2;value3") {
		t.Errorf("Invalid output string: %s", config)
	}
}
