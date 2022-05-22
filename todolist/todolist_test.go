package todolist

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode("test")
}

type FakeDatabase struct{}

func NewFakeDatabase() *FakeDatabase {
	return &FakeDatabase{}
}

func (d *FakeDatabase) Create(t Todolist) (Todolist, error) {
	return Todolist{}, nil
}

func (d *FakeDatabase) GetAll() (Todolists, error) {
	return Todolists{}, nil
}

func (d *FakeDatabase) Update(id string, t Todolist) (Todolist, error) {
	return Todolist{
		Id:   id,
		Task: t.Task,
	}, nil
}

func (d *FakeDatabase) Delete(id string) (bool, error) {
	return true, nil
}

func TestSuccessfulCreateTodolist(t *testing.T) {
	want := 201

	result := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(result)

	// Init http to avoi nil issue
	// ref : `https://stackoverflow.com/questions/57733801/how-to-set-mock-gin-context-for-bindjson/67034058#67034058`
	context.Request = &http.Request{}
	context.Set("tasks", "TEST")

	database := NewFakeDatabase()
	handler := NewTodolistHandler(database)
	handler.Create(context)

	if result.Code != want {
		t.Errorf("Expect %d but got %d", want, result.Code)
	}
}

func TestSuccessfulGetTodolists(t *testing.T) {
	want := 200

	result := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(result)

	context.Request = &http.Request{}

	database := NewFakeDatabase()
	handler := NewTodolistHandler(database)
	handler.GetAll(context)

	if result.Code != want {
		t.Errorf("Expect %d but got %d", want, result.Code)
	}
}

func TestSuccessfulUpdateTodolist(t *testing.T) {
	want := 200

	result := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(result)

	context.Request = &http.Request{}
	context.Set("id", "TEST")
	context.Set("task", "UPDATED_TASK")

	database := NewFakeDatabase()
	handler := NewTodolistHandler(database)
	handler.Update(context)

	if result.Code != want {
		t.Errorf("Expect %d but got %d", want, result.Code)
	}
}

func TestSuccessfulDeleteTodolist(t *testing.T) {
	want := 200

	result := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(result)

	context.Request = &http.Request{}
	context.Set("id", "TEST")

	database := NewFakeDatabase()
	handler := NewTodolistHandler(database)
	handler.Delete(context)

	if result.Code != want {
		t.Errorf("Expect %d but got %d", want, result.Code)
	}
}
