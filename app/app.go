// Package app 用于构建 Task Tracker 命令行程序
package app

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"task-tracker/fs"
	"task-tracker/task"
	"time"
)

var App = New()

var (
	ErrAppNil       = errors.New("app can't be nil")
	ErrTaskNotFound = errors.New("task not found")
)

type TaskApp interface {
	AddTask(status, description string) (err error)
	DeleteTask(id int) error
	UpdateTask(id int, status, description string) error
	DisplayTasks(status string)
}

type application struct {
	tasks []*task.Task
}

func (a *application) AddTask(status, description string) (err error) {
	if a == nil {
		return ErrAppNil
	}
	createAt := task.JSONTime{Time: time.Now()}
	id, err := a.generateID()
	if err != nil {
		return err
	}
	a.tasks = append(a.tasks, &task.Task{ID: id, Status: status, Description: description, CreateAt: createAt})
	return a.writeTasks()
}

func (a *application) DeleteTask(id int) error {
	if a == nil {
		return ErrAppNil
	}

	for i, t := range a.tasks {
		if t.ID == id {
			a.tasks = append(a.tasks[:i], a.tasks[i+1:]...)
			return a.writeTasks()
		}
	}
	return ErrTaskNotFound
}

func (a *application) UpdateTask(id int, status, description string) error {
	var err error
	for _, t := range a.tasks {
		if t.ID == id {
			err = t.Update(status, description)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return err
	}

	return a.writeTasks()
}

func (a *application) DisplayTasks(status string) {
	for _, t := range a.tasks {
		switch {
		case t.Status == status:
			fmt.Println(t)
		case status == "":
			fmt.Println(t)
		}
	}
}

func (a *application) generateID() (int, error) {
	if a == nil {
		return 0, ErrAppNil
	}

	if len(a.tasks) == 0 {
		return 1, nil
	}

	if err := a.sortTasks(); err != nil {
		return 0, err
	}

	minID := a.tasks[0].ID

	if minID != 1 {
		return 1, nil
	}

	for i, j := 0, 1; j < len(a.tasks); i, j = i+1, j+1 {
		if a.tasks[j].ID-a.tasks[i].ID != 1 {
			return a.tasks[i].ID + 1, nil
		}
	}

	return a.tasks[len(a.tasks)-1].ID + 1, nil
}

func (a *application) sortTasks() error {
	if a == nil {
		return ErrAppNil
	}
	sort.Slice(a.tasks, func(i, j int) bool {
		return a.tasks[i].ID < a.tasks[j].ID
	})
	return nil
}

func (a *application) readTasks() (err error) {

	if a == nil {
		return ErrAppNil
	}

	a.tasks, err = fs.ReadTaskFromFile()
	if err != nil {
		return err
	}
	return nil
}

func (a *application) writeTasks() error {
	if a == nil {
		return ErrAppNil
	}

	if err := a.sortTasks(); err != nil {
		return err
	}

	return fs.WriteTaskToFile(a.tasks)
}

func New() TaskApp {
	app := &application{}
	err := app.readTasks()

	if err != nil {
		log.Fatalf("Read tasks from task.json failed: %v\n", err)
	}

	return app
}
