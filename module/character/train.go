/*
 * train.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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

package character

import (
	"fmt"
	
	"github.com/isangeles/flame/module/train"
)

// Train trains character with specified training.
func (char *Character) Train(t train.Training) error {
	if !char.MeetReqs(t.Reqs()...) {
		return fmt.Errorf("reqs-not-meet")
	}
	char.ChargeReqs(t.Reqs()...)
	switch t := t.(type) {
	case *train.AttrsTraining:
		char.attributes.Str += t.Strenght()
		char.attributes.Con += t.Constitution()
		char.attributes.Dex += t.Dexterity()
		char.attributes.Wis += t.Wisdom()
		char.attributes.Int += t.Intelligence()
	}
	return nil
}
