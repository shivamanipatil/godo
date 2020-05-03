# Godo
Godo is simple CLI app to manage your todo tasks.

## Features
* Add, update, removea andread tasks.
* At jobs((man page)[https://linux.die.net/man/1/at]) can be created for sending remainder notification for a task at particular day and time.
* List of tasks, list of pending tasks, number of pending tasks.
* Tasks are stored in .json format.

## Steps 
1. Clone this repo and build the binary.
```git clone https://github.com/shivamanipatil/godo.git
go build -o godo
```
1. To make this binary executable from anywhere. Add following to .bashrc. Replace <..> with full path to the folder containing your binary created in step 1.
```
export PATH=$PATH:</full/path/to/binary>
```
 
# Usage
`godo help` will display help menu. 

## Made with
https://github.com/fatih/color
