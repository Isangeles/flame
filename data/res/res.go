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
	Effects map[string]EffectData
	skillsData  map[string]SkillData
	armorsData  map[string]ArmorData
	weaponsData map[string]WeaponData
	miscsData   map[string]MiscItemData
	charsData   map[string]CharacterData
	objectsData map[string]ObjectData
	dialogsData map[string]DialogData
	questsData  map[string]QuestData
	recipesData map[string]RecipeData
	areasData   map[string]AreaData
	racesData   map[string]RaceData
	langData    map[string]TranslationData
)

// On init.
func init() {
	Effects = make(map[string]EffectData)
	skillsData = make(map[string]SkillData)
	armorsData = make(map[string]ArmorData)
	weaponsData = make(map[string]WeaponData)
	miscsData = make(map[string]MiscItemData)
	charsData = make(map[string]CharacterData)
	objectsData = make(map[string]ObjectData)
	dialogsData = make(map[string]DialogData)
	questsData = make(map[string]QuestData)
	recipesData = make(map[string]RecipeData)
	areasData = make(map[string]AreaData)
	racesData = make(map[string]RaceData)
	langData = make(map[string]TranslationData)
}

// Skill returns resources for skill
// with specified ID or empty resource
// struct if data for specified skill ID
// was not found.
func Skill(id string) *SkillData {
	sd := skillsData[id]
	if len(sd.ID) < 1 {
		return nil
	}
	return &sd
}

// Armor returns resource for armor
// with specified ID or nil if data
// for specified ID was not found.
func Armor(id string) *ArmorData {
	ad := armorsData[id]
	if len(ad.ID) < 1 {
		return nil
	}
	return &ad
}

// Weapon returns weapon resource data
// for weapon with specified ID or nil if
// data for specified ID was not found.
func Weapon(id string) *WeaponData {
	wd := weaponsData[id]
	if len(wd.ID) < 1 {
		return nil
	}
	return &wd
}

// Misc returns misc item resource data
// for miscellaneous item with specified ID or
// nil if data for specified ID was not found.
func MiscItem(id string) *MiscItemData {
	md := miscsData[id]
	if len(md.ID) < 1 {
		return nil
	}
	return &md
}

// Item returns item resource data for item
// with specified ID or nil if data for
// specified ID was not found.
func Item(id string) ItemData {
	armor := Armor(id)
	if armor != nil {
		return armor
	}
	weapon := Weapon(id)
	if weapon != nil {
		return weapon
	}
	misc := miscsData[id]
	if len(misc.ID) < 1 {
		return nil
	}
	return &misc
}

// Character returns character resource data
// for character with specified ID or nil
// if data for specified ID was not found.
func Character(id string) *CharacterData {
	cd := charsData[id]
	if len(cd.ID) < 1 {
		return nil
	}
	return &cd
}

// Object returns object resource data
// for object with specified ID or nil
// if data for specified ID was not found.
func Object(id string) *ObjectData {
	od := objectsData[id]
	if len(od.ID) < 1 {
		return nil
	}
	return &od
}

