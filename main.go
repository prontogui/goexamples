package main

import (
	"flag"
	"fmt"
	"time"

	pg "github.com/prontogui/golib"
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
			ok, update := pgcomm.ExchangeUpdates([]byte{'p', 'i', 'n', 'g'})
			if !ok {
				fmt.Print("Error while exchanging updates.  Exiting.")
				return
			}
			fmt.Printf("Response: %s\n", string(update))
			time.Sleep(5 * time.Second)
		}

	} else {
		fmt.Println(pg.CallMe())
	}
}
