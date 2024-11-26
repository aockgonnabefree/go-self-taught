package main

import (
	"errors"
	"fmt"
	"os"
	"time"
	"strconv"
	"github.com/aquasecurity/table"
)

type Task struct {
	Description 	string
	Status			bool
	CreatedAt		time.Time
	CompletedAt		*time.Time
}

type Tasks []Task

func (tasks *Tasks) Add(description string) {
	task := Task{
		Description: description,
		Status: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*tasks = append(*tasks, task)
}

func (tasks *Tasks) ValidateIndexRangeBound(index int) error {
	if index < 0 || index > len(*tasks) {
		err := errors.New("INVALID INDEX")
		fmt.Println(err)
		return err
	}
	return nil
}

func (tasks *Tasks) Delete(index int) error{
	t := *tasks

	if err := t.ValidateIndexRangeBound(index) ; err != nil  {
		return err
	}
	
	*tasks = append(t[:index], t[index+1:]...)
	// *tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
	
	return nil
}

func (tasks *Tasks) Complete(index int) error {
	t := *tasks

	if err := t.ValidateIndexRangeBound(index) ; err != nil {
		return err	
	}

	isCompleted := t[index].Status

	if !isCompleted {
		t[index].Status = true
		completedTime := time.Now()
		t[index].CompletedAt = &completedTime
	}

	return nil
}

func (tasks *Tasks) PrintList() {
	table := table.New(os.Stdout)	

	table.SetHeaders("Index", "Description", "Status","Created At", "Completed At")

	for i, task := range *tasks {
		status := "❌"
		completedTime := "NOT FINISH YET" 
		if task.Status {
			status = "✅" 
			completedTime = task.CompletedAt.Format(time.RFC1123)
		}

		table.AddRow(strconv.Itoa(i), task.Description, status, task.CreatedAt.Format(time.RFC1123), completedTime)
	}

	table.Render()
}

func (tasks *Tasks) Modify(index int, newDescription string) error {
	t := *tasks

	if err := t.ValidateIndexRangeBound(index) ; err != nil {
		return err
	}

	t[index].Description = newDescription
	return nil
}

func (task *Task) Slice() []string{
	description := task.Description
	status := strconv.FormatBool(task.Status)
	createdAt := task.CreatedAt.Format(time.RFC1123)
	
	var completedAt string
	if c := task.CompletedAt ; c == nil {
		completedAt = "nil"
	} else {
		completedAt = c.Format(time.RFC1123)
	}

	return []string{
		description,
		status,
		createdAt,
		completedAt }
}