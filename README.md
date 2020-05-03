# Godo
Godo is simple CLI app to manage your todo tasks.

# Requirements
Go >= 1.13 recommended.[notify-send](http://manpages.ubuntu.com/manpages/xenial/man1/notify-send.1.html) and [at](https://linux.die.net/man/1/at) should be installed for scheduling the remainder notifications.

## Features
* Add, update, removea and read tasks.
* [At jobs](https://linux.die.net/man/1/at) can be created for sending remainder notification for a task at particular day and time.
* List of tasks, list of pending tasks, number of pending tasks.
* Tasks are stored in .json format.

## Steps 
1. Clone this repo and build the binary.
```git clone https://github.com/shivamanipatil/godo.git
go build -o godo
```
1. To make this binary executable from anywhere. Add following to .bashrc. E.g
```
export PATH="$PATH:/home/shivamani/funprojects/godo/"
```
Replace given path with full path to the folder containing your binary created in step 1.

 
# Usage
* `godo help` will display help menu.
* To add a task 
```
godo add -desc "Play football"
```
* To delete a task 
```
godo delete -id 1
```
* To update a task
```
godo update -id 1 -desc "This is new description"
```
* To get a task
```
godo get -id 1
```
* To list all tasks
```
godo list
```
* To list pending tasks
```
godo listPending
```
* To list pending number of tasks
```
godo pending
```
* To set complete a task
```
godo completed -id 1
```
* To schedule a task
```
godo scheduleAt -id 1 -time 17:19 -date 05/03/2020
```

 

## Made with
https://github.com/fatih/color
