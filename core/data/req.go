/*
 * req.go
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

package data

import (
	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/module/req"
)

// buildXMLReqs creates requirements from specified
// XML data.
func buildXMLReqs(xmlReqs *parsexml.ReqsNodeXML) []req.Requirement { 
	reqs := make([]req.Requirement, 0)
	// Level reqs.
	for _, xmlReq := range xmlReqs.LevelReqs {
		req := req.NewLevelReq(xmlReq.MinLevel)
		reqs = append(reqs, req)
	}
	return reqs
}
