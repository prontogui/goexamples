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

	if *pings > 0 {
		for i := 1; i <= *pings; i++ {
			fmt.Printf("Ping #%d", i)
			pgcomm.ExchangeUpdates([]byte{'p', 'i', 'n', 'g'})
			time.Sleep(5 * time.Second)
		}

	} else {
		fmt.Println(pg.CallMe())
	}
}
