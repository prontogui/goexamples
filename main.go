package main

import (
	"flag"
	"fmt"

	pg "github.com/prontogui/golib"
	"github.com/prontogui/golib/primitive"
)

var (
	pings = flag.Int("pings", 10, "Do testing of communication by sending N pings to app (0 = no pings).")
)

func main() {

	flag.Parse()

	pgui := pg.NewProntoGUI()
	err := pgui.StartServing("127.0.0.1", 50053)

	if err != nil {
		fmt.Printf("Error trying to start server:  %s", err.Error())
		return
	}

	// Build the GUI model
	txt := pg.TextWith{Content: "Nice to meet you."}.Make()
	cmd := pg.CommandWith{Label: "Click Me!"}.Make()
	choice := pg.ChoiceWith{Choices: []string{"Apple", "Orange", "Banana"}, Choice: "Apple"}.Make()
	chk := pg.CheckWith{Label: "Turn this on or off"}.Make()
	tri := pg.TristateWith{Label: "Vote for biden or trump."}.Make()
	grp := pg.GroupWith{GroupItems: []primitive.Interface{txt, cmd, choice, chk, tri}}.Make()
	pgui.SetGUI(grp)

	for {
		err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}

		if cmd.Issued() {
			fmt.Print("Command clicked.")
		}
		if choice.Changed() {
			fmt.Printf("Choice '%s' is a good one.", choice.Choice())
		}
		if chk.Changed() {
			if chk.Checked() {
				fmt.Print("Check is turned ON.")
			} else {
				fmt.Printf("Check is turned off.")
			}
		}
		if tri.Changed() {
			state := tri.State()
			if state == 0 {
				fmt.Print("Trump = 0")
			} else if state == 1 {
				fmt.Print("Biden = 1")
			} else {
				fmt.Print("Undecided = 2")
			}
		}
	}
	/*
		if *pings > 0 {
			for i := 1; i <= *pings; i++ {
				pingMsg := fmt.Sprintf("Ping #%d\n", i)
				fmt.Println(pingMsg)

				txt.SetContent(pingMsg)

				err := pgui.Wait()
				if err != nil {
					fmt.Printf("error from Wait() is:  %s\n", err.Error())
					break
				}
				time.Sleep(5 * time.Second)
			}
		}
	*/

	pgui.StopServing()
}
