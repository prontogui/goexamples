package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func snackbarDemo(pgui pg.ProntoGUI) {

	heading := pg.TextWith{
		Content:    "Consectetur Adipiscing Elit",
		Embodiment: "{\"fontSize\":\"25.0\"}",
	}.Make()

	cmd := pg.NewCommand("Show Snackbar")

	snackbar := pg.NewFrame(pg.NewText("This is an important message!")).SetEmbodiment("snackbar")

	pgui.SetGUI(heading, cmd, snackbar)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}

		if cmd.Issued() {
			snackbar.SetShowing(true)
		}
	}
}
