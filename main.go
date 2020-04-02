package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/shivamanipatil/GoTodo/taskmanager"
)

func drawTable(tasks *taskmanager.Tasks) {
	var checkString string
	for i := 0; i < len(*tasks); i++ {
		if (*tasks)[i].Completed {
			checkString = "[x]"
		} else {
			checkString = "[ ]"
		}
		yellow := color.New(color.FgYellow).SprintFunc()
		magenta := color.New(color.FgMagenta).SprintFunc()
		fmt.Printf("%d : %s %s %s\n", (*tasks)[i].Id, yellow(checkString), magenta((*tasks)[i].Created), (*tasks)[i].Description)
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
