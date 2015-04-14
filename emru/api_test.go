package main

import (
	"fmt"
	"testing"

	"github.com/zoli/emru/list"
	"github.com/zoli/emru/list/task"
)

func TestLists(t *testing.T) {

}

func TestList(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.NewList()}

	tsk := task.NewTask("test", "test body")
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

}

func TestDeleteList(t *testing.T) {
	lh := new(ListHandler)
	lh.ls = list.Lists{"a": list.NewList()}

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
