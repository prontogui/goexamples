package main

import (
	"flag"
	"fmt"
	"time"

	pg "github.com/prontogui/golib"
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
	txt := pg.Text{}
	cmd := pg.Command{}
	cmd.SetLabel("Click Me!")
	pgui.SetGUI(&txt, &cmd)

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

	pgui.StopServing()
}
