package main

import (
	"fmt"
)

func main() {
	fmt.Println("***** Another Obligatory To Do List App *****")
	fmt.Println()
	list := make(map[int]string)
	//Check if array is empty
	currentTask := ""

	fmt.Scan(currentTask);

	list[0] = addNewTask(currentTask)

	fmt.Print(list[0])

}

func addNewTask (currentTask string) string {
	return "Baboom"
}

