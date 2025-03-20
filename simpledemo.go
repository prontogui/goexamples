// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func simpleDemo(pgui pg.ProntoGUI) {

	heading := pg.TextWith{
		Content:    "Consectetur Adipiscing Elit",
		Embodiment: "{\"fontSize\":\"25.0\"}",
	}.Make()
	spacing := pg.TextWith{Content: "    "}.Make()

	txt1 := pg.TextWith{Content: "Bibendum X"}.Make()
	tf1 := pg.TextFieldWith{}.Make()
	grp1 := pg.GroupWith{GroupItems: []pg.Primitive{txt1, tf1, spacing}}.Make()

	txt2 := pg.TextWith{Content: "Bibendum Y"}.Make()
	tf2 := pg.TextFieldWith{}.Make()
	grp2 := pg.GroupWith{GroupItems: []pg.Primitive{txt2, tf2, spacing}}.Make()

	chk := pg.CheckWith{Label: "Magna aliqua"}.Make()
	grp3 := pg.GroupWith{GroupItems: []pg.Primitive{spacing, spacing, chk}}.Make()

	cmd1 := pg.CommandWith{Label: "Imperdiet"}.Make()
	cmd2 := pg.CommandWith{Label: "Suscipit"}.Make()
	grp4 := pg.GroupWith{GroupItems: []pg.Primitive{spacing, spacing, spacing, cmd1, spacing, cmd2}}.Make()

	pgui.SetGUI(heading, spacing, spacing, grp1, spacing, grp2, spacing, grp3, spacing, grp4)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}
