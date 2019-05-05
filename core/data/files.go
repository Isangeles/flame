/*
 * files.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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
	"fmt"
	"io/ioutil"
	"regexp"
)

// DirFilesNames returns names of all files matching specified
// file name pattern in directory with specified path.
func DirFilesNames(path, pattern string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	names := make([]string, 0)
	for _, info := range files {
		match, err := regexp.MatchString(pattern, info.Name())
		if err != nil {
			return nil, fmt.Errorf("fail_to_execute_pattern:%v", err)
		}
		if match {
			names = append(names, info.Name())
		}
	}
	return names, nil
}