package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func listDemo2(pgui pg.ProntoGUI) {

	item1 := pg.CardWith{
		LeadingItem:  pg.CheckWith{}.Make(),
		MainItem:     pg.TextWith{Content: "Wheel"}.Make(),
		SubItem:      pg.TextWith{Content: "6 with rubber tread"}.Make(),
		TrailingItem: pg.CommandWith{Label: "..."}.Make(),
	}.Make()

	item2 := pg.CardWith{
		LeadingItem:  pg.CheckWith{}.Make(),
		MainItem:     pg.TextWith{Content: "Brake"}.Make(),
		SubItem:      pg.TextWith{Content: "Carbide polymer"}.Make(),
		TrailingItem: pg.CommandWith{Label: "..."}.Make(),
	}.Make()

	item3 := pg.CardWith{
		LeadingItem:  pg.CheckWith{}.Make(),
		MainItem:     pg.TextWith{Content: "Engine"}.Make(),
		SubItem:      pg.TextWith{Content: "4 cyclinder gasoline"}.Make(),
		TrailingItem: pg.CommandWith{Label: "..."}.Make(),
	}.Make()

	item4 := pg.CardWith{
		LeadingItem:  pg.CheckWith{}.Make(),
		MainItem:     pg.TextWith{Content: "Transmission"}.Make(),
		SubItem:      pg.TextWith{Content: "3 speed manual"}.Make(),
		TrailingItem: pg.CommandWith{Label: "..."}.Make(),
	}.Make()

	list := pg.ListWith{
		ListItems: []pg.Primitive{item1, item2, item3, item4},
	}.Make()

	pgui.SetGUI(list)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}
