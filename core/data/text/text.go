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

// ReadDisplayText retrives text(one or more) with specified IDs from file 
// from sepcified path.
// Returns error if file/ID was not found.
func ReadConfigValue(filePath string, textIDs ...string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("CANT_OPEN:%s", filePath)
	}
	defer file.Close()
	
	var texts []string
	scann := bufio.NewScanner(file)
	for _, id := range textIDs {
		var found = false
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
			return nil, fmt.Errorf("TEXT_NOT_FONUD:%v", id)
		}
	}
	
	return texts, nil
}
