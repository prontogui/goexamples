package main

import (
	"flag"
	"fmt"

	pg "github.com/prontogui/golib"
)

var (
	demo = flag.String("demo", "timer", "Specify the demo to run.  Choices are:  frame, list1, list2, list3, table, textfield, timer, misc, blank, simple, ...")
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
	case "table":
		fmt.Print("Running table demo\n")
		tabledemo(pgui)
	case "timer":
		fmt.Print("Running timer demo\n")
		timerdemo(pgui)
	case "blank":
		fmt.Print("Running blank demo\n")
		blankDemo(pgui)
	case "simple":
		fmt.Print("Running simple demo\n")
		simpleDemo(pgui)
	default:
		fmt.Printf("Unknown demo selected at command line.  Exiting.")
	}

	pgui.StopServing()
}
