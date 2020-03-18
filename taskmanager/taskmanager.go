package taskmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
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

func (t *Tasks) Add(description, created string, completed bool) {
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
}

//GetTask using id
func (t *Tasks) GetTask(Id int) *Task {
	for _, v := range *t {
		if v.Id == Id {
			return &v
		}
	}
	fmt.Println("Task not found")
	os.Exit(1)
	return nil
}

//SetComplete flag of id
func (t *Tasks) SetComplete(Id int) {
	task := t.GetTask(Id)
	(*task).Completed = true

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

//GetLastId
func (t Tasks) GetLastId() int {
	totalTasks := len(t)
	if totalTasks <= 0 {
		return 0
	}
	id := t[0].Id
	for _, task := range t {
		if id >= task.Id {
			id = task.Id
		}
	}
	return id
}

func readDb() Tasks {
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
