package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func textFieldDemo(pgui pg.ProntoGUI) {

	txt := pg.TextWith{Content: "Enter here:"}.Make()
	tf := pg.TextFieldWith{TextEntry: "INITIAL TEXT"}.Make()

	pgui.SetGUI(txt, tf)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}

		fmt.Printf("Entered text:  %s\n", tf.TextEntry())
	}
}
