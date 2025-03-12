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

	fmt.Print("Enter task into the list: ")
	fmt.Scan(&currentTask)

	list[0] = currentTask

	fmt.Print(list[0])

}

func addNewTask(currentTask string) string {
	return "Baboom"
}
