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
	green   = color.New(color.FgHiGreen).SprintFunc()
)

func helpMenu() {
	fmt.Printf("  GoToDo - Simple CLI todo app in golang\n\n" +
		"  USAGE :- Use options without {} \n " +
		"\t1) help -\n" +
		"\t\tHelp menu\n\n" +
		"\t2) add -desc {description} -\n" +
		"\t\tadd task with description \n\n" +
		"\t3) delete -id {id} -\n" +
		"\t\tDelete task with given id\n\n" +
		"\t4) completed -id {id} -\n" +
		"\t\tSet completed for for given id\n\n" +
		"\t5) update -id {id} -desc {newDescriptionm} -\n" +
		"\t\tUpdates task with newDescriptionm\n\n" +
		"\t6) pending -\n" +
		"\t\tPrinting number of pending tasks\n\n" +
		"\t7) list -\n" +
		"\t\tPrints all tasks with id as first column\n\n" +
		"\t8) get -id {id}-\n" +
		"\t\tPrints task with given id\n\n" +
		"\t9) listPending -\n" +
		"\t\tPrints all pending tasks\n\n" +
		"\t10) scheduleAt -id {id} -time {time} -date {date}-\n" +
		"\t\tSchedules a at job at specified time(e.g 23:56) and date e.g(05/03/2020)(mm/dd/yyyy) for given id\n\n")

}

func main() {
	flag.CommandLine.SetOutput(os.Stdout)
	//Subcommands

	// 1)CRUD
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	delCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	// 2)Set Completed
	completeCommand := flag.NewFlagSet("completed", flag.ExitOnError)
	// 3)Schedule
	scheduleCommand := flag.NewFlagSet("scheduleAt", flag.ExitOnError)

	//Flags
	descriptionAPtr := addCommand.String("desc", "", "Description of task. (Required)")
	idGPtr := getCommand.Int("id", -1, "id of task. (Required)")
	idDPtr := delCommand.Int("id", -1, "id of task to be deleted. (Required)")
	idUPtr := updateCommand.Int("id", -1, "id of task to be deleted. (Required)")
	descriptionUPtr := updateCommand.String("desc", "", "New description of task. (Required)")
	idCPtr := completeCommand.Int("id", -1, "id of task to be completed. (Required)")
	idSPtr := scheduleCommand.Int("id", -1, "id of task to be scheduled. (Required)")
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
		if *descriptionAPtr == "" {
			fmt.Println(red("Give description!"))
			addCommand.PrintDefaults()
			os.Exit(1)
		}
		(&tasks).Add(*descriptionAPtr)
		fmt.Println(green("Successfully added task."))
	case "delete":
		delCommand.Parse(os.Args[2:])
		if *idDPtr == -1 {
			fmt.Println(red("Give id!"))
			delCommand.PrintDefaults()
			os.Exit(1)
		}
		(&tasks).Remove(*idDPtr)
		fmt.Println(green("Successfully deleted task."))
	case "update":
		updateCommand.Parse(os.Args[2:])
		if *idUPtr == -1 {
			fmt.Println(red("Give id!"))
			updateCommand.PrintDefaults()
			os.Exit(1)
		}
		if *descriptionUPtr == "" {
			fmt.Println(red("Give description!"))
			updateCommand.PrintDefaults()
			os.Exit(1)
		}
		err := tasks.Update(*idUPtr, *descriptionUPtr)
		if err != nil {
			fmt.Println(red(err))
			os.Exit(1)
		}
		fmt.Println(green("Successfully updated task."))
	case "get":
		getCommand.Parse(os.Args[2:])
		if *idGPtr == -1 {
			fmt.Println(red("Give id!"))
			getCommand.PrintDefaults()
			os.Exit(1)
		}
		task := tasks.GetTask(*idGPtr)
		task.DrawTask()
	case "completed":
		completeCommand.Parse(os.Args[2:])
		if *idCPtr == -1 {
			fmt.Println(red("Give id!"))
			completeCommand.PrintDefaults()
			os.Exit(1)
		}
		(&tasks).SetCompleted(*idCPtr)
		fmt.Println(green("Successfully completed task."))
	case "list":
		(&tasks).DrawTable()
	case "pending":
		fmt.Println((&tasks).Pending())
	case "listPending":
		pendingTasks := (&tasks).ListPendingTasks()
		(&pendingTasks).DrawTable()
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
		fmt.Printf("%v : %v\n", cyan("Your remainder is scheduled for"), magenta((*time)+" "+(*date)))
	default:
		fmt.Println(red("Incorrect command\n Use help subcommand for usage"))
		os.Exit(1)
	}

}
