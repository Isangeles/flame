/*
 * newcharacter.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/object/character"
)

// startNewCharacterDialog starts new CLI dialog to create new playable
// game character.
func newCharacterDialog() (*character.Character, error) {
	if flame.Mod() == nil {
		return nil, fmt.Errorf(lang.Text("ui", "cli_no_mod_err"))
	}
	var (
		name        string
		race        character.Race
		sex         character.Gender
		attrs       character.Attributes
		attrsPoints = flame.Mod().NewcharAttrsMax()
		c           *character.Character
	)

	scan := bufio.NewScanner(os.Stdin)

	// Character creation dialog
	var mainAccept = false
	for !mainAccept {
		// Name
		fmt.Printf("%s:", lang.Text("ui", "cli_newchar_name"))
		for scan.Scan() {
			name = scan.Text()
			if isCharNameValid(name) {
				break
			} else {
				fmt.Printf("%s\n",
					lang.Text("ui", "cli_newchar_invalid_name_err"))
				fmt.Printf("%s:", lang.Text("ui", "cli_newchar_name"))
			}
		}
		// Race.
		race = raceDialog()
		// Gender.
		sex = genderDialog()
		// Attributes.
		var accept = false
		for !accept {
			attrs = newAttributesDialog(attrsPoints)
			fmt.Printf("%s: %s\n",
				lang.Text("ui", "cli_newchar_attrs_summary"), attrs)
			fmt.Printf("%s:", lang.Text("ui", "cli_accept_dialog"))
			scan.Scan()
			input := scan.Text()
			if input != "r" {
				accept = true
			}
		}
		// Summary.
		charID := fmt.Sprintf("player_%s", name)
		charData := res.CharacterBasicData{
			ID: charID,
			Name: name,
			Level: 1,
			Sex: int(sex),
			Race: int(race),
			Attitude: int(character.Friendly),
			Alignment: int(character.True_neutral),
			Str: attrs.Str,
			Con: attrs.Con,
			Dex: attrs.Dex,
			Int: attrs.Int,
			Wis: attrs.Wis,
		}
		c = character.New(charData)
		fmt.Printf("%s: %s\n", lang.Text("ui", "cli_newchar_summary"),
			charDisplayString(c))
		fmt.Printf("%s:", lang.Text("ui", "cli_accept_dialog"))
		scan.Scan()
		input := scan.Text()
		if input != "r" {
			mainAccept = true
		}
	}

	return c, nil
}

// raceDialog starts CLI dialog for game character race.
// Returns character race.
func raceDialog() character.Race {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s:", lang.Text("ui", "cli_newchar_race"))
	racesNames := lang.Texts("ui", "race_human", "race_elf", "race_dwarf",
		"race_gnome")
	s := make([]interface{}, len(racesNames))
	for i, v := range racesNames {
		s[i] = v
	}
	for true {
		fmt.Printf("[1 - %s, 2 - %s, 3 - %s, 4 - %s]:", s...)
		scan.Scan()
		input := scan.Text()
		switch input {
		case "1":
			return character.Human
		case "2":
			return character.Elf
		case "3":
			return character.Dwarf
		case "4":
			return character.Gnome
		default:
			fmt.Printf("%s:%s\n", lang.Text("ui", "cli_newchar_invalid_value_err"),
				input)
		}
	}
	return character.Human
}

// genderDialog starts CLI dialog for game character gender.
// Returns character gender.
func genderDialog() character.Gender {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s:", lang.Text("ui", "cli_newchar_gender"))
	genderNames := lang.Texts("ui", "gender_male", "gender_female")
	s := make([]interface{}, len(genderNames))
	for i, v := range genderNames {
		s[i] = v
	}
	for true {
		fmt.Printf("[1 - %s, 2 - %s]:", s...)
		scan.Scan()
		input := scan.Text()
		switch input {
		case "1":
			return character.Male
		case "2":
			return character.Female
		default:
			fmt.Printf("%s:%s\n", lang.Text("ui", "cli_newchar_invalid_value_err"),
				input)
		}
	}
	return character.Male
}

// newAttributesDialog Starts CLI dialog for game character attributes.
// Returns character attributes.
func newAttributesDialog(attrsPoints int) (attrs character.Attributes) {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s:\n", lang.Text("ui", "cli_newchar_attrs"))
	for attrsPoints > 0 {
		// Strenght.
		for true {
			fmt.Printf("%s[%s = %d, %s = %d]+", lang.Text("ui", "attr_str"),
				lang.Text("ui", "cli_newchar_value"), attrs.Str,
				lang.Text("ui", "cli_newchar_points"), attrsPoints)
			scan.Scan()
			input := scan.Text()
			attr, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("%s:%s\n",
					lang.Text("ui", "cli_newchar_nan_error"), input)
			} else {
				if attrsPoints-attr >= 0 {
					attrs.Str += attr
					attrsPoints -= attr
					break
				} else {
					fmt.Printf("%s\n",
						lang.Text("ui", "cli_newchar_no_pts_error"))
				}
			}
		}
		// Constitution.
		for true {
			fmt.Printf("%s[%s = %d, %s = %d]+", lang.Text("ui", "attr_con"),
				lang.Text("ui", "cli_newchar_value"), attrs.Con,
				lang.Text("ui", "cli_newchar_points"), attrsPoints)
			scan.Scan()
			input := scan.Text()
			attr, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("%s:%s\n",
					lang.Text("ui", "cli_newchar_nan_error"), input)
			} else {
				if attrsPoints-attr >= 0 {
					attrs.Con += attr
					attrsPoints -= attr
					break
				} else {
					fmt.Printf("%s\n",
						lang.Text("ui", "cli_newchar_no_pts_error"))
				}
			}

		}
		// Dexterity.
		for true {
			fmt.Printf("%s[%s = %d, %s = %d]+", lang.Text("ui", "attr_dex"),
				lang.Text("ui", "cli_newchar_value"), attrs.Dex,
				lang.Text("ui", "cli_newchar_points"), attrsPoints)
			scan.Scan()
			input := scan.Text()
			attr, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("%s:%s\n",
					lang.Text("ui", "cli_newchar_nan_error"), input)
			} else {
				if attrsPoints-attr >= 0 {
					attrs.Dex += attr
					attrsPoints -= attr
					break
				} else {
					fmt.Printf("%s\n",
						lang.Text("ui", "cli_newchar_no_pts_error"))
				}
			}
		}
		// Wisdom.
		for true {
			fmt.Printf("%s[%s = %d, %s = %d]+", lang.Text("ui", "attr_wis"),
				lang.Text("ui", "cli_newchar_value"), attrs.Wis,
				lang.Text("ui", "cli_newchar_points"), attrsPoints)
			scan.Scan()
			input := scan.Text()
			attr, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("%s:%s\n",
					lang.Text("ui", "cli_newchar_nan_error"), input)
			} else {
				if attrsPoints-attr >= 0 {
					attrs.Wis += attr
					attrsPoints -= attr
					break
				} else {
					fmt.Printf("%s\n",
						lang.Text("ui", "cli_newchar_no_pts_error"))
				}
			}
		}
		// Inteligence.
		for true {
			fmt.Printf("%s[%s = %d, %s = %d]+", lang.Text("ui", "attr_int"),
				lang.Text("ui", "cli_newchar_value"), attrs.Int,
				lang.Text("ui", "cli_newchar_points"), attrsPoints)
			scan.Scan()
			input := scan.Text()
			attr, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("%s:%s\n", lang.Text("ui", "cli_newchar_nan_error"),
					input)
			} else {
				if attrsPoints-attr >= 0 {
					attrs.Int += attr
					attrsPoints -= attr
					break
				} else {
					fmt.Printf("%s\n",
						lang.Text("ui", "cli_newchar_no_pts_error"))
				}
			}
		}

	}
	return
}

// isCharNameVaild Checks if specified name is valid character name.
func isCharNameValid(name string) bool {
	return name != ""
}
