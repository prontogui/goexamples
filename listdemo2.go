package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func listDemo2(pgui pg.ProntoGUI) {

	item1 := pg.GroupWith{GroupItems: []pg.Primitive{
		pg.CheckWith{}.Make(),
		pg.TextWith{Content: "Wheel"}.Make(),
		pg.TextWith{Content: "6 with rubber tread"}.Make(),
		pg.CommandWith{Label: "...", Embodiment: "outlined-button"}.Make(),
	}}.Make()
	item2 := pg.GroupWith{GroupItems: []pg.Primitive{
		pg.CheckWith{}.Make(),
		pg.TextWith{Content: "Brake"}.Make(),
		pg.TextWith{Content: "Carbide polymer"}.Make(),
		pg.CommandWith{Label: "..."}.Make(),
	}}.Make()
	item3 := pg.GroupWith{GroupItems: []pg.Primitive{
		pg.CheckWith{}.Make(),
		pg.TextWith{Content: "Engine"}.Make(),
		pg.TextWith{Content: "4 cyclinder gasoline"}.Make(),
		pg.CommandWith{Label: "..."}.Make(),
	}}.Make()
	item4 := pg.GroupWith{GroupItems: []pg.Primitive{
		pg.CheckWith{}.Make(),
		pg.TextWith{Content: "Transmission"}.Make(),
		pg.TextWith{Content: "3 speed manual"}.Make(),
		pg.CommandWith{Label: "..."}.Make(),
	}}.Make()
	list := pg.ListWith{
		ListItems:  []pg.Primitive{item1, item2, item3, item4},
		Embodiment: "card-list",
	}.Make()
	pgui.SetGUI(list)
	pgui.Wait()

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}
