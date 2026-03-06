// Copyright 2024-2026 ProntoGUI, LLC
// ProntoGUI™ is a trademark of ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func tabledemo1(pgui pg.ProntoGUI) {

	txt1 := pg.TextWith{Content: "Default table embodiment"}.Make()

	t := pg.TableWith{
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
	}.Make().MakeHeadingsVA("Qty", "Description", "Price", "Subtotal", "Action")

	pgui.SetGUI(txt1, t)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}

func tabledemo2(pgui pg.ProntoGUI) {

	txt1 := pg.TextWith{Content: "Paged table embodiment"}.Make()

	t := pg.TableWith{
		Embodiment: "embodiment:paged, rowsPerPage:5",
		Rows: [][]pg.Primitive{
			{
				pg.TextWith{Content: "Apples"}.Make(),
				pg.TextWith{Content: "$10.00"}.Make(),
				pg.TextWith{Content: "In stock"}.Make(),
			},
			{
				pg.TextWith{Content: "Oranges"}.Make(),
				pg.TextWith{Content: "$5.00"}.Make(),
				pg.TextWith{Content: "In stock"}.Make(),
			},
			{
				pg.TextWith{Content: "Bananas"}.Make(),
				pg.TextWith{Content: "$3.00"}.Make(),
				pg.TextWith{Content: "Out of stock"}.Make(),
			},
			{
				pg.TextWith{Content: "Mangos"}.Make(),
				pg.TextWith{Content: "$10.00"}.Make(),
				pg.TextWith{Content: "In stock"}.Make(),
			},
			{
				pg.TextWith{Content: "Pears"}.Make(),
				pg.TextWith{Content: "$5.00"}.Make(),
				pg.TextWith{Content: "In stock"}.Make(),
			},
			{
				pg.TextWith{Content: "Grapes"}.Make(),
				pg.TextWith{Content: "$3.00"}.Make(),
				pg.TextWith{Content: "Out of stock"}.Make(),
			},
			{
				pg.TextWith{Content: "Raspberries"}.Make(),
				pg.TextWith{Content: "$10.00"}.Make(),
				pg.TextWith{Content: "In stock"}.Make(),
			},
			{
				pg.TextWith{Content: "Blueberries"}.Make(),
				pg.TextWith{Content: "$5.00"}.Make(),
				pg.TextWith{Content: "In stock"}.Make(),
			},
			{
				pg.TextWith{Content: "Plums"}.Make(),
				pg.TextWith{Content: "$5.00"}.Make(),
				pg.TextWith{Content: "Out of stock"}.Make(),
			},
		},
	}.Make().MakeHeadingsVA("Fruit", "Price/Lb.", "Stock")

	pgui.SetGUI(txt1, t)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}

func tabledemo3(pgui pg.ProntoGUI) {

	txt1 := pg.TextWith{Content: "DOS character matrix using table embodiment"}.Make()

	ibmScreen := []string{
		"The IBM Personal Computer DOS",
		"Version 2.00 (C)Copyright IBM Corp 1981, 1982, 1983",
		"",
		"A>_",
	}

	// Where the keyboard cursor is located
	cursorRow := 0
	cursorCol := 0

	// For a given DOS row and column, returns the character to display from ibmScreen information
	lookupChar := func(row int, col int) string {
		if row < 0 || row >= len(ibmScreen) {
			return ""
		}

		line := ibmScreen[row]

		if col < 0 || col >= len(line) {
			return ""
		}

		charText := line[col : col+1]

		// Keyboard cursor encountered?
		if charText == "_" {
			cursorRow = row
			cursorCol = col
		}

		return charText
	}

	// All the cells in the DOS screen, to be represented as Text primitives
	rows := make([][]pg.Primitive, 0)

	// Text primitive embodiment applied to each cell
	emb := "width:10,height:17.7,horizontalTextAlignment:center,color:#FFFFFFFF"

	// Fill in each cell
	for i := range 25 {
		cols := make([]pg.Primitive, 80)
		for j := range 80 {
			cols[j] = pg.TextWith{Content: lookupChar(i, j), Embodiment: emb}.Make()
		}
		rows = append(rows, cols)
	}

	// Example of mouse cursor
	p := rows[5][25].(*pg.Text)
	p.SetEmbodiment(emb + ",backgroundColor:#FFFFFFFF")

	// Use a table to display the cells
	t := pg.TableWith{
		Rows:       rows,
		Embodiment: "width:818,height:460.5,borderAll:2,marginAll:5,paddingAll:2,borderColor:#FF00FF00,backgroundColor:#FF000000",
	}.Make()

	// Timer for flashing the keyboard cursor
	cursorTimer := pg.TimerWith{PeriodMs: 500}.Make()
	tf := pg.TextFieldWith{TextEntry: "ABC", Embodiment: ""}.Make()

	// Construct the GUI
	pgui.SetGUI(txt1, t, tf, cursorTimer)

	// React to events
	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}

		if cursorTimer.Issued() {
			cursorText := rows[cursorRow][cursorCol].(*pg.Text)
			if cursorText.Content() == "" {
				cursorText.SetContent("_")
			} else {
				cursorText.SetContent("")
			}
		}
	}
}
