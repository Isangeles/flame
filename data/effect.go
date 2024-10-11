/*
 * effect.go
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
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/log"
)

// ImportEffects imports all JSON effects data from effects base
// with specified path.
func ImportEffects(path string) ([]res.EffectData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open effects data file: %v", err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file: %v", err)
	}
	data := new(res.EffectsData)
	err = unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal effects base: %v", err)
	}
	return data.Effects, nil
}

// ImportEffectsDir imports all effects from files in
// specified directory.
func ImportEffectsDir(dirPath string) ([]res.EffectData, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir: %v", err)
	}
	effects := make([]res.EffectData, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.FromSlash(dirPath + "/" + file.Name())
		effs, err := ImportEffects(filePath)
		if err != nil {
			log.Err.Printf("data: effects import: %s: unable to import base: %v",
				filePath, err)
			continue
		}
		for _, e := range effs {
			effects = append(effects, e)
		}
	}
	return effects, nil
}

// ExportEffects exports effects to the data file under specified path.
func ExportEffects(path string, effects ...res.EffectData) error {
	data := new(res.EffectsData)
	for _, e := range effects {
		data.Effects = append(data.Effects, e)
	}
	// Marshal effect data.
	json, err := marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal effects: %v", err)
	}
	// Create effect file.
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("unable to create effects file directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create effects file: %v", err)
	}
	defer file.Close()
	// Write data to file.
	writer := bufio.NewWriter(file)
	writer.Write(json)
	writer.Flush()
	return nil
}