// Dialog returns dialog resource data
// for dialog with specified ID or nil
// if data for specified ID was not found.
func Dialog(id string) *DialogData {
	dd := dialogsData[id]
	if len(dd.ID) < 1 {
		return nil
	}
	return &dd
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

// Characters returns all characters
// resources.
func Characters() (d []CharacterData) {
	for _, cd := range charsData {
		d = append(d, cd)
	}
	return
}

// Armors returns all  armors
// resources.
func Armors() (d []ArmorData) {
	for _, ad := range armorsData {
		d = append(d, ad)
	}
	return
}

// Weapons returns all weapons
// resources.
func Weapons() (d []WeaponData) {
	for _, wd := range weaponsData {
		d = append(d, wd)
	}
	return
}

// MiscItems returns all misc items
// resources.
func MiscItems() (d []MiscItemData) {
	for _, md := range miscsData {
		d = append(d, md)
	}
	return
}

// Objects returns all objects resources.
func Objects() (d []ObjectData) {
	for _, od := range objectsData {
		d = append(d, od)
	}
	return
}

// Dialogs returns all dialogs resources.
func Dialogs() (d []DialogData) {
	for _, dd := range dialogsData {
		d = append(d, dd)
	}
	return
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

// Skills returns all skills resources.
func Skills() (r []SkillData) {
	for _, sd := range skillsData {
		r = append(r, sd)
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

// SetSkillsData sets specified skills data
// as skills resources.
func SetSkillsData(data []SkillData) {
	skillsData = make(map[string]SkillData)
	for _, sd := range data {
		skillsData[sd.ID] = sd
	}
}

// SetArmorsData sets specified armors data as
// armors resources.
func SetArmorsData(data []ArmorData) {
	armorsData = make(map[string]ArmorData)
	for _, ad := range data {
		armorsData[ad.ID] = ad
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

// SetMiscItemsData sets specified misc items data as
// misc items resources.
func SetMiscItemsData(data []MiscItemData) {
	miscsData = make(map[string]MiscItemData)
	for _, md := range data {
		miscsData[md.ID] = md
	}
}

// SetCharactersData sets specified characters data as
// characters resources.
func SetCharactersData(data []CharacterData) {
	charsData = make(map[string]CharacterData)
	for _, cd := range data {
		charsData[cd.ID] = cd
	}
}

// SetObjectsData sets specified objects data as
// objects resources.
func SetObjectsData(data []ObjectData) {
	objectsData = make(map[string]ObjectData)
	for _, od := range data {
		objectsData[od.ID] = od
	}
}

// SetDialogsData sets specified dialogs data as
// dialogs resources.
func SetDialogsData(data []DialogData) {
	dialogsData = make(map[string]DialogData)
	for _, dd := range data {
		dialogsData[dd.ID] = dd
	}
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
	SetCharactersData(chars)
	races := append(mod.Resources.Races, mod.Chapter.Resources.Races...)
	SetRacesData(races)
	objects := append(mod.Resources.Objects, mod.Chapter.Resources.Objects...)
	SetObjectsData(objects)
	effects := append(mod.Resources.Effects, mod.Chapter.Resources.Effects...)
	for _, e := range effects {
		Effects[e.ID] = e
	}
	skills := append(mod.Resources.Skills, mod.Chapter.Resources.Skills...)
	SetSkillsData(skills)
	armors := append(mod.Resources.Armors, mod.Chapter.Resources.Armors...)
	SetArmorsData(armors)
	weapons := append(mod.Resources.Weapons, mod.Chapter.Resources.Weapons...)
	SetWeaponsData(weapons)
	miscs := append(mod.Resources.Miscs, mod.Chapter.Resources.Miscs...)
	SetMiscItemsData(miscs)
	dialogs := append(mod.Resources.Dialogs, mod.Chapter.Resources.Dialogs...)
	SetDialogsData(dialogs)
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
	SetCharactersData(append(Characters(), r.Characters...))
	SetRacesData(append(Races(), r.Races...))
	SetObjectsData(append(Objects(), r.Objects...))
	for _, e := range r.Effects {
		Effects[e.ID] = e
	}
	SetSkillsData(append(Skills(), r.Skills...))
	SetArmorsData(append(Armors(), r.Armors...))
	SetWeaponsData(append(Weapons(), r.Weapons...))
	SetMiscItemsData(append(MiscItems(), r.Miscs...))
	SetDialogsData(append(Dialogs(), r.Dialogs...))
	SetQuestsData(append(Quests(), r.Quests...))
	SetRecipesData(append(Recipes(), r.Recipes...))
	SetAreasData(append(Areas(), r.Areas...))
	SetTranslationData(append(Translations(), r.Translations...))
}
