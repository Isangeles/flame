/*
 * craft.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/config"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/craft"
)

// Craft attemps to make specified recipe.
func (c *Character) Craft(r *craft.Recipe) {
	langPath := config.LangPath()
	if !c.MeetReqs(r.Reqs()...) {
		msg := fmt.Sprintf("%s:%s:%s", c.Name(), r.ID(),
			lang.TextDir(langPath, "reqs_not_meet"))
		c.SendPrivate(msg)
		return
	}
	r.Cast()
	return
}
