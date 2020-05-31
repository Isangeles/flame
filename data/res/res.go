/*
 * res.go
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

package res

var (
	Effects      map[string]EffectData
	Skills       map[string]SkillData
	Armors       map[string]ArmorData
	Weapons      map[string]WeaponData
	Miscs        map[string]MiscItemData
	Characters   map[string]CharacterData
	Objects      map[string]ObjectData
	Dialogs      map[string]DialogData
	Quests       map[string]QuestData
	Recipes      map[string]RecipeData
	Areas        map[string]AreaData
	Races        map[string]RaceData
	Translations map[string]TranslationData
)

// On init.
func init() {
	Effects = make(map[string]EffectData)
	Skills = make(map[string]SkillData)
	Armors = make(map[string]ArmorData)
	Weapons = make(map[string]WeaponData)
	Miscs = make(map[string]MiscItemData)
	Characters = make(map[string]CharacterData)
	Objects = make(map[string]ObjectData)
	Dialogs = make(map[string]DialogData)
	Quests = make(map[string]QuestData)
	Recipes = make(map[string]RecipeData)
	Areas = make(map[string]AreaData)
	Races = make(map[string]RaceData)
	Translations = make(map[string]TranslationData)
}

// Item returns item resource data for item
// with specified ID or nil if data for
// specified ID was not found.
func Item(id string) ItemData {
	armor, ok := Armors[id]
	if ok {
		return armor
	}
	weapon, ok := Weapons[id]
	if ok {
		return weapon
	}
	misc, ok := Miscs[id]
	if ok {
		return misc
	}
	return nil
}

// SetModuleData sets resources from specified module data.
func SetModuleData(mod ModuleData) {
	chars := append(mod.Resources.Characters, mod.Chapter.Resources.Characters...)
	for _, c := range chars {
		Characters[c.ID] = c
	}
	races := append(mod.Resources.Races, mod.Chapter.Resources.Races...)
	for _, r := range races {
		Races[r.ID] = r
	}
	objects := append(mod.Resources.Objects, mod.Chapter.Resources.Objects...)
	for _, o := range objects {
		Objects[o.ID] = o
	}
	effects := append(mod.Resources.Effects, mod.Chapter.Resources.Effects...)
	for _, e := range effects {
		Effects[e.ID] = e
	}
	skills := append(mod.Resources.Skills, mod.Chapter.Resources.Skills...)
	for _, s := range skills {
		Skills[s.ID] = s
	}
	armors := append(mod.Resources.Armors, mod.Chapter.Resources.Armors...)
	for _, a := range armors {
		Armors[a.ID] = a
	}
	weapons := append(mod.Resources.Weapons, mod.Chapter.Resources.Weapons...)
	for _, w := range weapons {
		Weapons[w.ID] = w
	}
	miscs := append(mod.Resources.Miscs, mod.Chapter.Resources.Miscs...)
	for _, m := range miscs {
		Miscs[m.ID] = m
	}
	dialogs := append(mod.Resources.Dialogs, mod.Chapter.Resources.Dialogs...)
	for _, d := range dialogs {
		Dialogs[d.ID] = d
	}
	quests := append(mod.Resources.Quests, mod.Chapter.Resources.Quests...)
	for _, q := range quests {
		Quests[q.ID] = q
	}
	recipes := append(mod.Resources.Recipes, mod.Chapter.Resources.Recipes...)
	for _, r := range recipes {
		Recipes[r.ID] = r
	}
	areas := append(mod.Resources.Areas, mod.Chapter.Resources.Areas...)
	for _, a := range areas {
		Areas[a.ID] = a
	}
	translations := append(mod.Resources.Translations, mod.Chapter.Resources.Translations...)
	for _, t := range translations {
		Translations[t.ID] = t
	}
}

// SetModuleData sets resources from specified module data.
func AddResources(r ResourcesData) {
	for _, c := range r.Characters {
		Characters[c.ID] = c
	}
	for _, r := range r.Races {
		Races[r.ID] = r
	}
	for _, o := range r.Objects {
		Objects[o.ID] = o
	}
	for _, e := range r.Effects {
		Effects[e.ID] = e
	}
	for _, s := range r.Skills {
		Skills[s.ID] = s
	}
	for _, a := range r.Armors {
		Armors[a.ID] = a
	}
	for _, w := range r.Weapons {
		Weapons[w.ID] = w
	}
	for _, m := range r.Miscs {
		Miscs[m.ID] = m
	}
	for _, d := range r.Dialogs {
		Dialogs[d.ID] = d
	}
	for _, q := range r.Quests {
		Quests[q.ID] = q
	}
	for _, r := range r.Recipes {
		Recipes[r.ID] = r
	}
	for _, a := range r.Areas {
		Areas[a.ID] = a
	}
	for _, t := range r.Translations {
		Translations[t.ID] = t
	}
}
