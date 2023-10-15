/*
 * map_test.go
 *
 * Copyright 2023 Dariusz Sikora <ds@isangeles.dev>
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

package area

import (
	"fmt"
	"os"
	"testing"

	"github.com/isangeles/tmx"
)

// TestMapSize tests map size function.
func TestMapSize(t *testing.T) {
	// Create map
	data, err := testMap()
	if err != nil {
		t.Fatalf("Unable to get test map data: %v", err)
	}
	m := newMap(data)
	// Test
	width, height := m.Size()
	if width != 960 {
		t.Errorf("Invalid map width: %d != 100", width)
	}
	if height != 640 {
		t.Errorf("Invalid map height: %d != 100", height)
	}
}

// TestMapLayers tests map layers function.
func TestMapLayers(t *testing.T) {
	// Create map
	data, err := testMap()
	if err != nil {
		t.Fatalf("Unable to get test map data: %v", err)
	}
	m := newMap(data)
	// Test
	if len(m.Layers()) != 2 {
		t.Errorf("Invalid number of layers: %d != 2",
			len(m.Layers()))
	}
	var layer1, layer2 Layer
	for _, l := range m.Layers() {
		if l.Name() == "layer1" {
			layer1 = l
		}
		if l.Name() == "layer2" {
			layer2 = l
		}
	}
	if len(layer1.Name()) < 1 {
		t.Errorf("Layer 1 not found")
	}
	if len(layer2.Name()) < 1 {
		t.Errorf("Layer 2 not found")
	}
}

// testMap returns test map data.
func testMap() (*tmx.Map, error) {
	file, err := os.Open("testres/map.tmx")
	if err != nil {
		return nil, fmt.Errorf("Unable to open file: %v", err)
	}
	data, err := tmx.Read(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to read map file: %v", err)
	}
	return data, nil
}
