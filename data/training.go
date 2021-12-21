/*
 * training.go
 *
 * Copyright 2020-2021 Dariusz Sikora <dev@isangeles.pl>
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
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

const (
	TrainingsFileExt = ".trainings"
)

// ImportTrainings imports all trainings from data file with
// specified path.
func ImportTrainings(path string) ([]res.TrainingData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open data file: %v", err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.TrainingsData)
	err = xml.Unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal XML data: %v", err)
	}
	return data.Trainings, nil
}

// ImportTrainingsDir imports all trainings from data files in
// directory with specified path.
func ImportTrainingsDir(path string) ([]res.TrainingData, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	trainings := make([]res.TrainingData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), TrainingsFileExt) {
			continue
		}
		basePath := filepath.FromSlash(path + "/" + finfo.Name())
		dd, err := ImportTrainings(basePath)
		if err != nil {
			log.Err.Printf("data trainings import: %s: unable to import base: %v",
				basePath, err)
		}
		for _, d := range dd {
			trainings = append(trainings, d)
		}
	}
	return trainings, nil
}

// ExportTrainings exports trainings to the data file under specified path.
func ExportTrainings(path string, trainings ...res.TrainingData) error {
	data := new(res.TrainingsData)
	for _, t := range trainings {
		data.Trainings = append(data.Trainings, t)
	}
	// Marshal races data.
	xml, err := xml.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal trainings: %v", err)
	}
	// Create races file.
	if !strings.HasSuffix(path, TrainingsFileExt) {
		path += TrainingsFileExt
	}
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create trainings file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create trainings file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	writer := bufio.NewWriter(file)
	writer.Write(xml)
	writer.Flush()
	return nil
}
