package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/limetext/log4go"
	"github.com/zoli/emru/api"
)

func main() {
	// TODO: sock location
	l, err := net.Listen("unix", "/tmp/emru.sock")
	if err != nil {
		log.Critical(err)
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		os.Remove("/tmp/emru.sock")
		os.Exit(0)
	}()

	if err = http.Serve(l, api.NewHandler()); err != nil {
		log.Critical(err)
		panic(err)
	}
}
