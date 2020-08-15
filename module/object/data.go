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
	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/useaction"
	"github.com/isangeles/flame/log"
)

// Apply applies specifed data on the object.
func (o *Object) Apply(data res.ObjectData) {
	o.id = data.ID
	o.SetSerial(data.Serial)
	o.SetName(data.Name)
	o.SetMaxHealth(data.MaxHP)
	o.SetHealth(o.MaxHealth())
	o.SetPosition(data.PosX, data.PosY)
	o.Inventory().Apply(data.Inventory)
	o.action = useaction.New(data.UseAction)
	if len(o.Name()) < 1 {
		// Translate name.
		o.SetName(lang.Text(o.ID()))
	}
	// Restore.
	if data.Restore {
		o.SetHealth(data.HP)
	}
	// Add effects.
	for _, data := range data.Effects {
		effData := res.Effect(data.ID)
		if effData == nil {
			log.Err.Printf("Object: %s: Apply: effect data not found: %s",
				o.ID(), data.ID)
			continue
		}
		eff := effect.New(*effData)
		eff.SetSerial(data.Serial)
		eff.SetTime(data.Time)
		eff.SetSource(data.SourceID, data.SourceSerial)
		o.AddEffect(eff)
	}
}

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
