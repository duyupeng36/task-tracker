# Task Tracker CLI by RoadMap

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh).

## Run

```shell
# download
git clone https://github.com/duyupeng36/task-tracker.git

# change dir
cd roadmap-projects/task-tracker

# Installation dependency 
go mod tidy

# compile
go build .
```


## use example

```shell
# help information
task-tracker --help  

# Add 
task-tracker create -d "Learn Golang" # todo task
task-tracker create -s "in-progress" -d "Learn Golang" # in-progress task
task-tracker create -s "done" -d "Learn Golang" # Done task

# Delete
task-tracker delete -i 2 # delete task which id is 2

# Update
task-tracker update -i 2 -s "done" # update task status
task-tracker update -i 2 -d "Learn Python" # update task description

# List all tasks
task-tracker list

# list tasks which stats we give
task-tracker list -s "todo"
```
