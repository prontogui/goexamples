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

	cmd1 := pg.NewCommand("Show snackbar with text")
	cmd2 := pg.NewCommand("Show snackbar with text + button")
	msg := pg.NewText("")

	snackbar1 := pg.NewFrame(pg.NewText("This is an important message!")).SetEmbodiment("snackbar")

	detailsCmd := pg.NewCommand("Details")
	snackbar2 := pg.NewFrame(pg.NewText("Something didn't go as planned"), detailsCmd).SetEmbodiment("snackbar")

	pgui.SetGUI(heading, cmd1, cmd2, msg, snackbar1, snackbar2)

	detailClicks := 0

	for {

		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}

		if cmd1.Issued() {
			snackbar1.SetShowing(true)
		}

		if cmd2.Issued() {
			snackbar2.SetShowing(true)
		}

		if detailsCmd.Issued() {
			detailClicks = detailClicks + 1
			msg.SetContent(fmt.Sprintf("Details clicked %d times.", detailClicks))
		}
	}
}
