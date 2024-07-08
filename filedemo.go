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

	pgui.SetGUI(grp1, grp2)

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
