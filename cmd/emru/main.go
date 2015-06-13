package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/zoli/emru/emru"
)

type SocketTransport struct{ path string }

var (
	client   *http.Client
	bodyType = "application/json"

	magenta = color.New(color.FgMagenta, color.Bold)
	cyan    = color.New(color.FgCyan, color.Bold)
	white   = color.New(color.FgWhite)
	green   = color.New(color.FgGreen)

	sep   = "|"
	name  = "    name    "
	tasks = "    tasks    "
	id    = "    id    "
	title = "    title    "
	done  = "    done    "
)

func (d SocketTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	dial, err := net.Dial("unix", d.path)
	if err != nil {
		return nil, err
	}

	con := httputil.NewClientConn(dial, nil)
	defer con.Close()

	return con.Do(req)
}

func main() {
	client = &http.Client{Transport: SocketTransport{path: "/tmp/emru.sock"}}

	args := os.Args
	if len(args) == 1 {
		args = append(args, "")
	}
	args = args[1:]

	if args[0] == "lists" {
		if len(args) == 1 {
			if err := printLists(); err != nil {
				println(err.Error())
			}
			return
		}
		switch args[1] {
		case "add":
			newList(args[2])
		case "delete":
			deleteList(args[2])
		default:
		}
	} else {
		if len(args) == 1 {
			printList(args[0])
			return
		}
		switch args[1] {
		case "add":
			newTask(args[0], args[2])
		case "toggle":
			toggleTask(args[0], args[2])
		case "delete":
			deleteTask(args[0], args[2])
		default:
		}
	}
}

func printLists() error {
	resp, err := client.Get("unix:///lists")
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	var ls emru.Lists
	if err = decoder.Decode(&ls); err != nil {
		return err
	}

	white.Print("\n\n")
	magenta.Print(name)
	white.Print(sep)
	magenta.Println(tasks)
	for i := 0; i < len(name)+len(tasks)+len(sep); i++ {
		white.Print("-")
	}
	white.Print("\n")

	for n, l := range ls {
		margin := len(name) - len(n)
		for i := 0; i < margin/2; i++ {
			green.Print(" ")
		}
		if margin%2 != 0 {
			margin += 1
		}
		margin = margin / 2
		green.Print(n)
		for i := 0; i < margin; i++ {
			green.Print(" ")
		}
		white.Print(sep)

		num := strconv.Itoa(len(l.Tasks()))
		margin = len(tasks) - len(num)
		for i := 0; i < margin/2; i++ {
			green.Print(" ")
		}
		if margin%2 != 0 {
			margin += 1
		}
		margin = margin / 2
		green.Print(num)
		for i := 0; i < margin; i++ {
			green.Print(" ")
		}
		white.Print("\n")

		for i := 0; i < len(name)+len(tasks)+len(sep); i++ {
			white.Print("-")
		}
		white.Print("\n")
	}
	white.Print("\n\n")

	return nil
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

func deleteList(name string) error {
	req, err := http.NewRequest("DELETE", "unix:///lists/"+name, nil)
	if err != nil {
		return err
	}
	_, err = client.Do(req)

	return err
}

func printList(name string) error {
	resp, err := client.Get("unix:///lists/" + name)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	var list struct {
		Name  string       `json:"name"`
		Tasks []*emru.Task `json:"tasks"`
	}
	if err = decoder.Decode(&list); err != nil {
		return err
	}

	white.Print("\n")
	cyan.Print(name)
	white.Println(" tasks:\n")
	magenta.Print(id)
	white.Print(sep)
	magenta.Print(title)
	white.Print(sep)
	magenta.Println(done)
	for i := 0; i < len(id)+len(sep)*2+len(title)+len(done); i++ {
		white.Print("-")
	}
	white.Print("\n")

	for _, t := range list.Tasks {
		num := strconv.Itoa(t.ID)
		margin := len(id) - len(num)
		for i := 0; i < margin/2; i++ {
			green.Print(" ")
		}
		if margin%2 != 0 {
			margin += 1
		}
		margin = margin / 2
		green.Print(num)
		for i := 0; i < margin; i++ {
			green.Print(" ")
		}
		white.Print(sep)

		margin = len(title) - len(t.Title)
		for i := 0; i < margin/2; i++ {
			green.Print(" ")
		}
		if margin%2 != 0 {
			margin += 1
		}
		margin = margin / 2
		green.Print(t.Title)
		for i := 0; i < margin; i++ {
			green.Print(" ")
		}
		white.Print(sep)

		margin = len(done) - 1
		for i := 0; i < margin/2; i++ {
			green.Print(" ")
		}
		if margin%2 != 0 {
			margin += 1
		}
		margin = margin / 2
		if t.Done.Val() {
			green.Print("✔")
		} else {
			green.Print("✘")
		}
		for i := 0; i < margin; i++ {
			green.Print(" ")
		}
		white.Print("\n")
	}
	white.Print("\n\n")

	return nil
}

func newTask(name string, title string) error {
	task := emru.NewTask(title, "")
	buf, err := json.Marshal(&task)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("unix:///lists/%s/tasks", name)
	_, err = client.Post(url, bodyType, bytes.NewBuffer(buf))
	return err
}

func toggleTask(name string, id string) error {
	url := fmt.Sprintf("unix:///lists/%s/tasks/%s", name, id)
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

func deleteTask(name string, id string) error {
	url := fmt.Sprintf("unix:///lists/%s/tasks/%s", name, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = client.Do(req)

	return err
}
