// Copyright 2024-2026 ProntoGUI, LLC
// ProntoGUI™ is a trademark of ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func listDemo3(pgui pg.ProntoGUI) {

	// Build a file-explorer-like hierarchy:
	//
	// Documents/          (level 0, folder)
	//   Reports/          (level 1, folder)
	//     Q1 Report.pdf   (level 2, item)
	//     Q2 Report.pdf   (level 2, item)
	//   Notes/            (level 1, folder)
	//     Meeting.txt     (level 2, item)
	// Pictures/           (level 0, folder)
	//   Vacation.jpg      (level 1, item)
	//   Family.jpg        (level 1, item)

	folderList := pg.ListWith{
		Embodiment:    "folder-list",
		SelectionMode: 1,
		ModelFolder:   pg.FolderWith{Embodiment: "folderIcons:plusMinus"}.Make(),
		ListItems: []pg.Primitive{
			pg.FolderWith{LabelItem: pg.TextWith{Content: "Documents"}.Make(), Level: 0, Expanded: true}.Make(),
			pg.FolderWith{LabelItem: pg.TextWith{Content: "Reports"}.Make(), Level: 1, Expanded: true}.Make(),
			pg.FolderItemWith{Item: pg.TextWith{Content: "Q1 Report.pdf"}.Make(), Level: 2}.Make(),
			pg.FolderItemWith{Item: pg.TextWith{Content: "Q2 Report.pdf"}.Make(), Level: 2}.Make(),
			pg.FolderWith{LabelItem: pg.TextWith{Content: "Notes"}.Make(), Level: 1}.Make(),
			pg.FolderItemWith{Item: pg.TextWith{Content: "Meeting.txt"}.Make(), Level: 2}.Make(),
			pg.FolderWith{LabelItem: pg.TextWith{Content: "Pictures"}.Make(), Level: 0, Expanded: true}.Make(),
			pg.FolderItemWith{Item: pg.TextWith{Content: "Vacation.jpg"}.Make(), Level: 1}.Make(),
			pg.FolderItemWith{Item: pg.TextWith{Content: "Family.jpg"}.Make(), Level: 1}.Make(),
		},
	}.Make()

	pgui.SetGUI(folderList)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}
