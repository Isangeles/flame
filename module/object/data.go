/*
 * data.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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

package object

import (
	"github.com/isangeles/flame/data/res"
)

// Data return data resource for object.
func (o *Object) Data() res.ObjectData {
	data := res.ObjectData{
		ID:        o.ID(),
		Serial:    o.Serial(),
		Name:      o.Name(),
		MaxHP:     o.MaxHealth(),
		HP:        o.Health(),
		Inventory: o.Inventory().Data(),
		UseAction: o.UseAction().Data(),
		Restore:   true,
	}
	data.PosX, data.PosY = o.Position()
	for _, e := range o.Effects() {
		effData := res.ObjectEffectData{
			ID:     e.ID(),
			Serial: e.Serial(),
			Time:   e.Time(),
		}
		effData.SourceID, effData.SourceSerial = e.Source()
		data.Effects = append(data.Effects, effData)
	}
	return data
}
