/*
 * ai.go
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

package core

import (
	"fmt"
	
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/object/skill"
	"github.com/isangeles/flame/core/rng"
)

var (
	// Random actions frequences(in millis).
	move_freq int64 = 3000
	chat_freq int64 = 5000
)

// Struct for controlling non-player characters.
type AI struct {
	game      *Game
	npcs      map[string]*character.Character
	moveTimer int64
	chatTimer int64
}

// NewAI creates new AI.
func NewAI(g *Game) *AI {
	ai := new(AI)
	ai.game = g
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
		if ai.moveTimer >= move_freq {
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
		if ai.chatTimer >= chat_freq {
			if npc.Casting() || npc.Moving() || npc.Fighting() || npc.Agony() {
				continue
			}
			ai.saySomething(npc)
		}
		// Combat.
		tar := npc.Targets()[0]
		if tar == nil || npc.AttitudeFor(tar) != character.Hostile {
			area := ai.game.Module().Chapter().CharacterArea(npc)
			if area == nil {
				continue
			}
			for _, t := range area.NearTargets(npc, npc.SightRange()) {
				if t == npc || !t.Live() {
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
		skill := ai.combatSkill(npc, npc.Targets()[0])
		if skill == nil {
			continue
		}
		npc.UseSkill(skill)
		break
	}
	// Reset timers.
	if ai.moveTimer >= move_freq {
		ai.moveTimer = 0
	}
	if ai.chatTimer >= chat_freq {
		ai.chatTimer = 0
	}
}

// AddCharacter adds specified character to control by AI.
func (ai *AI) AddCharacter(char *character.Character) {
	ai.npcs[char.ID()+char.Serial()] = char
}

// RemoveCharacters removes specified character from AI control.
func (ai *AI) RemoveCharacter(char *character.Character) {
	delete(ai.npcs, char.ID()+char.Serial())
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
	switch npc.Race() {
	case character.Human:
		t := lang.AllText(ai.game.Module().Conf().LangPath(), "random_chat", npc.Race().ID())
		if len(t) < 1 {
			return
		}
		id := rng.RollInt(1, len(t))
		id -= 1
		npc.SendChat(t[id])
	}
}

// combatSkill selects NPC skill to use in combat or nil if specified
// NPC has no skills to use.
func (ai *AI) combatSkill(npc *character.Character, tar effect.Target) *skill.Skill {
	if len(npc.Skills()) < 1 {
		return nil
	}
	return npc.Skills()[0]
}
