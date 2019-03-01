/*
 * res.go
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

package res

var (
	effectsData map[string]EffectData
	skillsData  map[string]SkillData
	weaponsData map[string]WeaponData
	charsData   map[string]CharacterData
)

// Effect returns resources for effect
// with specified ID or empty resource
// struct if data for specified effect ID
// was not found.
func Effect(id string) EffectData {
	return effectsData[id]
}

// Skill returns resources for skill
// with specified iD or empty resource
// struct if data for specified skill ID
// was not found.
func Skill(id string) SkillData {
	return skillsData[id]
}

// Weapon returns weapon resource data
// for weapon with specified ID or empty
// weapon resurce struct if data for specified
// ID was not found.
func Weapon(id string) WeaponData {
	return weaponsData[id]
}

// Character returns character resource data
// for character with specified ID or empty
// character resurce struct if data for specified
// ID was not found. 
func Character(id string) CharacterData {
	return charsData[id]
}

// SetEffectsData sets specified effects data
// as effects resources.
func SetEffectsData(data []EffectData) {
	effectsData = make(map[string]EffectData)
	for _, ed := range data {
		effectsData[ed.ID] = ed
	}
}

// SetSkillsData sets specified skills data
// as skills resources.
func SetSkillsData(data []SkillData) {
	skillsData = make(map[string]SkillData)
	for _, sd := range data {
		skillsData[sd.ID] = sd
	}
}

// SetWeaponsData sets specified weapons data as
// weapons resources.
func SetWeaponsData(data []WeaponData) {
	weaponsData = make(map[string]WeaponData)
	for _, wd := range data {
		weaponsData[wd.ID] = wd
	}
}

// SetCharactersData sets specified characters data as
// characters resource.
func SetCharactersData(data []CharacterData) {
	charsData = make(map[string]CharacterData)
	for _, cd := range data {
		charsData[cd.BasicData.ID] = cd
	}
}
