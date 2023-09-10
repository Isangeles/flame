/*
 * map.go
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
	"github.com/salviati/go-tmx/tmx"
)

// Struct for area map.
type Map struct {
	width, height int
	layers        []Layer
	data          *tmx.Map
}

// Struct for area map layer.
type Layer struct {
	name  string
	tiles []Tile
}

// Struct for map tile.
type Tile struct {
	x, y, endX, endY float64
}

// newMap creates new area map.
func newMap(data *tmx.Map) Map {
	m := Map{data: data}
	m.width = data.TileWidth * data.Width
	m.height = data.TileHeight * data.Height
	for _, tmxLayer := range data.Layers {
		layer := Layer{name: tmxLayer.Name}
		var tileX, tileY int
		for _, dt := range tmxLayer.DecodedTiles {
			if dt.Tileset != nil {
				tilePosX := float64(int(data.TileWidth) * tileX)
				tilePosY := float64(int(data.TileHeight) * tileY)
				tilePosY = float64(m.height) - tilePosY
				tile := Tile{tilePosX, tilePosY, tilePosX + float64(data.TileWidth),
					tilePosY + float64(data.TileHeight)}
				layer.tiles = append(layer.tiles, tile)
			}
			tileX++
			if tileX > int(data.Width)-1 {
				tileX = 0
				tileY++
			}
		}
		m.layers = append(m.layers, layer)
	}
	return m
}

// Size returns map size.
func (m Map) Size() (int, int) {
	return m.width, m.height
}

// Layers returns all map layers.
func (m Map) Layers() []Layer {
	return m.layers
}

// PositionLayer returns visible layer on specified
// XY position on the map.
func (m Map) PositionLayer(x, y float64) (layer Layer) {
	for _, l := range m.layers {
		for _, t := range l.tiles {
			if t.Contains(x, y) {
				layer = l
			}
		}
	}
	return
}

// Data returns map data resource.
func (m Map) Data() *tmx.Map {
	return m.data
}

// Name returns layer name.
func (l Layer) Name() string {
	return l.name
}

// Tiles returns all layer tiles.
func (l Layer) Tiles() []Tile {
	return l.tiles
}

// Constains checks if specified XY posistion is contained
// inside the map tile.
func (t Tile) Contains(x, y float64) bool {
	return x >= t.x && y >= t.y && x <= t.endX && y <= t.endY
}
