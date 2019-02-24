/*
 * effect.go
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

package data

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"path/filepath"
	
	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/text"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/log"
)

const (
	EFFECTS_FILE_EXT = ".effects"
)

// Effect creates new instance of effect with specified ID
// for specified module, returns error if effect data with such ID
// was not found or module failed to assign serial value for
// effect.
func Effect(mod *module.Module, id string) (*effect.Effect, error) {
	data := res.Effect(id)
	if data.ID == "" {
		return nil, fmt.Errorf("effect_not_found:%s", id)
	}
	e := effect.NewEffect(data)
	effectsLangPath := filepath.FromSlash(mod.Conf().LangPath() + "/effects" +
		text.LANG_FILE_EXT)
	name := text.ReadDisplayText(effectsLangPath, e.ID())
	e.SetName(name[0])
	err := mod.AssignSerial(e)
	if err != nil {
		return nil, fmt.Errorf("fail_to_assing_serial_value:%v", err)
	}
	return e, nil
}

// ImportEffects imports all XML effects data from effects base
// with specified path.
func ImportEffects(basePath string) ([]res.EffectData, error) {
	effects := make([]res.EffectData, 0)
	doc, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_effects_base_file:%v", err)
	}
	defer doc.Close()
	xmlEffects, err := parsexml.UnmarshalEffectsBase(doc)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_effects_base:%v", err)
	}
	for _, xmlEffect := range xmlEffects {
		e := buildXMLEffectData(xmlEffect)
		effects = append(effects, e)
	}
	return effects, nil
}

// ImportEffectsDir imports all effects from files in
// specified directory.
func ImportEffectsDir(dirPath string) ([]res.EffectData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	effects := make([]res.EffectData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), EFFECTS_FILE_EXT) {
			continue
		}
		basePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		effs, err := ImportEffects(basePath)
		if err != nil {
			log.Err.Printf("data_effects_import:%s:fail_to_load_effects_file:%v",
				basePath, err)
			continue
		}
		for _, e := range effs {
			effects = append(effects, e)
		}
	}
	return effects, nil
}

// buildXMLEffect build effect from XML data.
func buildXMLEffectData(xmlEffect parsexml.EffectNodeXML) res.EffectData {
	mods := buildXMLModifiers(&xmlEffect.ModifiersNode)
	data := res.EffectData{
		ID: xmlEffect.ID,
		Duration: xmlEffect.Duration,
		Modifiers: mods,
		Subeffects: xmlEffect.Subeffects.Effects,
	}
	log.Dbg.Printf("subeffs_len:%d", len(data.Subeffects))
	return data
}
