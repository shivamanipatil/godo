package taskmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

type (
	//Task represents task object
	Task struct {
		Id          int    `json:"id"`
		Description string `json:"description"`
		Created     string `json:"created"`
		Completed   bool   `json:"completed"`
	}
	//multiple tasks
	Tasks []Task
)

const (
	dbFileName = ".taskdb.json"
	timeLayout = "Mon, 01/02/06, 03:04PM"
)

func (t *Tasks) Add(description string) {
	task := Task{Id: (*t).GetLastId() + 1, Description: description, Created: time.Now().Format(timeLayout), Completed: false}
	*t = append(*t, task)
	writeDb(t)
}

//Remove swap given element with last element and remove
func (t *Tasks) Remove(Id int) {
	if len(*t) == 0 {
		fmt.Print("No tasks present")
		os.Exit(1)
	}
	task := t.GetTask(Id)
	lastTask := &((*t)[len(*t)-1])
	task.Completed = lastTask.Completed
	task.Created = lastTask.Created
	task.Description = lastTask.Description
	(*t) = (*t)[:len(*t)-1]
	writeDb(t)
}

//Schedules task to at job
func (t *Tasks) ScheduleTask(Id int, dateTime string) error {
	magenta := color.New(color.FgMagenta).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	f, err := os.Create("t.txt")
	if err != nil {
		return fmt.Errorf("Couldn't create file for at job!")
	}
	//Close and remove temp file
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println("close:", err)
		}
		err = os.Remove(f.Name())
		if err != nil {
			log.Println("remove:", err)
		}
	}()
	task := t.GetTask(Id)
	if task == nil {
		return fmt.Errorf("Task not found!")
	}
	taskDescription := task.Description
	_, err = f.WriteString("notify-send Remainder " + "\"" + taskDescription + "\"")
	if err != nil {
		return fmt.Errorf("Couldn't write notification command to text file!")
	}
	_, err = exec.Command("at", "-f", "t.txt", dateTime).Output()
	if err != nil {
		return fmt.Errorf("Couldn't schedule at jobF!")
	}
	fmt.Printf("%v : %v\n", cyan("Your job is scheduled for"), magenta(dateTime))
	return nil
}

//GetTask using id
func (t *Tasks) GetTask(Id int) *Task {
	for i, v := range *t {
		if v.Id == Id {
			return &((*t)[i])
		}
	}
	return nil
}

//SetComplete flag of id
func (t *Tasks) SetComplete(Id int) {
	task := t.GetTask(Id)
	task.Completed = true
	writeDb(t)
}

//Pending number of tasks
func (t *Tasks) Pending() int {
	n := 0
	for _, v := range *t {
		if !v.Completed {
			n++
		}
	}
	return n
}

func (t *Tasks) ListPendingTasks() {
	checkString := "[ ]"
	yellow := color.New(color.FgYellow).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	for _, v := range *t {
		if !v.Completed {
			fmt.Printf("%d : %s %s %s\n", v.Id, yellow(checkString), magenta(v.Created), v.Description)
		}
	}
}

//GetLastId
func (t *Tasks) GetLastId() int {
	totalTasks := len(*t)
	if totalTasks <= 0 {
		return 0
	}
	id := (*t)[0].Id
	for _, task := range *t {
		if id <= task.Id {
			id = task.Id
		}
	}
	return id
}

func ReadDb() Tasks {
	dbFile, err := os.Open(dbFilePath())
	if err != nil {
		fmt.Println(err)
	}
	defer dbFile.Close()
	byteValue, err := ioutil.ReadAll(dbFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(byteValue) <= 0 {
		return nil
	}
	var tasks Tasks
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return tasks
}

func writeDb(tasks *Tasks) {
	removeDbFile()
	bytesArr, _ := json.Marshal(*tasks)
	err := ioutil.WriteFile(dbFilePath(), bytesArr, 0644)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

}

func removeDbFile() {
	if _, err := os.Stat(dbFilePath()); os.IsExist(err) {
		err := os.Remove(dbFilePath())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func dbFilePath() string {
	env := os.Getenv("TASK_DB_PATH")
	return filepath.Join(filepath.Clean(env), dbFileName)
}

func createDbFileIfNotExist() {
	if _, err := os.Stat(dbFilePath()); os.IsNotExist(err) {
		_, err := os.Create(dbFilePath())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

//checkEnv checks the environment variables
func checkEnv() {
	_, exists := os.LookupEnv("TASK_DB_PATH")
	if !exists {
		homePath, homePathExists := os.LookupEnv("HOME")
		if homePathExists {
			err := os.Setenv("TASK_DB_PATH", homePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Either set HOME env else set TASK_DB_PATH env varible")
		}
	}
}
func init() {
	checkEnv()
	createDbFileIfNotExist()
}
