/*
 * dialog.go
 *
 * Copyright 2019-2024 Dariusz Sikora <ds@isangeles.dev>
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
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

// ImportDialogs imports all dialogs from data file with
// specified path.
func ImportDialogs(path string) ([]res.DialogData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("unable to open data file: %v", err))
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.DialogsData)
	err = unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON data: %v", err)
	}
	return data.Dialogs, nil
}

// ImportDialogsDir imports all dialogs from data files in
// directory with specified path.
func ImportDialogsDir(path string) ([]res.DialogData, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("unable to read dir: %v", err))
	}
	dialogs := make([]res.DialogData, 0)
	for _, file := range files {
		filePath := filepath.FromSlash(path + "/" + file.Name())
		dd, err := ImportDialogs(filePath)
		if err != nil {
			log.Err.Printf("data dialogs import: %s: unable to import base: %v",
				filePath, err)
		}
		for _, d := range dd {
			dialogs = append(dialogs, d)
		}
	}
	return dialogs, nil
}

// ExportDialogs exports dialogs to the data file under specified path.
func ExportDialogs(path string, dialogs ...res.DialogData) error {
	data := new(res.DialogsData)
	for _, d := range dialogs {
		data.Dialogs = append(data.Dialogs, d)
	}
	// Marshal dialogs data.
	json, err := marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal dialogs: %v", err)
	}
	// Create dialogs data file.
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create dialogs file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create dialogs file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	writer := bufio.NewWriter(file)
	writer.Write(json)
	writer.Flush()
	return nil
}
