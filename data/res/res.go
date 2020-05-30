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
	questsData  map[string]QuestData
	recipesData map[string]RecipeData
	areasData   map[string]AreaData
	racesData   map[string]RaceData
	langData    map[string]TranslationData
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
	questsData = make(map[string]QuestData)
	recipesData = make(map[string]RecipeData)
	areasData = make(map[string]AreaData)
	racesData = make(map[string]RaceData)
	langData = make(map[string]TranslationData)
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

// Quest returns quest resource data
// for quest with specified ID or nil
// if data for specified ID was not found.
func Quest(id string) *QuestData {
	qd := questsData[id]
	if len(qd.ID) < 1 {
		return nil
	}
	return &qd
}

// Recipe returns recipe resource data
// for recipe with specified ID or nil
// if data for specified ID was not found.
func Recipe(id string) *RecipeData {
	rd := recipesData[id]
	if len(rd.ID) < 1 {
		return nil
	}
	return &rd
}

// Area returns area resource data
// for area with specified ID or nil
// if data for specified ID was not found.
func Area(id string) *AreaData {
	ad := areasData[id]
	if len(ad.ID) < 1 {
		return nil
	}
	return &ad
}

// Race returns area resource data
// for race with specified ID or nil
// if data for specified ID was not found.
func Race(id string) *RaceData {
	rd := racesData[id]
	if len(rd.ID) < 1 {
		return nil
	}
	return &rd
}

// Translation returns translation data
// texts for specified ID.
func Translation(id string) *TranslationData {
	ld := langData[id]
	if len(ld.ID) < 1 {
		return nil
	}
	return &ld
}

// Quests returns all quests resources.
func Quests() (d []QuestData) {
	for _, qd := range questsData {
		d = append(d, qd)
	}
	return
}

// Recipes returns all recipes resources.
func Recipes() (r []RecipeData) {
	for _, rd := range recipesData {
		r = append(r, rd)
	}
	return
}

// Areas returns all areas resources.
func Areas() (a []AreaData) {
	for _, ad := range areasData {
		a = append(a, ad)
	}
	return
}

// Races returns all races resources.
func Races() (r []RaceData) {
	for _, rd := range racesData {
		r = append(r, rd)
	}
	return
}

// Translations returns all translation resources.
func Translations() (t []TranslationData) {
	for _, td := range langData {
		t = append(t, td)
	}
	return
}

// SetQuestsData sets specified quests data as
// quests resources.
func SetQuestsData(data []QuestData) {
	questsData = make(map[string]QuestData)
	for _, qd := range data {
		questsData[qd.ID] = qd
	}
}

// SetRecipesData sets specified recipes data as
// recipes resources.
func SetRecipesData(data []RecipeData) {
	recipesData = make(map[string]RecipeData)
	for _, rd := range data {
		recipesData[rd.ID] = rd
	}
}

// SetAreasData sets specified data as
// areas resources.
func SetAreasData(data []AreaData) {
	areasData = make(map[string]AreaData)
	for _, ad := range data {
		areasData[ad.ID] = ad
	}
}

// SetRacesData sets specified data as
// races resources.
func SetRacesData(data []RaceData) {
	racesData = make(map[string]RaceData)
	for _, rd := range data {
		racesData[rd.ID] = rd
	}
}

// SetTranslationData sets specified data as
// translation resources.
func SetTranslationData(data []TranslationData) {
	langData = make(map[string]TranslationData)
	for _, td := range data {
		langData[td.ID] = td
	}
}

// SetModuleData sets resources from specified module data.
func SetModuleData(mod ModuleData) {
	chars := append(mod.Resources.Characters, mod.Chapter.Resources.Characters...)
	for _, c := range chars {
		Characters[c.ID] = c
	}
	races := append(mod.Resources.Races, mod.Chapter.Resources.Races...)
	SetRacesData(races)
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
	SetQuestsData(quests)
	recipes := append(mod.Resources.Recipes, mod.Chapter.Resources.Recipes...)
	SetRecipesData(recipes)
	areas := append(mod.Resources.Areas, mod.Chapter.Resources.Areas...)
	SetAreasData(areas)
	translations := append(Translations(), mod.Resources.Translations...)
	translations = append(translations, mod.Chapter.Resources.Translations...)
	SetTranslationData(translations)
}

// SetModuleData sets resources from specified module data.
func AddResources(r ResourcesData) {
	for _, c := range r.Characters {
		Characters[c.ID] = c
	}
	SetRacesData(append(Races(), r.Races...))
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
	SetQuestsData(append(Quests(), r.Quests...))
	SetRecipesData(append(Recipes(), r.Recipes...))
	SetAreasData(append(Areas(), r.Areas...))
	SetTranslationData(append(Translations(), r.Translations...))
}
