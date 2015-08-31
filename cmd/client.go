package cmd

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/zoli/emru/emru"
)

type SocketTransport struct{ path string }

func (d SocketTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	dial, err := net.Dial("unix", d.path)
	if err != nil {
		return nil, err
	}

	con := httputil.NewClientConn(dial, nil)
	defer con.Close()

	return con.Do(req)
}

var (
	bodyType = "application/json"
	client   = &http.Client{Transport: SocketTransport{path: "/tmp/emru.sock"}}
)

func getLists() (ls emru.Lists, err error) {
	resp, err := client.Get("unix:///lists")
	if err != nil {
		return
	}

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&ls); err != nil {
		return
	}

	return ls, nil
}

func newList(name string) error {
	var list struct {
		tasks []emru.Task
		Name  string
	}
	list.Name = name

	buf, err := json.Marshal(&list)
	if err != nil {
		return err
	}
	_, err = client.Post("unix:///lists", bodyType, bytes.NewBuffer(buf))

	return err
}

// TODO: change all delete to remove
func deleteList(name string) error {
	req, err := http.NewRequest("DELETE", "unix:///lists/"+name, nil)
	if err != nil {
		return err
	}
	_, err = client.Do(req)

	return err
}

func getTasks(list string) (tasks []emru.Task, err error) {
	resp, err := client.Get("unix:///lists/" + list + "/tasks")
	if err != nil {
		return
	}

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&tasks); err != nil {
		return
	}

	return tasks, nil
}

func newTask(list string, title string) error {
	task := emru.NewTask(title, "")
	url := "unix:///lists/" + list + "/tasks"

	buf, err := json.Marshal(&task)
	if err != nil {
		return err
	}
	_, err = client.Post(url, bodyType, bytes.NewBuffer(buf))

	return err
}

func toggleTask(list string, id string) error {
	url := "unix:///lists/" + list + "/tasks/" + id
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	var task emru.Task
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&task); err != nil {
		return err
	}

	task.Done.Toggle()
	buf, err := json.Marshal(&task)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	_, err = client.Do(req)

	return err
}

func deleteTask(list string, id string) error {
	url := "unix:///lists/" + list + "/tasks/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = client.Do(req)

	return err
}
