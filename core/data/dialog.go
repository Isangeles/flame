/*
 * dialog.go
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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/isangeles/flame/core/data/parsexml"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/module/object/dialog"
	"github.com/isangeles/flame/log"
)

const (
	DIALOGS_FILE_EXT = ".dialogs"
)

func Dialog(id string) (*dialog.Dialog, error) {
	data := res.Dialog(id)
	if data == nil {
		return nil, fmt.Errorf("dialog_not_found")
	}
	d, err := dialog.NewDialog(*data)
	if err != nil {
		return nil, fmt.Errorf("fail_to_create_dialog:%v", err)
	}
	return d, nil
}

// ImportDialogs imports all dialogs form base file with
// specified path.
func ImportDialogs(basePath string) ([]*res.DialogData, error) {
	doc, err := os.Open(basePath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_open_base_file:%v", err)
	}
	defer doc.Close()
	dialogs, err := parsexml.UnmarshalDialogsBase(doc)
	if err != nil {
		return nil, fmt.Errorf("fail_to_unmarshal_xml_data:%v", err)
	}
	return dialogs, nil
}

// ImportDialogsDir imports all dialogs form base files in
// directory with specified path.
func ImportDialogsDir(dirPath string) ([]*res.DialogData, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("fail_to_read_dir:%v", err)
	}
	dialogs := make([]*res.DialogData, 0)
	for _, finfo := range files {
		if !strings.HasSuffix(finfo.Name(), DIALOGS_FILE_EXT) {
			continue
		}
		basePath := filepath.FromSlash(dirPath + "/" + finfo.Name())
		dd, err := ImportDialogs(basePath)
		if err != nil {
			log.Err.Printf("data_dialogs_import:%s:fail_to_import_base:%v",
				basePath, err)
		}
		for _, d := range dd {
			dialogs = append(dialogs, d)
		}
	}
	return dialogs, nil
}
