/*
 * newcharacter.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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

package main

import (
	"bufio"
	"fmt"
	"os"
	
	"github.com/isangeles/flame/core/data/text/lang"
)

// newModDialog starts new CLI dialog to
// create new module.
func newModDialog() error {
	scan := bufio.NewScanner(os.Stdin)
	mainAccept := false
	for !mainAccept {
		// ID.
		id := ""
		fmt.Printf("%s:", lang.Text("cui", "newmod_name"))
		for scan.Scan() {
			id = scan.Text()
			if !modIDValid(id) {
				fmt.Printf("%s\n", lang.Text("cui", "newmod_invalid_id_err"))
				fmt.Printf("%s:", lang.Text("cui", "newmod_name"))
				continue
			}
			break
		}
 		err := NewModule(id)
		if err != nil {
			return fmt.Errorf("fail_to_create_module_dir:%v", err)
		}
		mainAccept = true
	}
	return nil
}

// modIDValid check if specified ID
// is valid ID for module.
func modIDValid(id string) bool {
	return len(id) > 0
}
