package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func frameDemo(pgui pg.ProntoGUI) {

	choice1 := pg.ChoiceWith{Choice: "GUI #1", Choices: []string{"GUI #1", "GUI #2"}}.Make()
	cmd1 := pg.CommandWith{Label: "Frame 1"}.Make()
	cmd2 := pg.CommandWith{Label: "Frame 2"}.Make()
	cmd3 := pg.CommandWith{Label: "Frame 3"}.Make()
	cmd4 := pg.CommandWith{Label: "1 Dialog"}.Make()
	cmd5 := pg.CommandWith{Label: "3 Dialogs"}.Make()
	cmd6 := pg.CommandWith{Label: "Frame A"}.Make()
	cmd7 := pg.CommandWith{Label: "Frame B"}.Make()
	cmd8 := pg.CommandWith{Label: "Show Full-View Again"}.Make()

	f1 := pg.FrameWith{
		Embodiment: "{\"embodiment\":\"full-view\"}",
		FrameItems: []pg.Primitive{
			pg.TextWith{
				Content: "Frame 1",
			}.Make(),
			choice1,
			cmd1,
			cmd2,
			cmd3,
			cmd4,
			cmd5,
		},
	}.Make()

	f2 := pg.FrameWith{
		Embodiment: "{\"embodiment\":\"full-view\"}",
		FrameItems: []pg.Primitive{
			pg.TextWith{
				Content: "Frame 2",
			}.Make(),
			pg.CheckWith{}.Make(),
			choice1,
			cmd1,
			cmd2,
			cmd3,
			cmd4,
			cmd5,
		},
	}.Make()

	f3 := pg.FrameWith{
		Embodiment: "{\"embodiment\":\"full-view\"}",
		FrameItems: []pg.Primitive{
			pg.TextWith{
				Content: "Frame 3",
			}.Make(),
			pg.TextWith{Content: "Some random text..."}.Make(),
			choice1,
			cmd1,
			cmd2,
			cmd3,
			cmd4,
			cmd5,
		},
		Showing: true,
	}.Make()

	fa := pg.FrameWith{
		Embodiment: "{\"embodiment\":\"full-view\"}",
		FrameItems: []pg.Primitive{
			pg.TextWith{
				Content: "Frame A",
			}.Make(),
			choice1,
			cmd6,
			cmd7,
		},
	}.Make()

	fb := pg.FrameWith{
		Embodiment: "{\"embodiment\":\"full-view\"}",
		FrameItems: []pg.Primitive{
			pg.TextWith{
				Content: "Frame B",
			}.Make(),
			pg.CheckWith{}.Make(),
			choice1,
			cmd6,
			cmd7,
		},
	}.Make()

	d1 := pg.FrameWith{
		Embodiment: "{\"embodiment\":\"dialog-view\"}",
		FrameItems: []pg.Primitive{
			pg.TextWith{
				Content: "Dialog 1",
			}.Make(),
			pg.CheckWith{}.Make(),
		},
	}.Make()

	d2 := pg.FrameWith{
		Embodiment: "{\"embodiment\":\"dialog-view\"}",
		FrameItems: []pg.Primitive{
			pg.TextWith{
				Content: "Dialog 2",
			}.Make(),
			pg.CheckWith{}.Make(),
		},
	}.Make()

	d3 := pg.FrameWith{
		Embodiment: "{\"embodiment\":\"dialog-view\"}",
		FrameItems: []pg.Primitive{
			pg.TextWith{
				Content: "Dialog 3",
			}.Make(),
			pg.CheckWith{}.Make(),
		},
	}.Make()

	pgui.SetGUI(f1, f2, f3, d1, d2, d3, cmd8)

	for {
		updatedPrimitive, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
		if updatedPrimitive == choice1 {
			switch choice1.Choice() {
			case "GUI #1":
				fmt.Printf("Switching to GUI #1")
				pgui.SetGUI(f1, f2, f3, d1, d2, d3, cmd8)
			case "GUI #2":
				fmt.Printf("Switching to GUI #2")
				fa.SetShowing(true)
				fb.SetShowing(false)
				pgui.SetGUI(fa, fb, d3)
			}
		}
		if updatedPrimitive == cmd1 {
			fmt.Print("command #1 clicked")
			f1.SetShowing(false)
			f2.SetShowing(false)
			f3.SetShowing(false)
			d1.SetShowing(false)
			d2.SetShowing(false)
			d3.SetShowing(false)
		}
		if updatedPrimitive == cmd2 {
			fmt.Print("command #2 clicked")
			f1.SetShowing(false)
			f2.SetShowing(true)
			f3.SetShowing(false)
			d1.SetShowing(false)
			d2.SetShowing(false)
			d3.SetShowing(false)
		}
		if updatedPrimitive == cmd3 {
			fmt.Print("command #3 clicked")
			f1.SetShowing(false)
			f2.SetShowing(false)
			f3.SetShowing(true)
			d1.SetShowing(false)
			d2.SetShowing(false)
			d3.SetShowing(false)
		}
		if updatedPrimitive == cmd4 {
			fmt.Print("command #4 clicked")
			d1.SetShowing(true)
		}
		if updatedPrimitive == cmd5 {
			fmt.Print("command #5 clicked")
			d1.SetShowing(true)
			d2.SetShowing(true)
			d3.SetShowing(true)
		}
		if updatedPrimitive == cmd6 {
			fmt.Print("command #6 clicked")
			fa.SetShowing(true)
			fb.SetShowing(false)
		}
		if updatedPrimitive == cmd7 {
			fmt.Print("command #7 clicked")
			fa.SetShowing(false)
			fb.SetShowing(true)
		}
		if updatedPrimitive == cmd8 {
			f1.SetShowing(true)
		}
		if updatedPrimitive == d1 {
			fmt.Print("Dialog #1 dismissed")
		}
		if updatedPrimitive == d2 {
			fmt.Print("Dialog #2 dismissed")
		}
		if updatedPrimitive == d3 {
			fmt.Print("Dialog #3 dismissed")
		}
	}
}
