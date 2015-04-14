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

}

func TestTask(t *testing.T) {

}

func TestNewTask(t *testing.T) {

}

func TestUpdateTask(t *testing.T) {

}

func TestDeleteTask(t *testing.T) {

}
