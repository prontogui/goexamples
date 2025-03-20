// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func miscDemo(pgui pg.ProntoGUI) {

	txt := pg.TextWith{Content: "Nice to meet you."}.Make()
	cmd := pg.CommandWith{Label: "Click Me!"}.Make()
	choice := pg.ChoiceWith{Choices: []string{"Apple", "Orange", "Banana"}, Choice: "Apple"}.Make()
	chk := pg.CheckWith{Label: "Turn this on or off"}.Make()
	tri := pg.TristateWith{Label: "Select all toppings"}.Make()
	spacer := pg.TextWith{Content: " "}.Make()

	grp := pg.GroupWith{GroupItems: []pg.Primitive{
		txt,
		spacer,
		cmd,
		spacer,
		choice,
		spacer,
		chk,
		spacer,
		tri,
	}}.Make()

	pgui.SetGUI(grp)

	for {
		updatedPrimitive, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}

		if updatedPrimitive == cmd {
			fmt.Println("Command clicked.")
		}
		if updatedPrimitive == choice {
			fmt.Printf("Choice '%s' is a good one.\n", choice.Choice())
		}
		if updatedPrimitive == chk {
			if chk.Checked() {
				fmt.Println("Check is turned ON.")
				txt.SetContent("Hello!")
			} else {
				fmt.Println("Check is turned off.")
				txt.SetContent("World!")
			}
		}
		if updatedPrimitive == tri {
			state := tri.State()
			if state == 0 {
				fmt.Println("Trump = 0")
			} else if state == 1 {
				fmt.Println("Biden = 1")
			} else {
				fmt.Println("Undecided = 2")
			}
		}
	}

}
