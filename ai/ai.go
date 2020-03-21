/*
 * ai.go
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

package ai

import (
	"fmt"

	"github.com/isangeles/flame/data/res/lang"
	"github.com/isangeles/flame/module"
	"github.com/isangeles/flame/module/character"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/skill"
	"github.com/isangeles/flame/rng"
)

var (
	// Random actions frequences(in millis).
	moveFreq int64 = 3000
	chatFreq int64 = 5000
)

// Struct for controlling non-player characters.
type AI struct {
	mod       *module.Module
	npcs      map[string]*character.Character
	moveTimer int64
	chatTimer int64
}

// New creates new AI.
func New(mod *module.Module) *AI {
	ai := new(AI)
	ai.mod = mod
	ai.npcs = make(map[string]*character.Character)
	return ai
}

// Update updates AI.
func (ai *AI) Update(delta int64) {
	ai.moveTimer += delta
	ai.chatTimer += delta
	// NPCs.
	for _, npc := range ai.npcs {
		// Move around.
		if ai.moveTimer >= moveFreq {
			if npc.Casting() || npc.Moving() || npc.Fighting() || npc.Agony() {
				continue
			}
			posX, posY := npc.Position()
			defX, defY := npc.DefaultPosition()
			if posX != defX || posY != defY {
				npc.SetDestPoint(defX, defY)
				continue
			}
			ai.moveAround(npc)
		}
		// Random chat.
		if ai.chatTimer >= chatFreq {
			if npc.Casting() || npc.Moving() || npc.Fighting() || npc.Agony() {
				continue
			}
			ai.saySomething(npc)
		}
		// Combat.
		tar := npc.Targets()[0]
		if tar == nil || npc.AttitudeFor(tar) != character.Hostile {
			// Look for hostile target.
			area := ai.mod.Chapter().CharacterArea(npc)
			if area == nil {
				continue
			}
			for _, t := range area.NearTargets(npc, npc.SightRange()) {
				if t == npc {
					continue
				}
				if npc.AttitudeFor(t) == character.Hostile {
					tar = t
					break
				}
			}
			if tar == nil {
				continue
			}
			npc.SetTarget(tar)
		}
		if npc.Fighting() {
			skill := ai.combatSkill(npc, npc.Targets()[0])
			if skill == nil {
				continue
			}
			npc.UseSkill(skill)
		}
		break
	}
	// Reset timers.
	if ai.moveTimer >= moveFreq {
		ai.moveTimer = 0
	}
	if ai.chatTimer >= chatFreq {
		ai.chatTimer = 0
	}
}

// Characters returns all characters under AI control.
func (ai *AI) Characters() (chars []*character.Character) {
	for _, c := range ai.npcs {
		chars = append(chars, c)
	}
	return
}

// AddCharacters adds specified characters to control by AI.
func (ai *AI) AddCharacters(chars ...*character.Character) {
	for _, c := range chars {
		ai.npcs[c.ID()+c.Serial()] = c
	}
}

// RemoveCharacters removes specified character from AI control.
func (ai *AI) RemoveCharacters(chars ...*character.Character) {
	for _, c := range chars {
		delete(ai.npcs, c.ID()+c.Serial())
	}
}

// moveAround moves specified character in random direction.
func (ai *AI) moveAround(npc *character.Character) {
	dir := rng.RollInt(1, 4)
	posX, posY := npc.Position()
	switch dir {
	case 1:
		posY += 1
	case 2:
		posX += 1
	case 3:
		posY -= 1
	case 4:
		posX -= 1
	}
	npc.SetDestPoint(posX, posY)
}

// saySomething sends random text on NPC chat channel.
func (ai *AI) saySomething(npc *character.Character) {
	if npc.Race() == nil {
		return
	}
	textID := fmt.Sprintf("random_chat_%s", npc.Race().ID())
	texts := lang.Texts(textID)
	if len(texts) < 1 {
		return
	}
	id := rng.RollInt(0, len(texts)-1)
	npc.SendChat(texts[id])
}

// combatSkill selects NPC skill to use in combat or nil if specified
// NPC has no skills to use.
func (ai *AI) combatSkill(npc *character.Character, tar effect.Target) *skill.Skill {
	if len(npc.Skills()) < 1 {
		return nil
	}
	return npc.Skills()[0]
}
