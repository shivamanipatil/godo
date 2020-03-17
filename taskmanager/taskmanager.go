package taskmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
)

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
	json.Unmarshal(byteValue, &tasks)
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
		os.Remove(dbFilePath())
	}
}

func dbFilePath() string {
	env := os.Getenv("TASK_DB_PATH")
	return filepath.Join(filepath.Clean(env), dbFileName)
}

func createDbFileIfNotExist() {
	if _, err := os.Stat(dbFilePath()); os.IsNotExist(err) {
		os.Create(dbFilePath())
	}
}

//checkEnv checks the environment variables
func checkEnv() {
	_, exists := os.LookupEnv("TASK_DB_PATH")
	if !exists {
		homePath, homePathExists := os.LookupEnv("HOME")
		if homePathExists {
			os.Setenv("TASK_DB_PATH", homePath)
		} else {
			fmt.Println("Either set HOME env else set TASK_DB_PATH env varible")
		}
	}
}
func init() {
	checkEnv()
	createDbFileIfNotExist()
}
