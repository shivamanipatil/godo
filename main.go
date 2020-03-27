package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shivamanipatil/GoTodo/taskmanager"
)

func drawTable(tasks *taskmanager.Tasks) {
	for i := 0; i < len(*tasks); i++ {
		fmt.Printf("%d, %s %s %t\n", (*tasks)[i].Id, (*tasks)[i].Description, (*tasks)[i].Created, (*tasks)[i].Completed)
	}
}

func helpMenu() {
	fmt.Printf("  GoToDo - Simple CLI todo app in golang\n\n" +
		"  USAGE\n" +
		"	1) h/help -\n" +
		"		Help menu\n\n" +
		"	2) add {description} -\n" +
		"		add task with description use without {}\n\n" +
		"	3) delete id -\n" +
		"		Delete task with given id\n\n" +
		"	4) completed id -\n" +
		"		Set completed for for given id\n\n" +
		"	5) pending -\n" +
		"		Printing number of pending tasks\n\n")

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Incorrect set of args")
		os.Exit(1)
	}
	commandName := os.Args[1]
	tasks := taskmanager.ReadDb()

	if commandName == "help" {
		helpMenu()
		os.Exit(0)
	} else if commandName == "list" {
		if tasks == nil {
			fmt.Println("No tasks")
			os.Exit(0)
		}
		drawTable(&tasks)
	} else if commandName == "pending" {
		fmt.Println((&tasks).Pending())
	} else {

		if len(os.Args) < 3 {
			fmt.Println("Give id")
			os.Exit(1)
		}
		arg := os.Args[2]
		switch commandName {
		case "add":
			description := arg
			(&tasks).Add(description)
		case "completed":
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			(&tasks).SetComplete(id)
		case "delete":
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(3)
			}
			(&tasks).Remove(id)
		}
	}

}
