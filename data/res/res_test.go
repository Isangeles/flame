/*
 * res_test.go
 *
 * Copyright 2022 Dariusz Sikora <dev@isangeles.pl>
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

package res

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Name of the directory with test resources.
const testDataDir = "testres"

// testData returns data buffer from file with specified name
// inside test resources directory.
func testData(name string) ([]byte, error) {
	path := filepath.Join(testDataDir, name)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to open data file: %v", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to read data file: %v", err)
	}
	return data, nil
}
