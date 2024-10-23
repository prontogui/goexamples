package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func tabledemo(pgui pg.ProntoGUI) {

	txt1 := pg.TextWith{Content: "Andy's Vending Machine Ordering Page"}.Make()

	t := pg.TableWith{
		Headings: []string{"Qty", "Description", "Price", "Subtotal", "Action"},
		Rows: [][]pg.Primitive{
			{
				pg.TextWith{Content: "10"}.Make(),
				pg.TextWith{Content: "Pepsi Max, 20 oz bottle"}.Make(),
				pg.TextWith{Content: "1.99"}.Make(),
				pg.TextWith{Content: "19.90"}.Make(),
				pg.CommandWith{Label: "Order Now"}.Make(),
			},
			{
				pg.TextWith{Content: "30"}.Make(),
				pg.TextWith{Content: "Snickers candy bar"}.Make(),
				pg.TextWith{Content: "1.99"}.Make(),
				pg.TextWith{Content: "59.70"}.Make(),
				pg.CommandWith{Label: "Order Now"}.Make(),
			},
			{
				pg.TextWith{Content: "12"}.Make(),
				pg.TextWith{Content: "PB and cheese crackers"}.Make(),
				pg.TextWith{Content: "1.29"}.Make(),
				pg.TextWith{Content: "15.48"}.Make(),
				pg.CommandWith{Label: "Order Now"}.Make(),
			},
			{
				pg.TextWith{Content: "5"}.Make(),
				pg.TextWith{Content: "Frito Lays corn chips, 2 oz bag"}.Make(),
				pg.TextWith{Content: "1.78"}.Make(),
				pg.TextWith{Content: "8.90"}.Make(),
				pg.CommandWith{Label: "Order Now"}.Make(),
			},
		},
	}.Make()

	pgui.SetGUI(txt1, t)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}
