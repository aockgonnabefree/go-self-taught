package main

import (
	"fmt"
)

func main() {
	tasks := Tasks{}

	tasks.Add("Create TODO CLI app")
	tasks.Add("Go to Central Lad Prao")

	tasks.PrintList()
 
	tasks.Complete(0)
	tasks.Modify(1, "GO CENTRAL NOWWWW")
	
	tasks.PrintList()
	tasks.writeCsvFile("tasks.csv")

	tasks, err := readCsvFile("tasks.csv")

	if err != nil {
		fmt.Println("Please Check file is valid.")
		return
	}

	tasks.PrintList()
}