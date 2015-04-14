package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/zoli/emru/list"
	"github.com/zoli/emru/list/task"
)

func TestLists(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.New(), "b": list.New()}

	tsk := task.New("test", "test body")
	lh.ls["a"].Add(tsk)

	tCreatedAt, _ := tsk.CreatedAt.MarshalJSON()
	laCreatedAt, _ := lh.ls["a"].CreatedAt.MarshalJSON()
	lbCreatedAt, _ := lh.ls["b"].CreatedAt.MarshalJSON()
	exp := fmt.Sprintf(`{"a":{"tasks":[{"id":%d,"title":"%s","body":"%s","done":%t,"created_at":%s}],"created_at":%s},"b":{"tasks":[],"created_at":%s}}`,
		tsk.Id, tsk.Title, tsk.Body, tsk.Done, string(tCreatedAt),
		string(laCreatedAt), string(lbCreatedAt))

	if err := lh.lists(); err != nil {
		t.Fatal(err)
	}
	if exp != string(lh.data) {
		t.Errorf("Expected %s but got %s", exp, lh.data)
	}
}

func TestList(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.New()}

	tsk := task.New("test", "test body")
	lh.ls["a"].Add(tsk)

	tCreatedAt, _ := tsk.CreatedAt.MarshalJSON()
	lCreatedAt, _ := lh.ls["a"].CreatedAt.MarshalJSON()
	exp := fmt.Sprintf(
		`{"tasks":[{"id":%d,"title":"%s","body":"%s","done":%t,"created_at":%s}],"created_at":%s}`,
		tsk.Id, tsk.Title, tsk.Body, tsk.Done, string(tCreatedAt),
		string(lCreatedAt))

	if err := lh.list("a"); err != nil {
		t.Fatal(err)
	}
	if string(lh.data) != exp {
		t.Errorf("Expected %s but got %s", exp, lh.data)
	}
}

func TestNewList(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{}

	exp := task.New("test", "test body")
	tskjs, _ := json.Marshal(exp)
	data := fmt.Sprintf(`{"name":"a","tasks":[%s]}`, string(tskjs))
	buf := bytes.NewBufferString(data)

	if req, err := http.NewRequest("POST", "/lists", buf); err != nil {
		t.Fatal(err)
	} else {
		lh.req = req
	}
	if err := lh.newList(); err != nil {
		t.Fatal(err)
	}

	if _, exist := lh.ls["a"]; !exist {
		t.Fatal("List with name a should exist")
	}
	if tsk, err := lh.ls["a"].Get(exp.Id); err != nil {
		t.Fatal(err)
	} else if tsk != *exp {
		t.Errorf("Expected %v but got %v", *exp, tsk)
	}
}

func TestDeleteList(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.New()}

	if err := lh.deleteList("b"); err == nil {
		t.Error("Expected error on deleting not existing list")
	}
	if err := lh.deleteList("a"); err != nil {
		t.Error(err)
	}
	if _, exist := lh.ls["a"]; exist {
		t.Error("Expected list with name a be deleted")
	}
}

func TestTasks(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.New()}

	tsk := task.New("title", "test body")
	lh.ls["a"].Add(tsk)

	if err := lh.tasks(lh.ls["a"]); err != nil {
		t.Fatal(err)
	}
	tskjs, _ := json.Marshal(tsk)
	exp := fmt.Sprintf(`[%s]`, string(tskjs))
	if string(lh.data) != exp {
		t.Errorf("Expected %s but got %s", exp, string(lh.data))
	}
}

func TestTask(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.New()}

	tsk := task.New("title", "test body")
	lh.ls["a"].Add(tsk)

	if err := lh.task(lh.ls["a"], tsk.Id); err != nil {
		t.Fatal(err)
	}
	tskjs, _ := json.Marshal(tsk)
	if string(lh.data) != string(tskjs) {
		t.Errorf("Expected %s but got %s", string(tskjs),
			string(lh.data))
	}
}

func TestNewTask(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.New()}

	tsk := task.New("title", "test body")
	tsk.Id = 1
	tskjs, _ := json.Marshal(tsk)
	buf := bytes.NewBufferString(string(tskjs))

	req, err := http.NewRequest("POST", "/lists/a/tasks", buf)
	if err != nil {
		t.Fatal(err)
	}
	lh.req = req

	if err := lh.newTask(lh.ls["a"]); err != nil {
		t.Fatal(err)
	}
	if tsk2, err := lh.ls["a"].Get(tsk.Id); err != nil {
		t.Fatal(err)
	} else if tsk2.Title != "title" {
		t.Errorf("Expected task title 'title' but got %s", tsk2.Title)
	}
}

func TestUpdateTask(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.New()}

	tsk := task.New("test", "test body")
	lh.ls["a"].Add(tsk)
	tsk.Title = "new"
	tskjs, _ := json.Marshal(tsk)
	buf := bytes.NewBufferString(string(tskjs))

	req, err := http.NewRequest("PUT", "/lists/a/tasks/0", buf)
	if err != nil {
		t.Fatal(err)
	}
	lh.req = req

	if err := lh.updateTask(lh.ls["a"], tsk.Id); err != nil {
		t.Fatal(err)
	}
	if tsk2, err := lh.ls["a"].Get(tsk.Id); err != nil {
		t.Fatal(err)
	} else if tsk2.Title != "new" {
		t.Errorf("Expected task title 'new' but got %s", tsk2.Title)
	}
}

func TestDeleteTask(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.New()}

	tsk := task.New("test", "test body")
	lh.ls["a"].Add(tsk)

	if err := lh.deleteTask(lh.ls["a"], tsk.Id); err != nil {
		t.Fatal(err)
	}
	if len(lh.ls["a"].Tasks()) != 0 {
		t.Error("Expected to tasks be empty")
	}
}
