/*
 * text.go
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

// text package provides functions to retrive parameters from engine 
// config files and text data from lang files.
// Config/lang file data format:
// [text ID]:[text];[new-line]
// ...
package text

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const (
	ID_TEXT_SEP = ":"
	LINE_TERMINATOR = ";"
	COMMENT_PREFIX = "#"
)

// ReadDisplayText retrives text(one or more) with specified IDs from file 
// from sepcified path.
// In case of error(file/ID not found) returns string with error message 
// instead of text. 
func ReadDisplayText(filePath string, textIDs ...string) (texts []string) {
	file, err := os.Open(filePath)
	if err != nil {
		texts = append(texts, fmt.Sprintf("CANT_OPEN:%s", filePath))
		return
	}
	defer file.Close()

	scann := bufio.NewScanner(file)
	for _, id := range textIDs {
		var found = false
		file.Seek(0, 0) // reset file pointer
		for scann.Scan() {
			line := scann.Text()
			if strings.HasPrefix(line, COMMENT_PREFIX) {
				continue
			}
			
			lineParts := strings.Split(line, ID_TEXT_SEP)
			if lineParts[0] == id {
				t := lineParts[1]
				texts = append(texts,
					strings.TrimSuffix(t, LINE_TERMINATOR))
				found = true
				break
			}
		}
		if !found {
			texts = append(texts, fmt.Sprintf("TEXT_NOT_FONUD:%v", id))
		}
	}
	
	return
}

// ReadConfigValue retrives text(one or more) with specified IDs from file 
// from sepcified path.
// Returns error if file or at least one speicfied ID was not found.
func ReadConfigValue(filePath string, textIDs ...string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("CANT_OPEN:%s", filePath)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	var texts []string
	for _, id := range textIDs {
		found := false
		file.Seek(0, 0) // reset file pointer
		for scan.Scan() {
			line := scan.Text()
			if strings.HasPrefix(line, COMMENT_PREFIX) {
				continue
			}

			lineParts := strings.Split(line, ID_TEXT_SEP)
			if lineParts[0] == id {
				t := lineParts[1]
				texts = append(texts,
					strings.TrimSuffix(t, LINE_TERMINATOR))
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("TEXT_NOT_FONUD:%v", id)
		}
	}
	return texts, nil
}

// ReadConfigInt retrives integer(one or more) with specified IDs from file 
// from sepcified path.
// Returns error if file or at least one speicfied ID was not found, or
// value for at least one specified ID was not parseable to integer.
func ReadConfigInt(filePath string, ids ...string) ([]int, error) {
	vals, err := ReadConfigValue(filePath, ids...)
	if err != nil {
		return nil, err
	}
	var ints []int
	for _, v := range vals {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
