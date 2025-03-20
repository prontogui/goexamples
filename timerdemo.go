// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func timerdemo(pgui pg.ProntoGUI) {

	timer := pg.TimerWith{PeriodMs: 200}.Make()
	txt := pg.TextWith{Content: "Reading"}.Make()
	reading := pg.TextWith{Content: ""}.Make()
	cmd := pg.CommandWith{Label: "+100"}.Make()

	// Set the GUI.  Note:  we wrap the reading in a group to improve the rendering
	// performance.  This is because the reading is updated frequently and we don't
	// want to re-render the entire GUI each time it changes.
	pgui.SetGUI(timer, txt, pg.NewGroup(reading), cmd)

	ticker := 0.0

	for {
		updatedPrimitive, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}

		if updatedPrimitive == cmd {
			ticker = ticker + 100
		} else if updatedPrimitive == timer {
			ticker = ticker + 1
			reading.SetContent(fmt.Sprintf("%f", ticker))
			fmt.Printf("Tick\n")
		}
	}
}
