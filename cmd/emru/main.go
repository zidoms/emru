package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"

	"github.com/fatih/color"
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

var client *http.Client

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

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var ls emru.Lists
	if err = json.Unmarshal(data, &ls); err != nil {
		return err
	}

	magenta := color.New(color.FgMagenta, color.Bold)
	white := color.New(color.FgWhite)
	green := color.New(color.FgGreen)

	name := "    name    "
	tasks := "    tasks    "
	sep := "|"

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
	}
	white.Print("\n\n")

	return nil
}

func newList(title string) error {
	s := fmt.Sprintf(`{"name":"%s","tasks":[]}`, string(title))
	_, err := client.Post("unix:///lists", "application/json",
		bytes.NewBufferString(s))
	if err != nil {
		return err
	}
	return nil
}

func deleteList(id string) {}

func printList(name string)             {}
func newTask(name string, title string) {}
func toggleTask(name string, id string) {}
func deleteTask(name string, id string) {}
