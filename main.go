package main

import (
	"flag"
	"fmt"
	"time"

	pgcomm "github.com/prontogui/golib/pgcomm"
)

var (
	pings = flag.Int("pings", 0, "Do testing of communication by sending N pings to app (0 = no pings).")
)

func main() {

	flag.Parse()

	err := pgcomm.StartServing("127.0.0.1", 50053)

	if err != nil {
		fmt.Printf("Error trying to start server:  %s", err.Error())
		return
	}

	if *pings > 0 {
		for i := 1; i <= *pings; i++ {
			fmt.Printf("Ping #%d\n", i)
			update, err := pgcomm.ExchangeUpdates("ping")
			if err != nil {
				fmt.Printf("Error while exchanging updates: %s.  Exiting.", err.Error())
				return
			}
			fmt.Printf("Response: %s\n", update.(string))
			time.Sleep(5 * time.Second)
		}

	}
}
