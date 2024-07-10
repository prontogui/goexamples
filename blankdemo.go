package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func blankDemo(pgui pg.ProntoGUI) {

	pgui.SetGUI(pg.GroupWith{}.Make())

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}
