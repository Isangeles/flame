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
	Effects          []EffectData
	Skills           []SkillData
	Armors           []ArmorData
	Weapons          []WeaponData
	Miscs            []MiscItemData
	Characters       []CharacterData
	Objects          []ObjectData
	Dialogs          []DialogData
	Quests           []QuestData
	Recipes          []RecipeData
	Areas            []AreaData
	Races            []RaceData
	Trainings        []TrainingData
	TranslationBases []*TranslationBaseData
)

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
	misc := Misc(id)
	if misc != nil {
		return misc
	}
	return nil
}

// Effect returns effect data for specified ID.
func Effect(id string) *EffectData {
	for _, d := range Effects {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Skill returns skill data for specified ID.
func Skill(id string) *SkillData {
	for _, d := range Skills {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Armor returns armor data for specified ID.
func Armor(id string) *ArmorData {
	for _, d := range Armors {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Weapon returns weapon data for specified ID.
func Weapon(id string) *WeaponData {
	for _, d := range Weapons {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Misc returns misc data for specified ID.
func Misc(id string) *MiscItemData {
	for _, d := range Miscs {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Character returns character data for specified ID
// and serial value.
func Character(id, serial string) *CharacterData {
	for _, d := range Characters {
		if d.ID == id && d.Serial == serial {
			return &d
		}
	}
	return nil
}

// Object returns skill data for specified ID and
// serial value.
func Object(id, serial string) *ObjectData {
	for _, d := range Objects {
		if d.ID == id && d.Serial == serial {
			return &d
		}
	}
	return nil
}

// Dialog returns dialog data for specified ID.
func Dialog(id string) *DialogData {
	for _, d := range Dialogs {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Quest returns quest data for specified ID.
func Quest(id string) *QuestData {
	for _, d := range Quests {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Recipe returns recipe data for specified ID.
func Recipe(id string) *RecipeData {
	for _, d := range Recipes {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Area returns area data for specified ID.
func Area(id string) *AreaData {
	for _, d := range Areas {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Race returns race data for specified ID.
func Race(id string) *RaceData {
	for _, d := range Races {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// Training returns training data for specified ID.
func Training(id string) *TrainingData {
	for _, d := range Trainings {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

// TranslationBase returns translation base for specified ID.
func TranslationBase(id string) *TranslationBaseData {
	for _, b := range TranslationBases {
		if b.ID == id {
			return b
		}
	}
	return nil
}

// Clear removes all resources from base.
func Clear() {
	Effects = make([]EffectData, 0)
	Skills = make([]SkillData, 0)
	Armors = make([]ArmorData, 0)
	Weapons = make([]WeaponData, 0)
	Miscs = make([]MiscItemData, 0)
	Characters = make([]CharacterData, 0)
	Objects = make([]ObjectData, 0)
	Dialogs = make([]DialogData, 0)
	Quests = make([]QuestData, 0)
	Recipes = make([]RecipeData, 0)
	Areas = make([]AreaData, 0)
	Races = make([]RaceData, 0)
	Trainings = make([]TrainingData, 0)
	TranslationBases = make([]*TranslationBaseData, 0)
}

// Add adds resources from specified module data
// to resources base.
func Add(r ResourcesData) {
	Characters = append(Characters, r.Characters...)
	Races = append(Races, r.Races...)
	Objects = append(Objects, r.Objects...)
	Effects = append(Effects, r.Effects...)
	Skills = append(Skills, r.Skills...)
	Armors = append(Armors, r.Armors...)
	Weapons = append(Weapons, r.Weapons...)
	Miscs = append(Miscs, r.Miscs...)
	Dialogs = append(Dialogs, r.Dialogs...)
	Quests = append(Quests, r.Quests...)
	Recipes = append(Recipes, r.Recipes...)
	Trainings = append(Trainings, r.Trainings...)
	Areas = append(Areas, r.Areas...)
	for _, b := range r.TranslationBases {
		TranslationBases = append(TranslationBases, &b)
	}
}
