package main

import (
	"fmt"
	//"maps"
)

func main() {
	fmt.Println("***** Another Obligatory To Do List App *****")
	fmt.Println()
	//list := make(map[int][]string)
	keyword := map[string][]any{
		"help": {
		},
		"add": {
			"task",
			"function",
		},
		"remove": {
			"task",
		},
		"update": {
			"task",
		},
	}
	//Check if array is empty
	//currentTask := ""
	
	var command string

	fmt.Printf("What do you want to do? We have %d functions (Type \"help\" to get command list): ", len(keyword))
	fmt.Scan(&command)

	fmt.Println(keyword[command])
}

func helpFunctions() string {
	return "add - Add a new task\nremove - Remove a task\nupdate - Update a task"
}

func addTask(task string) string {
	return "You added a " + task;
}

func removeTask() string {
	return "You removed a task!"
}

func updateTask() string {
	return "You updated a task!"
}

func returnNumber() int {
	return 5
}