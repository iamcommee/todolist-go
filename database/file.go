package database

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"todolist/todolist"

	"github.com/google/uuid"
)

type File struct{}

func NewFileDatabase() *File {
	return &File{}
}

const FILE = "database.json"

func (d *File) Create(t todolist.Todolist) (todolist.Todolist, error) {
	data, _ := ioutil.ReadFile(FILE)

	var todolists todolist.Todolists
	json.Unmarshal([]byte(data), &todolists)

	t.Id = uuid.NewString()

	todolists.Todolists = append(todolists.Todolists, t)

	newData, _ := json.MarshalIndent(todolists, "", " ")

	err := ioutil.WriteFile(FILE, newData, 0644)

	return t, err
}

func (d *File) GetAll() (todolist.Todolists, error) {
	data, _ := ioutil.ReadFile(FILE)

	var todolists todolist.Todolists

	json.Unmarshal([]byte(data), &todolists)

	return todolists, nil
}

func (d *File) Update(id string, t todolist.Todolist) (todolist.Todolist, error) {
	data, _ := ioutil.ReadFile(FILE)

	var todolists todolist.Todolists
	json.Unmarshal([]byte(data), &todolists)

	// @todo : improve performance
	isUpdated := false
	for i := range todolists.Todolists {
		item := &todolists.Todolists[i]
		if item.Id == id {
			item.Task = t.Task
			isUpdated = true
			break
		}
	}

	if !isUpdated {
		return t, errors.New("id not found")
	}

	newData, _ := json.MarshalIndent(todolists, "", " ")

	ioutil.WriteFile(FILE, newData, 0644)

	return t, nil
}

func (d *File) Delete(id string) (bool, error) {
	data, _ := ioutil.ReadFile(FILE)

	var todolists todolist.Todolists
	json.Unmarshal([]byte(data), &todolists)

	// @todo : improve performance
	var newTodolists todolist.Todolists
	isDeleted := false
	for i := range todolists.Todolists {
		item := &todolists.Todolists[i]
		if item.Id == id {
			// item = nil
			newTodolists.Todolists = append(todolists.Todolists[:i], todolists.Todolists[i+1:]...)
			isDeleted = true
		}
	}

	if !isDeleted {
		return isDeleted, errors.New("id not found")
	}

	newData, _ := json.MarshalIndent(newTodolists, "", " ")

	ioutil.WriteFile(FILE, newData, 0644)

	return isDeleted, nil
}
