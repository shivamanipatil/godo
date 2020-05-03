package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/shivamanipatil/GoTodo/taskmanager"
)

var (
	magenta = color.New(color.FgMagenta).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	white   = color.New(color.FgHiWhite).SprintFunc()
	red     = color.New(color.FgHiRed).SprintFunc()
)

func drawTable(tasks *taskmanager.Tasks) {
	var checkString string
	for i := 0; i < len(*tasks); i++ {
		if (*tasks)[i].Completed {
			checkString = "[x]"
		} else {
			checkString = "[ ]"
		}
		fmt.Printf("%d : %s %s %s\n", (*tasks)[i].Id, cyan(checkString), magenta((*tasks)[i].Created), white((*tasks)[i].Description))
	}
}

func helpMenu() {
	fmt.Printf("  GoToDo - Simple CLI todo app in golang\n\n" +
		"  USAGE :- Use options without {} \n " +
		"\t1) help -\n" +
		"\t\tHelp menu\n\n" +
		"\t2) add {description} -\n" +
		"\t\tadd task with description \n\n" +
		"\t3) delete {id} -\n" +
		"\t\tDelete task with given id\n\n" +
		"\t4) completed {id} -\n" +
		"\t\tSet completed for for given id\n\n" +
		"\t5) pending -\n" +
		"\t\tPrinting number of pending tasks\n\n" +
		"\t6) list -\n" +
		"\t\tPrints all tasks with id as first column\n\n" +
		"\t7) listPending -\n" +
		"\t\tPrints all pending tasks\n\n" +
		"\t8) schedule {id} {time} {date}-\n" +
		"\t\tSchedules a at job at specified time(e.g 23:56) and date(mm/dd/yyyy) for given id\n\n")

}

func main() {
	flag.CommandLine.SetOutput(os.Stdout)
	//Subcommands

	// 1)CRUD
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	delCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	//getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	//updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	// 2)Set Completed
	completeCommand := flag.NewFlagSet("completed", flag.ExitOnError)
	// 3)Schedule
	scheduleCommand := flag.NewFlagSet("scheduleAt", flag.ExitOnError)

	//Flags
	descriptionPtr := addCommand.String("desc", "", "Description of task. (Required)")
	idDPtr := delCommand.Int("id", -1, "id of task to be deleted. (Required)")
	idCPtr := completeCommand.Int("id", -1, "id of task to be deleted. (Required)")
	idSPtr := scheduleCommand.Int("id", -1, "id of task to be deleted. (Required)")
	time := scheduleCommand.String("time", "", "time at which notification should be sent.(required")
	date := scheduleCommand.String("date", "", "date at which notification should be sent.(required")

	flag.CommandLine.SetOutput(os.Stdout)
	tasks, err := taskmanager.ReadDb()
	if err != nil {
		fmt.Println(red(err))
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println(red("Enter subcommand!"))
		os.Exit(1)
	}
	switch os.Args[1] {
	case "help":
		helpMenu()
	case "add":
		addCommand.Parse(os.Args[2:])
		if *descriptionPtr == "" {
			fmt.Println(red("Give description!"))
			addCommand.PrintDefaults()
			os.Exit(1)
		}
		(&tasks).Add(*descriptionPtr)
	case "delete":
		delCommand.Parse(os.Args[2:])
		if *idDPtr == -1 {
			fmt.Println(red("Give id!"))
			delCommand.PrintDefaults()
			os.Exit(1)
		}
		(&tasks).Remove(*idDPtr)
	case "completed":
		completeCommand.Parse(os.Args[2:])
		if *idCPtr == -1 {
			fmt.Println(red("Give id!"))
			completeCommand.PrintDefaults()
			os.Exit(1)
		}
		(&tasks).SetCompleted(*idCPtr)
	case "list":
		drawTable(&tasks)
	case "pending":
		fmt.Println((&tasks).Pending())
	case "listPending":
		pendingTasks := (&tasks).ListPendingTasks()
		drawTable(&pendingTasks)
	case "scheduleAt":
		scheduleCommand.Parse(os.Args[2:])

		if (*time) == "" || (*date) == "" {
			fmt.Println(red("Provide time and date!"))
			scheduleCommand.PrintDefaults()
			os.Exit(1)
		}
		if *idSPtr == -1 {
			fmt.Println(red("Give id!"))
			scheduleCommand.PrintDefaults()
			os.Exit(1)
		}
		err := (&tasks).ScheduleTask(*idSPtr, (*time)+" "+(*date))
		if err != nil {
			fmt.Println(err)
			scheduleCommand.PrintDefaults()
			os.Exit(1)
		}
	default:
		fmt.Println(red("Incorrect command\n Use help subcommand for usage"))
		os.Exit(1)
	}

}
