// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func filedemo(pgui pg.ProntoGUI) {

	txt1 := pg.TextWith{Content: "Import file"}.Make()
	importFile := pg.ImportFileWith{ValidExtensions: []string{"md", "txt", "go"}}.Make()
	grp1 := pg.GroupWith{GroupItems: []pg.Primitive{txt1, importFile}}.Make()

	txt2 := pg.TextWith{Content: "Export file"}.Make()
	exportFile := &pg.ExportFile{}
	grp2 := pg.GroupWith{GroupItems: []pg.Primitive{txt2, exportFile}}.Make()

	spacer := pg.TextWith{Content: " "}.Make()

	pgui.SetGUI(spacer, grp1, spacer, grp2)

	for {
		updatedPrimitive, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}

		if updatedPrimitive == importFile {
			exportFile.SetData(importFile.Data())
			exportFile.SetName("EXPORTED " + importFile.Name())
		}
	}
}
