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
		"\t1) help -\n" +
		"\t\tHelp menu\n\n" +
		"\t2) add {description} -\n" +
		"\t\tadd task with description use without {}\n\n" +
		"\t3) delete id -\n" +
		"\t\tDelete task with given id\n\n" +
		"\t4) completed id -\n" +
		"\t\tSet completed for for given id\n\n" +
		"\t5) pending -\n" +
		"\t\tPrinting number of pending tasks\n\n" +
		"\t6) list -\n" +
		"\t\tPrints all tasks with id as first column\n\n" +
		"\t7) listPending -\n" +
		"\t\tPrints all pending tasks\n\n")

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

	} else if commandName == "listPending" {
		(&tasks).ListPendingTasks()
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
