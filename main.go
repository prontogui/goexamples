// Copyright 2024-2026 ProntoGUI, LLC
// ProntoGUI™ is a trademark of ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"fmt"

	pg "github.com/prontogui/golib"
)

var (
	demo = flag.String("demo", "bingo", "Specify the demo to run.  Choices are:  frame, list1, list2, list3, table1, table2, table3, bingo, textfield, timer, misc, blank, snackbar, image, or simple.")
)

func main() {

	flag.Parse()

	pgui := pg.NewProntoGUI()
	err := pgui.StartServing("127.0.0.1", 50053)

	if err != nil {
		fmt.Printf("Error trying to start server:  %s", err.Error())
		return
	}

	switch *demo {
	case "frame":
		fmt.Printf("Running frame demo\n")
		frameDemo(pgui)
	case "list1":
		fmt.Printf("Running list1 demo\n")
		listDemo1(pgui)
	case "list2":
		fmt.Printf("Running list2 demo\n")
		listDemo2(pgui)
	case "list3":
		fmt.Printf("Running list3 demo\n")
		listDemo3(pgui)
	case "textfield":
		fmt.Printf("Running textfield demo\n")
		textFieldDemo(pgui)
	case "misc":
		fmt.Printf("Running misc demo\n")
		miscDemo(pgui)
	case "file":
		fmt.Print("Running file demo\n")
		filedemo(pgui)
	case "table1":
		fmt.Print("Running table demo 1\n")
		tabledemo1(pgui)
	case "table2":
		fmt.Print("Running table demo 2\n")
		tabledemo2(pgui)
	case "table3":
		fmt.Print("Running table demo 3\n")
		tabledemo3(pgui)
	case "bingo":
		fmt.Print("Running bingo board demo 3\n")
		bingoDemo(pgui)
	case "timer":
		fmt.Print("Running timer demo\n")
		timerdemo(pgui)
	case "blank":
		fmt.Print("Running blank demo\n")
		blankDemo(pgui)
	case "simple":
		fmt.Print("Running simple demo\n")
		simpleDemo(pgui)
	case "snackbar":
		fmt.Print("Running snackbar demo\n")
		snackbarDemo(pgui)
	case "image":
		fmt.Print("Running image demo\n")
		imageDemo(pgui)
	default:
		fmt.Printf("Unknown demo selected at command line.  Exiting.")
	}

	pgui.StopServing()
}
